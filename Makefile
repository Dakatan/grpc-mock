ECHO_IMAGE_NAME	= grpc-echo:latest
SESSION_IMAGE_NAME = grpc-session:latest
APP = app

build-echo:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(APP) cmd/echo/echo_server.go

build-echo-docker:
	docker build -t $(ECHO_IMAGE_NAME) . --build-arg build_target=build-echo

build-session:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(APP) cmd/session/session_server.go

build-session-docker:
	docker build -t $(SESSION_IMAGE_NAME) . --build-arg build_target=build-session

protoc:
	protoc --go_out=plugins=grpc:./pkg/proto ./pkg/proto/*.proto

clean:
	rm app
