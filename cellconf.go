//|------------------------------------------------------------------
//|        __
//|     __/  \
//|  __/  \__/_
//| /  \__/    \
//|/\__/CellGo /_
//|\/_/NetFW__/  \
//|  /\__ _/  \__/
//|  \/_/  \__/_/
//|    /\__/_/
//|    \/_/
//|------------------------------------------------------------------
//| Cellgo Framework conf file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-18

package cellgo

// Conf is the main struct for Config
type Conf struct {
	NetName    string //Net name
	ServerName string
	Listen
	SiteConfig
}

// Listen holds for http/https/websocket related config
type Listen struct {
	ServerTimeOut int64
	HTTPAddr      string //http  conn addr
	HTTPPort      int    //http  conn Port
	HTTPSAddr     string //https conn addr
	HTTPSPort     int    //https conn Port
	HTTPSCertFile string //https conn certfile
	HTTPSKeyFile  string //https conn keyfile
	WEBSOCKETAddr string //websocket conn Port
	WEBSOCKETPort int    //websocket conn Port
	SOCKETIOAddr  string //socketio conn Port
	SOCKETIOPort  int    //socketio conn Port
}

// SiteConfig holds Site related config
type SiteConfig struct {
	AutoDisplay       bool
	DefaultBeforeAct  string
	DefaultAfterAct   string
	DefaultController string
	DefaultAction     string
	Dynamic           string
	StaticDir         string
	StaticRouter      []string
	LabLeft           string
	LabRight          string
	TemplateExt       string
	TemplatePath      string
	IsViewFilter      bool //Open html filter * 1. 对多人合作，安全性可控比较差的项目建议开启 2. 对HTML进行转义，预防XSS攻击
	IsUri             bool //Open rewrite path: /user/add/username/jack
	SessionOn         bool
	SessionName       string
	SessionMaxage     int64
	CookieOn          bool
	CookieName        string
	CookieMaxage      string
	CookieHashKey     string //Encryption interference
	CookieSecure      string //https security
}

// Version number of the cellgo.
const (
	VERSION  = "0.2.2"
	LASTDATE = "August 20, 2016"
)

var (
	// CellConf is the default config for Cellgo
	CellConf *Conf
)

func init() {
	CellConf = &Conf{
		NetName:    "cellgo",
		ServerName: "CellgoService_" + VERSION,
		Listen: Listen{
			ServerTimeOut: 0,
			HTTPAddr:      "",
			HTTPPort:      8001,
			HTTPSAddr:     "",
			HTTPSPort:     10443,
			HTTPSCertFile: "",
			HTTPSKeyFile:  "",
			WEBSOCKETAddr: "",
			WEBSOCKETPort: 8088,
			SOCKETIOAddr:  "",
			SOCKETIOPort:  5000,
		},
		SiteConfig: SiteConfig{
			AutoDisplay:       true,
			DefaultBeforeAct:  "Before",
			DefaultAfterAct:   "After",
			DefaultController: "Index",
			DefaultAction:     "Run",
			Dynamic:           "/",
			StaticDir:         "static",
			StaticRouter:      []string{"/css/", "/js/", "/images/"},
			LabLeft:           "{{",
			LabRight:          "}}",
			TemplateExt:       "html",
			TemplatePath:      "template",
			IsViewFilter:      false,
			IsUri:             false,
			SessionOn:         true,
			SessionName:       "cellsession",
			SessionMaxage:     3600,
			CookieOn:          true,
			CookieName:        "cellcookie",
			CookieMaxage:      "86400",
			CookieHashKey:     "9597f4KpYTsJ5tD6",
			CookieSecure:      "false",
		},
	}

}
