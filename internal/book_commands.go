package internal

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

// ListBooksCmd list all available books
func (env *Env) ListBooksCmd() *cli.Command {
	return &cli.Command{
		Name:        "ls",
		ArgsUsage:   "",
		Description: "Lists all books",
		Usage:       "Lists books",
		Action: func(c *cli.Context) error {
			var books []Book
			env.DB.Find(&books)
			if len(books) == 0 {
				fmt.Println("It seems you have no books yet. Try to parse clippings from your Kindle.")
				return nil
			}

			templates := &promptui.SelectTemplates{
				Inactive: "{{ .ID }} | {{ .Name }}",
				Active:   "{{ .ID | cyan }} | {{ .Name | cyan }}",
				Selected: "{{ .Name | bold }}",
			}

			prompt := promptui.Select{
				Label:     "Your books",
				Items:     books,
				Templates: templates,
				Size:      8,
			}

			idx, _, err := prompt.Run()
			if err != nil {
				return err
			}
			env.showNotes(books[idx].ID, true)
			return nil
		},
	}
}

// DeleteBookCmd deletes selected book by ID
func (env *Env) DeleteBookCmd() *cli.Command {
	return &cli.Command{
		Name:        "rm",
		Description: "Deletes book by ID",
		ArgsUsage:   "BOOK_ID",
		Usage:       "Deletes a book",
		Action: func(c *cli.Context) error {
			bookID, err := strToUint(c.Args().First())
			if err != nil {
				fmt.Println("ID of a book have to be a integer")
				return err
			}
			env.removeBook(bookID)
			return nil
		},
	}
}