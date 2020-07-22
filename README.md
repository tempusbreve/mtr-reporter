# mtr-reporter

Command line tool to collect MTR report metrics into InfluxDB

Requires that the tool `mtr` is installed and on the PATH


```sh
Usage:

  -db string
    	InfluxDB DB (default "db0")
  -destination string
    	Destination to report on (default "1.1.1.1")
  -url string
    	InfluxDB reporting URL (default "http://localhost:8086")
  -verbose
    	Verbose logging
```
