# k8-go-api

compile project 
go build -o server ./app
./server

docker build 
docker build -t k8-go-api .
docker run -it -p 8100:8100 k8-go-api
