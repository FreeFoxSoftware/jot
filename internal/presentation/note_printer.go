package presentation

import (
	"fmt"
	"jot/internal/store"
	"os"
	"text/tabwriter"
)

func PrintNotesList(notes []store.Note) {
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

func PrintContents(notes ...store.Note) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	_, _ = fmt.Fprintln(w, "ID\tTITLE\tDESCRIPTION")
	_, _ = fmt.Fprintln(w, "--\t-----\t-----------")

	for _, note := range notes {
		_, _ = fmt.Fprintf(w, "%d\t%s\t%s\n", note.ID, note.Title, note.Description)
	}

	err := w.Flush()
	if err != nil {
		return
	}
}
