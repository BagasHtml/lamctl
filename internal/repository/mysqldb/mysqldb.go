package mysqldb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"stackctl/internal/entity"
)

type MySQLRepository struct {
	db *sql.DB
}

func Connect(cred *entity.Credential) (*MySQLRepository, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cred.User, cred.Password, cred.Host, cred.Port, cred.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke MySQL: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping MySQL: %w", err)
	}

	return &MySQLRepository{db: db}, nil
}

func (r *MySQLRepository) ListDatabases() ([]entity.Database, error) {
	rows, err := r.db.Query("SHOW DATABASES")
	if err != nil {
		return nil, fmt.Errorf("gagal list database: %w", err)
	}
	defer rows.Close()

	var databases []entity.Database
	for rows.Next() {
		var db entity.Database
		if err := rows.Scan(&db.Name); err != nil {
			return nil, err
		}
		databases = append(databases, db)
	}

	return databases, rows.Err()
}

func (r *MySQLRepository) QueryRows(query string) ([][]string, error) {
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal menjalankan query: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]any, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var result [][]string
	result = append(result, columns)

	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		row := make([]string, len(columns))
		for i, v := range values {
			row[i] = string(v)
		}
		result = append(result, row)
	}

	return result, rows.Err()
}

func (r *MySQLRepository) CreateDatabase(name string) error {
	if _, err := r.db.Exec("CREATE DATABASE IF NOT EXISTS " + name); err != nil {
		return fmt.Errorf("gagal membuat database: %w", err)
	}
	return nil
}

func (r *MySQLRepository) DropDatabase(name string) error {
	if _, err := r.db.Exec("DROP DATABASE IF EXISTS " + name); err != nil {
		return fmt.Errorf("gagal menghapus database: %w", err)
	}
	return nil
}

func (r *MySQLRepository) Close() error {
	return r.db.Close()
}
