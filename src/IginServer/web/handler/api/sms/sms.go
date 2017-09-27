package sms

import (
	"fmt"
	// "net/http"
	"IginServer/framework/API/R"
	// "IginServer/lib/Imartini"
	"IginServer/lib/md5"
	"IginServer/lib/other"
	"IginServer/lib/redis"
	// "IginServer/lib/session"
	"github.com/gin-gonic/gin"
	"time"
)

/**
 * @api {post} /api/sms/send 短信发送接口
 * @apiName 短信发送接口
 * @apiGroup 短信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} phone 手机号
 * @apiParam {String} mode 发送的类型，比如register
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
 *
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
 * @apiSampleRequest http://121.41.116.104:8003/api/sms/send
 */
func SendMSG(c *gin.Context) {
	phone := c.PostForm("phone")
	mode := c.PostForm("mode")
	r := redis.Get()
	defer r.Close()

	if phone == "" || mode == "" {
		R.Api404(c.Writer)
		return
	}
	fmt.Println(phone)
	if Send(phone, mode, r) == "" {
		R.Error(c.Writer, 201, "调用太频繁")
	} else {
		R.Success(c.Writer, "")
	}
}

func Send(phone, mode string, r *redis.Redis) string {
	if !other.IsPhone(phone) {
		return "-1"
	}
	i, _ := r.TTL("code:" + mode + ":" + phone)
	if 600-i < 10 {
		return ""
	}
	code := fmt.Sprintf("%04d", time.Now().Unix()*3%10000)
	v, _ := r.Get("code:" + mode + ":" + phone)
	if v != "" {
		code = v
	}

	key := "16f2ffeb7d3e410abe10be54a25ddcb9%s,SHTsms,"
	url := "http://www.shanghaitong.biz/api/v1/sms/sendSMS?MOBILE_PHONE=%s&VALID_CODE=%s&SHTKEY=%s"

	y, m, d := time.Now().Date()
	now := fmt.Sprintf("%d%02d%02d", y, m, d)

	key = fmt.Sprintf(key, now)
	url = fmt.Sprintf(url, phone, code, md5.Md5(key))

	bret := string(other.Httpget(url))
	fmt.Println(bret, code)

	r.Set("code:"+mode+":"+phone, code)
	r.Expire("code:"+mode+":"+phone, 600)
	return bret
}

func Validate(phone, mode, code string, r *redis.Redis) int {
	v, _ := r.Get("code:" + mode + ":" + phone)
	if v == code {
		r.Del("code:" + mode + ":" + phone)
		return 0
	} else {
		return 1
	}
}
