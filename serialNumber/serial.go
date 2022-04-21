package serialNumber

import (
	"fmt"
	"os/exec"
	"strings"
)

func SerialNumber() {
	out, _ := exec.Command("/usr/sbin/ioreg", "-l").Output() // err ignored for brevity
	for _, l := range strings.Split(string(out), "\n") {
		if strings.Contains(l, "IOPlatformSerialNumber") {
			s := strings.Split(l, " ")
			fmt.Printf("%s\n", s[len(s)-1])
		}
	}
}
