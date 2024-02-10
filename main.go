package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type Configuration struct {
	Port  int  `mapstructure:"port" json:"port" yaml:"port"`
	Debug bool `mapstructure:"debug" json:"debug" yaml:"debug"`
}

func setupConfig() *Configuration {
	var config Configuration

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/config")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("app")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(&config); err != nil {
			log.Fatalf("Unable to decode into struct, %s", err)
		}
	})
	viper.WatchConfig()

	return &config
}

func main() {
	config := setupConfig()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/config", func(c *fiber.Ctx) error {
		return c.JSON(config)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Port)))
}
