# Kubernetes Auto-Scaling CI/CD Pipeline

> Production-grade CI/CD infrastructure with Kubernetes HPA, Docker multi-stage builds, Terraform EKS provisioning, and Prometheus monitoring.

## 🚀 Overview

A complete infrastructure-as-code project demonstrating how to build, test, containerize, deploy, and monitor a Go application on Kubernetes with automated horizontal scaling. Includes GitHub Actions pipelines, Terraform for AWS EKS provisioning, and Prometheus/Grafana monitoring.

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🚀 CI/CD Pipeline | GitHub Actions: test → build → push → deploy |
| 🐳 Multi-Stage Docker | Optimized Go binary (< 20MB final image) |
| ☸️ Kubernetes Deployment | Rolling updates with zero downtime |
| 📈 HPA Auto-Scaling | Scale 2-10 pods based on CPU/memory (70%/80%) |
| 🏗️ Terraform IaC | AWS EKS cluster provisioning |
| 📊 Prometheus Monitoring | Metrics collection + Grafana dashboards |
| 🛡️ Health Checks | Liveness + readiness probes |
| ⚡ Go HTTP Server | Demo application with /health, /ready, / endpoints |

## 🛠️ Tech Stack

| Layer | Technology |
|-------|-----------|
| Application | Go 1.22 |
| Container | Docker (multi-stage) |
| Orchestration | Kubernetes (EKS) |
| CI/CD | GitHub Actions |
| IaC | Terraform |
| Monitoring | Prometheus + Grafana |

## 📁 Project Structure

```
kubernetes-cicd-pipeline/
├── src/main.go              # Go HTTP server
├── Dockerfile               # Multi-stage build
├── Makefile                 # Build commands
├── k8s/
│   ├── deployment.yaml      # 3-replica deployment with probes
│   └── hpa.yaml            # Horizontal Pod Autoscaler
├── terraform/
│   └── main.tf             # EKS cluster + VPC
├── docs/
│   └── ci-pipeline.yml     # GitHub Actions workflow
└── go.mod
```

## ⚡ Quick Start

```bash
# Build locally
make build && make run

# Docker
make docker-build
docker run -p 8080:8080 demo-app:latest

# Deploy to Kubernetes
kubectl apply -f k8s/

# Terraform (AWS EKS)
cd terraform && terraform init && terraform apply
```

### Endpoints

| Path | Description |
|------|-------------|
| `GET /` | App info + version |
| `GET /health` | Liveness check (JSON) |
| `GET /ready` | Readiness check |

## 📄 License

MIT
