// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "custom-api/pkg"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type HttpApiServer struct {
	pb.UnimplementedHttpApiServer
}

func (s *HttpApiServer) Echo(ctx context.Context, req *pb.UserMessage) (*pb.UserMessage, error) {
	//log.Fatalln("hej", cli.ExecuteGet(ctx))
	return req, nil
}

func main() {
	go func() {
		// mux (multiplexer)
		mux := runtime.NewServeMux()

		// register server
		pb.RegisterHttpApiHandlerServer(context.Background(), mux, &HttpApiServer{})

		// http server
		log.Fatalln(http.ListenAndServe(":8069", mux))
	}()

	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHttpApiServer(grpcServer, &HttpApiServer{})

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Println(err)
	}
}
