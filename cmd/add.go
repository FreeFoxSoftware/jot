package cmd

import (
	"fmt"
	"jot/internal/store"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new note",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := store.Save(title, description, isGlobal)
		if err != nil {
			return fmt.Errorf("could not save note: %w", err)
		}
		fmt.Println("saved note:", id)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the note (required)")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Body of the note (required)")
	addCmd.Flags().BoolVarP(&isGlobal, "global", "g", false, "Save note globally")

	_ = addCmd.MarkFlagRequired("title")
}
