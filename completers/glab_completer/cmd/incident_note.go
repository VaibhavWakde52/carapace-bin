package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var incident_noteCmd = &cobra.Command{
	Use:     "note <incident-id>",
	Short:   "Comment on an incident in GitLab",
	Aliases: []string{"comment"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_noteCmd).Standalone()

	incident_noteCmd.Flags().StringP("message", "m", "", "Comment/Note message")
	incidentCmd.AddCommand(incident_noteCmd)

	// TODO positional completion
}
