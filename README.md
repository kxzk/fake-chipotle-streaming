# What's Poppin?

> Demo project to mess around.

# Project Overview

# Architecture Overview

```text
+--------+     +---------+     +---------+     +-------------+
|        | --> |         |     |         |     |             |
|  SONG  | --> |  KAFKA  | --> |  SPARK  | --> |  CASSANDRA  |
|        | --> |         |     |         |     |             |
+--------+     +---------+     +---------+     +-------------+
```

# Data

[Hip Hop Song Data](https://github.com/pdp2600/chartscraper/blob/master/ChartScraper_data/All_Hip_Hop_Songs_from_1958-10-20_to_2018-12-31.csv)

* Filtered to songs post 1993 using [visidata](https://www.visidata.org).
* Removed duplicate entries - since data was originally for weekly charting - using awk.

```bash
awk '{if (!($0 in x)) {print $0; x[$0]=1}}' input.csv > output.csv
```

# Up & Running
