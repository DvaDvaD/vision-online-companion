package config

import (
	"github.com/portierglobal/keyonthego-service/src/utils"
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in
	viper.AutomaticEnv()          // read environment variables that match

	// Set default value for MODE
	viper.SetDefault("MODE", "on-prem")
	viper.SetDefault("TMP_FOLDER", "tmp")

	// Read the environment variable
	mode := viper.GetString("MODE")
	if mode != "cloud" && mode != "on-prem" {
		log.Fatal().Msgf("Invalid MODE value: %s. Must be 'cloud' or 'on-prem'.", mode)
	}
	if mode == "cloud" {
		viper.SetDefault("TMP_FOLDER", "/tmp/requests")
	} else {
		viper.SetDefault("TMP_FOLDER", utils.GetTMPPath()+"\\tmp\\requests")
	}
}
