package godaemon

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func init() {
	goDaemon := flag.Bool("d", false, "run app as a daemon with -d=true.")
	flag.Parse()

	if *goDaemon {
		cmd := exec.Command(os.Args[0], flag.Args()...)
		cmd.Start()
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		os.Exit(0)
	}
}
