# Kubernetes Auto-Scaling CI/CD Pipeline

Production-grade CI/CD pipeline with Kubernetes, Docker, GitHub Actions, Terraform, and monitoring.

## Features
- 🚀 Automated CI/CD with GitHub Actions
- 🐳 Multi-stage Docker builds
- ☸️ Kubernetes deployment with HPA auto-scaling
- 🏗️ Infrastructure as Code with Terraform (EKS)
- 📊 Monitoring with Prometheus + Grafana
- 🔄 Rolling updates with zero downtime
- 🛡️ Health checks and readiness probes

## Tech Stack
- **Container**: Docker, Kubernetes
- **CI/CD**: GitHub Actions
- **IaC**: Terraform (AWS EKS)
- **Monitoring**: Prometheus, Grafana
- **App**: Go HTTP server (demo)

## Getting Started
```bash
make docker-build
make deploy
# Or with Terraform:
cd terraform && terraform init && terraform apply
```

## License
MIT
