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

// Config is the main struct for BConfig
type Conf struct {
	NetName          string //Application name
	ServerName       string
	DefaultBeforeAct string
	DefaultAfterAct  string
	Listen           map[string]string
}

// Version number of the cellgo.
const (
	VERSION  = "0.0.3"
	LASTDATE = "July 19, 2016"
)

var (
	// BConfig is the default config for Cellgo
	CellConf *Conf
)

func init() {
	CellConf = &Conf{
		NetName:          "cellgo",
		ServerName:       "CellgoService_" + VERSION,
		DefaultBeforeAct: "Before",
		DefaultAfterAct:  "After",
		Listen:           map[string]string{},
	}

}
