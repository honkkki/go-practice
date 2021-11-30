BINARY="goweb"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./main.go

clean:
	@if [ -f ${BINARY} ]; then rm ${BINARY}; fi

gotool:
	go fmt ./



