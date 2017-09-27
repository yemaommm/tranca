package auth

import (
	// "IginServer/lib/Imartini"
	// "IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	// "IginServer/conf"
	"IginServer/framework/API/R"
	"IginServer/web/handler/api/sms"
	"IginServer/web/model/tranca/wap_user"
	// "encoding/json"
	// w "IginServer/framework/weixin"
	"IginServer/lib/md5"
	"IginServer/lib/other"
	"IginServer/lib/redis"
	"IginServer/lib/session"
	"fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	// "net/http"
	// "time"
)

/**
 * @api {post} /api/auth/register 注册接口
 * @apiName 注册接口
 * @apiGroup auth
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} password 密码
 * @apiParam {String} code 验证码
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
 *
 *
 *
 * @apiSuccess (Reponse 201) {Number} status  201
 * @apiSuccess (Reponse 201) {String} error 验证码不正确
 *
 * @apiSuccess (Reponse 202) {Number} status  202
 * @apiSuccess (Reponse 202) {String} error 注册失败
 *
 * @apiSuccess (Reponse 203) {Number} status  203
 * @apiSuccess (Reponse 203) {String} error 用户名已存在
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
 * @apiSampleRequest http://121.41.116.104:8003/api/auth/register
 */
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	code := c.PostForm("code")

	if username == "" || password == "" || code == "" {
		R.Api404(c.Writer)
		return
	}

	r := redis.Get()
	defer r.Close()

	if sms.Validate(username, "register", code, r) == 1 {
		R.Error(c.Writer, 201, "验证码不正确")
		return
	}
	if wap_user.FindUserForUsername(username) != nil {
		R.Error(c.Writer, 203, "用户名已存在")
		return
	}
	if 0 == wap_user.Register(username, password) {
		R.Error(c.Writer, 202, "注册失败")
		return
	}
	R.Success(c.Writer, "")
}

/**
 * @api {post} /api/auth/login 登录接口
 * @apiName 登录接口
 * @apiGroup auth
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} password 密码
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
 * {"data":{"token":"7cbcc1ff5912891228f3e0e6eb56077b","uid":"123","username":"username"},"status":200}
 * token和username需要保存
 *
 * @apiSuccess (Reponse 300) {Number} status  300
 * @apiSuccess (Reponse 300) {String} error 用户名或密码不正确
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
 * @apiSampleRequest http://121.41.116.104:8003/api/auth/login
 */
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	userinfo := wap_user.Login(username, password)

	if len(userinfo) <= 0 {
		R.Error(c.Writer, 300, "用户名或密码不正确")
	} else {
		userinfo["token"] = fmt.Sprintf("%v|%v|%v", username, password, md5.GetUUID().Hex())
		userinfo["token"] = md5.Md5(userinfo["token"])
		SaveToken(username, userinfo["token"], userinfo)
		R.Success(c.Writer, userinfo)
	}
}

/**
 * @api {post} /api/auth/logout 退出登录接口
 * @apiName 退出登录接口
 * @apiGroup auth
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} token 登录返回的token
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
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
 * @apiSampleRequest http://121.41.116.104:8003/api/auth/logout
 */
func Logout(c *gin.Context) {
	username := c.PostForm("username")
	token := c.PostForm("token")

	if username == "" || token == "" {
		R.Api404(c.Writer)
		return
	}

	DelToken(username, token)
	R.Success(c.Writer, "")
}

/**
 * @api {post} /api/auth/ 登录验证接口
 * @apiName 登录验证接口
 * @apiGroup auth
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} token 登录返回的token
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
 * {"data":{"token":"7cbcc1ff5912891228f3e0e6eb56077b","uid":"123","username":"username"},"status":200}
 * token和username需要保存
 *
 * @apiSuccess (Reponse 400) {Number} status  400
 * @apiSuccess (Reponse 400) {String} error 登录验证失败
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
 * @apiSampleRequest http://121.41.116.104:8003/api/auth/
 */
func Auth(c *gin.Context) {
	username := c.PostForm("username")
	token := c.PostForm("token")

	userinfo := GetTokenInfo(username, token)

	if userinfo == nil {
		R.Error(c.Writer, 400, "登录验证失败")
	} else {
		R.Success(c.Writer, userinfo)
	}
}

func AuthHandler(c *gin.Context) {
	username := c.PostForm("username")
	token := c.PostForm("token")

	userinfo := GetTokenInfo(username, token)

	if userinfo == nil {
		R.Error(c.Writer, 400, "登录验证失败")
		c.AbortWithStatus(200)
		return
	}
	c.Set("userinfo", userinfo)
}

func CreateTokenKey(username, token string) string {
	return fmt.Sprintf("%v:%v:%v", "token", username, token)
}

func SaveToken(username, token string, v interface{}) {
	str := CreateTokenKey(username, token)

	r := redis.Get()
	defer r.Close()
	r.Set(str, other.Map2Json(v))
	ExpireToken(username, token)
}

func DelToken(username, token string) {
	str := CreateTokenKey(username, token)

	r := redis.Get()
	defer r.Close()
	r.Del(str)
}

func ExpireToken(username, token string) {
	r := redis.Get()
	defer r.Close()

	r.Expire(CreateTokenKey(username, token), session.SESSION_TIME)
}

func GetTokenInfo(username, token string) map[string]interface{} {
	r := redis.Get()
	defer r.Close()

	str, _ := r.Get(CreateTokenKey(username, token))

	if str == "" {
		return nil
	}
	return other.Json2Map(str)
}
