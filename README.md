# Twitch Stats Producer

A Dockerized Go server that produces Twitch chat statistics to a Kafka cluser via IRC.

# Cloning and Running

To clone the repository:

```bash
git clone https://github.com/michaelfioretti/twitch-stats-producer.git
cd twitch-stats-producer
make init && make dev
```

# Configuration
You will need to generate a couple of secrets that are used between the two containers to
authenticate with Twitch, as well as connect to Kafka.

```bash
make secrets
```

This will create a `secrets` directory with the following files:

- `twitch_client_id`: Your Twitch client ID
- `twitch_client_secret`: Your Twitch client secret
- `kafka_brokers`: A comma separated list of Kafka brokers

# Note
This project is still a work in progress, and was a great way for
me to get my hands dirty using Go and Kafka. I'm sure there are optimizations that could have been made in some of the code, but I'm happy with the way it turned out!
