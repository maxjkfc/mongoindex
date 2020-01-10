package env

import (
	"fmt"

	"github.com/spf13/viper"
)

//ReadConfig -
func ReadConfig() {
	fmt.Println("read config")
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	} else {
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.SetConfigName("index")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	Body = &IndexYaml{}

	if err := viper.Unmarshal(Body); err != nil {
		panic(err)
	}

	fmt.Println("check finished")
}
