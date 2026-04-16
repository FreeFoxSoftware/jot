package cmd

import (
	"fmt"
	"jot/internal/store"

	"github.com/spf13/cobra"
)

var reindexCmd = &cobra.Command{
	Use:   "reindex",
	Short: "Reindex all notes sequentially from 1",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		count, err := store.Reindex()
		if err != nil {
			return err
		}
		fmt.Printf("Reindexed %d notes (1 through %d)\n", count, count)
		return nil
	},
}
