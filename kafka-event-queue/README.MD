## Kafka Setup

Now start the Kafka broker with the new `docker compose` command:
```
$ docker compose up -d
```
## Publishing & Subscribing
Try and run this basic skeleton of Pub-Sub Event Queue:

Compile the publisher & subscriber:
```
$ go build -o build-out/publisher util.go publisher.go
$ go build -o build-out/subscriber util.go subscriber.go
```

Run the publisher:
```
$ ./out/publisher kafka-sv.properties
``` 

Run the subscriber:
```
$ ./out/subscriber kafka-sv.properties
```

~
Check [AIBlockchain] (https://github.com/ai-coin/KafkaBlockchain) for a tamper immune pub-sub blockchian architecture. Also resolves the workflow for on/off chain data.


## Gin Server

Install & setup Gin Server:
```
$ go get -u github.com/gin-gonic/gin
```

`server.go` will serve as the gin server to facilitate any requests to add/remove users or event queues to the server's knowledge. It will also facilitate the addition/removal of user's access to event queues. 
> TODO: * Adding publish/subscribe calls & kafka integration to those calls

`