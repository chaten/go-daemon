package main
import (
	"daemon"
	"time"
)
func init() {
	daemon.Daemonize()
}
func main() {
	for {
		time.Sleep(1000)
		println("tick");
	}
}
