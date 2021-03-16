package cmd

import (
	"context"
	"os"

	"github.com/gaius-qi/honk/internal/config"
	"github.com/gaius-qi/honk/pkg/stock"
	"github.com/jedib0t/go-pretty/v6/table"

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
	Args:         cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cfg.Number = args[0]
		logrus.Debugf("load config success: %#v", cfg)

		s := stock.NewStockContext(ctx, cfg.Platform, cfg)
		data, err := s.Get()
		logrus.Debugf("get stock data success: %#v", data)
		if err != nil {
			logrus.Errorf("get stock data failed")
			return err
		}

		prettyPrint(data)
		return nil
	},
}

// Execute is the entry point of the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	// init config
	cfg = config.New()

	// initialize cobra
	cobra.OnInitialize(initConfig)
	addFlags(rootCmd, cfg)
}

// initConfig reads in config file and ENV variables if set
func initConfig() {
	// allow to read in from environment
	viper.SetEnvPrefix("honk")
	viper.AutomaticEnv()

	for _, e := range []string{"index", "platform"} {
		if err := viper.BindEnv(e); err != nil {
			logrus.Fatalf(errors.Wrap(err, "cannot bind environment variable").Error())
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		logrus.Fatalf(errors.Wrap(err, "cannot unmarshal config").Error())
	}

	// config logger
	logConfig(cfg)
}

func addFlags(cmd *cobra.Command, cfg *config.Config) {
	rootCmd.PersistentFlags().VarP(&cfg.Platform, "platform", "p", "set the source platform for stock data")
	rootCmd.PersistentFlags().VarP(&cfg.Index, "index", "i", "set the stock market index")
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

func prettyPrint(s *stock.Stock) {
	timeLayout := "2006-01-02 15:04:05"

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBlackOnMagentaWhite)

	t.AppendHeader(table.Row{"Number", "Current Price", "Opening Price", "Previous Closing Price", "High Price", "Low Price", "Date"})
	t.AppendRows([]table.Row{
		{s.Number, s.CurrentPrice, s.OpeningPrice, s.PreviousClosingPrice, s.HighPrice, s.LowPrice, s.Date.Format(timeLayout)},
	})
	t.AppendSeparator()

	t.Render()
}
