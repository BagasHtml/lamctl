package dbmanager

import (
	"stackctl/internal/entity"
	"stackctl/internal/repository/mysqldb"
)

type UseCase struct {
	cred *entity.Credential
}

func NewUseCase(cred *entity.Credential) *UseCase {
	return &UseCase{cred: cred}
}

func (u *UseCase) ListDatabases() ([]entity.Database, error) {
	repo, err := mysqldb.Connect(u.cred)
	if err != nil {
		return nil, err
	}
	defer repo.Close()

	return repo.ListDatabases()
}

func (u *UseCase) Query(sql string) ([][]string, error) {
	repo, err := mysqldb.Connect(u.cred)
	if err != nil {
		return nil, err
	}
	defer repo.Close()

	return repo.QueryRows(sql)
}

func (u *UseCase) Create(name string) error {
	repo, err := mysqldb.Connect(u.cred)
	if err != nil {
		return err
	}
	defer repo.Close()

	return repo.CreateDatabase(name)
}

func (u *UseCase) Drop(name string) error {
	repo, err := mysqldb.Connect(u.cred)
	if err != nil {
		return err
	}
	defer repo.Close()

	return repo.DropDatabase(name)
}

func (u *UseCase) Shell(clientPath string) error {
	return mysqldb.Shell(u.cred, clientPath)
}
