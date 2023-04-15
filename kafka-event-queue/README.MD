## Kafka Setup

Now start the Kafka broker with the new `docker compose` command:
```
$ docker compose up -d
```

## Gin Server

Install & setup the dependencies then run the server: 
```
$ go run *
```

`server.go` will serve as the gin server to facilitate any requests to add/remove users or event queues to the server's knowledge. It will also facilitate the addition/removal of user's access to event queues. 