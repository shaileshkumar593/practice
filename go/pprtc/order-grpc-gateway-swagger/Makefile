PROTO_DIR=./proto
OUT_DIR=.

PROTO_FILES=$(PROTO_DIR)/order.proto

proto:
	protoc -I ${PROTO_DIR} \
		--go_out ${OUT_DIR} --go_opt paths=source_relative \
		--go-grpc_out ${OUT_DIR} --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ${OUT_DIR} --grpc-gateway_opt paths=source_relative \
		--openapiv2_out ${OUT_DIR} ${PROTO_FILES}

clean:
	rm -f proto/*.pb.go proto/*.pb.gw.go proto/*.swagger.json
