package config

import "github.com/urfave/cli/v3"

var settingsPath string = "./settings"

func LoadConfigFromCliFlags(cmd *cli.Command) {
	// SettingsPath
	if cmd.String("settings-path") != "" {
		settingsPath = cmd.String("settings-path")
	}
}

func GetSettingsPath() string {
	return settingsPath
}
