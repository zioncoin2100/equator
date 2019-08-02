---
title: Administration
---

Equator is responsible for providing an HTTP API to data in the Zion network. It ingests and re-serves the data produced by the zion network in a form that is easier to consume than the performance-oriented data representations used by zion-core.

## Why run equator?

The zion development foundation runs two equator servers, one for the public network and one for the test network, free for anyone's use at https://zionc.info and https://zionc.info.  These servers should be fine for development and small scale projects, but is not recommended that you use them for production services that need strong reliability.  By running equator within your own infrastructure provides a number of benefits:

  - Multiple instances can be run for redundancy and scalability.
  - Request rate limiting can be disabled.
  - Full operational control without dependency on the Zion Development Foundations operations.

## Prerequisites

Equator is a dependent upon a zion-core server.  Equator needs access to both the SQL database and the HTTP API that is published by zion-core. See [the administration guide](http://zionc.info/developers/zion-core/learn/admin.html
) to learn how to set up and administer a zion-core server.  Secondly, equator is dependent upon a postgresql server, which it uses to store processed core data for ease of use. Equator requires postgres version >= 9.3.

In addition to the two required prerequisites above, you may optionally install a redis server to be used for rate limiting requests.

## Installing

To install equator, you have a choice: either downloading a [prebuilt release for your target architecture](https://github.com/zion2100/equator/releases) and operation system, or [building equator yourself](#Building).  When either approach is complete, you will find yourself with a directory containing a file named `equator`.  This file is a native binary.

After building or unpacking equator, you simply need to copy the native binary into a directory that is part of your PATH.  Most unix-like systems have `/usr/local/bin` in PATH by default, so unless you have a preference or know better, we recommend you copy the binary there.

To test the installation, simply run `equator --help` from a terminal.  If the help for equator is displayed, your installation was successful. Note: some shells, such as zsh, cache PATH lookups.  You may need to clear your cache  (by using `rehash` in zsh, for example) before trying to run `equator --help`.


## Building

Should you decide not to use one of our prebuilt releases, you may instead build equator from source.  To do so, you need to install some developer tools:

- A unix-like operating system with the common core commands (cp, tar, mkdir, bash, etc.)
- A compatible distribution of go (we officially support go 1.6 and later)
- [gb](https://getgb.io/)
- [git](https://git-scm.com/)

Provided your workstation satisfies the requirements above, follow the steps below:

1. Clone equator's source:  `git clone https://github.com/zion2100/equator.git && cd equator`
2. Download external dependencies: `gb vendor restore`
3. Build the binary: `gb build`

After running the above commands have succeeded, the built equator will have be written into the `bin` subdirectory of the current directory.

Note:  Building directly on windows is not supported.


## Configuring

Equator is configured using command line flags or environment variables.  To see the list of command line flags that are available (and their default values) for your version of equator, run:

`equator --help`

As you will see if you run the command above, equator defines a large number of flags, however only three are required:

| flag                    | envvar                      | example                              |
|-------------------------|-----------------------------|--------------------------------------|
| `--db-url`              | `DATABASE_URL`              | postgres://localhost/equator_testnet |
| `--zion-core-db-url` | `ZION_CORE_DATABASE_URL` | postgres://localhost/core_testnet    |
| `--zion-core-url`    | `ZION_CORE_URL`          | http://localhost:11626               |

`--db-url` specifies the equator database, and its value should be a valid [PostgreSQL Connection URI](http://www.postgresql.org/docs/9.2/static/libpq-connect.html#AEN38419).  `--zion-core-db-url` specifies a zion-core database which will be used to load data about the zion ledger.  Finally, `--zion-core-url` specifies the HTTP control port for an instance of zion-core.  This URL should be associated with the zion-core that is writing to the database at `--zion-core-db-url`.

Specifying command line flags every time you invoke equator can be cumbersome, and so we recommend using environment variables.  There are many tools you can use to manage environment variables:  we recommend either [direnv](http://direnv.net/) or [dotenv](https://github.com/bkeepers/dotenv).  A template configuration that is compatible with dotenv can be found in the [equator git repo](https://github.com/zion2100/equator/blob/master/.env.template).



## Preparing the database

Before the equator server can be run, we must first prepare the equator database.  This database will be used for all of the information produced by equator, notably historical information about successful transactions that have occurred on the zion network.  

To prepare a database for equator's use, first you must ensure the database is blank.  It's easiest to simply create a new database on your postgres server specifically for equator's use.  Next you must install the schema by running `equator db init`.  Remember to use the appropriate command line flags or environment variables to configure equator as explained in [Configuring ](#Configuring).  This command will log any errors that occur.

## Running

Once your equator database is configured, you're ready to run equator.  To run equator you simply run `equator` or `equator serve`, both of which start the HTTP server and start logging to standard out.  When run, you should see some output that similar to:

```
INFO[0000] Starting equator on :8000                     pid=29013
```

The log line above announces that equator is ready to serve client requests. Note: the numbers shown above may be different for your installation.  Next we can confirm that equator is responding correctly by loading the root resource.  In the example above, that URL would be [http://127.0.0.1:8000/] and simply running `curl http://127.0.0.1:8000/` shows you that the root resource can be loaded correctly.


## Ingesting zion-core data

Equator provides most of its utility through ingested data.  Your equator server can be configured to listen for and ingest transaction results from the connected zion-core.  We recommend that within your infrastructure you run one (and only one) equator process that is configured in this way.   While running multiple ingestion processes will not corrupt the equator database, your error logs will quickly fill up as the two instances race to ingest the data from zion-core.  We may develop a system that coordinates multiple equator processes in the future, but we would also be happy to include an external contribution that accomplishes this.

To enable ingestion, you must either pass `--ingest=true` on the command line or set the `INGEST` environment variable to "true".

### Managing storage for historical data

Given an empty equator database, any and all available history on the attached zion-core instance will be ingested. Over time, this recorded history will grow unbounded, increasing storage used by the database.  To keep you costs down, you may configure equator to only retain a certain number of ledgers in the historical database.  This is done using the `--history-retention-count` flag or the `HISTORY_RETENTION_COUNT` environment variable.  Set the value to the number of recent ledgers you with to keep around, and every hour the equator subsystem will reap expired data.  Alternatively, you may execute the command `equator db reap` to force a collection.

### Surviving zion-core downtime

Equator tries to maintain a gap-free window into the history of the zion-network.  This reduces the number of edge cases that equator-dependent software must deal with, aiming to make the integration process simpler.  To maintain a gap-free history, equator needs access to all of the metadata produced by zion-core in the process of closing a ledger, and there are instances when this metadata can be lost.  Usually, this loss of metadata occurs because the zion-core node went offline and performed a catchup operation when restarted.

To ensure that the metadata required by equator is maintained, you have several options: You may either set the `CATCHUP_COMPLETE` zion-core configuration option to `true` or configure `CATCHUP_RECENT` to determine the amount of time your zion-core can be offline without having to rebuild your equator database.

We _do not_ recommend using the `CATCHUP_COMPLETE` method, as this will force zion-core to apply every transaction from the beginning of the ledger, which will take an ever increasing amount of time.  Instead, we recommend you set the `CATCHUP_RECENT` config value.  To do this, determine how long of a downtime you would like to survive (expressed in seconds) and divide by ten.  This roughly equates to the number of ledgers that occur within you desired grace period (ledgers roughly close at a rate of one every ten seconds).  With this value set, zion-core will replay transactions for ledgers that are recent enough, ensuring that the metadata needed by equator is present.

### Correcting gaps in historical data

In the section above, we mentioned that equator _tries_ to maintain a gap-free window.  Unfortunately, it cannot directly control the state of zion-core and so gaps may form due to extended down time.  When a gap is encountered, equator will stop ingesting historical data and complain loudly in the log with error messages (log lines will include "ledger gap detected").  To resolve this situation, you must re-establish the expected state of the zion-core database and purge historical data from equator's database.  We leave the details of this process up to the reader as it is dependent upon your operating needs and configuration, but we offer one potential solution:

We recommend you configure the HISTORY_RETENTION_COUNT in equator to a value less than or equal to the configured value for CATCHUP_RECENT in zion-core.  Given this situation any downtime that would cause a ledger gap will require a downtime greater than the amount of historical data retained by equator.  To re-establish continuity, simply:

1.  Stop equator.
2.  run `equator db reap` to clear the historical database.
3.  Clear the cursor for equator by running `zion-core -c "dropcursor?id=HORIZON"` (ensure capitilization is maintained).
4.  Clear ledger metadata from before the gap by running `zion-core -c "maintenance?queue=true"`.
5.  Restart equator.    

## Managing Stale Historical Data

Equator ingests ledger data from a connected instance of zion-core.  In the event that zion-core stops running (or if equator stops ingesting data for any other reason), the view provided by equator will start to lag behind reality.  For simpler applications, this may be fine, but in many cases this lag is unacceptable and the application should not continue operating until the lag is resolved.

To help applications that cannot tolerate lag, equator provides a configurable "staleness" threshold.  Given that enough lag has accumulated to surpass this threshold (expressed in number of ledgers), equator will only respond with an error: [`stale_history`](./errors/stale-history.md).  To configure this option, use either the `--history-stale-threshold` command line flag or the `HISTORY_STALE_THRESHOLD` environment variable.  NOTE:  non-historical requests (such as submitting transactions or finding payment paths) will not error out when the staleness threshold is surpassed.

## Monitoring

To ensure that your instance of equator is performing correctly we encourage you to monitor it, and provide both logs and metrics to do so.  

Equator will output logs to standard out.  Information about what requests are coming in will be reported, but more importantly and warnings or errors will also be emitted by default.  A correctly running equator instance will not ouput any warning or error log entries.

Metrics are collected while a equator process is running and they are exposed at the `/metrics` path.  You can see an example at (https://zionc.info/metrics).

## I'm Stuck! Help!

If any of the above steps don't work or you are otherwise prevented from correctly setting up equator, please come to our community and tell us.  Either [post a question at our Stack Exchange](https://zion.stackexchange.com/) or [chat with us on slack](http://slack.zion.org/) to ask for help.
