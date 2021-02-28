# k8-go-api

### Build

Clone the repo then

```
$ cd k8-go-api
$ go build -o server ./cmd
$ ./server
```

The server will start at:

- Local: http://localhost:8100

### Docker build

```
$ docker build -t k8-go-api .
$ docker run -it -p 8100:8100 k8-go-api
```

## End points:

```
1- /api/rebuild/file
2- /api/rebuild/base64
3- /api/rebuild/zip
```

## Postman Collections link:

```
https://www.getpostman.com/collections/78fd72df0d74b4c5e849

```
