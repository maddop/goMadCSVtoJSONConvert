dir# GoMadCSVtoJSONConvert

**Warning** I'm new to go - feel free to improve/fix!

Problem: I had a vast number of reports output from over night batch runs, the data in each report needed to be uploaded to a DB. 
My Solution: Create this golang application to parse a given directory, find all reports (csv or any specified delimeter) for a given date and present the contents as JSON output to be read by a client.

## Getting Started

You can modify it to suit your report requirements. Currently its configured to expect 11 fields that have been seperated by a space (you can change this to any delimeter). It expects the report names to start with a hostname and also contain the date in format 20060102 anywhere within the report name (e.g. hostname.date.report)

### Prerequisites

Reports should reside within a directory relative to where the go program is execute from (you can modify this)

```
goMadCSVtoJSONConvert
data/hosta.20180721.reporta.csv
data/hostb.20180721.reporta.csv
```

### Installing

A quick up and running:

```
go run goMadCSVtoJSONConvert.go --recorddate 20180721 --root data
```

## Authors

* **Paul Maddocks** - *Initial work* - [maddop](https://github.com/maddop)
