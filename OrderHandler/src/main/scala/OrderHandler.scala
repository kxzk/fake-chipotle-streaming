import org.apache.spark.sql._
import org.apache.spark.sql.functions._
import org.apache.spark.sql.streaming._
import org.apache.spark.sql.types._
import org.apache.spark.sql.cassandra._

case class Order(time: String, name: String, age: Integer, item: String, price: Double, quantity: Integer)

object OrderHandler {
  def main(args: Array[String]) {
    val spark = SparkSession
      .builder
      .appName("Order Handler")
      .config("spark.cassandra.connection.host", "localhost")
      .getOrCreate()

    import spark.implicits._

    val inputDF = spark
      .readStream
      .format("kafka")
      .option("kafka.bootstrap.servers", "localhost:9092")
      .option("subscribe", "orders")
      .load()

    val rawDF = inputDF.selectExpr("CAST(value as STRING)").as[String]

    val expandedDF = rawDF.map(row => row.split(","))
      .map(row => Order(
        row(0),
        row(1),
        row(2).toInt,
        row(3),
        row(4).toDouble,
        row(5).toInt,
      ))

    val summaryDF = expandedDF
      .groupBy("item")
      .agg(sum("quantity").as("item_count"))

    val query = summaryDF
      .writeStream
      .trigger(Trigger.ProcessingTime("10 seconds"))
      .foreachBatch{ (batchDF: DataFrame, batchID: Long) =>
        println(s"Writing to Cassandra $batchID")
        batchDF.write
          .cassandraFormat("user_prefs_v1", "chipotle")
          .mode("append")
          .save()
      }
      .outputMode("update")
      .start()

    query.awaitTermination()
  }
}
