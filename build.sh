# Build the protocol buffers
protoc --go_out=network/ --go-grpc_out=network/ proto/*.proto --plugin=grpc:
# Tidy up our go mod
go mod tidy
# Run go build
go build -x