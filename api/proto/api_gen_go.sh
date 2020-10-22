# API Source Code Generator for MASS Client API Developers
# Run this script ONLY on Linux-like systems (including macOS)
# See README.md for details

protoc -I ./api/proto \
    -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.0/third_party/googleapis \
    -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.0 \
    --go_out=plugins=grpc:./api/proto \
    ./api/proto/api.proto

protoc -I ./api/proto \
    -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.0/third_party/googleapis \
    -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.0 \
    --grpc-gateway_out=logtostderr=true:./api/proto \
    ./api/proto/api.proto

protoc -I ./api/proto \
    -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.0/third_party/googleapis \
    -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.0 \
    --swagger_out=logtostderr=true:./api/proto \
    ./api/proto/api.proto
