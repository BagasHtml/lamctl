package lamppctrl

import (
	"lamctl/internal/entity"
	"lamctl/internal/repository/lampp"
)

type UseCase struct {
	repo *lampp.LamppRepository
}

func NewUseCase(cfg entity.LamppConfig) *UseCase {
	return &UseCase{repo: lampp.New(cfg.Path)}
}

func (u *UseCase) Start() error {
	return u.repo.Start()
}

func (u *UseCase) Stop() error {
	return u.repo.Stop()
}

func (u *UseCase) Restart() error {
	return u.repo.Restart()
}

func (u *UseCase) Status() (string, error) {
	return u.repo.Status()
}

func (u *UseCase) ServiceStatus(service string) (string, error) {
	output, err := u.repo.Status()
	if err != nil {
		return "", err
	}
	return lampp.ParseServiceStatus(output, service)
}
