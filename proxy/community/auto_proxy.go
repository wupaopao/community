package community

import (
	"github.com/mz-eco/mz/http"
)

type Proxy struct {
	http.Proxy
}

func NewProxy(service string) *Proxy {
	return &Proxy{
		Proxy: http.Proxy{
			Name: service,
		},
	}

}
