package options

import "github.com/spf13/cobra"

type TableWriter interface {
	GetAllColumnsOption() *bool
	GetColumnsOption() *[]string
	GetNoDecorateOption() *bool
	GetShowLargeValuesOption() *bool
}

func AddTableWriterFlags(cmd *cobra.Command, opts TableWriter) {
	cmd.Flags().BoolVar(
		opts.GetNoDecorateOption(),
		"no-decorate",
		false,
		"Display output table without field names and separators",
	)

	cmd.PersistentFlags().BoolVar(
		opts.GetShowLargeValuesOption(),
		"show-large-values",
		false,
		"Prevent trimming of large cell values",
	)
}
