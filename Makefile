init:
	make clean
	make proto

docker-build:
	docker build -t twichchatstatsproducer .

test:
	go test -v ./...

build:
	go build -o dist/main cmd/twitchstatsproducer/main.go

coverage:
	go test -coverprofile=coverage.txt

proto:
	make clean
	@echo "Generating proto files"
	mkdir -p internal/models/proto
	protoc --proto_path=internal/models --go_out=internal/models/proto --go_opt=paths=source_relative internal/models/*.proto

clean:
	@echo "Cleaning up proto files"
	if [ -d internal/models/proto ]; then \
		rm -rf internal/models/proto; \
	fi
