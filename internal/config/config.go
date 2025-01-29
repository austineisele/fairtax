package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const rate_levy = "RateLevy"
const equalization_rates = "EqualizationRates"

type Config struct {
	API struct {
		RateLevy          string `yaml:"levy_data"`
		EqualizationRates string `yaml:"equalization_data"`
	} `yaml:"api"`
}

func GetApiUrl(api string) string {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	switch api {
	case rate_levy:
		return config.API.RateLevy
	default:
		return config.API.EqualizationRates
	}

}
