all: setup

.PHONY: githooks
githooks:
	git config --local core.hooksPath .githooks/

.PHONY: setup
setup: githooks protoc evans golangci-lint cfssl
	@echo "Installed dependencies..."

.PHONY: protoc
protoc: 
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5

.PHONY: golangci-lint
golangci-lint:
	@if ! (command -v golangci-lint &> /dev/null) ; then \
		echo "Installing golangci-lint...";\
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.61.1;\
	fi;

.PHONY: cfssl
cfssl:
	@if ! (command -v cfssl &> /dev/null) ; then \
		echo "Installing cfssl...";\
		go install github.com/cloudflare/cfssl/cmd/cfssl@latest;\
	fi;
	@if ! (command -v cfssljson &> /dev/null) ; then \
		echo "Installing cfssljson...";\
		go install github.com/cloudflare/cfssl/cmd/cfssljson@latest;\
	fi;

.PHONY: protoc-compile
protoc-compile:
	@echo "Generating gRPC code..."
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative backend/pkg/agentservice/*.proto

.PHONY: evans
evans:
	@if command -v brew &> /dev/null && ! (command -v evans &> /dev/null) ; then \
		echo "Installing evans (grpc cli tool)...";\
		brew tap ktr0731/evans;\
		brew install evans;\
	fi;

.PHONY: generate-certificates
generate-certificates:
	@echo "Generating certificates..."
	sh scripts/generate-certificates.sh

.PHONY: loc
loc:
	cloc --3 --exclude-list-file=.clocignore .

.PHONY: crq
crq:
	docker exec -it katalog-postgres-1 postgres sql --insecure