package main

import (
	"fmt"
	"os"

	"github.com/jayunit100/blackduck-ctl/pkg/interactive2"
	"github.com/spf13/viper"
)

type Blackduckctl struct {
	LaunchUI bool
}

func main() {
	viper.ReadInConfig()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	config := &Blackduckctl{}
	viper.Unmarshal(config)

	if config.LaunchUI {
		interactive2.Launch()
	}

}
