TEST?=$$(go list ./... | grep -v 'vendor')
BINARY=terraform-provider-policyguru

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

default: install

build:
	go build -o ${BINARY}

install: build
	mv ${BINARY} ~/.terraform.d/plugins

fmt:
	gofmt -l -w $(SRC)

lint:
	golangci-lint run

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
	xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

clean:
	rm -rf examples/.terraform && rm -rf examples/terraform.tfstate

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

validate-modules:
	@echo "- Verifying that the dependencies have expected content..."
	go mod verify
	@echo "- Checking for any unused/missing packages in go.mod..."
	go mod tidy
	@echo "- Checking for unused packages in vendor..."
	go mod vendor
	@git diff --exit-code -- go.sum go.mod vendor/

terraform-demo: install
	cd examples && terraform init && terraform apply