package codespace

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type editOptions struct {
	codespaceName string
	displayName   string
	idleTimeout   time.Duration
	machine       string
}

func newEditCmd(app *App) *cobra.Command {
	opts := editOptions{}

	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a codespace",
		Args:  noArgsConstraint,
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Edit(cmd.Context(), opts)
		},
	}

	editCmd.Flags().StringVarP(&opts.codespaceName, "codespace", "c", "", "Name of the codespace")
	editCmd.Flags().StringVarP(&opts.displayName, "displayName", "d", "", "display name")
	editCmd.Flags().DurationVar(&opts.idleTimeout, "idle-timeout", 0, "allowed inactivity before codespace is stopped, e.g. \"10m\", \"1h\"")
	editCmd.Flags().StringVarP(&opts.machine, "machine", "m", "", "hardware specifications for the VM")

	return editCmd
}

// Edits a codespace
func (a *App) Edit(ctx context.Context, opts editOptions) error {
	fmt.Println("HEY PIZZA")
	return nil
}
