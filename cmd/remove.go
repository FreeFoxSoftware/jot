package cmd

import (
	"fmt"
	"jot/internal/store"
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "rm [id]",
	Short: "Remove a note ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID %q — must be a number", args[0])
		}

		err = store.Delete(id)
		if err != nil {
			return err
		}

		return nil
	},
}
