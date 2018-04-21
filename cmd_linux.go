package main

import (
	"fmt"
	"os/exec"
)

func volumeCommand(v int) *exec.Cmd {
	return exec.Command("/usr/bin/pactl", "set-sink-volume", "1", fmt.Sprintf("%d%%", v))
}

func ttyGlob() string {
	return "/dev/ttyACM*"
}
