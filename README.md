# caf

A Go microservice built on Apache Kafka using the IBM Sarama client.

## Tech Stack

- **Go 1.25** — core language
- **IBM/sarama** — Kafka client library
- **Docker / Docker Compose** — containerization
- **UUID** — event/message identification

## Project Structure

```
caf/
├── cmd/        # Application entrypoints
├── config/     # Configuration (broker addresses, topics, etc.)
├── internal/   # Business logic and Kafka producers/consumers
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Getting Started

**Prerequisites:** Docker, Docker Compose, Go 1.25+

```bash
git clone https://github.com/F-YOHS/caf.git
cd caf

# Start Kafka infrastructure
docker-compose up -d

# Run the service
go run ./cmd/...
```

## Configuration

Edit `config/` or set environment variables for:
- Kafka broker address (default: `localhost:9092`)
- Topic names
- Consumer group ID
