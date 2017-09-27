package other

import (
	"fmt"
	// "os"
	"os/exec"
)

func Py(str string) []byte {
	cmd := exec.Command("python", "-c", str)
	// cmd.Stderr = os.Stderr
	ret, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}

func PyFile(filename string, arg ...string) []byte {
	args := []string{filename}
	args = append(args, arg...)
	cmd := exec.Command("python", args...)
	// cmd.Stderr = os.Stderr
	ret, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}
