package cmd

import (
	"fmt"
	"jot/internal/presentation"
	"jot/internal/store"
	"strconv"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read [id]",
	Short: "Read a note by ID",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {

		if all {
			notes, err := store.List(false, true)
			if err != nil {
				return err
			}

			presentation.PrintContents(notes...)
			return nil
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID %q — must be a number", args[0])
		}

		note, err := store.Read(id)
		if err != nil {
			return err
		}

		presentation.PrintContents(note)
		return nil
	},
}

func init() {
	readCmd.Flags().BoolVarP(&all, "all", "a", false, "List all notes")
}
