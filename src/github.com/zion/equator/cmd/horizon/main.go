package main

import (
	"log"
	"runtime"

	"github.com/PuerkitoBio/throttled"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zion/equator"
	hlog "github.com/zion/equator/log"
)

var app *equator.App
var config equator.Config
var version string

var rootCmd *cobra.Command

func main() {
	if version != "" {
		equator.SetVersion(version)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd.Execute()
}

func init() {
	viper.SetDefault("port", 8000)
	viper.SetDefault("history-retention-count", 0)

	viper.BindEnv("port", "PORT")
	viper.BindEnv("db-url", "DATABASE_URL")
	viper.BindEnv("zion-core-db-url", "ZION_CORE_DATABASE_URL")
	viper.BindEnv("zion-core-url", "ZION_CORE_URL")
	viper.BindEnv("friendbot-secret", "FRIENDBOT_SECRET")
	viper.BindEnv("per-hour-rate-limit", "PER_HOUR_RATE_LIMIT")
	viper.BindEnv("redis-url", "REDIS_URL")
	viper.BindEnv("ruby-equator-url", "RUBY_HORIZON_URL")
	viper.BindEnv("log-level", "LOG_LEVEL")
	viper.BindEnv("sentry-dsn", "SENTRY_DSN")
	viper.BindEnv("loggly-token", "LOGGLY_TOKEN")
	viper.BindEnv("loggly-host", "LOGGLY_HOST")
	viper.BindEnv("tls-cert", "TLS_CERT")
	viper.BindEnv("tls-key", "TLS_KEY")
	viper.BindEnv("ingest", "INGEST")
	viper.BindEnv("network-passphrase", "NETWORK_PASSPHRASE")
	viper.BindEnv("history-retention-count", "HISTORY_RETENTION_COUNT")
	viper.BindEnv("history-stale-threshold", "HISTORY_STALE_THRESHOLD")
	viper.BindEnv("skip-cursor-update", "SKIP_CURSOR_UPDATE")

	rootCmd = &cobra.Command{
		Use:   "equator",
		Short: "client-facing api server for the zion network",
		Long:  "client-facing api server for the zion network",
		Run: func(cmd *cobra.Command, args []string) {
			initApp(cmd, args)
			app.Serve()
		},
	}

	rootCmd.Flags().String(
		"db-url",
		"",
		"equator postgres database to connect with",
	)

	rootCmd.Flags().String(
		"zion-core-db-url",
		"",
		"zion-core postgres database to connect with",
	)

	rootCmd.Flags().String(
		"zion-core-url",
		"",
		"zion-core to connect with (for http commands)",
	)

	rootCmd.Flags().Int(
		"port",
		8000,
		"tcp port to listen on for http requests",
	)

	rootCmd.Flags().Int(
		"per-hour-rate-limit",
		3600,
		"max count of requests allowed in a one hour period, by remote ip address",
	)

	rootCmd.Flags().String(
		"redis-url",
		"",
		"redis to connect with, for rate limiting",
	)

	rootCmd.Flags().String(
		"log-level",
		"info",
		"Minimum log severity (debug, info, warn, error) to log",
	)

	rootCmd.Flags().String(
		"sentry-dsn",
		"",
		"Sentry URL to which panics and errors should be reported",
	)

	rootCmd.Flags().String(
		"loggly-token",
		"",
		"Loggly token, used to configure log forwarding to loggly",
	)

	rootCmd.Flags().String(
		"loggly-host",
		"",
		"Hostname to be added to every loggly log event",
	)

	rootCmd.Flags().String(
		"friendbot-secret",
		"",
		"Secret seed for friendbot functionality. When empty, friendbot will be disabled",
	)

	rootCmd.Flags().String(
		"tls-cert",
		"",
		"The TLS certificate file to use for securing connections to equator",
	)

	rootCmd.Flags().String(
		"tls-key",
		"",
		"The TLS private key file to use for securing connections to equator",
	)

	rootCmd.Flags().Bool(
		"ingest",
		false,
		"causes this equator process to ingest data from zion-core into equator's db",
	)

	rootCmd.Flags().String(
		"network-passphrase",
		"",
		"Override the network passphrase",
	)

	rootCmd.Flags().Uint(
		"history-retention-count",
		0,
		"the minimum number of ledgers to maintain within equator's history tables.  0 signifies an unlimited number of ledgers will be retained",
	)

	rootCmd.Flags().Uint(
		"history-stale-threshold",
		0,
		"the maximum number of ledgers the history db is allowed to be out of date from the connected zion-core db before equator considers history stale",
	)

	rootCmd.AddCommand(dbCmd)

	viper.BindPFlags(rootCmd.Flags())
}

func initApp(cmd *cobra.Command, args []string) {
	initConfig()

	var err error
	app, err = equator.NewApp(config)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func initConfig() {
	if viper.GetString("db-url") == "" {
		log.Fatal("Invalid config: db-url is blank.  Please specify --db-url on the command line or set the DATABASE_URL environment variable.")
	}

	if viper.GetString("zion-core-db-url") == "" {
		log.Fatal("Invalid config: zion-core-db-url is blank.  Please specify --zion-core-db-url on the command line or set the ZION_CORE_DATABASE_URL environment variable.")
	}

	if viper.GetString("zion-core-url") == "" {
		log.Fatal("Invalid config: zion-core-url is blank.  Please specify --zion-core-url on the command line or set the ZION_CORE_URL environment variable.")
	}

	ll, err := logrus.ParseLevel(viper.GetString("log-level"))

	if err != nil {
		log.Fatalf("Could not parse log-level: %v", viper.GetString("log-level"))
	}

	hlog.DefaultLogger.Level = ll

	cert, key := viper.GetString("tls-cert"), viper.GetString("tls-key")

	switch {
	case cert != "" && key == "":
		log.Fatal("Invalid TLS config: key not configured")
	case cert == "" && key != "":
		log.Fatal("Invalid TLS config: cert not configured")
	}

	config = equator.Config{
		DatabaseURL:            viper.GetString("db-url"),
		ZioncoreDatabaseURL: viper.GetString("zion-core-db-url"),
		ZioncoreURL:         viper.GetString("zion-core-url"),
		Port:                   viper.GetInt("port"),
		RateLimit:              throttled.PerHour(viper.GetInt("per-hour-rate-limit")),
		RedisURL:               viper.GetString("redis-url"),
		LogLevel:               ll,
		SentryDSN:              viper.GetString("sentry-dsn"),
		LogglyToken:            viper.GetString("loggly-token"),
		LogglyHost:             viper.GetString("loggly-host"),
		FriendbotSecret:        viper.GetString("friendbot-secret"),
		TLSCert:                cert,
		TLSKey:                 key,
		Ingest:                 viper.GetBool("ingest"),
		HistoryRetentionCount:  uint(viper.GetInt("history-retention-count")),
		StaleThreshold:         uint(viper.GetInt("history-stale-threshold")),
		SkipCursorUpdate:       viper.GetBool("skip-cursor-update"),
	}
}
