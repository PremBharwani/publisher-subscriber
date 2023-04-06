## Kafka Setup

Now start the Kafka broker with the new `docker compose` command:
```
$ docker compose up -d
```
## Publishing & Subscribing
Try and run this basic skeleton of Pub-Sub Event Queue:

Compile the publisher & subscriber:
```
$ go build -o out/producer util.go producer.go
$ go build -o out/producer util.go consumer.go
```

Run the publisher:
```
$ ./out/producer kafka-sv.properties
``` 

Run the subscriber:
```
$ ./out/consumer kafka-sv.properties
```

~
Check [AIBlockchain] (https://github.com/ai-coin/KafkaBlockchain) for a tamper immune pub-sub blockchian architecture. Also resolves the workflow for on/off chain data.
