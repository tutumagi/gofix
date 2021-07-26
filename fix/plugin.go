package fix

import (
	"fmt"
	"plugin"
	"plugin_fix/route"
)

const (
	EntryFuncName = "InitPlugin"
	PluginPostfix = "so"
)

type EntryFunc = func(r *route.Route)

type Plugin struct {
	Name string

	EntrySymbol EntryFunc

	p           *plugin.Plugin
	load        bool
	initVersion string
	curVersion  string
}

func NewPlugin(name string, initVersion string) *Plugin {
	p := &Plugin{}
	p.Name = name
	p.initVersion = initVersion
	p.curVersion = initVersion

	return p
}

func (pl *Plugin) curFilePath() string {
	return fmt.Sprintf("%s%s.%s", pl.Name, pl.curVersion, PluginPostfix)
}

func (pl *Plugin) Load() error {
	if pl.load {
		return nil
	}
	var ok bool

	p, err := plugin.Open(pl.curFilePath())
	if err != nil {
		return err
	}

	// load entry symbol
	entrySymbol, err := p.Lookup(EntryFuncName)
	if err != nil {
		return err
	}
	pl.EntrySymbol, ok = entrySymbol.(EntryFunc)
	if !ok {
		return fmt.Errorf("entry func symbol invalid. should be func(r *route.Route)")
	}

	pl.p = p
	pl.load = true
	return nil
}

func (pl *Plugin) Reload(newVersion string) error {
	fmt.Printf("reload plugin(name:%s) newVer:%s\n", pl.Name, newVersion)
	pl.load = false
	pl.curVersion = newVersion
	return pl.Load()
}
