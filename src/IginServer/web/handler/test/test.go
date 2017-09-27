package test

import (
	"IginServer/lib/Imartini"
	"IginServer/lib/mygin"
	"IginServer/lib/mysqldb"
	"IginServer/web"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Middleware(c *gin.Context) {
	// fmt.Println("this is a middleware!")
	c.Next()
	// fmt.Println("over")
}
func Middleware2(c *gin.Context) {
	// fmt.Println("this is a middleware!222")
	c.Next()
	// fmt.Println("over222")
}

func overtest(c *gin.Context) {
	// fmt.Println("this is a middleware!222")
	c.String(200, "123")
	c.Abort() //结束逻辑，不往下再执行
	c.Next()
	c.String(200, "123")
	// c.Next()
	// fmt.Println("over222")
}

func GetHandler(c *gin.Context) {
	// value, exist := c.GetQuery("key")
	// if !exist {
	//  value = "the key is not exist!"
	// }
	// c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))

	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "my name"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
	// fmt.Println("body")
	return
}
func PostHandler(c *gin.Context) {
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "my name222"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
	return
}
func PutHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
	return
}
func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}

func testdb(c *mygin.IContext) {

	data := c.DB.SetTable("test").FindAll()

	fmt.Println(c.GetAllSession())
	c.SetSession("data", mysqldb.List(data))
	c.JSON(200, mysqldb.List(data))
}

//模板测试
func testTemplates(c *gin.Context) {
	Imartini.HTML(c.Writer, c.Request, "test.html", map[string]interface{}{})
}

//模板测试
func testTemplates2(c *mygin.IContext) {
	c.HTML("test.html", map[string]interface{}{})
}

//流式响应
func Chunked(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	c.Writer.WriteHeader(http.StatusOK)
	type (
		Geolocation struct {
			Altitude  float64
			Latitude  float64
			Longitude float64
		}
	)

	var (
		locations = []Geolocation{
			{-97, 37.819929, -122.478255},
			{1899, 39.096849, -120.032351},
			{2619, 37.865101, -119.538329},
			{42, 33.812092, -117.918974},
			{15, 37.77493, -122.419416},
		}
	)
	for _, l := range locations {
		if err := json.NewEncoder(c.Writer).Encode(l); err != nil {
			fmt.Println(err)
		}
		c.Writer.Flush() //先把一部分消息推送出去
		time.Sleep(1 * time.Second)
	}
}

func init() {
	r := web.Router.Group("/test")
	//添加中间件
	r.Use(Middleware)
	r.GET("/testTemplates", testTemplates)
	r.GET("/testTemplates2", mygin.Handler(testTemplates2))
	r.GET("/chunked", Chunked)
	r.GET("/t", GetHandler)
	r.GET("/testdb", mygin.Handler(testdb))

	d := web.Router.Group("/test")
	//注册接口
	d.GET("/simple/server/get", overtest, GetHandler)
	d.POST("/simple/server/post", PostHandler)
	d.PUT("/simple/server/put", PutHandler)
	d.DELETE("/simple/server/delete", DeleteHandler)
}
