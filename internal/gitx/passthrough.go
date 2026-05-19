package gitx

import (
	"os"
	"os/exec"
)

// Passthrough runs the real git binary with the given arguments,
// inheriting stdin/stdout/stderr so output streams directly to the user.
// Returns git's exit code, or an error if git could not be started.
func Passthrough(args []string) (int, error) {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		return 1, err
	}

	cmd := exec.Command(gitPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		if cmd.ProcessState != nil {
			return cmd.ProcessState.ExitCode(), nil
		}
		return 1, err
	}
	return 0, nil
}
