APP = k8s-set-configuration
all: clean k8s-set-configuration
clean:
	rm -rfv ../../tools/$(APP).*
	rm -f $(APP)
k8s-set-configuration:
	echo $(APP)
	GOARCH=arm64 GOOS=darwin go build -o ../../tools/$(APP).arm64.darwin
	GOARCH=amd64 GOOS=darwin go build -o ../../tools/$(APP).amd64.darwin
	GOARCH=amd64 GOOS=linux go build -o ../../tools/$(APP).amd64.linux

