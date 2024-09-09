# API em Golang utilizando Gin e Redis

## Tecnologias utilizadas

- Golang
- Gin
- Redis
- Docker
- Docker-Compose

## Construir a imagem do Docker
```
docker build -t gin-redis-api .
```

## Executar o container do Docker
```
docker run -p 8080:8080 gin-redis-api
```

## Executar o projeto sem docker
```
go run main.go
```

## Para rodar o docker-compose
```
docker-compose up --build
```

## Para parar o docker-compose
```
docker-compose down
```

## Endpoints

1. SET ​http://127.0.0.1:8080/set

2. GET http://127.0.0.1:8080/get/nome_chave