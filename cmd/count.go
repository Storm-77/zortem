package cmd

import (
	"fmt"
	"zortem/core"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count how many files there are with at least one duplicate",
	Run:   CountDuplicates,
}

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().BoolP("json", "j", false, "output as json")
}

func CountDuplicates(cmd *cobra.Command, args []string) {

	jsonFlag, _ := cmd.Flags().GetBool("json")

	maps := core.GetMapOfDuplicates(workDir)
	count := len(maps)

	if jsonFlag {
		fmt.Printf("{\"count\":%v}\n", count)
	} else {
		fmt.Println(count)
	}

}
