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
//| Cellgo Framework tool/tool file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-03

package tool

var TOOL = &Tool{
	Rsa:     Rsa,
	Base64:  Base64,
	Urlcode: Urlcode,
	Json:    Json,
	Map2:    Map,
}

type Tool struct {
	Rsa     *rsaTool     //rsa Encrypt and Decrypt
	Base64  *base64Tool  //Base64 Encrypt and Decrypt
	Urlcode *urlcodeTool //Urlcode Encrypt and Decrypt
	Json    *jsonTool    //Json Transformation
	Map2    *mapTool     //2Map Transformation
}
