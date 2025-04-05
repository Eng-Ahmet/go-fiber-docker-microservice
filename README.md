# Go Microservices Architecture

This project demonstrates a microservices architecture built using Go with the Fiber framework and Docker. It consists of two microservices and an API Gateway, all orchestrated using Docker Compose.

## Project Structure

```
go-microservices
├── service1          # First microservice
│   ├── main.go      # Entry point for service1
│   ├── go.mod       # Module definition for service1
│   ├── go.sum       # Dependency checksums for service1
│   └── Dockerfile    # Dockerfile for building service1
├── service2          # Second microservice
│   ├── main.go      # Entry point for service2
│   ├── go.mod       # Module definition for service2
│   ├── go.sum       # Dependency checksums for service2
│   └── Dockerfile    # Dockerfile for building service2
├── api-gateway       # API Gateway service
│   ├── main.go      # Entry point for the API Gateway
│   ├── go.mod       # Module definition for the API Gateway
│   ├── go.sum       # Dependency checksums for the API Gateway
│   └── Dockerfile    # Dockerfile for building the API Gateway
├── docker-compose.yml # Docker Compose configuration
└── README.md         # Project documentation
```

## Getting Started

### Prerequisites

- Go (version 1.20 or later)
- Docker and Docker Compose

### Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   cd go-microservices
   ```

2. Build the Docker images:

   ```
   docker-compose build
   ```

### Running the Application

To start the microservices and the API Gateway, run:

```
docker-compose up
```

The services will be accessible at the following endpoints:

- Service 1: `http://localhost:3001`
- Service 2: `http://localhost:3002`
- API Gateway: `http://localhost:8080`

### Scaling Services

You can scale the microservices by modifying the `docker-compose.yml` file to specify the number of replicas for each service.

### Logging and Monitoring

Consider integrating logging and monitoring solutions such as Prometheus and Grafana for better observability of your microservices.

## Conclusion

This project provides a basic skeleton for building a microservices architecture using Go with Fiber and Docker. You can extend it by adding more services, implementing service discovery, and integrating logging and monitoring solutions as needed.