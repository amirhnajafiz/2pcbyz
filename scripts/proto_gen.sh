#!/bin/bash

protoc -I=proto/ --go_out=cluster/ proto/pbft.proto
protoc -I=proto/ --go-grpc_out=cluster/ proto/pbft.proto
protoc -I=proto/ --go_out=cluster/ proto/database.proto
protoc -I=proto/ --go-grpc_out=cluster/ proto/database.proto
protoc -I=proto/ --go_out=cluster/ proto/api.proto
protoc -I=proto/ --go-grpc_out=cluster/ proto/api.proto
protoc -I=proto/ --go_out=client/ proto/database.proto
protoc -I=proto/ --go-grpc_out=client/ proto/database.proto
protoc -I=proto/ --go_out=client/ proto/api.proto
protoc -I=proto/ --go-grpc_out=client/ proto/api.proto
