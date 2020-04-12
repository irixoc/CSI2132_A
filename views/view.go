package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

//LayoutDir gives a string to be used by layoutFiles function
const (
	LayoutDir string = "views/layouts/"
	TempExt   string = ".html"
)

// View struct
type View struct {
	Template *template.Template
	Layout   string
}

// NewView returns a new view
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Render is called to create a view
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// ServeHTTP is implemented by View type from Handler interface
// This allows certain views to be able to be passed to the r.Handle
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// layoutFiles() globs through the layouts folder
// and returns a slice of all the filespaths for the layout files
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TempExt)
	if err != nil {
		panic(err)
	}
	return files
}
