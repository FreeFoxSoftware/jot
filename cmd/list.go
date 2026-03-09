package cmd

import (
	"fmt"
	"jot/internal/store"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists notes",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		notes, err := store.List(isGlobal)
		if err != nil {
			return err
		}

		printNotes(notes)
		return nil
	},
}

func printNotes(notes []store.Note) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	_, err := fmt.Fprintln(w, "ID\tTITLE\tCREATED")
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(w, "--\t-----\t-------")
	if err != nil {
		return
	}
	for _, n := range notes {
		_, err := fmt.Fprintf(w, "%d\t%s\t%s\n", n.ID, n.Title, n.CreatedAt.Format("2006-01-02"))
		if err != nil {
			return
		}
	}
	err = w.Flush()
	if err != nil {
		return
	}
}

func init() {
	listCmd.Flags().BoolVarP(&isGlobal, "global", "g", false, "Whether to get notes from working directory or global")
}
