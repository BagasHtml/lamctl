package mysqldb

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"stackctl/internal/entity"
)

func FindClientPath(xamppPath string) string {
	bundled := filepath.Join(filepath.Dir(xamppPath), "bin", "mysql")
	if _, err := os.Stat(bundled); err == nil {
		return bundled
	}
	if p, err := exec.LookPath("mysql"); err == nil {
		return p
	}
	return bundled
}

func Shell(cred *entity.Credential, clientPath string) error {
	args := []string{"-h", cred.Host, "-P", cred.Port, "-u", cred.User}
	if cred.DBName != "" {
		args = append(args, cred.DBName)
	}

	cmd := exec.Command(clientPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "MYSQL_PWD="+cred.Password)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("gagal menjalankan mysql client: %w", err)
	}
	return nil
}