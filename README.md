# Baetyl

Fully containerized Go web dev kit. Preconfigured with Docker, Postgres, Go, React, TypeScript, and Sass.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Development Environment Setup

### Clone Repository

$ `cd <projects-parent-directory> && git clone https://github.com/pascalallen/Baetyl.git`

### Add Entry to `/etc/hosts` File

$ `127.0.0.1       local.baetyl.com`

### Copy & Modify .env File

$ `cp .env.example .env`

### Bring Up Environment

$ `bin/up`

You will find the site running at [http://local.baetyl.com/](http://local.baetyl.com/)

### Install JavaScript Dependencies

$ `bin/yarn ci`

### Compile Assets For Development

$ `bin/yarn dev`

### Watch For Frontend Changes

$ `bin/yarn watch`

### Take Down Environment

$ `bin/down`
