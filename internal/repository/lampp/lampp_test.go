package lampp

import "testing"

func TestParseServiceStatus(t *testing.T) {
	output := `Starting XAMPP for Linux 8.2.12...
XAMPP:  Starting Apache with SSL...
XAMPP:  Starting MySQL...
XAMPP:  Starting ProFTPD...`

	tests := []struct {
		name    string
		out     string
		service string
		want    string
		wantErr bool
	}{
		{
			name:    "found service",
			out:     output,
			service: "MySQL",
			want:    "XAMPP:  Starting MySQL...",
		},
		{
			name:    "case insensitive match",
			out:     output,
			service: "mysql",
			want:    "XAMPP:  Starting MySQL...",
		},
		{
			name:    "service not present",
			out:     output,
			service: "PostgreSQL",
			wantErr: true,
		},
		{
			name:    "empty output",
			out:     "",
			service: "MySQL",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseServiceStatus(tt.out, tt.service)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("ParseServiceStatus() = %q, want %q", got, tt.want)
			}
		})
	}
}
