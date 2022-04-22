customer:
	cd customers && make dev

product:
	cd products && make dev

order:
	cd order && make dev

generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/customer.proto

# protoc -I=. --go_out=./ proto/customer.proto