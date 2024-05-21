# Twitch Stats Producer

A dockerized Go server that listens to Twitch streams via IRC and pushes messages into a Kafka cluster.

# Cloning and Testing

To clone the repository:

```bash
git clone https://github.com/michaelfioretti/twitch-stats-producer.git
cd twitch-stats-producer
go test -v ./...
```
