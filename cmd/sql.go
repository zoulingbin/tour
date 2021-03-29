package cmd

import "github.com/spf13/cobra"

var username 	string
var password 	string
var host 		string
var charset 	string
var dbType, dbName, tableName string

var sqlCmd = &cobra.Command{
	Use: "sql",
	Short: "sql转换和处理",
	Long: "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {},
}