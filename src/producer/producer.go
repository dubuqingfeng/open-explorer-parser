package main

import (
	"github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"

	"fmt"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
	"github.com/dubuqingfeng/explorer-parser/src/producer/processors"
	"time"
	"path"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"os"
	"flag"
)

func init() {
	// TODO flag
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

	initProducerLog()
}

func main() {
	log.Info("producer start")
	fmt.Printf("config: %#v\n", config.Config.APPName)
	// Load the processor by configuration
	fmt.Printf("enable processor: %#v\n", config.Config.EnableCoin)

	multiCoin := make([]processors.Proccessor, 0)

	for _, value := range config.Config.EnableCoin {
		processor := newProccessor(value)
		if processor != nil {
			multiCoin = append(multiCoin, processor)
		}
	}

	// WaitGroup
	for index, value := range multiCoin {
		// go func
		fmt.Printf("arr[%d]=%d \n", index, value.Parse("test"))
		// select
		// write to kafka
	}
}

func newProccessor(coin string) processors.Proccessor {
	switch coin {
	case "btc":
		return processors.NewBTCProcessor()
	case "eth":
		return processors.NewETHProcessor()
	}
	return nil;
}

// Init Log
func initProducerLog() {
	level, err := log.ParseLevel(config.Config.Log.Level)
	if err != nil {
		log.Errorf("init producer logger error. %+v", errors.WithStack(err))
	}

	log.SetLevel(level)

	ConfigLocalFilesystemLogger(config.Config.Log.Path, config.Config.Log.Filename, 7*time.Hour*24, time.Second*20)
}

// Rotate Log
func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		//rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)

	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
	log.AddHook(lfHook)
}
