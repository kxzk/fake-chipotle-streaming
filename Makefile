.PHONY: spark sparkdev package

KAFKA_PKG = org.apache.spark:spark-sql-kafka-0-10_2.12:3.2.0
CASS_PKG = com.datastax.spark:spark-cassandra-connector_2.12:3.2.0

spark: package
	@echo "== INFO == submitting spark job"
	cd aggregate-orders && spark-submit --class AggOrders --master local[*] --packages $(KAFKA_PKG),$(CASS_PKG) target/scala-2.12/aggregate-orders_2.12-1.0.jar

sparkdev: package
	@echo "== INFO == submitting spark job"
	cd aggregate-orders && spark-submit --class AggOrders --master local[*] --packages $(KAFKA_PKG) target/scala-2.12/aggregate-orders_2.12-1.0.jar

package:
	@echo "== INFO == packaging spark code using sbt"
	cd aggregate-orders/ && sbt package
