package cmd

import (
	"bytes"
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/n6o/go-grpc-imit-gen/templates"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

var rootCmd = &cobra.Command{
	Use:   "go-grpc-imit-gen",
	Short: "gRPC service imitation generator for golang",
	Long: `go-grpc-imit-gen generates gRPC service imitation.
This imitation runs as gRPC service and can be scheduled arbitrary responses.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := exec()
		if err != nil {
			// TODO: Use logger
			fmt.Printf("err: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var repo string
var dstDir string

func init() {
	// TODO: logger
	rootCmd.Flags().StringVar(&repo, "repo", "", "Repository file imported grpc services to imitate")
	rootCmd.Flags().StringVar(&dstDir, "dstDir", ".", "Destination directory to generate imitations")
}

func exec() error {
	if repo == "" {
		// TODO: Use logger
		return fmt.Errorf("--repo flag is required.")
	}

	dstDir = strings.TrimSuffix(dstDir, "/")

	pkgs, err := getTargetPackages()
	if err != nil {
		return fmt.Errorf("Failed to get target packages: %w", err)
	}

	for _, pkg := range pkgs {
		material, err := createMaterial(pkg)
		if err != nil {
			return fmt.Errorf("Failed to create materials: %w", err)
		}

		imit, err := genImit(material)
		if err != nil {
			return err
		}

		imitPath := fmt.Sprintf("%s/%s/%s", dstDir, material.PkgName, material.FileName)
		err = os.MkdirAll(filepath.Dir(imitPath), 0700)
		if err != nil {
			return fmt.Errorf("Failed to create dir: %w", err)
		}
		err = os.WriteFile(imitPath, imit, 0666)
		if err != nil {
			return fmt.Errorf("Failed to write file: %w", err)
		}
	}

	return nil
}

func getTargetPackages() ([]*build.Package, error) {
	paths, err := extractImportPathsFromRepo()
	if err != nil {
		return nil, err
	}

	return getPackages(paths)
}

func extractImportPathsFromRepo() ([]string, error) {
	f, err := parser.ParseFile(token.NewFileSet(), repo, nil, parser.ImportsOnly)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse repo file[%s]: %w", repo, err)
	}

	paths := []string{}
	for _, imprt := range f.Imports {
		path, err := strconv.Unquote(imprt.Path.Value)
		if err != nil {
			return nil, fmt.Errorf("Failed to unquote[%s]: %w", imprt.Path.Value, err)
		}
		paths = append(paths, path)
	}

	return paths, nil
}

func getPackages(paths []string) ([]*build.Package, error) {
	pkgs := []*build.Package{}
	for _, path := range paths {
		pkg, err := build.Import(path, "", build.FindOnly)
		if err != nil {
			return nil, fmt.Errorf("Failed to find package: %w", err)
		}

		pkgs = append(pkgs, pkg)
	}

	return pkgs, nil
}

type Material struct {
	FileName   string
	PkgName    string
	Imports    []*Imprt
	Interfaces []*ImitInterface
}

type Imprt struct {
	Alias string
	Path  string
}

type ImitInterface struct {
	Name       string
	UnimplName string
	RecvName   string
	FuncSigs   []*ImitFuncSig
}

type ImitFuncSig struct {
	Name    string
	Params  []*Param
	Results []*Param
}

type Param struct {
	Name string
	Type string
}

func createMaterial(pkg *build.Package) (*Material, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		Dir:  pkg.Dir,
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to load package[%s]: %w", pkg.Name, err)
	}
	if len(pkgs) != 1 {
		// TODO: log all pkgs
		return nil, fmt.Errorf("Unexpedted: package is not found or found more than one[%s]", pkg.Dir)
	}

	targetPkg := pkgs[0]
	pathToUsed, pathToAlias, err := prepareImports(targetPkg)
	if err != nil {
		return nil, err
	}

	targetInterfaceNames, err := findUnimplementedServer(targetPkg)
	if err != nil {
		return nil, err
	}

	ifcs, err := createImitInterfaces(targetPkg, targetInterfaceNames, pathToUsed, pathToAlias)
	if err != nil {
		return nil, err
	}

	imports := createImports(pathToUsed, pathToAlias)

	return &Material{
		FileName:   "imit_" + targetPkg.Name + ".go",
		PkgName:    "imit_" + targetPkg.Name,
		Imports:    imports,
		Interfaces: ifcs,
	}, nil
}

func prepareImports(pkg *packages.Package) (map[string]bool, map[string]string, error) {
	pathToUsed := map[string]bool{}
	pathToAlias := map[string]string{}
	for _, syntax := range pkg.Syntax {
		for _, imprt := range syntax.Imports {
			// alias isn't defined
			if imprt.Name == nil {
				unquoted, err := strconv.Unquote(imprt.Path.Value)
				if err != nil {
					return nil, nil, fmt.Errorf("Failed to unquote: %w", err)
				}
				pathToUsed[unquoted] = false
				continue
			}

			//  alias is blank.
			if imprt.Name.Name == "_" {
				continue
			}

			unquoted, err := strconv.Unquote(imprt.Path.Value)
			if err != nil {
				return nil, nil, fmt.Errorf("Failed to unquote: %w", err)
			}

			pathToAlias[unquoted] = imprt.Name.Name
			pathToUsed[unquoted] = false
		}
	}

	// Set target pkg import because imitation code uses that.
	pathToAlias[pkg.PkgPath] = pkg.Name
	pathToUsed[pkg.PkgPath] = true

	return pathToUsed, pathToAlias, nil
}

func findUnimplementedServer(pkg *packages.Package) ([]string, error) {
	targetInterfaceNames := []string{}
	for _, name := range pkg.Types.Scope().Names() {
		obj := pkg.Types.Scope().Lookup(name)
		if obj == nil {
			return nil, fmt.Errorf("Unexpected: pkg may be broken")
		}
		// TODO: Make it more robust.
		if strings.HasPrefix(obj.Name(), "Unimplemented") {
			targetInterfaceNames = append(targetInterfaceNames, strings.Replace(obj.Name(), "Unimplemented", "", 1))
		}
	}

	return targetInterfaceNames, nil
}

func createImitInterfaces(pkg *packages.Package, names []string, pathToUsed map[string]bool, pathToAlias map[string]string) ([]*ImitInterface, error) {
	imits := []*ImitInterface{}

	for _, name := range names {
		obj := pkg.Types.Scope().Lookup(name)
		if obj == nil {
			return nil, fmt.Errorf("Unexpected: pkg may be broken")
		}

		if !types.IsInterface(obj.Type()) {
			continue
		}

		// XXX: So far only for interface which suffix is 'Server'.
		if !strings.HasSuffix(obj.Name(), "Server") {
			continue
		}

		ifc := obj.Type().Underlying().(*types.Interface).Complete()
		funcSigs := createImitFuncSigs(ifc, pathToUsed, pathToAlias)

		imits = append(imits, &ImitInterface{
			Name:       fmt.Sprintf("%s.%s", obj.Pkg().Name(), obj.Name()),
			UnimplName: fmt.Sprintf("%s.Unimplemented%s", obj.Pkg().Name(), obj.Name()),
			RecvName:   fmt.Sprintf("Imit%s", obj.Name()),
			FuncSigs:   funcSigs,
		})
	}

	return imits, nil
}

func createImitFuncSigs(ifc *types.Interface, pathToUsed map[string]bool, pathToAlias map[string]string) []*ImitFuncSig {
	funcSigs := []*ImitFuncSig{}
	for j := 0; j < ifc.NumMethods(); j++ {
		method := ifc.Method(j)
		methodName := method.Name()

		sig := method.Type().(*types.Signature)
		params := createParams(sig, pathToUsed, pathToAlias, sig.Params().Len(), lookupRaram, true)
		results := createParams(sig, pathToUsed, pathToAlias, sig.Results().Len(), lookupResult, false)

		funcSigs = append(funcSigs, &ImitFuncSig{
			Name:    methodName,
			Params:  params,
			Results: results,
		})
	}

	return funcSigs
}

var lookupRaram = func(sig *types.Signature, i int) *types.Var {
	return sig.Params().At(i)
}
var lookupResult = func(sig *types.Signature, i int) *types.Var {
	return sig.Results().At(i)
}

func createParams(sig *types.Signature, pathToUsed map[string]bool, pathToAlias map[string]string, n int, lookup func(sig *types.Signature, i int) *types.Var, isParam bool) []*Param {
	params := []*Param{}

	for i := 0; i < n; i++ {
		vr := lookup(sig, i)
		paramName := vr.Name()
		if paramName == "" && isParam {
			paramName = fmt.Sprintf("P%d", i)
		}
		if !isParam {
			paramName = ""
		}

		paramType := vr.Type().String()

		// TODO: switch

		named, ok := vr.Type().(*types.Named)
		if ok {
			namedPkg := named.Obj().Pkg()
			if namedPkg != nil {
				namedPath := named.Obj().Pkg().Path()
				pathToUsed[namedPath] = true

				paramType = fmt.Sprintf("%s.%s", named.Obj().Pkg().Name(), named.Obj().Name())
				pathToAlias[namedPath] = named.Obj().Pkg().Name()
			} else {
				paramType = named.Obj().Name()
			}
		}

		ptr, ok := vr.Type().(*types.Pointer)
		if ok {
			ptrNamed, ok := ptr.Elem().(*types.Named)
			if ok {
				namedPath := ptrNamed.Obj().Pkg().Path()
				pathToUsed[namedPath] = true

				paramType = fmt.Sprintf("*%s.%s", ptrNamed.Obj().Pkg().Name(), ptrNamed.Obj().Name())
				pathToAlias[namedPath] = ptrNamed.Obj().Pkg().Name()
			}
		}

		params = append(params, &Param{
			Name: paramName,
			Type: paramType,
		})
	}

	return params
}

func createImports(pathToUsed map[string]bool, pathToAlias map[string]string) []*Imprt {
	imprts := []*Imprt{}
	for k, v := range pathToUsed {
		if !v {
			continue
		}

		imprts = append(imprts, &Imprt{
			Alias: pathToAlias[k],
			Path:  strconv.Quote(k),
		})
	}
	return imprts
}

func genImit(m *Material) ([]byte, error) {
	// TODO: switchable template
	tmpl, err := template.New("server.tmpl").Parse(string(templates.DefaultServerTemplate))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse template: %w", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, m)
	if err != nil {
		return nil, fmt.Errorf("Failed to write template: %w", err)
	}

	formatted, err := imports.Process(m.FileName, buf.Bytes(), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to format code: %w", err)
	}

	return formatted, nil
}
