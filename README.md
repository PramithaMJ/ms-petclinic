# Microservice Pet Clinic

Microservice Pet Clinic is a fully distributed, cloud-native application designed to demonstrate the power of microservices architecture using multiple technologies like Java, Golang, Node.js, and Python. It follows best practices for service discovery, API management, monitoring, tracing, authentication, and CI/CD.

## 📌 Features
- **Microservices Architecture**: Independently deployed services
- **Service Discovery**: Eureka / Consul for dynamic service registration
- **API Gateway**: Kong / Traefik for centralized request handling
- **Configuration Management**: Spring Cloud Config / HashiCorp Vault
- **Asynchronous Messaging**: Kafka / RabbitMQ for event-driven architecture
- **Authentication & Authorization**: Keycloak with OAuth2 & JWT
- **Monitoring & Tracing**: Prometheus, Grafana, Zipkin / Jaeger
- **Centralized Logging**: ELK Stack (Elasticsearch, Logstash, Kibana)
- **CI/CD Pipeline**: GitHub Actions, Jenkins, Kubernetes, Helm
- **Containerization & Orchestration**: Docker, Kubernetes, Helm, ArgoCD

## 🏗 Architecture
```plaintext
                      +------------------------------------------------+
                      |               API Gateway (Kong)              |
                      |   - Handles Routing, Authentication (JWT)     |
                      |   - Rate Limiting, Load Balancing             |
                      +--------------------+---------------------------+
                                           |
      +---------------------+--------------+-----------------------+
      |                     |                                      |
      v                     v                                      v
+------------+     +----------------+                    +------------------+
| Service Reg|<--> | Config Server  |                    | Monitoring & Logs|
| Discovery  |     | (Spring Config)|                    | Prometheus + ELK |
| Eureka/Consul|   | HashiCorp Vault|                    | Zipkin / Jaeger  |
+------------+     +----------------+                    +------------------+
      |                     |                                      |
      |                     v                                      |
      |             +------------------+                           |
      |             |  Auth Service     |   <----->   Keycloak     |
      |             |  (Java + Spring)  |   (OAuth2 + JWT)         |
      |             +------------------+                           |
      |                     |                                      |
+----------------+     +----------------+    +------------------+   +----------------+
|  Pet Service  |<--> |  Owner Service  |<-> |  Visit Service   |<->| Billing Service |
|  Golang + Gin |     | Python + FastAPI|    | Java + Spring    |   | Golang + Fiber  |
+----------------+     +----------------+    +------------------+   +----------------+
      |                     |                                      |
      |                     v                                      |
      |             +-------------------+                          |
      |             | Event Bus (Kafka) |                          |
      |             |  Async Messaging  |                          |
      |             +-------------------+                          |
      |                     |                                      |
      |                     v                                      |
      |            +----------------------+                        |
      |            | Notification Service |                        |
      |            |  (Python + Celery)   |                        |
      |            +----------------------+                        |
      |                                                              |
      +--------------------------------------------------------------+
                                      |
        +--------------------------------------------------+
        |  CI/CD: GitHub Actions + Jenkins + Kubernetes    |
        |  Docker, Helm, ArgoCD for Automated Deployment  |
        +--------------------------------------------------+
```

## 📂 Repository Structure
```plaintext
microservice-pet-clinic/
│── api-gateway/       → Kong / Traefik
│── service-registry/  → Eureka / Consul
│── config-server/     → Spring Cloud Config
│── services/
│   ├── auth-service/      → Java (Spring Boot)
│   ├── pet-service/       → Golang (Gin)
│   ├── owner-service/     → Python (FastAPI)
│   ├── vet-service/       → Node.js (Express)
│   ├── visit-service/     → Java (Spring Boot)
│   ├── billing-service/   → Golang (Fiber)
│   ├── notification-service/ → Python (Celery)
│── event-streaming/ → Kafka / RabbitMQ
│── monitoring/      → Prometheus + Grafana
│── tracing/         → Zipkin / Jaeger
│── logging/         → ELK Stack (Elasticsearch, Logstash, Kibana)
│── k8s/            → Kubernetes manifests
│── ci-cd/          → GitHub Actions / Jenkins pipelines
│── README.md
```

## 🚀 Getting Started

### **1️⃣ Prerequisites**
- Docker & Docker Compose
- Kubernetes & Helm
- Java 17, Golang, Node.js, Python
- PostgreSQL / MongoDB
- Kafka / RabbitMQ

### **2️⃣ Running Services Locally**
#### **With Docker Compose**
```sh
docker-compose up -d
```

#### **With Kubernetes**
```sh
kubectl apply -f k8s/
```

### **3️⃣ API Endpoints**
| **Service** | **Endpoint** | **Method** |
|------------|-------------|------------|
| **Auth Service** | `/auth/login` | POST |
| **Pet Service** | `/pets` | GET, POST |
| **Owner Service** | `/owners` | GET, POST |
| **Vet Service** | `/vets` | GET, POST |
| **Visit Service** | `/visits` | GET, POST |
| **Billing Service** | `/billing` | POST |

## 🛠 CI/CD Pipeline
1. **GitHub Actions** - Code build & testing
2. **Jenkins** - Continuous Integration
3. **ArgoCD** - Kubernetes GitOps deployment
4. **Helm** - Kubernetes package manager

## 🔍 Monitoring & Logging
| **Component** | **Purpose** |
|------------|------------|
| **Zipkin / Jaeger** | Distributed Tracing |
| **Prometheus + Grafana** | Metrics & Alerts |
| **ELK Stack (Elasticsearch, Logstash, Kibana)** | Centralized Logging |

## 🤝 Contributing
Feel free to contribute to this project! Open an issue or create a pull request.

## 📜 License
MIT License - Feel free to use and modify this project.
