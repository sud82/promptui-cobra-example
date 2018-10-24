package interactive2

import (
	"github.com/manifoldco/promptui"
)

func Launch() {
	prompt := promptui.Select{
		Label: "Select operation",
		Items: []string{
			"opssight",
			"blackduck",
			"api",
		},
	}
	_, result, _ := prompt.Run()
	if result == "cluster" {
		prompt := promptui.Select{
			Label: "Select operation",
			Items: []string{
				"status",
				"hubs",
			},
		}
		prompt.Run()

	}
	if result == "api" {
		prompt := promptui.Select{
			Label: "Select operation",
			Items: []string{
				"all-components",
				"all-projects",
				"list-hubs",
			},
		}
		prompt.Run()
	}
	if result == "opssight" {
		prompt := promptui.Select{
			Label: "Select operation",
			Items: []string{
				"status",
				"throughput",
				"ns-vuln",
				"deploy",
			},
		}
		prompt.Run()
	}
	if result == "blackduck" {
		prompt := promptui.Select{
			Label: "Select operation",
			Items: []string{
				"status",
				"deploy",
			},
		}
		prompt.Run()
	}
}
