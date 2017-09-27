package Imartini

import (
	// "log"
	// "fmt"
	"net/http"
	// "server/conf"
	// "server/lib/redis"
	// "server/lib/session"
	// "server/lib/templatefunc"
	"IginServer/conf"
	"IginServer/lib/render"
	"IginServer/lib/templatefunc"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var M *Imartini
var PgRender *pongo2.TemplateSet
var DEBUG = conf.GetBool("config", "DEBUG")

var (
	LOGO_SIZE, _  = strconv.Atoi(conf.GET["config"]["LOGO_SIZE"])
	LOGO_PATH     = conf.GET["config"]["LOGO_PATH"]
	STATIC_ROUTER = conf.GET["config"]["STATIC_ROUTER"]
	STATIC_PATH   = conf.GET["config"]["STATIC_PATH"]
	GZIP, _       = strconv.ParseBool(conf.GET["config"]["GZIP"])
	LOGO, _       = strconv.ParseBool(conf.GET["config"]["LOGO"])
	CATURL, _     = strconv.ParseBool(conf.GET["config"]["CATURL"])
	SESSION, _    = strconv.ParseBool(conf.GET["config"]["SESSION"])
	RENDER, _     = strconv.ParseBool(conf.GET["config"]["RENDER"])
	STATIC, _     = strconv.ParseBool(conf.GET["config"]["STATIC"])
	IREDIS, _     = strconv.ParseBool(conf.GET["config"]["REDIS"])
)

type Imartini struct {
	*gin.Engine
}

func ServerInit() *Imartini {

	return M
}

func CreateStaticHandler(relativePath string, fs http.FileSystem) gin.HandlerFunc {
	// absolutePath := group.calculateAbsolutePath(relativePath)
	absolutePath := relativePath
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *gin.Context) {
		// if nolisting {
		// 	c.Writer.WriteHeader(404)
		// }
		file := filepath.Join(STATIC_PATH, strings.Split(c.Request.RequestURI, "?")[0][1:])
		fi, err := os.Stat(file)
		if os.IsNotExist(err) {
			return
		}
		if fi.IsDir() {
			file = filepath.Join(file, "index.html")
			_, err = os.Stat(file)
			if os.IsNotExist(err) {
				return
			}
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

func HTML(res http.ResponseWriter, req *http.Request, str string, context pongo2.Context) {
	res.WriteHeader(200)
	context["req"] = req
	render.New(PgRender, res, req).HTML(str, context)
}

func init() {
	// gin.DefaultWriter = os.Stdout
	// gin.DefaultErrorWriter = os.Stderr
	if DEBUG {
		gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	} else {
		gin.SetMode(gin.ReleaseMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	}
	engine := gin.New()
	// engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(Logger())
	// engine.Use(gin.Recovery())

	if STATIC {
		// engine.Static(STATIC_ROUTER, STATIC_PATH)
		engine.Use(CreateStaticHandler("/", gin.Dir(STATIC_PATH, false)))
	}

	if RENDER {
		options := render.Options{
			Debug:   DEBUG,
			Base:    "templates",
			Globals: templatefunc.Temfunc,
		}
		PgRender = pongo2.NewSet("render", pongo2.MustNewLocalFileSystemLoader(options.Base))
		PgRender.Debug = options.Debug
		PgRender.Globals = options.Globals
	}

	M = &Imartini{engine}
}
