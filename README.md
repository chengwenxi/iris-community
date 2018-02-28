# iris-community
Website for IRIS Community Activities


### Installation

Please make sure that glide is installed.

```
go get github.com/irisnet/iris-community
cd $GOPATH/src/github.com/irisnet/iris-community
glide install
go run main.go
```

We use PostgreSQL and Redis for repository, you can use them with docker.

```
docker-compose -f docker/postgre.yml up -d
docker-compose -f docker/redis.yml up -d
```

You will visit [API document](http://127.0.0.1:8080/swagger/index.html) 