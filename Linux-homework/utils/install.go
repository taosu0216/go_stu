package utils

import (
	"fmt"
	"os/exec"
)

func Install() {
	var commands []*exec.Cmd
	var err error
	cmd1 := exec.Command("sh", "-c", "apt install cpuinfo -y")
	cmd2 := exec.Command("sh", "-c", "apt install iotop -y")
	commands = append(commands, cmd1, cmd2)
	for i, cmd := range commands {
		err = cmd.Run()
		if err != nil {
			fmt.Println("指令", i, "错误,错误为: ", err)
			err = nil
		}
	}
}
