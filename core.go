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
//| ------------------------------------------------------------------
//| Cellgo Framework core type
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-15

package cellgo

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mrkt/cellgo/session"
)

var (
	// CellCore is an core instance
	CellCore *Core
	SESSION  *session.Handle
)

func init() {
	// create cellgo core
	CellCore = NewCore()
	RegisterSession()
}

type Core struct {
	Handlers *ControllerRegister
	Server   *http.Server
}

func (core *Core) defalutHandler(w http.ResponseWriter, r *http.Request) {
	if core.CheckParam(w, r) {
		CellCore.Handlers.workHTTP(w, r)
	}
	return
}

// NewCore returns a new cellgo core.
func NewCore() *Core {
	cr := NewControllerRegister()
	core := &Core{Handlers: cr, Server: &http.Server{}}
	return core
}

// Run cellgo core.
func (core *Core) Run() {
	core.RegisterHttp()
}

//register http service
func (core *Core) RegisterHttp() {

	for _, v := range CellConf.SiteConfig.StaticRouter {
		http.Handle(v, http.FileServer(http.Dir(CellConf.SiteConfig.StaticDir)))
	}
	http.HandleFunc(CellConf.SiteConfig.Dynamic, core.defalutHandler)
	server := http.Server{
		Addr:        CellConf.Listen.HTTPAddr + ":" + fmt.Sprintf("%d", CellConf.Listen.HTTPPort),
		Handler:     nil,
		ReadTimeout: time.Duration(CellConf.Listen.ServerTimeOut) * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

//Check default parameters
//From the configuration file options
func (core *Core) CheckParam(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm() //Analytical parameters, the default is not resolved
	//Split URL, the backslash separated
	urlSlice := strings.Split(r.URL.Path, "/")
	if CellConf.SiteConfig.IsUri { //Open rewrite path
		switch len(urlSlice) {
		case 2:
			if urlSlice[1] == "" {
				break
			}
			r.Form["c"] = []string{urlSlice[1]}
		default:
			var tempKey string
			for k, url := range urlSlice {
				if k > 0 {
					if k == 1 {
						r.Form["c"] = []string{url}
					} else if k == 2 {
						r.Form["a"] = []string{url}
					} else {
						if k%2 != 0 {
							tempKey = url
						} else {
							r.Form[tempKey] = []string{url}
						}
					}
				}
			}
		}
	} else {
		if urlSlice[1] != "" {
			CellError.ErrMaps["404"].handler(w, r)
			return false
		}
	}

	if r.Form["c"] == nil {
		r.Form["c"] = []string{strings.ToLower(CellConf.SiteConfig.DefaultController)} //default controller
	}
	if r.Form["a"] == nil {
		r.Form["a"] = []string{strings.ToLower(CellConf.SiteConfig.DefaultAction)} //default action
	}
	return true
}

//register Session
func RegisterSession() {
	SESSION, _ = session.NewHandle("memorySess", "cellgosid", 3600)
	go SESSION.GC()
}

//register https service
func (core *Core) RegisterHttps() {}

//register websocket service
func (core *Core) RegisterWebsocket() {}

func (core *Core) RegisterController(title string, c ControllerInterface, param []string) *Core {
	core.Handlers.Add(title, c, param)
	return core
}
