package main

import (
	"bufio"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/tarm/serial"
)

var mutex = sync.Mutex{}
var latestVolume int

func setVolume(v int) {
	mutex.Lock()
	if v == latestVolume {
		cmd := volumeCommand(v)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Print("[error] setting volume: ", err)
			log.Println(string(output))
		}
	}
	mutex.Unlock()
}

func main() {
	ttys, _ := filepath.Glob(ttyGlob())
	c := &serial.Config{Name: ttys[0], Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(s)
	for {
		line, err := r.ReadString('\n')
		if err == nil {
			log.Print("[input] ", line)
			if strings.HasPrefix(line, "Percent:") {
				volume := strings.TrimSpace(line[8:])
				log.Print("[volume] ", volume)

				v, err := strconv.Atoi(volume)
				if err != nil {
					log.Print("[error] parsing volume: ", err)
					continue
				}

				latestVolume = v

				go setVolume(v)
			}
		} else {
			log.Print("[error] reading line: ", err)
		}
	}
}
