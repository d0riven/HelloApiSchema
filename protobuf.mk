DOCKER := docker
PROTOC := protoc

.PHONY: FORCE
FORCE:

.PHONY: generate generate/server generate/client
generate: generate/server generate/client
generate/server:
	$(PROTOC) -I api/proto \
		-I $$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:server/golang/proto api/proto/user.proto \
		--grpc-gateway_out=server/golang/proto api/proto/user.proto
generate/client:
	$(PROTOC) -I api/proto \
		-I $$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:client/golang/proto api/proto/user.proto
