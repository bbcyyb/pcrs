package cmd

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadYamlContentRecursive(t *testing.T) {
	ass := assert.New(t)
	allSettings := generateAllConfigSettings()
	formattedContent := readYamlContent(allSettings)

	for _, v := range formattedContent {
		fmt.Println(v)
	}
	ass.Equal(1, 1)

	//TODO: consider use suite component to define the Setup and TearDown functions.

}

func generateAllConfigSettings() map[string]interface{} {
	root := make(map[string]interface{})
	root["hacker"] = true
	root["Name"] = "steve"
	root["age"] = 35
	root["createOn"] = time.Now()
	root["iso8601"] = time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local)

	bbcyyb := make(map[string]interface{})
	aaa := make(map[string]interface{})
	aaa["xxx"] = 22.98
	bbb := []string{"mm", "nn"}
	ccc := make(map[string]interface{})
	ddd := []int{111, 222, 333}
	ccc["ddd"] = ddd
	bbcyyb["aaa"] = aaa
	bbcyyb["bbb"] = bbb
	bbcyyb["ccc"] = ccc

	hobbies := []string{"stkateboarding", "snowboarding", "go"}

	clothing := map[string]interface{}{
		"jacket":   "leather",
		"trousers": "denim",
	}

	root["bbcyyb"] = bbcyyb
	root["hobbies"] = hobbies
	root["chothing"] = clothing

	return root
}
