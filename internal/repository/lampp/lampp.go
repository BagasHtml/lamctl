package lampp

import (
	"errors"
	"fmt"
	"os"
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
	out, err := exec.Command(r.path, "status").CombinedOutput()	

	if err != nil {
		return string(out), fmt.Errorf("gagal cek status: %w\n%s", err, out)
	}
	
	return string(out), nil
}

func (r *LamppRepository) command(action string) *exec.Cmd {
	if os.Geteuid() == 0 {
		return exec.Command(r.path, action)
	}

	cmd := exec.Command("sudo", r.path, action)
	cmd.Stdin = os.Stdin
	
	return cmd
}

func (r *LamppRepository) run(action string) error {
	out, err := r.command(action).CombinedOutput()

	if err != nil {
		return fmt.Errorf("gagal %s: %w\n%s", action, err, out)
	}
	fmt.Print(string(out))

	return nil
}

func ParseServiceStatus(out, service string) (string, error) {
	for _, line := range strings.Split(out, "\n") {
		if strings.Contains(strings.ToLower(line), strings.ToLower(service)) {
			return strings.TrimSpace(line), nil
		}
	}
	return "", errors.New("service tidak ditemukan")
}