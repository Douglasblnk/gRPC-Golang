customer:
	cd customers && make dev

product:
	cd products && make dev

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

# protoc -I=. --go_out=./ proto/customer.proto