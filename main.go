//go:generate protoc --proto_path=model --go_out=model model/webpage.proto
package main

import (
	"blahProtobuf/model"
	"github.com/ligato/cn-infra/rpc/grpc"
)

func main()  {
p := &webPage.ListofLists{}
a := &webPage.List{}
	a.NumIterms = 1
	a.Title = "John"

	config := &grpc.Config{
		Endpoint:             "0.0.0.0:18202",
	}
c := grpc.Server()
c.Send(a)





}