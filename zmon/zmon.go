package main
import (
	"log"
	//"log/syslog"
	linuxproc "github.com/c9s/goprocinfo/linux"
)

func main() {
	log.Print("hello world from zmon")

	/*logwriter, e := syslog.New(syslog.LOG_NOTICE, "myprog")
	if e == nil {
		log.SetOutput(logwriter)
	} else {
		log.Print("cannot make log write to syslog")
	}

	log.Print("hello world should go to syslog")*/

	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}

	for _, s := range stat.CPUStats {
		log.Print("idle: ", s.Idle)
	}
}
