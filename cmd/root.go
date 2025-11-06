package cmd

import (
	"fmt"
	"os"

	"github.com/marie-mattheo/crm-project/internal/storage"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gocrm",
		Short: "Gocrm est un outil pour gérer les contacts",
		Long:  "Un outil CLI en Go pour gérer les contacts dans un CRM",
	}

	store storage.Storer
)

func Execute(s storage.Storer) {
	store = s
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func init() {

}
