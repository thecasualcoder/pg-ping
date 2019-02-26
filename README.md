# pg-ping

`pg-ping` is a command line utility to continously ping your postgres. This is useful to check if there is downtime when doing changes to your postgres instance.

## Installation

Using Homebrew

```bash
brew tap thecasualcoder/stable
brew install pg-ping
```

## Usage

```bash
NAME:
   pg-ping - Ping your postgres continously

USAGE:
   pg-ping [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --once                       Ping only once and quit [$PGPING_ONCE]
   --debug                      Print debug logs [$PGPING_DEBUG]
   --username value, -U value   Username to connect to postgres [$PGPING_USERNAME]
   --password value, -p value   Password to connect to postgres [$PGPING_PASSWORD]
   --host value, -h value       Host of postgres server [$PGPING_HOST]
   --dbname value, -d value     DBName to connect to [$PGPING_DBNAME]
   --frequency value, -f value  Frequency to ping (default: 0) [$PGPING_FREQUENCY_IN_MS]
   --query value, -c value      Query to run (default: "SELECT 1") [$PGPING_QUERY]
   --help                       
   --version, -v                print the version
```

## Example

```bash
$ pg-ping -U myuser -h myhost -d mydb -c 'SELECT 1'
{"Value":"1","TimeTakenInMS":0.408563,"Failed":false,"FailureMessage":""}
{"Value":"1","TimeTakenInMS":0.46392500000000003,"Failed":false,"FailureMessage":""}
{"Value":"","TimeTakenInMS":0.634099,"Failed":true,"FailureMessage":"dial tcp 10.134.125.111:5432: connect: connection refused"} # Downtime
{"Value":"","TimeTakenInMS":0.402107,"Failed":true,"FailureMessage":"dial tcp 10.134.125.111:5432: connect: connection refused"}
{"Value":"1","TimeTakenInMS":10.726904000000001,"Failed":false,"FailureMessage":""}
{"Value":"1","TimeTakenInMS":0.372709,"Failed":false,"FailureMessage":""}
{"Value":"1","TimeTakenInMS":0.429533,"Failed":false,"FailureMessage":""}
```

## Build locally

```
git clone https://github.com/thecasualcoder/pg-ping.git
make compile
```