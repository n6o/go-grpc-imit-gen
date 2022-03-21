package example

import (
	"context"
	"net"
	"reflect"
	"testing"

	"github.com/n6o/go-grpc-imit-gen/example/imit/imit_spanner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"

	spanner "google.golang.org/genproto/googleapis/spanner/v1"
)

func runImitation(t *testing.T) (*imit_spanner.ImitSpannerServer, spanner.SpannerClient, *grpc.Server) {
	s := grpc.NewServer()
	l := bufconn.Listen(20)

	imit := &imit_spanner.ImitSpannerServer{}
	spanner.RegisterSpannerServer(s, imit)

	go s.Serve(l)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(c context.Context, s string) (net.Conn, error) {
		return l.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Error grpc dial with bufconn: %v", err)
	}

	return imit, spanner.NewSpannerClient(conn), s
}

func Test_Example_SpannerServer_ExecuteSql(t *testing.T) {
	imit, client, server := runImitation(t)
	defer server.GracefulStop()

	type scenario struct {
		req interface{}
		res interface{}
	}
	type args struct {
		scenarios []*scenario
	}

	tests := []struct {
		name     string
		setupper func(imit *imit_spanner.ImitSpannerServer)
		args     args
		wantErr  bool
	}{
		{
			name: "example",
			setupper: func(imit *imit_spanner.ImitSpannerServer) {
				imit.EnqueueCreateSessionMock(func(P0 context.Context, P1 *spanner.CreateSessionRequest) (*spanner.Session, error) {
					return &spanner.Session{
						Name: "Example",
						Labels: map[string]string{
							"key1": "value1",
							"key2": "value2",
						},
						CreateTime: &timestamppb.Timestamp{
							Seconds: 10,
							Nanos:   20,
						},
						ApproximateLastUseTime: &timestamppb.Timestamp{
							Seconds: 100,
							Nanos:   200,
						},
					}, nil
				})
			},
			args: args{
				scenarios: []*scenario{
					{
						req: &spanner.CreateSessionRequest{
							Database: "ExampleDatabase",
						},
						res: &spanner.Session{
							Name: "Example",
							Labels: map[string]string{
								"key1": "value1",
								"key2": "value2",
							},
							CreateTime: &timestamppb.Timestamp{
								Seconds: 10,
								Nanos:   20,
							},
							ApproximateLastUseTime: &timestamppb.Timestamp{
								Seconds: 100,
								Nanos:   200,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupper(imit)
			for _, scenario := range tt.args.scenarios {
				got, err := client.CreateSession(context.Background(), scenario.req.(*spanner.CreateSessionRequest))
				if (err != nil) != tt.wantErr {
					t.Errorf("ImitSpannerServer.ExecuteSql() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				res := scenario.res.(*spanner.Session)
				if got.Name != res.Name {
					t.Errorf("got = %v, want %v", got.Name, res.Name)
				}
				if !reflect.DeepEqual(got.Labels, res.Labels) {
					t.Errorf("got = %v, want %v", got.Labels, res.Labels)
				}
				if !reflect.DeepEqual(got.CreateTime, res.CreateTime) {
					t.Errorf("got = %v, want %v", got.CreateTime, res.CreateTime)
				}
				if !reflect.DeepEqual(got.ApproximateLastUseTime, res.ApproximateLastUseTime) {
					t.Errorf("got = %v, want %v", got.ApproximateLastUseTime, res.ApproximateLastUseTime)
				}

				records := imit.TakeCreateSessionRecords()
				if len(records) != 1 {
					t.Errorf("got = %v, want %v", len(records), len(tt.args.scenarios))
				}

				req := scenario.req.(*spanner.CreateSessionRequest)
				record := records[0]
				if record.P1.Database != req.Database {
					t.Errorf("recorded = %v, want %v", record.P1.Database, req.Database)
				}
			}
		})
	}
}
