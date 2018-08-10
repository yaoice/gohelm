package main

import (
	"time"
	"context"
	"google.golang.org/grpc"
	tpb "github.com/yaoice/gohelm/hapi/services"
	"fmt"
	"log"
	"google.golang.org/grpc/metadata"
)

const (
	TILLER_HOST = "10.125.233.67"
	TILLER_PORT = 44134
	TILLER_VERSION = "2.9.1"
	TILLER_TIMEOUT = 50 * time.Second
)

type Tiller struct {
	host string
	port int32
	Channel *grpc.ClientConn
	Timeout time.Duration
}

func NewTiller() (*Tiller) {
	address := fmt.Sprintf("%s:%d", TILLER_HOST, TILLER_PORT)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return &Tiller{
		host: TILLER_HOST,
		port: TILLER_PORT,
		Channel: conn,
		Timeout: TILLER_TIMEOUT,
	}
}

func (tiller *Tiller) String() string {
	return fmt.Sprintf("x-helm-api-client: %s", TILLER_VERSION)
}

func (tiller *Tiller) ListReleases() {

	// 调用前，需要传带x-helm-api-client的context，不然会报client server版本不兼容
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("x-helm-api-client", TILLER_VERSION),
	)

	conn := tiller.Channel
	defer conn.Close()
	c := tpb.NewReleaseServiceClient(conn)
	req := &tpb.ListReleasesRequest{}
	releaseList, err := c.ListReleases(ctx, req)
	if err != nil {
		log.Fatalf("could not releaseList: %v", err)
	}
	resp, err := releaseList.Recv()
	if err != nil {
		log.Fatalf("xx xx %v", err)
	}
	for i, release := range resp.Releases {
		fmt.Println(i, release.Name)
	}
}

func main() {
	NewTiller().ListReleases()
}
