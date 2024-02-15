run:
	cd cmd/api && go run main.go

proto admin:
	cd pkg/pb && protoc --go_out=. --go-grpc_out=. admin_airline.proto

proto2 booking:
	cd pkg/pb && protoc --go_out=. --go-grpc_out=. booking.proto