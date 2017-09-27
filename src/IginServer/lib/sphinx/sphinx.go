package sphinx

import (
	"fmt"
	"github.com/yunge/sphinx"
	"server/conf"
	"server/lib/mysqldb/obj"
	"strconv"
)

var (
	ISPHINX, _  = strconv.ParseBool(conf.GET["config"]["SPHINX"])
	host        = conf.GET["config"]["SPHINX_HOST"]
	port, _     = strconv.Atoi(conf.GET["config"]["SPHINX_PORT"])
	timeout, _  = strconv.Atoi(conf.GET["config"]["SPHINX_TIMEOUT"])
	sqlport, _  = strconv.Atoi(conf.GET["config"]["SPHINX_SQLPORT"])
	max_idle, _ = strconv.Atoi(conf.GET["config"]["SPHINX_MAX_IDLE"])
	max_open, _ = strconv.Atoi(conf.GET["config"]["SPHINX_MAX_OPEN"])
)

// var search *Isphinx
var searchql = obj.Create("mysql", "tcp("+host+":"+conf.GET["config"]["SPHINX_SQLPORT"]+")/", max_idle, max_open)

func GetSearchql() *obj.DB {
	return obj.CreateMyDB(searchql)
}

func Create() *Isphinx {
	// 链接参数
	opts := &sphinx.Options{
		// Limit: 20,
		Host:    host,
		Port:    port,
		SqlPort: sqlport,
		// MaxMatches: 1000,
		MatchMode: sphinx.SPH_MATCH_EXTENDED2, //SPH_MATCH_ANY,
		// Index: "test1",
		// Socket:  "/var/run/searchd.sock",
		Timeout: timeout,
	}

	// 创建客户端
	spClient := sphinx.NewClient(opts)
	// defer spClient.Close()
	if err := spClient.Error(); err != nil {
		fmt.Printf("err1:%v", err)
	}

	// 打开链接
	if err := spClient.Open(); err != nil {
		fmt.Printf("err2:%v", err)
	}
	ret := &Isphinx{spClient}
	ret.SetLimits(0, 100, 1000, 0)
	return ret
}

func GET(ret *sphinx.Result) []map[string]string {
	AttrNames := ret.AttrNames
	Matches := ret.Matches
	iret := make([]map[string]string, len(Matches))
	for i, j := range Matches {
		iret[i] = map[string]string{}
		stmp := j.AttrValues
		for x, y := range AttrNames {
			iret[i][y] = fmt.Sprintf("%s", stmp[x])
		}
	}
	return iret
}

type Isphinx struct {
	*sphinx.Client
}

func GetClient() *Isphinx {
	// return Create()
	return Create()
}

func init() {
	// if ISPHINX {
	// 	// search = Create()
	// }
}
