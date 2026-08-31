package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"stackctl/internal/entity"
	"stackctl/internal/usecase/lamppctrl"
)

var statusCmd = &cobra.Command{
	Use:   "status [service]",
	Short: "Cek status layanan XAMPP",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uc := lamppctrl.NewUseCase(entity.LamppConfig{Path: settings.GetLamppPath()})

		if len(args) == 0 {
			status, err := uc.Status()
			if err != nil {
				return err
			}
			fmt.Print(status)
			return nil
		}

		service, err := uc.ServiceStatus(args[0])
		if err != nil {
			return err
		}

		fmt.Println(service)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
