proto:
	@echo "Generating proto files"
	protoc --proto_path=internal/models --go_out=internal/models --go_opt=paths=source_relative internal/models/*.proto
