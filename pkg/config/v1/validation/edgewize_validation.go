package validation

import splugin "github.com/fatedier/frp/pkg/plugin/server"

func init() {
	SupportedHTTPPluginOps = []string{
		splugin.OpLogin,
	}
}
