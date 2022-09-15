import org.apache.spark.sql._
import org.apache.spark.sql.functions._
import org.apache.spark.sql.streaming._
import org.apache.spark.sql.types._
// import org.apache.spark.sql.cassandra._

import java.sql.Timestamp

case class Customer(name: String, age: Integer, city: String, state: String)

case class Item(name: String, price: Double)

case class Order(time: Timestamp, customer: Customer, order: Array[Item], card: String)

object AggOrders {
  def main(args: Array[String]) {
    val spark = SparkSession
      .builder
      .appName("AggOrders")
      // .config("spark.cassandra.connection.host", "localhost")
      .getOrCreate()

    import spark.implicits._

    val inputDF = spark
      .readStream
      .format("kafka")
      .option("kafka.bootstrap.servers", "localhost:9092")
      .option("subscribe", "orders")
      .load()

    val rawString = inputDF.selectExpr("cast(value as string)")

    val orderSchema = Encoders.product[Order].schema

    val orderStream = rawString
      .select(from_json($"value", orderSchema) as "record")
      .select("record.*")
      .as[Order]

    val query = orderStream
      .writeStream
      .outputMode("update")
      .format("console")
      .option("truncate", "false")
      .start()

//    val query = summaryDF
//      .writeStream
//      .trigger(Trigger.ProcessingTime("10 seconds"))
//      .foreachBatch{ (batchDF: DataFrame, batchID: Long) =>
//        println(s"Writing to Cassandra $batchID")
//        batchDF.write
//          .cassandraFormat("user_prefs_v1", "chipotle")
//          .mode("append")
//          .save()
//      }
//      .outputMode("update")
//      .start()

    query.awaitTermination()
  }
}
