APP = k8s-set-configuration
all: $(APP)

k8s-set-configuration:
	GOARCH=arm64 GOOS=darwin go build -o ../../tools/$(APP).arm64.Darwin
	GOARCH=amd64 GOOS=darwin go build -o ../../tools/$(APP).amd64.Darwin
	GOARCH=amd64 GOOS=linux go build -o ../../tools/$(APP).amd64.Linux