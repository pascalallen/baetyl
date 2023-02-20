# Baetyl

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pascalallen/baetyl)
[![Go Report Card](https://goreportcard.com/badge/github.com/pascalallen/baetyl)](https://goreportcard.com/report/github.com/pascalallen/baetyl)
![GitHub Workflow Status (with branch)](https://img.shields.io/github/actions/workflow/status/pascalallen/baetyl/go.yml?branch=main)
![GitHub](https://img.shields.io/github/license/pascalallen/baetyl)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/pascalallen/baetyl)

A lightweight SDK for Go. Preconfigured with:

- Kubernetes (coming soon)
- Docker
- Postgres
- Go
- Adminer
- React
- TypeScript
- Sass
- Mercure (coming soon)
- Command, and event bus (coming soon)
- Models, database migrations, and seeders for User, SecurityToken, Role, and Permission

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Development Environment Setup

### Clone Repository

```bash
cd <projects-parent-directory> && git clone https://github.com/pascalallen/baetyl.git
```

### Add Entry to `/etc/hosts` File

```bash
127.0.0.1       local.baetyl.com
```

### Copy & Modify `.env` File

```bash
cp .env.example .env
```

### Bring Up Environment

```bash
bin/up
``` 

or (to watch for backend changes)

```bash
bin/watch
```

You will find the site running at [http://local.baetyl.com/](http://local.baetyl.com/)

### Install JavaScript Dependencies

```bash
bin/yarn ci
```

### Compile Assets For Development

```bash
bin/yarn dev
```

### Watch For Frontend Changes

```bash
bin/yarn watch
```

### Take Down Environment

```bash
bin/down
```

## Deploying To Kubernetes

### Update Container Registry

```bash
bin/update-registry Dockerfile ghcr.io/pascalallen/baetyl
```

### Deploy To Kubernetes

```bash
kubectl apply -f etc/k8s/deployment.yaml
kubectl apply -f etc/k8s/service.yaml
```

### Check Deployment Status

```bash
kubectl get deployments
```

### Check Service Status

```bash
kubectl get services
```

### Check All Resources

```bash
kubectl get all
```

### Follow Logs

```bash
kubectl logs -f baetyl-app
```

### Delete Kubernetes Resources

```bash
kubectl delete -f etc/k8s/deployment.yaml
kubectl delete -f etc/k8s/service.yaml
```
