package cmd

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type yamlContent struct {
	content string
	prefix  string
	space   string
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the site configuration",
	Long:  `Print the site configuration, both default and custom settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		allSettings := viper.AllSettings()

		content := readYamlContentRecursive(allSettings, 0)

		for _, v := range content {
			fmt.Printf("%s%s%s\n", v.space, v.prefix, v.content)
		}
	},
}

func readYamlContentRecursive(settings map[string]interface{}, hierarchy int) []*yamlContent {
	var content []*yamlContent
	var keys []string
	space := ""

	for i := 0; i < hierarchy; i++ {
		space += "  "
	}

	for k := range settings {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	hierarchy += 1
	for _, k := range keys {
		kv := reflect.ValueOf(settings[k])
		kind := kv.Kind()
		if kind == reflect.Slice {
			c := &yamlContent{
				content: fmt.Sprintf("%s:", k),
				prefix:  "",
				space:   space,
			}

			content = append(content, c)
			for _, val := range settings[k].([]interface{}) {
				c := &yamlContent{
					content: fmt.Sprintf("%v", val),
					prefix:  "- ",
					space:   space,
				}

				content = append(content, c)
			}
		} else if kind == reflect.Map {
			c := &yamlContent{
				content: fmt.Sprintf("%s:", k),
				prefix:  "",
				space:   space,
			}

			content = append(content, c)

			returnedContent := readYamlContentRecursive(settings[k].(map[string]interface{}), hierarchy)
			content = append(content, returnedContent...)
		} else {
			c := &yamlContent{
				content: fmt.Sprintf("%s: \"%+v\"", k, settings[k]),
				prefix:  "",
				space:   space,
			}

			content = append(content, c)
		}
	}
	return content
}

func init() {
	rootCmd.AddCommand(configCmd)
}
