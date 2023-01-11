package formatter

import (
	"encoding/json"
	"github.com/TylerBrock/colorjson"
	"go-http-cli/color"

	"io/ioutil"
	"net/http"
	"strings"
)

type formatter struct {
	resp *http.Response
}

func NewFormatter(resp *http.Response) *formatter {
	return &formatter{resp: resp}
}

func (f *formatter) GetBodyResponse() string {
	body, _ := ioutil.ReadAll(f.resp.Body)
	var obj map[string]interface{}
	json.Unmarshal(body, &obj)
	jsonformat := colorjson.NewFormatter()
	jsonformat.Indent = 4
	s, _ := jsonformat.Marshal(obj)
	return string(s)
}

func (f *formatter) GetHeaderResponse() string {
	var headerResponse = []string{}
	for headerName, headerValue := range f.resp.Header {
		temp := append(headerResponse, color.Blue+headerName+": "+color.Reset+strings.Join(headerValue, ", "))
		headerResponse = temp
	}
	return strings.Join(headerResponse, "\n")
}

func (f *formatter) GetStatusResponse() string {
	return color.Green + f.resp.Status + " " + color.White + f.resp.Proto
}
