# Baetyl

Fully containerized Go web dev kit. Preconfigured with:

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

$ `cd <projects-parent-directory> && git clone https://github.com/pascalallen/Baetyl.git`

### Add Entry to `/etc/hosts` File

$ `127.0.0.1       local.baetyl.com`

### Copy & Modify `.env` File

$ `cp .env.example .env`

### Bring Up Environment

$ `bin/up` or $ `bin/watch` (to watch for backend changes)

You will find the site running at [http://local.baetyl.com/](http://local.baetyl.com/)

### Install JavaScript Dependencies

$ `bin/yarn ci`

### Compile Assets For Development

$ `bin/yarn dev`

### Watch For Frontend Changes

$ `bin/yarn watch`

### Take Down Environment

$ `bin/down`

## Deploying To Kubernetes

### Update Container Registry

$ `bin/update-registry Dockerfile ghcr.io/pascalallen/baetyl`

### Deploy To Kubernetes

$ `kubectl apply -f etc/k8s/deployment.yaml`
$ `kubectl apply -f etc/k8s/service.yaml`

### Check Deployment Status

$ `kubectl get deployments`

### Check Service Status

$ `kubectl get services`

### Check All Resources

$ `kubectl get all`

### Follow Logs

$ `kubectl logs -f baetyl-app`

### Delete Kubernetes Resources

$ `kubectl delete -f etc/k8s/deployment.yaml`
$ `kubectl delete -f etc/k8s/service.yaml`
