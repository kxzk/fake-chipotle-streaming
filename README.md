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

# Up & Running

* [Install Apache Kafka](https://tecadmin.net/how-to-install-apache-kafka-on-ubuntu-22-04/)
