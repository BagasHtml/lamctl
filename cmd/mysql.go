package cmd

import (
	"github.com/spf13/cobra"

	"lamctl/internal/repository/mysqldb"
	"lamctl/internal/usecase/dbmanager"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Buka mysql client interaktif",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		uc := dbmanager.NewUseCase(settings.GetCredential())
		clientPath := mysqldb.FindClientPath(settings.GetLamppPath())
		return uc.Shell(clientPath)
	},
}

func init() {
	RootCmd.AddCommand(mysqlCmd)
}
