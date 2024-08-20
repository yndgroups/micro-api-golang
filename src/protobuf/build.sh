protoc --go_out=. --go-grpc_out=.  ./protos/**/*.proto && protoc-go-inject-tag -input="./build/**/*.pb.go"
rm -rf ./build/global && mv  ./protobuf/build/global ./build/global && rm -rf ./protobuf