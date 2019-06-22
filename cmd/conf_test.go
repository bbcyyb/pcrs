package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadYamlContentRecursive(t *testing.T) {
	ass := assert.New(t)
	allSettings := generateAllConfigSettings()
	readYamlContentRecursive(allSettings, 0)

	//TODO: need to have a mechanism to allow redirecting fmt print to buffer.
	// Reference: https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string

	//TODO: consider use suite component to define the Setup and TearDown functions.

}

func generateAllConfigSettings() map[string]interface{} {
	root := make(map[string]interface{})
	root["hacker"] = true
	root["Name"] = "steve"
	root["age"] = 35

	bbcyyb := make(map[string]interface{})
	aaa := make(map[string]interface{})
	aaa["xxx"] = 22.98
	bbb := []string{"mm", "nn"}
	ccc := make(map[string]interface{})
	ddd := []int{111, 222, 333}
	ccc["ddd"] = ccc
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
	root["chothing"] = chothing

	return root
}
