package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-grpc-imit-gen",
	Long:  "Print the version number of go-grpc-imit-gen",
	Run: func(cmd *cobra.Command, args []string) {
		if info, ok := debug.ReadBuildInfo(); ok {
			fmt.Printf("version %s\n", info.Main.Version)
		} else {
			fmt.Println("version unknown")
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
