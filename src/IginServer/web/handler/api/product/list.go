package product

import (
	// "IginServer/lib/Imartini"
	// "IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	// "IginServer/conf"
	"IginServer/framework/API/R"
	// "IginServer/web/handler/api/sms"
	"IginServer/web/model/tranca/product"
	// "encoding/json"
	// w "IginServer/framework/weixin"
	// "IginServer/lib/md5"
	// "IginServer/lib/other"
	// "IginServer/lib/redis"
	// "IginServer/lib/session"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	// "net/http"
	// "time"
)

/**
 * @api {get} /api/product/list 获取商品列表
 * @apiName 获取商品列表
 * @apiGroup 产品
 *
 * @apiVersion 1.0.0
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
 * {
 *  "data":[
 *   {
 *    "createtime":"0",
 *    "id":"1",
 *    "name":"8999",
 *    "price":"0.01",
 *    "product_id":"8999",
 *    "type":"0",
 *    "updatetime":"2017-05-03 14:43:45"
 *   }
 *  ],
 *  "status":200
 * }
 *
 *
 *
 * @apiSuccess (Reponse 404) {Number} status  404
 * @apiSuccess (Reponse 404) {String} error 缺少必要参数
 *
 * @apiSuccess (Reponse 403) {Number} status  403
 * @apiSuccess (Reponse 403) {String} error 参数不正确
 *
 * @apiError (Reponse 500) {Number} status 500
 * @apiError (Reponse 500) {String} error 系统错误
 *
 * @apiSampleRequest http://121.41.116.104:8003/api/product/list
 */
func GetList(c *gin.Context) {
	list := product.ProductFindAll()
	R.Success(c.Writer, list)
}
