//go:generate go-grpc-imit-gen --repo=$GOFILE --dstDir=../imit/

package gen

import (
	_ "google.golang.org/genproto/googleapis/firestore/v1"
	_ "google.golang.org/genproto/googleapis/pubsub/v1"
	_ "google.golang.org/genproto/googleapis/spanner/v1"
)
