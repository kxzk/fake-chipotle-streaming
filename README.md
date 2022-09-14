# Fake Chipotle Streaming

> Demo project to mess around.

# Project Overview


# Architecture Overview

```text
+---------+     +---------+     +---------+     +-------------+
|         | --> |         |     |         |     |             |
|  ORDER  | --> |  KAFKA  | --> |  SPARK  | --> |  CASSANDRA  |
|         | --> |         |     |         |     |             |
+---------+     +---------+     +---------+     +-------------+
```
# Data

* [Chipotle Prices](https://www.fastfoodmenuprices.com/chipotle-prices/)
* [GoFakeIt](https://github.com/brianvoe/gofakeit)

# Chipotle Order Schema

```json
{
  "time": "2022-09-14 22:03:15",
  "customer": {
    "name": "Yoshiko McGlynn",
    "age": 43,
    "city": "Chula Vista",
    "state": "Texas"
  },
  "order": [
    {
      "name": "Patron Margarita",
      "price": 7.15
    },
    {
      "name": "Salad (Chicken)",
      "price": 6.5
    }
  ],
  "creditCard": "Mastercard"
}
```

# Versions

* Spark -> v 3.2.0
* Scala -> v 2.12.15
* SBT -> v 1.7.1
* Cassandra -> v 4.0.5
* Kafka -> v 3.2.1

# Installation

* [Install Apache Kafka](https://tecadmin.net/how-to-install-apache-kafka-on-ubuntu-22-04/)
* [Install Apache Spark](https://www.vultr.com/docs/install-apache-spark-on-ubuntu-20-04/)
* [Install SBT](https://www.scala-sbt.org/1.x/docs/Installing-sbt-on-Linux.html)
* [Install Cassandra](https://www.digitalocean.com/community/tutorials/how-to-install-cassandra-and-run-a-single-node-cluster-on-ubuntu-22-04)

# Reference

* [Structured Streaming + Kafka Integration](https://spark.apache.org/docs/latest/structured-streaming-kafka-integration.html)

# Up & Running

* Submit Spark Job + Streaming
```bash
make spark
```

* Create Fake Order Data
```
cd src && go run *.go
```
