cd product && protoc --go_out=../services product.proto && protoc --go-grpc_out=../services product.proto
protoc --go_out=../services models.proto