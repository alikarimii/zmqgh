SERILIZATION_DIR := zero/infrastructure/serialization/proto
PROTO_GRPC_DIR := zero/infrastructure/adapter/grpc/proto
PROTO_OUT := pb
SERILIZATION_OUT := pkg/shared
clean:
	@[ -d $(PROTO_OUT) ] && rm -r $(PROTO_OUT) || true
	mkdir pb
ser:
	protoc -I broker/$(SERILIZATION_DIR) broker.proto --go_out=paths=source_relative:$(SERILIZATION_OUT)
genb:
	protoc -I broker/$(PROTO_GRPC_DIR) broker.proto --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:$(PROTO_OUT) --go_out=paths=source_relative:$(PROTO_OUT)
genp:
	protoc -I publisher/$(PROTO_GRPC_DIR) publisher.proto --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:$(PROTO_OUT) --go_out=paths=source_relative:$(PROTO_OUT)
gens:
	protoc -I subscriber/$(PROTO_GRPC_DIR) subscriber.proto --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:$(PROTO_OUT) --go_out=paths=source_relative:$(PROTO_OUT)
