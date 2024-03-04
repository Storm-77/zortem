package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete duplicate files",
	Long:  `Delete all duplicate files by default the one that isn't deleted is the one with shorter path, but you can make a choice in interactive mode`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

}

func deleteFile(paths []string, toRemain string) {

	for _, file := range paths {
		if file == toRemain {
			continue
		}

		err := os.Remove(file)
		if err != nil {
            os.Stderr.WriteString(err.Error())
		}

	}
}
