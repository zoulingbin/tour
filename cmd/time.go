package cmd

import "github.com/spf13/cobra"

var timeCmd = &cobra.Command{
	Use: "time",
	Short: "时间格式处理",
	Long: "时间格式处理",
	Run:timeRun,
}



var calculateTime string
var duration string



func timeRun(cmd *cobra.Command, args []string) {

}
