package main

// This file used to build files from protobuf. such like pb.go, pb.gw.go...

//go:generate go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
//go:generate go install github.com/envoyproxy/protoc-gen-validate@latest
