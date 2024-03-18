package main

import (
	"github.com/ZKProofAggregator-RPC/aggregator"
	aggregatorServer "github.com/ZKProofAggregator-RPC/aggregator/interface"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	GitVersion = ""
	GitCommit  = ""
	GitDate    = ""
)

func main() {
	//app := NewCli(GitVersion, GitCommit, GitDate)
	//if err := app.Run(os.Args); err != nil {
	//	log.Crit("Application failed", "message", err)
	//}
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}
	grpcserver := grpc.NewServer()
	aggregatorServer.RegisterAggregatorServiceServer(grpcserver, new(aggregator.Aggregator))
	err = grpcserver.Serve(listener)
	if err != nil {
		log.Panic(err)
	}
}
