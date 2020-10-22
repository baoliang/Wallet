# API Source Code Generator for MASS Client API Developers
# Run this script ONLY on Linux-like systems (including macOS)
# See README.md for details

protoc -I ./api/proto \
    -I $GOPATH/pkg/mode/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I $GOPATH/pkg/mode/github.com/grpc-ecosystem/grpc-gateway \
    --go_out=plugins=grpc:. \
    ./api/proto/api.proto

protoc -I ./api/proto \
    -I $GOPATH/pkg/mode/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I $GOPATH/pkg/mode/github.com/grpc-ecosystem/grpc-gateway \
    --grpc-gateway_out=logtostderr=true:. \
    ./api/proto/api.proto

protoc -I $GOPATH/src/massnet.org/mass-wallet/api/proto \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    --swagger_out=logtostderr=true:. \
    $GOPATH/src/massnet.org/mass-wallet/api/proto/api.proto
