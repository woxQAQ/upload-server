.PHONY: fmt
fmt: 
	go fmt ./...
	swag fmt


.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: module
module: ## Run go mod tidy->verify against go modules.
	go mod tidy
	go mod verify