package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"lamctl/internal/usecase/dbmanager"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Operasi database MySQL",
}

var dbListCmd = &cobra.Command{
	Use:   "list",
	Short: "List semua database",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		uc := dbmanager.NewUseCase(settings.GetCredential())
		databases, err := uc.ListDatabases()
		if err != nil {
			return err
		}

		for _, db := range databases {
			fmt.Println(db.Name)
		}
		return nil
	},
}

var dbQueryCmd = &cobra.Command{
	Use:   "query [SQL]",
	Short: "Jalankan query SQL",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uc := dbmanager.NewUseCase(settings.GetCredential())
		rows, err := uc.Query(args[0])
		if err != nil {
			return err
		}

		for _, row := range rows {
			fmt.Println(strings.Join(row, " | "))
		}
		return nil
	},
}

var dbCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Buat database baru",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uc := dbmanager.NewUseCase(settings.GetCredential())
		if err := uc.Create(args[0]); err != nil {
			return err
		}

		fmt.Printf("Database %q berhasil dibuat\n", args[0])
		return nil
	},
}

var dbDropCmd = &cobra.Command{
	Use:   "drop [name]",
	Short: "Hapus database",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uc := dbmanager.NewUseCase(settings.GetCredential())
		if err := uc.Drop(args[0]); err != nil {
			return err
		}

		fmt.Printf("Database %q berhasil dihapus\n", args[0])
		return nil
	},
}

func init() {
	dbCmd.AddCommand(dbListCmd, dbQueryCmd, dbCreateCmd, dbDropCmd)
	RootCmd.AddCommand(dbCmd)
}
