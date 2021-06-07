#!/bin/bash

protoc --go_out=./golang --go_opt=paths=source_relative \
--go-grpc_out=golang --go-grpc_opt=paths=source_relative \
$(find . -name "*.proto")

protoc --java_out=./java/src/main/java \
--plugin=protoc-gen-grpc-java=/usr/local/bin/protoc-gen-grpc-java \
--grpc-java_out=java/src/main/java \
$(find . -name "*.proto")