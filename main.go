package main

import (
	"context"
	"os"
	"time"

	"github.com/kayac/go-config"
	notifier "github.com/kunihiko-t/google-home-notifier-go"
	"github.com/sirupsen/logrus"
)

// Config このアプリで使用できる設定値
type Config struct {
	Global globalConfig `toml:"global"`
	Notify notifyConfig `toml:"notify"`
	Play   playConfig   `toml:"play"`
}

type globalConfig struct {
	Host     string        `toml:"host"`
	Port     int           `toml:"port"`
	Interval time.Duration `toml:"interval"`
}

type notifyConfig struct {
	Text     string `toml:"text"`
	Language string `toml:"lang"`
}

type playConfig struct {
	URL string `toml:"url"`
}

var defaultConfig = Config{
	Global: globalConfig{
		Host:     "127.0.0.1",
		Port:     8009,
		Interval: 3,
	},
}

func main() {
	if len(os.Args) != 2 {
		logrus.Fatal("config file is missing")
	}
	path := os.Args[1]

	conf := defaultConfig
	if err := config.LoadWithEnvTOML(&conf, path); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"file":  path,
		}).Fatal("failed to parse config file")
	}

	client, err := notifier.NewClient(context.Background(), conf.Global.Host, conf.Global.Port)
	if err != nil {
		logrus.WithError(err).Fatal("unexpected error")
	}
	defer client.Close()

	// send notify
	if text := conf.Notify.Text; text != "" {
		lang := conf.Notify.Language
		if lang == "" {
			lang = "en"
		}
		if err := client.Notify(text, lang); err != nil {
			logrus.WithError(err).Error("unexpected error")
		}
		time.Sleep(conf.Global.Interval * time.Second)
	}

	// send play
	if url := conf.Play.URL; url != "" {
		if err := client.Play(url); err != nil {
			logrus.WithError(err).Error("unexpected error")
		}
		time.Sleep(conf.Global.Interval * time.Second)
	}

	// quit application
	if err := client.Quit(); err != nil {
		logrus.WithError(err).Error("unexpected error")
	}
}
