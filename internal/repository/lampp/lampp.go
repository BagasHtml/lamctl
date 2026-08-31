package lampp

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type LamppRepository struct {
	path string
}

func New(path string) *LamppRepository {
	return &LamppRepository{path: path}
}

func (r *LamppRepository) Start() error {
	return r.run("start")
}

func (r *LamppRepository) Stop() error {
	return r.run("stop")
}

func (r *LamppRepository) Restart() error {
	return r.run("restart")
}

func (r *LamppRepository) Status() (string, error) {
	output, err := exec.Command(r.path, "status").Output()
	if err != nil {
		return string(output), fmt.Errorf("gagal cek status: %w", err)
	}
	return string(output), nil
}

func (r *LamppRepository) run(action string) error {
	if err := exec.Command(r.path, action).Run(); err != nil {
		return fmt.Errorf("gagal %s: %w", action, err)
	}
	return nil
}

func ParseServiceStatus(output, service string) (string, error) {
	for _, line := range strings.Split(output, "\n") {
		if strings.Contains(strings.ToLower(line), strings.ToLower(service)) {
			return strings.TrimSpace(line), nil
		}
	}
	return "", errors.New("service tidak ditemukan")
}
