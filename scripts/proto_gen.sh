#!/bin/bash

protoc -I=proto/ --go_out=cluster/pkg/ proto/pbft.proto
protoc -I=proto/ --go-grpc_out=cluster/pkg/ proto/pbft.proto
protoc -I=proto/ --go_out=cluster/pkg/ proto/database.proto
protoc -I=proto/ --go-grpc_out=cluster/pkg/ proto/database.proto
protoc -I=proto/ --go_out=cluster/pkg/ proto/api.proto
protoc -I=proto/ --go-grpc_out=cluster/pkg/ proto/api.proto
protoc -I=proto/ --go_out=client/pkg/ proto/database.proto
protoc -I=proto/ --go-grpc_out=client/pkg/ proto/database.proto
protoc -I=proto/ --go_out=client/pkg/ proto/api.proto
protoc -I=proto/ --go-grpc_out=client/pkg/ proto/api.proto
