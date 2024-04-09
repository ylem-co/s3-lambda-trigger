package config

import (
	"encoding/json"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type config struct {
	LogLevel string `split_words:"true" default:"info"`
	Env      string `split_words:"true" default:"production"`
	API      api
	S3       s3
}

type api struct {
	BasicAuthUsername string `split_words:"true"`
	BasicAuthPassword string `split_words:"true"`
	ClientId          string `split_words:"true"`
	ClientSecret      string `split_words:"true"`
}

type s3 struct {
	// "<bucket_name>/path/to/directory/*.json:wf_uuid1,wf_uuid2;"
	Mapping string
}

func Cfg() config {
	return c
}

var c config

func init() {
	c = new()

	lvl, err := log.ParseLevel(c.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(lvl)

	formattedConfig, _ := json.MarshalIndent(c, "", "    ")
	log.Debug("Configuration: ", string(formattedConfig))
}

func new() config {
	godotenv.Load(".env.local")
	godotenv.Load()
	var c config
	err := envconfig.Process("DTMN", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return c
}
