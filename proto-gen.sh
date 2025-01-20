set -x
find ./proto -name "*.proto" | xargs -I {} protofmt -w {}
protoc --proto_path=/usr/include --go_out=/tmp --go-grpc_out=/tmp -I . $(find ./proto -name "*.proto")
protoc-go-inject-tag -input="/tmp/pb/*.pb.go"

rm -rf ./pb
mv /tmp/pb ./pb
chmod -R 777 ./pb
