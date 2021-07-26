package main

import (
	"fmt"
	"os"
	"os/signal"
	"plugin_fix/fix"
	"plugin_fix/route"
	"syscall"
	"time"
)

// type routers struct {
// 	rts map[string]func()
// }

// func (r *routers) Register

const (
	initVer   = "1"
	hotFixVer = "10"
)

var routes = route.NewRoute()

func main() {
	pp := fix.NewPlugin("plugin", initVer)

	if err := pp.Load(); err != nil {
		panic(err)
	}

	pp.EntrySymbol(routes)

	dosomelogic()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, os.Interrupt)

	counter := 0
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			newVersion := fmt.Sprintf("%d", counter)
			if newVersion == hotFixVer {
				if err := pp.Reload(newVersion); err != nil {
					fmt.Printf("reload plugin(newVer:%s) err %s\n", newVersion, err)
				} else {
					pp.EntrySymbol(routes)
				}
			}

			fmt.Println(counter)
			dosomelogic()
			counter++
		case <-sig:
			return
		}
	}
}

func dosomelogic() {
	routes.Get("error")()
	routes.Get("info")()
}
