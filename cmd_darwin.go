package main

import (
	"fmt"
	"os/exec"
)

func volumeCommand(v int) *exec.Cmd {
	return exec.Command("/usr/bin/osascript", "-e", fmt.Sprintf("set volume output volume %d", v))
}

func ttyGlob() string {
	return "/dev/tty.usbmodem*"
}
