package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli" // imports as package "cli"
	"os"
	"strings"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(c.String("config")) > 0 {
			// Load Config.
			loadConfig(c.String("config"))
			log.Info(fmt.Sprintf("Config '%v.json' loaded...", c.String("config")))
		} else {
			log.Error("No config found. Please use config.")
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

