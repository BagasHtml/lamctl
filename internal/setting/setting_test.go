package setting

import (
	"os"
	"path/filepath"
	"testing"

	"lamctl/internal/entity"
)

func tempEnvPath(t *testing.T) string {
	t.Helper()
	return filepath.Join(t.TempDir(), EnvFile)
}

func TestLoadReadsEnvWithDefaults(t *testing.T) {
	// Empty keys in the environment must fall back to defaults.
	t.Setenv("LAMCTL_DB_HOST", "")
	t.Setenv("LAMCTL_DB_PORT", "")
	t.Setenv("LAMCTL_DB_USER", "")

	r := New(tempEnvPath(t))
	if err := r.Load(); err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	cred := r.GetCredential()
	if got, want := cred.Host, "localhost"; got != want {
		t.Errorf("Host = %q, want %q", got, want)
	}
	if got, want := cred.Port, "3306"; got != want {
		t.Errorf("Port = %q, want %q", got, want)
	}
	if got, want := cred.User, "root"; got != want {
		t.Errorf("User = %q, want %q", got, want)
	}
}

func TestLoadReadsEnvValues(t *testing.T) {
	t.Setenv("LAMCTL_DB_HOST", "db.internal")
	t.Setenv("LAMCTL_DB_PORT", "5433")
	t.Setenv("LAMCTL_DB_USER", "app")
	t.Setenv("LAMCTL_DB_PASS", "secret")
	t.Setenv("LAMCTL_DB_NAME", "orders")
	t.Setenv("LAMCTL_DB_ENGINE", "postgres")

	r := New(tempEnvPath(t))
	if err := r.Load(); err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	got := *r.GetCredential()
	want := entity.Credential{
		Host: "db.internal", Port: "5433", User: "app",
		Password: "secret", DBName: "orders", DBEngine: "postgres",
	}
	if got != want {
		t.Errorf("Load() = %+v, want %+v", got, want)
	}
}

func TestApplyFlagsOnlyOverridesProvidedValues(t *testing.T) {
	t.Setenv("LAMCTL_DB_HOST", "")
	t.Setenv("LAMCTL_DB_PORT", "")
	t.Setenv("LAMCTL_DB_USER", "")

	r := New(tempEnvPath(t))
	if err := r.Load(); err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	r.ApplyFlags(entity.Credential{
		Host: "192.168.1.10",
		Port: "3307",
	})

	cred := r.GetCredential()
	if got, want := cred.Host, "192.168.1.10"; got != want {
		t.Errorf("Host = %q, want %q", got, want)
	}
	if got, want := cred.Port, "3307"; got != want {
		t.Errorf("Port = %q, want %q", got, want)
	}
	// Empty fields in the flag struct must not clobber loaded defaults.
	if got, want := cred.User, "root"; got != want {
		t.Errorf("User = %q, want %q (should be untouched)", got, want)
	}
}

func TestSaveUpdatesInMemory(t *testing.T) {
	r := New(tempEnvPath(t))

	in := entity.Credential{
		Host: "h", Port: "3306", User: "u", DBName: "d", DBEngine: "mysql",
	}
	if err := r.Save(in); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	// In-memory state must reflect what was just saved.
	got := *r.GetCredential()
	if got != in {
		t.Errorf("Save() did not update in-memory state: got %+v, want %+v", got, in)
	}
}

func TestSavePersistsFile(t *testing.T) {
	path := tempEnvPath(t)
	r := New(path)

	in := entity.Credential{
		Host: "h", Port: "5432", User: "u",
		Password: "pw", DBName: "app", DBEngine: "postgres",
	}
	if err := r.Save(in); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading saved file: %v", err)
	}
	want := "LAMCTL_DB_HOST=h\nLAMCTL_DB_PORT=5432\nLAMCTL_DB_USER=u\nLAMCTL_DB_PASS=pw\nLAMCTL_DB_NAME=app\nLAMCTL_DB_ENGINE=postgres\n"
	if string(data) != want {
		t.Errorf("saved file content:\n got %q\nwant %q", string(data), want)
	}
}
