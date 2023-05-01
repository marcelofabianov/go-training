# Golang

Início de estudo GoLang

Objetivo é criar uma api rest de crud de customers.

## Instalação

```bash
rm -rf .git && cp docker/local/.env.example .env && cp docker/local/docker-compose.example.yml docker-compose.yml
```

## Uso

```bash
docker-compose up -d
```

Hello World

```bash
curl --location 'http://localhost/8000' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json'
```
