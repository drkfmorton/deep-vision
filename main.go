package main

import (
	"deep-vision/liveStream"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli" // imports as package "cli"
	"os"
	"strconv"
	"strings"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
			&cli.BoolFlag{
				Name:  "live",
				Aliases: []string{"s"},
				Usage: "Enables liveStream.",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(c.String("config")) > 0 {
			// Load Config.
			loadConfig(c.String("config"))
			log.Info(fmt.Sprintf("Config '%v.json' loaded...", c.String("config")))
		} else {
			log.Fatal("No config found. Please use config.")
		}

		if b, _ := strconv.ParseBool(c.String("live")); b == true {
			log.Info("Starting live stream...")
			startStreaming()
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// Check and Load dependencies

	str := fmt.Sprintf("%v", viper.Get("Dependencies"))
	s := strings.Split((str) , ",")
	for index, element := range s{
		fmt.Println(index)
		fmt.Println(element)
	}

	log.Info("Start Streaming...")
	log.SetLevel(log.InfoLevel)

}

func loadConfig(fileName string) {
	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.SetConfigType("json") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./configs/") // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func startStreaming() {
	w, _ := strconv.Atoi(fmt.Sprintf("%v", viper.Get("video_width")))
	h, _ := strconv.Atoi(fmt.Sprintf("%v", viper.Get("video_height")))
	r, _ := strconv.Atoi(fmt.Sprintf("%v", viper.Get("video_rotation")))
	v, _ := strconv.ParseBool(fmt.Sprintf("%v", viper.Get("is_verbose")))

	liveStream.Init(
		fmt.Sprintf("%v", viper.Get("youtube_stream_key")),
		w,
		h,
		r,
		fmt.Sprintf("%v", viper.Get("video_exposure")),
		fmt.Sprintf("%v", viper.Get("video_awb")),
		v,
	)

	liveStream.Start()

	log.Info("Starting live stream...")
}

