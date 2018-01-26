# iris-community
Website for IRIS Community Activities


### Installation

Please make sure that glide is installed.

```
go get github.com/irisnet/iris-community
cd $GOPATH/src/github.com/irisnet/iris-community
glide install
go run app.go
```

We use PostgreSQL and Redis for repository, you can use them with docker.

```
docker-compose -f docker\postgre.yml up -d
docker-compose -f docker\redis.yml up -d
```