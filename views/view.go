package views

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/csrf"
)

// View type contains a generic template.
type View struct {
	Template *template.Template
}

// NewView takes passed template file from handler, all layout files,
// parses them and checks for errors. Layout template files are
// non-specific templates, like "base", "navbar" or "footer".
func NewView(files ...string) *View {
	// Gather all the layout files.
	layoutFiles, err := filepath.Glob("views/templates/layouts/*.html")
	if err != nil {
		log.Fatal("Error finding layout files:", err)
	}

	// Take passed specific template from the handler, append
	// to the layout template files, define csrfField function and then
	// parse all layout templates. Template Func csrfField here is only definition,
	// implementation is done in the Render method. If csrfField function
	// returns an error, function stops execution of the template immediately.
	files = append(files, layoutFiles...)
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", errors.New("CSRF is not defined")
		},
	}).ParseFiles(files...))

	// Pass parsed template and layouts to the View.
	return &View{
		Template: tmpl,
	}
}

// Render sets header, executes passed template as base (b string)
// with the passed view data (vd interface{}) and checks for errors.
func (v *View) Render(w http.ResponseWriter, r *http.Request, b string, vd interface{}) {
	// Set header as "text/html".
	w.Header().Set("Content-Type", "text/html")

	// csrfField function implementation. Adds CSRF protection to templates.
	// Add {{csrfField}} in the template form.
	t := v.Template.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
	})

	// Execute template with data, if there is passed any data.
	if err := t.ExecuteTemplate(w, b, vd); err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong!", http.StatusInternalServerError)
	}
}
