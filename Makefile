dev:
	go run cmd/twitchstatsproducer/main.go

proto:
	make clean
	echo "Generating proto files"
	mkdir -p internal/models/proto
	protoc --proto_path=internal/models --go_out=internal/models/proto --go_opt=paths=source_relative internal/models/*.proto

clean:
	echo "Cleaning up proto files"
	if [ -d internal/models/proto ]; then \
		rm -rf internal/models/proto; \
	fi
