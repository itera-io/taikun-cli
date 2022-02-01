package docs

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/spf13/cobra"
)

func PrintCommandTree(cmd *cobra.Command) {
	tree := list.NewWriter()
	tree.SetStyle(list.StyleMarkdown)

	buildTree(tree, cmd)

	fmt.Println(tree.Render())
}

func buildTree(tree list.Writer, cmd *cobra.Command) {
	cmdDescription := cmd.Name() + " (" + cmd.Short + ")"
	tree.AppendItem(cmdDescription)
	tree.Indent()
	for _, childCmd := range cmd.Commands() {
		buildTree(tree, childCmd)
	}
	tree.UnIndent()
}
