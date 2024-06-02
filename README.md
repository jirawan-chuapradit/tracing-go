# tracing-go


## Introduction

This project demonstrates how to implement distributed tracing using OpenTelemetry in a Go application.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Step 1: Clone the Repository

```sh
git clone https://github.com/jirawan-chuapradit/tracing-go.git
cd tracing-go
```
### Step 2: Run Docker Compose
```sh
docker-compose up
```

This command will start all the necessary services defined in the docker-compose.yml file.

## Project Structure
- main.go: Entry point of the application
- handlers/: Contains HTTP handlers
- otel/: OpenTelemetry configuration
- docker-compose.yml: Docker Compose configuration
- Dockerfile: Docker configuration
## Usage
After running docker-compose up, the application will be available at http://localhost:8080.
