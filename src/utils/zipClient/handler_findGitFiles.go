package zipClient

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/zerops-io/zcli/src/utils/cmdRunner"
)

func (h *Handler) FindGitFiles(workingDir string) (res []File, _ error) {

	workingDir, err := filepath.Abs(workingDir)
	if err != nil {
		return nil, err
	}

	createFile := func(filePath string) File {
		return File{
			SourcePath:  path.Join(workingDir, filePath),
			ArchivePath: filePath,
		}
	}

	createCmd := func(name string, arg ...string) *exec.Cmd {
		cmd := exec.Command(name, arg...)
		cmd.Dir = workingDir
		return cmd
	}

	// find excluded files
	excludedFiles := make(map[string]struct{})
	if err := h.listFiles(
		createCmd("git", "ls-files", "--deleted", "--exclude-standard"),
		func(path string) error {
			excludedFiles[path] = struct{}{}

			return nil
		},
	); err != nil {
		return nil, err
	}

	if err := h.listFiles(
		createCmd("git", "ls-files", "--others", "--ignored", "--exclude-standard"),
		func(path string) error {
			excludedFiles[path] = struct{}{}

			return nil
		},
	); err != nil {
		return nil, err
	}

	// add all non deleted
	if err := h.listFiles(
		createCmd("git", "ls-files", "--exclude-standard"),
		func(path string) error {
			if _, exists := excludedFiles[path]; !exists {
				res = append(res, createFile(path))
			}
			return nil
		},
	); err != nil {
		return nil, err
	}

	// add untracked files
	if err := h.listFiles(
		createCmd("git", "ls-files", "--others", "--exclude-standard"),
		func(path string) error {
			if _, exists := excludedFiles[path]; !exists {
				res = append(res, createFile(path))
			}
			return nil
		},
	); err != nil {
		return nil, err
	}

	return
}

func (h *Handler) listFiles(cmd *exec.Cmd, fn func(path string) error) error {

	output, err := cmdRunner.Run(cmd)
	if err != nil {
		return err
	}

	rd := bufio.NewReader(bytes.NewReader(output))
	for {
		lineB, _, err := rd.ReadLine()
		line := string(lineB)

		if err == io.EOF {
			if line != "" {
				err := fn(line)
				if err != nil {
					return err
				}
			}
			break
		}
		if err != nil {
			return err
		}

		err = fn(line)
		if err != nil {
			return err
		}
	}

	return nil
}
