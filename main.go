package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CheckErr to handle errors
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetNIST reads the nist.json file
func GetNIST(file string) (raw []byte) {
	raw, err := ioutil.ReadFile(file)
	CheckErr(err)
	return raw
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Nist is a placeholder type for the nist.json to parse it
type Nist []interface{}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	e := echo.New()
	//e.Static("/static", "static")
	e.File("/favicon.ico", "favicon.ico")
	//e.File("/common.css", "common.css")
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET}, //echo.PUT, echo.POST, echo.DELETE},
	}))

	var b []byte
	var n []Nist

	b = GetNIST("nist.json")
	json.Unmarshal(b, &n)
	log.Println(b)

	e.Logger.Fatal(e.Start(":3000"))
	//

}
