# 2PCByz: Byzantine Fault-Tolerant Two-Phase Commit

2PCByz is an implementation of the Two-Phase Commit (2PC) protocol with Byzantine fault tolerance. This project explores distributed transaction commitment under adversarial conditions, ensuring correctness despite Byzantine failures.

## Features

- Implements the standard Two-Phase Commit (2PC) protocol.
- Enhances resilience by incorporating Byzantine fault tolerance.
- Simulates different failure scenarios to evaluate robustness.
- Developed in Go for efficient distributed systems handling.

## Prerequisites

Ensure you have the following installed:
- Go 1.23+
- MongoDB

## Installation

Clone the repository:

```sh
git clone https://github.com/amirhnajafiz/2pcbyz.git
cd 2pcbyz
```

## Usage

Run the coordinator:

```sh
go run cmd/coordinator/main.go
```

Run a participant:

```sh
go run cmd/participant/main.go
```

Modify configurations in `config.json` for custom setups.

## Configuration

The system's behavior can be adjusted using the `config.json` file:

```json
{
  "coordinator": "localhost:5000",
  "participants": [
    "localhost:5001",
    "localhost:5002",
    "localhost:5003"
  ],
  "byzantine_behavior": "none" // Options: none, crash, malicious
}
```

## Testing

To run tests:
```sh
go test ./...
```
