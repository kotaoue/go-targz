package targz

import (
	"fmt"
	"os/exec"

	"github.com/kotaoue/go-wd"
)

func Decompress(dir, name string) error {
	path := fullPath(dir)

	cmd := exec.Command(
		"tar",
		"-zxvf", fmt.Sprintf("%s/%s", path, name),
		"-C", path,
	)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func fullPath(dir string) string {
	current, err := wd.Get()
	if err != nil {
		panic(err)
	}
	return current + "/" + dir
}
