package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"

	"stackctl/internal/entity"
	"stackctl/internal/setting"
)

const logo = `.__                         __  .__
|  | _____    _____   _____/  |_|  |
|  | \__  \  /     \_/ ___\   __\  |
|  |__/ __ \|  Y Y  \  \___|  | |  |__
|____(____  /__|_|  /\___  >__| |____/
          \/      \/     \/`

var settings *setting.SettingRepository

var RootCmd = &cobra.Command{
	Use:   "lamctl",
	Short: "Helper untuk XAMPP/LAMPP dan database MySQL",
	Long: `lamctl adalah CLI helper untuk menyalakan XAMPP/LAMPP
dan mengelola database MySQL dengan mudah.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("%s\n\n", logo)
		cmd.Help()
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		settings = setting.New(filepath.Join(".", setting.EnvFile))
		if err := settings.Load(); err != nil {
			return err
		}

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		dbName, _ := cmd.Flags().GetString("db")
		dbEngine, _ := cmd.Flags().GetString("dbEngine")

		settings.ApplyFlags(entity.Credential{
			Host: host, Port: port, User: user, Password: password, DBName: dbName, DBEngine: dbEngine,
		})
		return nil
	},
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().String("host", "", "Host database (default dari .env)")
	RootCmd.PersistentFlags().String("port", "", "Port database (default dari .env)")
	RootCmd.PersistentFlags().String("user", "", "Username database (default dari .env)")
	RootCmd.PersistentFlags().String("password", "", "Password database")
	RootCmd.PersistentFlags().String("db", "", "Nama database")
	RootCmd.PersistentFlags().String("db_engine", "", "Database yang kamu gunakan")
}
