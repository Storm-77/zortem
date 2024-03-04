/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
    "path/filepath"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "zortem",
	Short: "Simple utility to find and delete file duplicates",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var workDir string;
func init() {

    wd, _ := os.Getwd()
    rootCmd.PersistentFlags().StringVarP(&workDir, "path","p",wd , "path to dir to work on")
	workDir = filepath.Clean(wd)

}


