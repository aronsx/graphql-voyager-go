package voyager

import (
	_ "embed"
	"net/http"
	"text/template"
)

var voyagerTemplate *template.Template

//go:embed html/voyager.html
var htmlTemplate string

//go:embed html/voyager.css
var css string

type voyagerOptions struct {
	Protocol string
	Endpoint string
	Host     string
	Headers  string
	CSS      string
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
		Endpoint: endpoint,
		Headers:  "{'Accept': 'application/json', 'Content-Type': 'application/json'}",
		Host:     "window.location.host",
		Protocol: "window.location.protocol",
		CSS:      css,
	}
	return v
}

// Serves the Voyager UI
func (v voyagerOptions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	voyagerTemplate.Execute(w, &v)
}
