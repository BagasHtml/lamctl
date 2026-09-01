package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"lamctl/internal/entity"
	"lamctl/internal/usecase/lamppctrl"
)

func newLamppCommand(action, description string) *cobra.Command {
	return &cobra.Command{
		Use:   action,
		Short: description,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			uc := lamppctrl.NewUseCase(entity.LamppConfig{Path: settings.GetLamppPath()})

			var err error
			switch action {
			case "start":
				err = uc.Start()
			case "stop":
				err = uc.Stop()
			case "restart":
				err = uc.Restart()
			}

			if err != nil {
				return err
			}

			fmt.Printf("XAMPP %s berhasil\n", action)
			return nil
		},
	}
}

var startCmd = newLamppCommand("start", "Start semua layanan XAMPP")
var stopCmd = newLamppCommand("stop", "Stop semua layanan XAMPP")
var restartCmd = newLamppCommand("restart", "Restart semua layanan XAMPP")

func init() {
	RootCmd.AddCommand(startCmd, stopCmd, restartCmd)
}
