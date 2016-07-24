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
//| Cellgo Framework net type
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-15

package cellgo

import (
	"html/template"
	//"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// NewContext return the Context with Input and Output
func NewNetInfo() *Netinfo {
	return &Netinfo{
		Input:  NewInput(),
		Output: NewOutput(),
	}
}

type Netinfo struct {
	Input    *CellInput
	Output   *CellOutput
	Request  *http.Request
	Response http.ResponseWriter
}

// Reset init Netinfo, CellInput and CellOutput
func (ni *Netinfo) Reset(w http.ResponseWriter, r *http.Request) {
	ni.Request = r
	ni.Response = w
	ni.Input.Reset(ni)
	ni.Output.Reset(ni)
}

// ------------------------------------------------------------------
// CellInput type
//-------------------------------------------------------------------
var (
	acceptsHTMLRegex = regexp.MustCompile(`(text/html|application/xhtml\+xml)(?:,|$)`)
	acceptsXMLRegex  = regexp.MustCompile(`(application/xml|text/xml)(?:,|$)`)
	acceptsJSONRegex = regexp.MustCompile(`(application/json)(?:,|$)`)
	maxParam         = 50
)

type CellInput struct {
	Netinfo *Netinfo
	//session!
	pnames      []string
	pvalues     []string
	data        map[interface{}]interface{}
	RequestBody []byte //Reserve
}

// NewInput return CellInput generated by Netinfo.
func NewInput() *CellInput {
	return &CellInput{
		pnames:  make([]string, 0, maxParam),
		pvalues: make([]string, 0, maxParam),
		data:    make(map[interface{}]interface{}),
	}
}

// Reset init the CellInput
func (input *CellInput) Reset(ni *Netinfo) {
	input.Netinfo = ni
	//session!
	//input.Session = nil
	input.pnames = input.pnames[:0]
	input.pvalues = input.pvalues[:0]
	input.data = nil
	input.RequestBody = []byte{}
}

// Protocol returns request protocol name, such as HTTP/1.1 .
func (input *CellInput) Protocol() string {
	return input.Netinfo.Request.Proto
}

// URI returns full request url with query string, fragment.
func (input *CellInput) URI() string {
	return input.Netinfo.Request.RequestURI
}

// URL returns request url path (without query string, fragment).
func (input *CellInput) URL() string {
	return input.Netinfo.Request.URL.Path
}

// Site returns base site url as scheme://domain type.
func (input *CellInput) Site() string {
	return input.Scheme() + "://" + input.Domain()
}

// Domain returns host name.
// Alias of Host method.
func (input *CellInput) Domain() string {
	return input.Host()
}

// Host returns host name.
// if no host info in request, return localhost.
func (input *CellInput) Host() string {
	if input.Netinfo.Request.Host != "" {
		hostParts := strings.Split(input.Netinfo.Request.Host, ":")
		if len(hostParts) > 0 {
			return hostParts[0]
		}
		return input.Netinfo.Request.Host
	}
	return "localhost"
}

// Scheme returns request scheme as "http" or "https".
func (input *CellInput) Scheme() string {
	if input.Netinfo.Request.URL.Scheme != "" {
		return input.Netinfo.Request.URL.Scheme
	}
	if input.Netinfo.Request.TLS == nil {
		return "http"
	}
	return "https"
}

// Method returns http request method.
func (input *CellInput) Method() string {
	return input.Netinfo.Request.Method
}

// Is returns boolean of this request is on given method, such as Is("POST").
func (input *CellInput) Is(method string) bool {
	return input.Method() == method
}

// IsGet Is this a GET method request?
func (input *CellInput) IsGet() bool {
	return input.Is("GET")
}

// IsPost Is this a POST method request?
func (input *CellInput) IsPost() bool {
	return input.Is("POST")
}

// IsHead Is this a Head method request?
func (input *CellInput) IsHead() bool {
	return input.Is("HEAD")
}

// IsOptions Is this a OPTIONS method request?
func (input *CellInput) IsOptions() bool {
	return input.Is("OPTIONS")
}

// IsPut Is this a PUT method request?
func (input *CellInput) IsPut() bool {
	return input.Is("PUT")
}

// IsDelete Is this a DELETE method request?
func (input *CellInput) IsDelete() bool {
	return input.Is("DELETE")
}

// IsPatch Is this a PATCH method request?
func (input *CellInput) IsPatch() bool {
	return input.Is("PATCH")
}

// IsAjax returns boolean of this request is generated by ajax.
func (input *CellInput) IsAjax() bool {
	return input.Header("X-Requested-With") == "XMLHttpRequest"
}

// IsSecure returns boolean of this request is in https.
func (input *CellInput) IsSecure() bool {
	return input.Scheme() == "https"
}

// IsWebsocket returns boolean of this request is in webSocket.
func (input *CellInput) IsWebsocket() bool {
	return input.Header("Upgrade") == "websocket"
}

// IsUpload returns boolean of whether file uploads in this request or not..
func (input *CellInput) IsUpload() bool {
	return strings.Contains(input.Header("Content-Type"), "multipart/form-data")
}

// AcceptsHTML Checks if request accepts html response
func (input *CellInput) AcceptsHTML() bool {
	return acceptsHTMLRegex.MatchString(input.Header("Accept"))
}

// AcceptsXML Checks if request accepts xml response
func (input *CellInput) AcceptsXML() bool {
	return acceptsXMLRegex.MatchString(input.Header("Accept"))
}

// AcceptsJSON Checks if request accepts json response
func (input *CellInput) AcceptsJSON() bool {
	return acceptsJSONRegex.MatchString(input.Header("Accept"))
}

// UserAgent returns request client user agent string.
func (input *CellInput) UserAgent() string {
	return input.Header("User-Agent")
}

// ParamsLen return the length of the params
func (input *CellInput) ParamsLen() int {
	return len(input.pnames)
}

// GetGP returns router param by a given key.
func (input *CellInput) GetGP(key string, isfilter bool) string {
	var res string
	for i, v := range input.pnames {
		if v == key && i <= len(input.pvalues) {
			if isfilter {
				res = template.HTMLEscapeString(input.pvalues[i])
			} else {
				res = input.pvalues[i]
			}
			return res
		}
	}
	return ""
}

// Params returns the map[key]value.
func (input *CellInput) GetGPs(isfilter bool) map[string]string {
	m := make(map[string]string)
	var tempVal string
	for i, v := range input.pnames {
		if i <= len(input.pvalues) {
			if isfilter {
				tempVal = template.HTMLEscapeString(input.pvalues[i])
			} else {
				tempVal = input.pvalues[i]
			}
			m[v] = tempVal
		}
	}
	return m
}

// SetParam will set the param with key and value
func (input *CellInput) SetParam(key, val string) {
	// check if already exists
	for i, v := range input.pnames {
		if v == key && i <= len(input.pvalues) {
			input.pvalues[i] = val
			return
		}
	}
	input.pvalues = append(input.pvalues, val)
	input.pnames = append(input.pnames, key)
}

// Query returns input data item string by a given string.
func (input *CellInput) Query(key string) string {
	if val := input.GetGP(key, false); val != "" {
		return val
	}
	if input.Netinfo.Request.Form == nil {
		input.Netinfo.Request.ParseForm()
	}
	return input.Netinfo.Request.Form.Get(key)
}

// Header returns request header item string by a given string.
// if non-existed, return empty string.
func (input *CellInput) Header(key string) string {
	return input.Netinfo.Request.Header.Get(key)
}

// Cookie returns request cookie item string by a given key.
// if non-existed, return empty string.
func (input *CellInput) Cookie(key string) string {
	ck, err := input.Netinfo.Request.Cookie(key)
	if err != nil {
		return ""
	}
	return ck.Value
}

// Session returns current session item value by a given key.
// if non-existed, return empty string.
//func (input *CellInput) Session(key interface{}) interface{} {
//	return input.Session.Get(key)
//}

// Data return the implicit data in the input
func (input *CellInput) Data() map[interface{}]interface{} {
	if input.data == nil {
		input.data = make(map[interface{}]interface{})
	}
	return input.data
}

// GetData returns the stored data in this context.
func (input *CellInput) GetData(key interface{}) interface{} {
	if v, ok := input.data[key]; ok {
		return v
	}
	return nil
}

// SetData stores data with given key in this context.
// This data are only available in this context.
func (input *CellInput) SetData(key, val interface{}) {
	if input.data == nil {
		input.data = make(map[interface{}]interface{})
	}
	input.data[key] = val
}

// ------------------------------------------------------------------
// CellOutput type
//-------------------------------------------------------------------

type CellOutput struct {
	Netinfo    *Netinfo
	Status     int
	EnableGzip bool
}

// NewOutput returns new CellOutput.
// it contains nothing now.
func NewOutput() *CellOutput {
	return &CellOutput{}
}

// Reset init CellOutput
func (output *CellOutput) Reset(ni *Netinfo) {
	output.Netinfo = ni
	output.Status = 0
}

// Header sets response header item string via given key.
func (output *CellOutput) Header(key, val string) {
	output.Netinfo.Response.Header().Set(key, val)
}

// SetStatus sets response status code.
// It writes response header directly.
func (output *CellOutput) SetStatus(status int) {
	output.Status = status
}

// IsCachable returns boolean of this request is cached.
// HTTP 304 means cached.
func (output *CellOutput) IsCachable(status int) bool {
	return output.Status >= 200 && output.Status < 300 || output.Status == 304
}

// IsEmpty returns boolean of this request is empty.
// HTTP 201，204 and 304 means empty.
func (output *CellOutput) IsEmpty(status int) bool {
	return output.Status == 201 || output.Status == 204 || output.Status == 304
}

// IsOk returns boolean of this request runs well.
// HTTP 200 means ok.
func (output *CellOutput) IsOk(status int) bool {
	return output.Status == 200
}

// IsSuccessful returns boolean of this request runs successfully.
// HTTP 2xx means ok.
func (output *CellOutput) IsSuccessful(status int) bool {
	return output.Status >= 200 && output.Status < 300
}

// IsRedirect returns boolean of this request is redirection header.
// HTTP 301,302,307 means redirection.
func (output *CellOutput) IsRedirect(status int) bool {
	return output.Status == 301 || output.Status == 302 || output.Status == 303 || output.Status == 307
}

// IsForbidden returns boolean of this request is forbidden.
// HTTP 403 means forbidden.
func (output *CellOutput) IsForbidden(status int) bool {
	return output.Status == 403
}

// IsNotFound returns boolean of this request is not found.
// HTTP 404 means forbidden.
func (output *CellOutput) IsNotFound(status int) bool {
	return output.Status == 404
}

// IsClientError returns boolean of this request client sends error data.
// HTTP 4xx means forbidden.
func (output *CellOutput) IsClientError(status int) bool {
	return output.Status >= 400 && output.Status < 500
}

// IsServerError returns boolean of this server handler errors.
// HTTP 5xx means server internal error.
func (output *CellOutput) IsServerError(status int) bool {
	return output.Status >= 500 && output.Status < 600
}

func stringsToJSON(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}

// Session sets session item value with given key.
//func (output *CellOutput) Session(name interface{}, value interface{}) {
//	output.Netinfo.Input.Session.Set(name, value)
//}
