package ntpsvr

import (
	"time"

	ntp "github.com/lixiangyun/go_ntp"
)

//Start start ntp server
func Start(port string) {

	go func(p string) {
		// Listens to addresses and ports
		ntps := ntp.NewNTPS(p)
		// Start the service, then coroutine process is created in the background.
		ntps.Start()
		for {
			time.Sleep(60 * time.Second)
		}

		// Stop the service.
		//ntps.Stop()
	}(port)

}
