package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"stackctl/internal/entity"
	"stackctl/internal/setting"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup kredensial database secara interaktif",
	RunE: func(cmd *cobra.Command, args []string) error {
		answers := struct {
			Host     string
			Port     string
			User     string
			Password string
			DBName   string
		}{}

		questions := []*survey.Question{
			{
				Name:     "host",
				Prompt:   &survey.Input{Message: "Database host", Default: "localhost"},
				Validate: survey.Required,
			},
			{
				Name:     "port",
				Prompt:   &survey.Input{Message: "Database port", Default: "3306"},
				Validate: survey.Required,
			},
			{
				Name:     "user",
				Prompt:   &survey.Input{Message: "Database username", Default: "root"},
				Validate: survey.Required,
			},
			{
				Name:   "password",
				Prompt: &survey.Password{Message: "Database password (kosongkan jika tidak ada)"},
			},
			{
				Name:   "dbname",
				Prompt: &survey.Input{Message: "Database name (kosongkan untuk koneksi umum)"},
			},
		}

		if err := survey.Ask(questions, &answers); err != nil {
			return err
		}

		cred := entity.Credential{
			Host: answers.Host, Port: answers.Port, User: answers.User,
			Password: answers.Password, DBName: answers.DBName,
		}

		if err := settings.Save(cred); err != nil {
			return err
		}

		abs, err := filepath.Abs(setting.EnvFile)
		if err != nil {
			abs = setting.EnvFile
		}

		_ = os.Setenv("LAMCTL_DB_HOST", cred.Host)
		_ = os.Setenv("LAMCTL_DB_PORT", cred.Port)
		_ = os.Setenv("LAMCTL_DB_USER", cred.User)
		_ = os.Setenv("LAMCTL_DB_PASS", cred.Password)
		_ = os.Setenv("LAMCTL_DB_NAME", cred.DBName)

		fmt.Printf("Kredensial tersimpan di %s\n", abs)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
