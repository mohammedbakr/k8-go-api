<h1 align="center">k8-go-api</h1>

<p align="center">
    <a href="https://github.com/k8-proxy/k8-go-api/actions/workflows/build.yml">
        <img src="https://github.com/k8-proxy/k8-go-api/actions/workflows/build.yml/badge.svg"/>
    </a>
    <a href="https://codecov.io/gh/k8-proxy/k8-go-api">
        <img src="https://codecov.io/gh/k8-proxy/k8-go-api/branch/main/graph/badge.svg"/>
    </a>	    
    <a href="https://goreportcard.com/report/github.com/k8-proxy/k8-go-api">
      <img src="https://goreportcard.com/badge/k8-proxy/k8-go-api" alt="Go Report Card">
    </a>
	<a href="https://github.com/k8-proxy/k8-go-api/pulls">
        <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="Contributions welcome">
    </a>
    <a href="https://opensource.org/licenses/Apache-2.0">
        <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="Apache License, Version 2.0">
    </a>
    <a href="https://github.com/k8-proxy/k8-go-api/releases/latest">
        <img src="https://img.shields.io/github/release/k8-proxy/k8-go-api.svg?style=flat"/>
    </a>
</p>

# k8-go-api

- Go api that connects with rebuild engine service that recieves an infected file and rebuild it using its content.
- It uses [comm package](https://github.com/k8-proxy/k8-go-comm) that wraps MinIO and RabbitMQ together.
- for the moment we use [k8-go-echo](https://github.com/k8-proxy/k8-go-echo) that mimic the rebuild engine

### Steps of processing

- User post request with pdf file to k8-go-api.
- The k8-go-api endpoint uploads the file to MinIO and returns a URL.
- The endpoint publishes a message to RabbitMQ queue that contains Minio Pre-signed URL .
- The k8-go-echo ( mimic processing pod ) consumes the message and download pdf file from MinIO and add watermark to it and then upload it to MinIO and get a URL.
- The k8-go-echo Then publish the processed pdf file URL to the queue.
- k8-go-api consumes the message that contains URL to the processed pdf file and download it and write it to the HTTP response.


## Info

- only the rebuildfile enpoint working , the other still under development
- the docker build have some issue we are working on

## Build

- For quick start using docker to run containers for RabbitMQ and MinIO.
- Run Standalone MinIO on Docker.

```
docker run -d -p 9000:9000 -e "MINIO_ROOT_USER=<minio_root_user_name>" -e "MINIO_ROOT_PASSWORD=<minio_root_password>" minio/minio server /data
```

- Run RabbitMQ on Docker.

```
docker run -d --hostname <host_name> --name <container_name> -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

- Clone the repo.
- Rename/copy .env.example file to .env
- Modify .env file according to your needs.
- Open the terminal in your current directory then run:

```
cd k8-go-api
go build -o server ./cmd
./server
```

- The server will start at:

- Local: http://localhost:8100

### Docker build

```
docker build -t k8-go-api .
docker run -it -p 8100:8100 k8-go-api
```

## Test

For testing, run `go test ./...`

## k8-go-echo 
## Build

- Clone [k8-go-echo](https://github.com/k8-proxy/k8-go-echo) then run:

```
cd k8-go-echo
go build
./k8-go-echo
```

## End points:

1- `/api/rebuild/file`<br>
<strong>Method</strong>: POST.<br>
<strong>Description: </strong>This endpoint accepts requests to rebuild a file with Glasswall d-FIRST™ Engine. Both the file and the Content Management Policy are sent in the request body with a 'Content-Type' of 'multipart/form-data'. The Rebuilt file is then returned in the response body with a 'Content-Type' of 'application/octet-stream'.<br>

2- `/api/rebuild/base64`<br>
<strong>Method</strong> POST.<br>
<strong>Description: </strong>This endpoint accepts requests to rebuild a file with Glasswall d-FIRST™ Engine. The request body contains the Base64 representation of the file and Glasswall Content Management Flags with a 'Content-Type' of 'application/json'. A Base64 Representation of the rebuilt file is then returned in the response with a 'Content-Type' of 'text/plain'.<br>

Select a file below to copy its Base64 Encoded representation to clipboard. The Total supported request size of the API gateway is 6MB, therefore the base64 encoded string must also be less than 6MB.<br>

3- `/api/rebuild/zip`<br>
<strong>Method</strong>: POST.<br>
<strong>Description: </strong>This endpoint accepts requests to rebuild a zip file with Glasswall d-FIRST™ Engine. Both the file and the Content Management Policy are sent in the request body with a 'Content-Type' of 'multipart/form-data'. The Rebuilt file is then returned in the response body with a 'Content-Type' of 'application/octet-stream'.

## Postman Collections link:

https://www.getpostman.com/collections/78fd72df0d74b4c5e849

## Video Demo

https://www.youtube.com/watch?v=TlXwsJrXe68&amp;feature=youtu.be
