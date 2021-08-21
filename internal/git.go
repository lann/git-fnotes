package internal

import "os/exec"

func Git(args ...string) ([]byte, error) {
	return exec.Command("git", args...).Output()
}
