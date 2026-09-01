package setting

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"lamctl/internal/entity"
)

const EnvFile = ".env"

type SettingRepository struct {
	credential entity.Credential
	path       string
}

func New(path string) *SettingRepository {
	return &SettingRepository{path: path}
}

func (r *SettingRepository) Load() error {
	if _, err := os.Stat(r.path); err == nil {
		if err := godotenv.Load(r.path); err != nil {
			return fmt.Errorf("gagal memuat file .env: %w", err)
		}
	}

	r.credential = entity.Credential{
		Host:     envOrDefault("LAMCTL_DB_HOST", "localhost"),
		Port:     envOrDefault("LAMCTL_DB_PORT", "3306"),
		User:     envOrDefault("LAMCTL_DB_USER", "root"),
		Password: os.Getenv("LAMCTL_DB_PASS"),
		DBName:   os.Getenv("LAMCTL_DB_NAME"),
		DBEngine: os.Getenv("LAMCTL_DB_ENGINE"),
	}

	return nil
}

func (r *SettingRepository) GetCredential() *entity.Credential {
	return &r.credential
}

func (r *SettingRepository) GetLamppPath() string {
	return envOrDefault("LAMCTL_XAMPP_PATH", "/opt/lampp/lampp")
}

func (r *SettingRepository) ApplyFlags(cred entity.Credential) {
	if cred.Host != "" {
		r.credential.Host = cred.Host
	}
	if cred.Port != "" {
		r.credential.Port = cred.Port
	}
	if cred.User != "" {
		r.credential.User = cred.User
	}
	if cred.Password != "" {
		r.credential.Password = cred.Password
	}
	if cred.DBName != "" {
		r.credential.DBName = cred.DBName
	}
	if cred.DBEngine != "" {
		r.credential.DBEngine = cred.DBEngine
	}
}

func (r *SettingRepository) Save(cred entity.Credential) error {
	content := fmt.Sprintf(
		"LAMCTL_DB_HOST=%s\nLAMCTL_DB_PORT=%s\nLAMCTL_DB_USER=%s\nLAMCTL_DB_PASS=%s\nLAMCTL_DB_NAME=%s\nLAMCTL_DB_ENGINE=%s\n",
		cred.Host, cred.Port, cred.User, cred.Password, cred.DBName, cred.DBEngine,
	)

	if err := os.WriteFile(r.path, []byte(content), 0o600); err != nil {
		return fmt.Errorf("gagal menyimpan %s: %w", r.path, err)
	}

	r.credential = cred
	return nil
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
