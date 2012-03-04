/*
  This Source Code Form is subject to the terms of the Mozilla Public
  License, v. 2.0. If a copy of the MPL was not distributed with this
  file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
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
