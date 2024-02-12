## Docs

- [Login](docs/login.md)
- [Validate Token](docs/0.general.md#validate-token)

## Setup

go version 1.20

```shell
cd $(go env GOPATH)/src/github.com/go-chi-base

git clone git@github.com:go-chi-base/zeus.git

cd zeus

# install all dependencies
go mod tidy

go run main.go
```


## build

```shell
# update environment then run
docker-compose up --build -d
```