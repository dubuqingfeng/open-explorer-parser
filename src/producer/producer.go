package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
	"strings"

	"flag"
	"fmt"
	"github.com/dubuqingfeng/explorer-parser/producer/config"
	"github.com/dubuqingfeng/explorer-parser/producer/processors"
	"github.com/dubuqingfeng/explorer-parser/utils"
	"os"
	"time"
)

func init() {
	app := &cli.App{
		Name:  "Producer",
		Usage: "",
		Action: func(c *cli.Context) error {
			set := flag.NewFlagSet("contrive", 0)
			nc := cli.NewContext(c.App, set, c)
			fmt.Printf("Load config from %#v \n", nc.String("config"))
			config.InitConfig(nc.String("config"))
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./config/producer.json",
				Usage:   "Load configuration from `FILE`",
			},
		},
	}
	app.Version = "0.0.1"
	app.Name = "Explorer Parser Producer"
	app.Run(os.Args)

	utils.InitLog(config.Config.Log.Level, config.Config.Log.Path, config.Config.Log.Filename)
}

func main() {
	log.Infof("producer start, app name: %#v", config.Config.APPName)
	// Load the processor by configuration.
	log.Infof("enable processor: %#v", config.Config.EnableCoin)

	multiCoin := make([]processors.Processor, 0)

	// So it need to restart to reload the configuration.
	for _, value := range config.Config.EnableCoin {
		processor := newProcessor(strings.ToLower(value))
		if processor != nil {
			multiCoin = append(multiCoin, processor)
		}
	}

	for {
		for _, value := range multiCoin {
			go value.Parse("test")
		}
		time.Sleep(1 * time.Second)
	}
}

func newProcessor(coin string) processors.Processor {
	switch coin {
	case "btc":
		return processors.NewBTCProcessor()
	case "eth":
		return processors.NewETHProcessor()
	case "xmr":
		return processors.NewXMRProcessor()
	}
	return nil
}
