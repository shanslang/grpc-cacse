cd product && protoc --go_out=../services product.proto && protoc --go-grpc_out=../services product.proto
protoc --go_out=../services models.proto
protoc --go_out=../services orders.proto
protoc --go-grpc_out=../services orders.proto
protoc --grpc-gateway_out=logtostderr=true:../services product.proto
protoc --grpc-gateway_out=logtostderr=true:../services orders.proto