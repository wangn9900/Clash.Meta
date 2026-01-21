package tunnel

import (
	C "github.com/metacubex/mihomo/constant"
	P "github.com/metacubex/mihomo/constant/provider"
)

var (
	allProxies = make(map[string]C.Proxy)
)

func AllProxies() map[string]C.Proxy {
	return allProxies
}

func UpdateAllProxies(proxies map[string]C.Proxy, providers map[string]P.ProxyProvider) {
	var allProxiesTemp = make(map[string]C.Proxy)
	for name, proxy := range proxies {
		allProxiesTemp[name] = proxy
	}
	for _, p := range providers {
		for _, proxy := range p.Proxies() {
			name := proxy.Name()
			allProxies[name] = proxy
		}
	}
	allProxies = allProxiesTemp
}
