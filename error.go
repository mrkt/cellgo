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
//| Cellgo Framework error type
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-23

package cellgo

import (
	"html/template"
	"net/http"
	"reflect"
	"strings"
)

//Type numbers
const (
	errHandlerType = iota
	errControllerType
	errOtherType
)

// Error type.
type Error struct {
	// ErrMaps holds map of http handlers for each errorcode string.
	// error(40x and 50x)
	ErrMaps map[string]*errInfo
}

// Error Handler type or controller type or Other type.
type errInfo struct {
	controllerType reflect.Type
	otherType      reflect.Type
	handler        http.HandlerFunc
	errCode        string
	errType        int
}

var (
	// Error is the Error handling type for Cellgo
	CellError *Error
)

func init() {
	CellError = &Error{ErrMaps: make(map[string]*errInfo, 10)}
	CellError.registerHandlers()
}

// register Command and handle function
func (err *Error) registerHandlers() {
	m := map[string]func(http.ResponseWriter, *http.Request){
		"404": err.notFound,
		"500": err.serverInternalError,
	}
	for e, h := range m {
		if _, ok := err.ErrMaps[e]; !ok {
			err.ErrHandler(e, h)
		}
	}
}

//Set errorCode with errorFunc
// 	CellError.ErrHandler("404",notFound)
//	CellError.ErrHandler("500",serverInternalError)
func (err *Error) ErrHandler(code string, h http.HandlerFunc) *Core {
	err.ErrMaps[code] = &errInfo{
		errType: errHandlerType,
		handler: h,
		errCode: code,
	}
	return CellCore
}

//Set controllerFunc with errorFunc
// 	CellError.ErrController(&controllers.ErrController{})
func (err *Error) ErrController(c ControllerInterface) *Core {
	reflectVal := reflect.ValueOf(c)
	rt := reflectVal.Type()
	ct := reflect.Indirect(reflectVal).Type()
	for i := 0; i < rt.NumMethod(); i++ {
		methodName := rt.Method(i).Name
		if strings.HasPrefix(methodName, "Error") {
			errName := strings.TrimPrefix(methodName, "Error")
			err.ErrMaps[errName] = &errInfo{
				errType:        errControllerType,
				controllerType: ct,
				errCode:        methodName,
			}
		}
	}
	return CellCore
}

var errtpl = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<div id="wrapper">
			<div id="container">
				<div class="navtop">
					<h1>{{.Title}}</h1>
				</div>
				<div id="content">
					{{.Content}}
					<a href="/" title="Home" class="button">Go Home</a><br />
					<br>Powered by cellgo {{.Version}}
				</div>
			</div>
		</div>
	</body>
</html>
`

// show 404 notfound error.
func (err *Error) notFound(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("cellgoerrortemp").Parse(errtpl)
	data := map[string]interface{}{
		"Title":   http.StatusText(404),
		"Version": VERSION,
	}
	data["Content"] = template.HTML("<br>The page you have requested has flown the coop." +
		"<br>Perhaps you are here because:" +
		"<br><br><ul>" +
		"<br>The page has moved" +
		"<br>The page no longer exists" +
		"<br>You were looking for your puppy and got lost" +
		"<br>You like 404 pages" +
		"</ul>")
	t.Execute(rw, data)
}

// show 500 internal server error.
func (err *Error) serverInternalError(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("beegoerrortemp").Parse(errtpl)
	data := map[string]interface{}{
		"Title":   http.StatusText(500),
		"Version": VERSION,
	}
	data["Content"] = template.HTML("<br>The page you have requested is down right now." +
		"<br><br><ul>" +
		"<br>Please try again later and report the error to the website administrator" +
		"<br></ul>")
	t.Execute(rw, data)
}
