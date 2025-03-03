init:
	make clean
	make proto
	make network

dev:
	docker compose -f docker-compose-dev.yml up --build

test:
	go test -v ./...

build:
	go build -o dist/main cmd/twitchstatsproducer/main.go

coverage:
	go test ./... -coverprofile=coverage.out

view-coverage:
	go tool cover -html=coverage.out

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


# Helpers
network:
	docker network create twitch-chat-stats
	docker network connect twitch-chat-stats

secrets:
secrets:
	if [ ! -d secrets ]; then \
		mkdir secrets; \
	fi
	echo "twitch_client_id" > secrets/twitch_client_id.txt
	echo "twitch_client_secret" > secrets/twitch_client_secret.txt
	echo "twitch_oauth_token" > secrets/twitch_oauth_token.txt

make k8-secrets:
	kubeseal --controller-name sealed-secrets --controller-namespace kube-system --format=yaml < twitch-chat-stats-secrets.yml > kubernetes/sealed-twitch-chat-stats-secrets.yml
	kubeseal --controller-name sealed-secrets --controller-namespace kube-system --format=yaml < kafka-jaas-config.yml > kubernetes/sealed-kafka-jaas-config.yml
