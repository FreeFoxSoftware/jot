package cmd

import (
	"fmt"
	"jot/internal/store"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read [id]",
	Short: "Read a note by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID %q — must be a number", args[0])
		}

		note, err := store.Read(id)
		if err != nil {
			return err
		}

		printNote(note)
		return nil
	},
}

func printNote(note store.Note) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	_, err := fmt.Fprintln(w, "TITLE\tDESCRIPTION")
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(w, "-----\t-------")
	if err != nil {
		return
	}

	_, err = fmt.Fprintf(w, "%s\t%s\n", note.Title, note.Description)
	if err != nil {
		return
	}

	err = w.Flush()
	if err != nil {
		return
	}
}
