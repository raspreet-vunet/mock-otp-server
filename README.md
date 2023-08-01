# Mock OTP Server

This is a mock OTP (One-Time Password) server written in Go. It reads a JSON file for the seed data (username
and OTP pairs) and provides an HTTP endpoint to check if a given OTP for a specific user is valid.

## Seed Data

The seed data should be a JSON file containing an array of objects, where each object represents a user and
their OTP. Each object should have a `username` field and an `otp` field. The `username` field should be a
string and the `otp` field should be an integer.

Here's an example of what the seed data might look like:

```json
[
  {
    "username": "alice",
    "otp": 123456
  },
  {
    "username": "bob",
    "otp": 654321
  }
]
```

## Quick Start

### Prerequisites

- Go (1.16 or later)
- Docker

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/raspreet-vunet/mock-otp-server.git
   cd mock-otp-server
   ```

2. Build the Docker image:

   ```bash
   docker build -t mock-otp-server .
   ```

3. Run the Docker container:

   ```bash
   docker run -d -p 8085:8085 -v $(pwd)/data:/opt/mock-otp-server/data mock-otp-server
   ```

### Customization

- Change the HTTP port by setting the `HTTP_PORT` environment variable.
- Change the data directory by setting the `DATA_DIR` environment variable.

Example:

```bash
docker run -d -p 8085:8085 -v $(pwd)/data:/custom/data/directory -e DATA_DIR=/custom/data/directory -e HTTP_PORT=8085 mock-otp-server
```

### Endpoint

The server provides the `/otp` endpoint that accepts POST requests. The request body should be a JSON object
that includes `username` and `otp` fields.

Example request:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"testuser", "otp":123456}' http://localhost:8085/otp
```
