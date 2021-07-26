package route

import (
	"fmt"
	"strings"
)

type Handler func()

type Route struct {
	rt map[string]Handler
}

func NewRoute() *Route {
	return &Route{
		rt: make(map[string]Handler),
	}
}

func (r *Route) Register(name string, handle Handler) {
	r.rt[name] = handle
}

func (r *Route) Get(name string) Handler {
	if h, ok := r.rt[name]; ok {
		return h
	}
	return nil
}

func (r *Route) Dump() {
	sb := &strings.Builder{}
	for k := range r.rt {
		sb.WriteString(k)
		sb.WriteString(", ")
	}

	fmt.Println(strings.TrimRight(sb.String(), ", "))
}
