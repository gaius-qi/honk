package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/gaius-qi/honk/internal/config"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "honk",
	Short: "Show stock real-time data tools",
	Long: `A command line tool to display real-time stock 
information and analysis results.
Complete documentation is available at https://github.com/gaius-qi/honk`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		fmt.Println(ctx)
		return nil
	},
}

// Execute is the entry point of the command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// init config
	cfg = config.New()

	// initialize cobra
	cobra.OnInitialize(initConfig)
	addFlags(rootCmd, cfg)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// allow to read in from environment
	viper.SetEnvPrefix("honk")
	viper.AutomaticEnv()

	for _, e := range []string{"index", "platform"} {
		if err := viper.BindEnv(e); err != nil {
			log.Fatalf(errors.Wrap(err, "cannot bind environment variable").Error())
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf(errors.Wrap(err, "cannot unmarshal config").Error())
	}

	// config logger
	logConfig(cfg)
}

func addFlags(cmd *cobra.Command, cfg *config.Config) {
	rootCmd.PersistentFlags().StringVarP(&cfg.Platform, "platform", "p", config.DefaultIndex, "set the source platform for stock data")
	rootCmd.PersistentFlags().StringVarP(&cfg.Index, "index", "i", config.DefaultPlatform, "set the stock market index")
}

func logConfig(cfg *config.Config) {
	// reset log format
	if cfg.LogFormat == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	// set debug log level
	if cfg.Debug {
		cfg.LogLevel = "debug"
	}

	// set the configured log level
	if level, err := logrus.ParseLevel(cfg.LogLevel); err == nil {
		logrus.SetLevel(level)
	}
}
