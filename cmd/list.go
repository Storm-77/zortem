package cmd

import (
	"encoding/json"
	"fmt"
	"zortem/core"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all paths that contain identical file brief description of your command",
	Run:   list,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("json", "j", false, "output as json")
}

func list(cmd *cobra.Command, args []string) {
	jsonFlag, _ := cmd.Flags().GetBool("json")
	hashPathsMap := core.GetMapOfDuplicates(workDir)

	if jsonFlag {
		concatenated := [][]string{}
		for _, paths := range hashPathsMap {
			concatenated = append(concatenated, paths)
		}
		jsonData, err := json.Marshal(concatenated)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println(string(jsonData))

	} else {

		for _, paths := range hashPathsMap {
			fmt.Println(paths)
		}
	}
}
