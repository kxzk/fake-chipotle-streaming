name := "Aggregate Orders"

version := "1.0"

scalaVersion := "2.12.15"

libraryDependencies ++= Seq(
    "org.apache.spark" %% "spark-core" % "3.2.0" % "provided",
    "org.apache.spark" %% "spark-sql" % "3.2.0" % "provided",
    "com.datastax.spark" %% "spark-cassandra-connector" % "3.2.0"
)
