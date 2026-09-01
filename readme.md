# GoRide - Microservics Ride-Hailing Platform

![Java](https://img.shields.io/badge/Java-21-orange.svg)
![Spring Boot](https://img.shields.io/badge/Spring_Boot-3.x-brightgreen.svg)
![Gin](https://img.shields.io/badge/Gin-v1.12.0-blue.svg)
![Golang](https://img.shields.io/badge/Go-1.27-blue.svg)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED.svg)
![CI/CD](https://img.shields.io/badge/CI%2FCD-GitHub_Actions-2088FF.svg)

## About the Project

GoRide is a highly scalable, event-driven backend architecture for a ride-hailing system. It is built using the **Microservices Architecture**, separating core domains into independent services to ensure high availability, fault tolerance, and independent deployment.

## System Architecture

- **API Gateway (Nginx):** Reverses proxy and routes external HTTP requests to internal microservices via Port 80.
- **Trip Service (Java/Spring Boot):** Manages ride requests, pricing, matching logic, and trip states. Backed by **PostgreSQL** for ACID compliance.
- **Location Service (Golang):** High-throughput service handling real-time driver coordinates and geo-hashing. Backed by **Redis** for ultra-fast spatial queries.
- **Message Broker (RabbitMQ):** Handles asynchronous events (e.g., `trip_created_events`) to ensure loose coupling between services.
- **gRPC:** Used for low-latency, internal synchronous communication between `location-service` and `trip-service`.

## Tech Stack

| Category                | Technologies                                       |
| :---------------------- | :------------------------------------------------- |
| **API Gateway**         | Nginx                                              |
| **Backend Services**    | Java 21 (Spring Boot), Golang 1.27 (Gin Framework) |
| **Databases**           | PostgreSQL 15, Redis 7.2                           |
| **Message Broker**      | RabbitMQ 4.0                                       |
| **RPC & Serialization** | gRPC, Protocol Buffers (Protobuf)                  |
| **DevOps & CI/CD**      | Docker, Docker Compose, GitHub Actions             |

## Project Structure (Monorepo)

```text
goride-backend
 ┣ api-gateway       # Nginx configurations & templates
 ┣ location-service  # Golang service for real-time tracking
 ┣ trip-service      # Java Spring Boot service for trip management
 ┣ proto             # Shared Protocol Buffers definitions
 ┣ .github/workflows # CI/CD pipelines for Java & Go
 ┣ docker-compose.yml # Root orchestration file
 ┗ .env               # Environment variables
```

# Getting Started

## Prerequisites

- Docker & Docker Compose V2
- Git

## Installation & Run

1. Clone the repository:

```bash
git clone [https://github.com/your-username/goride.git](https://github.com/your-username/goride.git)
cd goride
```

2. Configure Environment Variables:
   Update the .env file in the root directory with your preferred database credentials and service ports.

3. Start the Infrastructure & Services:
   We use Docker Compose include to modularize our infrastructure. Run the following command from the root directory:

```bash
docker-compose up -d --build
```

4. Verify the deployment:
   The API Gateway should now be running on http://localhost:80. All internal ports (8080, 8081, 9090) are completely isolated and secured.

## API reference (Via API Gateway)

All external client requests must pass through the Nginx Gateway on port 80.

- Trip Service: http://localhost/api/v1/trips/...
- Location Service: http://localhost/api/v1/locations/...

## CI/CD pipelines

This repository implements Automated Testing and Continuous Integration via GitHub Actions.

- Go CI: Triggers automatically on changes inside location-service/. Runs unit tests and verifies Docker build.
- Java CI: Triggers automatically on changes inside trip-service/ or proto/. Caches Maven dependencies, executes Unit Tests, and verifies Docker build.

If any test fails, the deployment pipeline is halted to prevent bugs from reaching production.
