package entity

type Credential struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	DBEngine string
}

type Database struct {
	Name string
}

type LamppConfig struct {
	Path string
}
