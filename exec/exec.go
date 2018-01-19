package exec

import (
	"os/exec"
)


func Exec(command string){
	exec.Command("ls" +
		"-l")
}
