// Package render is a middleware for Martini that provides easy JSON serialization and HTML template rendering.
//
//  package main
//
//  import (
//    "encoding/xml"
//
//    "github.com/go-martini/martini"
//    "github.com/martini-contrib/render"
//  )
//
//  type Greeting struct {
//    XMLName xml.Name `xml:"greeting"`
//    One     string   `xml:"one,attr"`
//    Two     string   `xml:"two,attr"`
//  }
//
//  func main() {
//    m := martini.Classic()
//    m.Use(render.Renderer()) // reads "templates" directory by default
//
//    m.Get("/html", func(r render.Render) {
//      r.HTML(200, "mytemplate", nil)
//    })
//
//    m.Get("/json", func(r render.Render) {
//      r.JSON(200, "hello world")
//    })
//
//    m.Get("/xml", func(r render.Render) {
//      r.XML(200, Greeting{One: "hello", Two: "world"})
//    })
//
//    m.Run()
//  }
package render

import (
	// "bytes"
	// "encoding/json"
	// "encoding/xml"
	"fmt"
	// "html/template"
	// "io"
	// "io/ioutil"
	"net/http"
	// "os"
	// "path/filepath"
	// "strings"

	"github.com/flosch/pongo2"
	// "github.com/go-martini/martini"
)

type Options struct {
	Base    string
	Debug   bool
	Globals pongo2.Context
}

type HTMLOptions struct {
	Layout string
}

type Render struct {
	*pongo2.TemplateSet
	res http.ResponseWriter
	req *http.Request
}

func (r *Render) HTML(str string, context pongo2.Context) {
	body, err := r.RenderTemplateFile(str, context)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	r.res.Write([]byte(body))
}

func (r *Render) String(str string, context pongo2.Context) {
	body, err := r.RenderTemplateString(str, context)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	r.res.Write([]byte(body))
}

func (r *Render) Template(str string, context pongo2.Context) {
	body, err := r.RenderTemplateFile(str, context)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	r.res.Write([]byte(body))
}

func (r *Render) TemplateStr(str string, context pongo2.Context) {
	body, err := r.RenderTemplateString(str, context)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	r.res.Write([]byte(body))
}

var render *pongo2.TemplateSet

func New(pt *pongo2.TemplateSet, res http.ResponseWriter, req *http.Request) *Render {
	return &Render{pt, res, req}
}

// func Renderer(options Options) martini.Handler {
// 	render = pongo2.NewSet("render", pongo2.MustNewLocalFileSystemLoader(options.Base))
// 	render.Debug = options.Debug
// 	render.Globals = options.Globals
// 	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
// 		options.Globals["req"] = req
// 		c.Map(New(render, res, req))
// 	}
// }
