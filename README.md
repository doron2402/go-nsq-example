# go-nsq-example
Golang example of nsq writer and reader

## Write message
- `go run wrtier.go`

## Run docker compose
- In order to run only the nsq deamons (lookupd, admin, nsqd) - `docker-compose up -d`
- To view the running containers status and mapped ports. - `docker-compose ps`
- To see the logs from the running containers. - `docker-compose logs`
