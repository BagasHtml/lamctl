package mysqldb

import "testing"

func TestValidateDBName(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantErr bool
	}{
		{name: "valid simple", in: "mydb"},
		{name: "valid with digits", in: "db2"},
		{name: "valid with underscore", in: "my_db_1"},
		{name: "valid uppercase", in: "MyDB"},
		{name: "empty", in: "", wantErr: true},
		{name: "trailing newline", in: "mydb\n", wantErr: true},
		{name: "semicolon injection", in: "mydb; DROP DATABASE x", wantErr: true},
		{name: "space", in: "my db", wantErr: true},
		{name: "hyphen", in: "my-db", wantErr: true},
		{name: "backtick", in: "my`db", wantErr: true},
		{name: "quote", in: "my'db", wantErr: true},
		{name: "comment marker", in: "db--x", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDBName(tt.in)
			if tt.wantErr && err == nil {
				t.Fatalf("validateDBName(%q) = nil, want error", tt.in)
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("validateDBName(%q) = %v, want nil", tt.in, err)
			}
		})
	}
}
