package cmd

import (
	"os"
	"runtime"

	"github.com/Liangyu-Zhou/registry-demo/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	log         = logrus.WithField("prefix", "cmd")
	cfgFilePath string
	rootCmd     = &cobra.Command{
		Use:   "faucet",
		Short: "Run a faucet server for Ethereum using captcha",
		RunE: func(command *cobra.Command, args []string) error {
			runtime.GOMAXPROCS(runtime.NumCPU())
			customFormatter := new(logrus.TextFormatter)
			customFormatter.TimestampFormat = "2006-01-02 15:04:05"
			logrus.SetFormatter(customFormatter)
			customFormatter.FullTimestamp = true

			var cfg *internal.Config
			if err := viper.Unmarshal(&cfg); err != nil {
				log.Fatal(err)
			}

			srv, err := internal.NewServer(cfg)
			if err != nil {
				log.WithError(err).Fatal("Could not initialize faucet server")
			}
			srv.Start()
			return nil
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVar(&cfgFilePath, "config", "", "Flag config yaml file path (optional)")
	rootCmd.Flags().Int("grpc-port", 5000, "Port to serve gRPC requests")
	rootCmd.Flags().String("grpc-host", "127.0.0.1", "Host to serve gRPC requests")
	rootCmd.Flags().Int("http-port", 8000, "Port to serve REST http requests")
	rootCmd.Flags().String("http-host", "127.0.0.1", "Host to serve REST http requests")
	rootCmd.Flags().StringSlice("allowed-origins", []string{"*"}, "Allowed origins for REST http requests, comma-separated")

	// Bind all flags to a viper configuration.
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		log.Fatal(err)
	}
	viper.SetDefault("author", "Raul Jordan <raul@prysmaticlabs.com>")
	viper.SetDefault("license", "MIT")
}

func initConfig() {
	// Use config file from the flag.
	viper.SetConfigFile(cfgFilePath)
	viper.AutomaticEnv()
	if cfgFilePath != "" {
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Fatalf("Could not read config file: %s", viper.ConfigFileUsed())
		}
	}
}

// Execute the faucet root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("Could not execute root command")
		os.Exit(1)
	}
}
