# kvne

kvne is a simple key-value store written in Go. It is compatible with the `redis-cli` client.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.20 or higher

### Installing

Clone the repository:

```sh
git clone https://github.com/knvi/kvne.git
cd kvne
```

Build the project:

```sh
go build ./cmd/main.go
```

Run the project:

```sh
./main
```

## Usage

### list of commands
- `set key value` - set the string value of a key
- `get key` - get the value of a key
- `del key [key ...]` - delete one or more keys
- `ping [message]` - ping the server
- `ttl key` - get the time to live for a key
- `expire key seconds` - set a key's time to live in seconds
- `bgrewriteaof` - asynchronously rewrite the append-only file