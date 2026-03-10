package cmd

import (
	"jot/internal/presentation"
	"jot/internal/store"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists notes",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		notes, err := store.List(isGlobal, all)
		if err != nil {
			return err
		}

		presentation.PrintNotesList(notes)
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVarP(&isGlobal, "global", "g", false, "Whether to get notes from working directory or global")
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "List all notes")
}
