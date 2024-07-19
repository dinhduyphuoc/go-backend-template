package initialize

import (
	"fmt"

	"github.com/dinhduyphuoc/go-backend-template/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	// Config directory
	viper.AddConfigPath("config")
	// Config file name
	viper.SetConfigName("local")
	// Config file type
	viper.SetConfigType("yaml")

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Unmarshal config failed: ", err)
	}

	fmt.Println("Config loaded successfully.")
}
