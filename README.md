# pg-ping

`pg-ping` is a command line utility to continously ping your postgres. This is useful to check if there is downtime when doing changes to your postgres instance.

## Configuration

```bash
export PGPING_USERNAME=your-username
export PGPING_PASSWORD=your-password
export PGPING_HOST=your-host
export PGPING_DBNAME=your-dbname
export PGPING_QUERY=your-query # default is select 1
export PGPING_FREQUENCYINMS=300 # default is 1 second
```

## Usage

```bash
$ pg-ping
{"Value":"1","TimeTakenInMS":0.408563,"Failed":false,"FailureMessage":""}
{"Value":"1","TimeTakenInMS":0.46392500000000003,"Failed":false,"FailureMessage":""}
{"Value":"","TimeTakenInMS":0.634099,"Failed":true,"FailureMessage":"dial tcp 10.134.125.111:5432: connect: connection refused"} # Downtime
{"Value":"","TimeTakenInMS":0.402107,"Failed":true,"FailureMessage":"dial tcp 10.134.125.111:5432: connect: connection refused"}
{"Value":"1","TimeTakenInMS":10.726904000000001,"Failed":false,"FailureMessage":""}
{"Value":"1","TimeTakenInMS":0.372709,"Failed":false,"FailureMessage":""}
{"Value":"1","TimeTakenInMS":0.429533,"Failed":false,"FailureMessage":""}
```

## Build

```
go build -o pg-ping main.go
```