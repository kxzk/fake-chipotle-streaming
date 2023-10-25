# Fake Chipotle Streaming

> Demo project to mess around.

![header](https://www.chipotle.co.uk/content/dam/chipotle/pages/app/uk/web-banner.jpg)

![Scala](https://img.shields.io/badge/scala-%23DC322F.svg?style=for-the-badge&logo=scala&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Apache Kafka](https://img.shields.io/badge/Apache%20Kafka-000?style=for-the-badge&logo=apachekafka)
![ApacheCassandra](https://img.shields.io/badge/cassandra-%231287B1.svg?style=for-the-badge&logo=apache-cassandra&logoColor=white)

# Project Overview

Initially, fake chipotle order data is generated via Go using the script under `/src`. Within the script we serialize the order into JSON and push to a predefined Kafka topic `orders`. From there, we grab the topic within Spark, parse the data back to JSON (since must be []byte to send via Kafka) and explode nested columns. Finally, we aggregate order items on a 5 minute basis and push the micro-batch result to a table in Cassandra.

This workflow lighlty ressembles a real-world real-time (ish) reporting scenario. Ideally there would be some sort of BI tool that sits on top of Cassandra to query the data. Additionally, the configuration and replication of the services is overly simplified since I did this project on a singular [DigitalOcean droplet](https://www.digitalocean.com/products/droplets).

# Architecture Overview

```text
+---------+     +---------+     +---------+     +-------------+
|         | --> |         |     |         |     |             |
|  ORDER  | --> |  KAFKA  | --> |  SPARK  | --> |  CASSANDRA  |
|         | --> |         |     |         |     |             |
+---------+     +---------+     +---------+     +-------------+
```
# Installation

* [Install Apache Kafka](https://tecadmin.net/how-to-install-apache-kafka-on-ubuntu-22-04/)
* [Install Apache Spark](https://www.vultr.com/docs/install-apache-spark-on-ubuntu-20-04/)
* [Install SBT](https://www.scala-sbt.org/1.x/docs/Installing-sbt-on-Linux.html)
* [Install Cassandra](https://www.digitalocean.com/community/tutorials/how-to-install-cassandra-and-run-a-single-node-cluster-on-ubuntu-22-04)

# Up & Running

* Create Kafka Topic
```bash
kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partition 1 --topic orders
```

* Set Up Cassandra
```sql
create keyspace chipotle with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

create table top_orders_every_5_mins (
    start timestamp,
    end timestamp,
    name text,
    count int,
    primary key (start, name)
);
```

* Submit Spark Job + Streaming
```bash
make spark
```

* Create Fake Order Data
```
cd src && go run *.go
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
* Go -> v 1.18.1

# Reference

* [Structured Streaming + Kafka Integration](https://spark.apache.org/docs/latest/structured-streaming-kafka-integration.html)
