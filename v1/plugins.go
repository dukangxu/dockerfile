// +build ignore

package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

var plugins string
var telemetry bool

func main() {
	flag.StringVar(&plugins, "plugins", "", "plugins separate by comma")
	flag.BoolVar(&telemetry, "telemetry", false, "enable telemetry")
	flag.Parse()

	d := &data{
		EnableTelemetry: telemetry,
	}
	plugins := strings.FieldsFunc(plugins, func(c rune) bool { return c == ',' })
	d.Plugins = make([]string, 0, len(plugins))
	for _, plugin := range plugins {
		value, ok := pluginRepo[plugin]
		if !ok {
			log.Printf("plugin %s not support\n", plugin)
			continue
		}
		d.Plugins = append(d.Plugins, value)
		log.Printf("Additional plugins: %v\n", plugin)
	}
	log.Printf("Enabled telemetry: %v\n", d.EnableTelemetry)
	f, err := os.Create("caddy.go")
	defer f.Close()
	if err != nil {
		log.Fatal("Unable to open file")
	}

	t := template.Must(template.New("caddy.go").Parse(caddyTemplate))
	t.Execute(f, d)
}

type data struct {
	Plugins         []string
	EnableTelemetry bool
}

var caddyTemplate = `
package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"

	// plug in plugins here
	{{range $plugin := .Plugins}}
	_ "{{$plugin}}"
	{{end}}
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = {{.EnableTelemetry}}
	caddymain.Run()
}
`

var pluginRepo = map[string]string{
	"realip":        "github.com/captncraig/caddy-realip",
	"git":           "github.com/abiosoft/caddy-git",
	"proxyprotocol": "github.com/mastercactapus/caddy-proxyprotocol",
	"locale":        "github.com/simia-tech/caddy-locale",
	"cache":         "github.com/nicolasazrak/caddy-cache",
	"minify":        "github.com/hacdias/caddy-minify",
	"geoip":         "github.com/kodnaplakal/caddy-geoip",
	"authz":         "github.com/casbin/caddy-authz",
	"filter":        "github.com/echocat/caddy-filter",
	"ipfilter":      "github.com/pyed/ipfilter",
	"ratelimit":     "github.com/xuqingfeng/caddy-rate-limit",
	"recaptcha":     "github.com/defund/caddy-recaptcha",
	"expires":       "github.com/epicagency/caddy-expires",
	"forwardproxy":  "github.com/caddyserver/forwardproxy",
	"cors":          "github.com/captncraig/cors/caddy",
	"s3browser":     "github.com/techknowlogick/caddy-s3browser",
	"nobots":        "github.com/Xumeiquer/nobots",
	"login":         "github.com/tarent/loginsrv/caddy",
	"reauth":        "github.com/freman/caddy-reauth",
	"extauth":       "github.com/BTBurke/caddy-extauth",
	"jwt":           "github.com/BTBurke/caddy-jwt",
	"permission":    "github.com/dhaavi/caddy-permission",
	"jsonp":         "github.com/pschlump/caddy-jsonp",
	"upload":        "blitznote.com/src/caddy.upload",
	"multipass":     "//github.com/namsral/multipass/caddy",
	"datadog":       "github.com/payintech/caddy-datadog",
	"prometheus":    "github.com/miekg/caddy-prometheus",
	"pubsub":        "github.com/jung-kurt/caddy-pubsub",
	"cgi":           "github.com/jung-kurt/caddy-cgi",
	"filebrowser":   "github.com/filebrowser/caddy",
	"webdav":        "github.com/hacdias/caddy-webdav",
	"mailout":       "//github.com/SchumacherFM/mailout",
	"awses":         "github.com/miquella/caddy-awses",
	"awslambda":     "github.com/coopernurse/caddy-awslambda",
	"grpc":          "github.com/pieterlouw/caddy-grpc",
	"gopkg":         "github.com/zikes/gopkg",
	"restic":        "github.com/restic/caddy",
	"wkd":           "github.com/emersion/caddy-wkd",
	"dyndns":        "github.com/linkonoid/caddy-dyndns",
}
