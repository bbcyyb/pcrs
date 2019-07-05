package cmd

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ConfTestSuite struct {
	suite.Suite

	Configs map[string]interface{}
}

func TestConfSuite(t *testing.T) {
	suite.Run(t, new(ConfTestSuite))
}

func (suite *ConfTestSuite) SetupTest() {
	suite.Configs = testGenerateAllConfigSettings()
}

func (suite *ConfTestSuite) TestReadYamlContent() {
	ass := suite.Assert()

	formattedContent := readYamlContent(suite.Configs)

	for _, v := range formattedContent {
		fmt.Println(v)
	}

	ass.Equal("age: 35", formattedContent[0])
	ass.Equal("bbcyyb:", formattedContent[1])
	ass.Equal("  aaa:", formattedContent[2])
	ass.Equal("    xxx: 22.98", formattedContent[3])
	ass.Equal("  ccc:", formattedContent[7])
	ass.Equal("    ddd:", formattedContent[8])
	ass.Equal("    - 111", formattedContent[9])
	ass.Equal("    - 222", formattedContent[10])
	ass.Equal("    - 333", formattedContent[11])
	ass.Equal("  jacket: leather", formattedContent[13])
	ass.Equal("isBoolean: true", formattedContent[21])

	slice := strings.Split(formattedContent[22], ": ")
	ass.Equal(2, len(slice))
	createOn, err := time.Parse(time.RFC3339, slice[1])
	if ass.Nil(err) {
		ass.Equal(2020, createOn.Year())
		ass.Equal(12, (int)(createOn.Month()))
		ass.Equal(31, createOn.Day())
		ass.Equal(23, createOn.Hour())
		ass.Equal(59, createOn.Minute())
		ass.Equal(59, createOn.Second())
	}
}

func testGenerateAllConfigSettings() map[string]interface{} {
	root := make(map[string]interface{})
	root["hacker"] = true
	root["name"] = "steve"
	root["age"] = 35
	root["createOn"] = time.Now().Format(time.RFC3339)
	root["iso8601"] = time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local).Format(time.RFC3339)
	root["isBoolean"] = true

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
