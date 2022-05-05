customer-rest:
	cd customers && make rest

customer-grpc:
	cd customers && make grpc

product-rest:
	cd products && make rest

product-grpc:
	cd products && make grpc

order:
	cd orders && make dev

generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/customer.proto && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/product.proto
	cd ./proto && \
	cp *.go ../customers/proto && \
	cp *.go ../orders/proto && \
	cp *.go ../products/proto && \
	rm *.go