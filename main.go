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
	"net/http"

	pb "custom-api/pkg"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type HttpApiServer struct {
	pb.UnimplementedHttpApiServer
}

// Return the input message
func (s *HttpApiServer) Echo(ctx context.Context, req *pb.UserMessage) (*pb.UserMessage, error) {
	return req, nil
}

// Run a http server that the user can communicate with
func main() {
	// Make a HTTP request multiplexer that matches the URL of incoming req.
	// (Use the NewSurveMux() function to create an empty servemux)
	mux := runtime.NewServeMux()
	// register the server
	pb.RegisterHttpApiHandlerServer(context.Background(), mux, &HttpApiServer{})

	//Create a new server and start listening for incoming requests
	// with the http.ListenAndServe() function, passing in servemux for it to
	// match reqs against as the second parameter.
	// ListenAndServe = Listen on TCP network address and handle requests on incoming connections
	log.Fatalln(http.ListenAndServe(":8080", mux))
}
