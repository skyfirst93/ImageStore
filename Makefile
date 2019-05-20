#SUBDIRS := $(wildcard */.)
SUBDIRS := cmd pkg
IMAGE_NAME := imagestore
IMAGE_LABEL := latest


# Static code analyzers
GOLINT := golint
GOFMT := gofmt
GOIMPORTS := goimports
STATIC_TOOLS := $(GOLINT) $(GOFMT) $(GOIMPORTS)
GO_FILES := `find . -name *.go`

exe:
	$(MAKE) -C cmd exe
image:	exe
	@echo Building Docker $(IMAGE_NAME)....
	sudo docker build -t $(IMAGE_NAME):$(IMAGE_LABEL) -f Dockerfile.$(IMAGE_NAME) .
run:
	go run cmd/app.go
docker-run:
	docker run -it --env KAFKA_SERVICE=127.0.0.1 --env STORAGE_PATH=/app -p 8081:8081 --name c2 imagestore:latest
mtest: 
	sh test.sh 

unit-test:
	go test -cover ./...

static: $(STATIC_TOOLS)

$(GOLINT):
	@echo "Running golint on all the files present"
	@$(GOLINT) ./... |tee golint.log

$(GOFMT):
	@echo "Running gofmt on all the files present"
	@$(GOFMT) -s -w . |tee gofmt.log;

$(GOIMPORTS):
	@echo "Running goimports on all the files present"
	@$(GOIMPORTS) -w $(GO_FILES) |tee goimports.log;

.PHONY: $(SUBDIRS) static $(STATIC_TOOLS)
