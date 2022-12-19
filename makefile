run_protoc:
	protoc -I . \
		--go_out ./ --go_opt paths=source_relative \
		--go-grpc_out ./ --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./gen/go \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		./*.proto

buf_generate:
	buf generate