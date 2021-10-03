# GORM CRUD

A GO simple CRUD example made with [GORM](https://gorm.io/index.html)

## Requirements
Git, docker, docker-compose & go.

## Usage
First build & up containers:
```shell
docker-compose up -d --build
```

Checkout to the following branches:

### Create
```shell
git checkout create
go run main.go
```

### Select all
```shell
git checkout select
go run main.go
```

### Select by ID
```shell
git checkout select-id
go run main.go
```

### Update
```shell
git checkout update
go run main.go
```

### Delete
```shell
git checkout update
go run main.go
```

### Trasaction
```shell
git checkout tx
go run main.go
```
