.PHONY: fmt
fmt: 
	go fmt ./...
	cd proto && buf format -w

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: lint-buf
lint-buf:
	cd proto && buf lint

.PHONY: module
module: ## Run go mod tidy->verify against go modules.
	go mod tidy
	go mod verify

.PHONY: generate
generate: 
	cd proto && buf format -w && buf generate 
