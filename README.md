# k8-go-api

### Build

Clone the repo then

```
cd k8-go-api
cd cmd
go get
cd ..
go build -o server ./cmd
./server
```

The server will start at:

- Local: http://localhost:8000

### Docker build

```shell
$ docker build -t k8-go-api .
$ docker run -it -p 8000:8000 k8-go-api
```

## End points:

```
1- /api/rebuild/file
2- /api/rebuild/base64
```
