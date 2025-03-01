package voyager

import (
	_ "embed"
	"net/http"
	"text/template"
)

var voyagerTemplate *template.Template

//go:embed html/voyager.html
var htmlTemplate string

//go:embed html/content/voyager.css
var css string

//go:embed html/content/voyager.standalone.js
var voyagerStandalone string

type voyagerOptions struct {
	Protocol string
	Endpoint string
	Host     string
	Headers  string

	CSS               string
	VoyagerStandalone string
}

func init() {
	vt, err := template.New("voyager").Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	voyagerTemplate = vt
}

// NewVoyagerHandler Creates a new Voyager Options object
// Example:
//
//	http.Handle("/voyager", voyagerHandler)
func NewVoyagerHandler(endpoint string) http.Handler {
	v := voyagerOptions{
		Endpoint:          endpoint,
		Headers:           "{'Accept': 'application/json', 'Content-Type': 'application/json'}",
		Host:              "window.location.host",
		Protocol:          "window.location.protocol",
		CSS:               css,
		VoyagerStandalone: voyagerStandalone,
	}
	return v
}

// Serves the Voyager UI
func (v voyagerOptions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	voyagerTemplate.Execute(w, &v)
}
