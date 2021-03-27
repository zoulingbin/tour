package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zoulingbin/tour/internal/word"
	"log"
	"strings"
)

const (
	MODE_UPPER	= iota + 1
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERSCORE
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
},"\n")

var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: "支持多种单词格式转换",
	Run: run,
}

var mode int8
var str string

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}

func run(cmd *cobra.Command, args []string) {
	content := ""
	switch mode {
	case MODE_UPPER:
		content = word.ToUpper(str)
	case MODE_LOWER:
		content = word.ToLower(str)
	case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
		content = word.UnderscoreToUpperCamelCase(str)
	case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
		content = word.UnderscoreToLowerCamelCase(str)
	case 	MODE_CAMELCASE_TO_UNDERSCORE:
		content = word.CamelCaseToUnderscore(str)
	default:
		log.Fatalf("暂不支持改转换模式")
	}
	log.Printf("result: %s",content)
}