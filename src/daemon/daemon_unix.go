package daemon
import (
	"syscall"
	"os"
	"log"
)

func doChildDaemon() {
	sid,err := syscall.Setsid();
	if(sid < 0 || err != nil) {
		log.Println("Error setting sid ",err);
		os.Exit(-1);
	}
}
