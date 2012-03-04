package daemon
import (
	"log"
	"os"
	"procinfo"
	"path/filepath"
)
func IsDaemon() bool {
	return os.Getppid() == 1;
}
func Daemonize() {
	DaemonizeWithFiles(nil,nil,nil)
}
func DaemonizeWithStdFiles() {
	DaemonizeWithFiles(os.Stdin,os.Stdout,os.Stderr)
}
func DaemonizeWithFiles(in *os.File,out *os.File,err *os.File) {
	filename,_ := procinfo.GetProcessExecutable()
	DaemonizeFile(filename,os.Args[1:],&os.ProcAttr{Files:[]*os.File{in,out,err}})
}
func DaemonizeFile(filename string,args []string,procAttr *os.ProcAttr) {
	filename,_ = filepath.Abs(filename);
	if(IsDaemon()) {
		doChildDaemon()
	} else {
		_,err := os.StartProcess(filename,append([]string{filename},args...),procAttr);
		if(err != nil) {
			log.Println("Error creating child processing: ",err);
		}
		os.Exit(0);
	}
}
