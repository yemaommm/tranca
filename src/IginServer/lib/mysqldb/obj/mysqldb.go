package obj

import (
	"IginServer/conf"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/widuu/goini"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var TxList = make(map[*Transaction]int64)

type MyDB struct {
	Conn       *sql.DB
	driver, co string
}

func execute(tdb interface{}, sqls string, params ...interface{}) ([]map[string]string, error) {
	var rows *sql.Rows
	var err error
	switch tdb.(type) {
	case *sql.DB:
		rows, err = tdb.(*sql.DB).Query(sqls, params...)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer rows.Close()
	case *sql.Tx:
		rows, err = tdb.(*sql.Tx).Query(sqls, params...)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer rows.Close()
	default:
		fmt.Println("db type error")
	}

	var result = make([]map[string]string, 0)
	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanargs := make([]interface{}, len(values))
	for i := range values {
		scanargs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanargs...)

		if err != nil {
			fmt.Println(err)
		}
		sdata := make(map[string]string)
		for i, v := range values {
			sdata[columns[i]] = string(v)
		}
		result = append(result, sdata)
	}
	return result, err
}

func update(tdb interface{}, sqls string, params ...interface{}) (int64, error) {
	var res sql.Result
	var err error
	switch tdb.(type) {
	case *sql.DB:
		res, err = tdb.(*sql.DB).Exec(sqls, params...)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	case *sql.Tx:
		res, err = tdb.(*sql.Tx).Exec(sqls, params...)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	default:
		fmt.Println("db type error")
	}
	i, _ := res.RowsAffected()
	return i, err
}

func del(tdb interface{}, sqls string, params ...interface{}) (int64, error) {
	var res sql.Result
	var err error
	switch tdb.(type) {
	case *sql.DB:
		res, err = tdb.(*sql.DB).Exec(sqls, params...)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	case *sql.Tx:
		res, err = tdb.(*sql.Tx).Exec(sqls, params...)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	default:
		fmt.Println("db type error")
	}
	i, _ := res.RowsAffected()
	return i, err
}

func insert(tdb interface{}, sqls string, params ...interface{}) (int64, error) {
	var res sql.Result
	var err error
	switch tdb.(type) {
	case *sql.DB:
		res, err = tdb.(*sql.DB).Exec(sqls, params...)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	case *sql.Tx:
		res, err = tdb.(*sql.Tx).Exec(sqls, params...)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	default:
		fmt.Println("db type error")
	}
	i, _ := res.LastInsertId()
	return i, err
}

func insertone(tdb interface{}, table string, m map[string]interface{}) (int64, error) {
	var res sql.Result
	var err error
	params := make([]interface{}, 0)
	var columns []string
	cs := ""
	vs := ""
	for i, j := range m {
		params = append(params, j)
		columns = append(columns, i)
		cs += ", " + i
		vs += ", ?"
	}
	update_sql := "INSERT INTO %s (%s) VALUES (%s)"
	update_sql = fmt.Sprintf(update_sql, table, cs[2:], vs[2:])
	fmt.Println(update_sql)

	switch tdb.(type) {
	case *sql.DB:
		res, err = tdb.(*sql.DB).Exec(update_sql, params...)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	case *sql.Tx:
		res, err = tdb.(*sql.Tx).Exec(update_sql, params...)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	default:
		fmt.Println("db type error")
	}
	count, _ := res.LastInsertId()

	return count, err
}

func insertmanay(tdb interface{}, table string, m ...map[string]interface{}) (int64, error) {
	var res *sql.Stmt
	var err error

	var columns []string
	cs := ""
	vs := ""
	for i, _ := range m[0] {
		columns = append(columns, i)
		cs += ", " + i
		vs += ", ?"
	}
	update_sql := "INSERT INTO %s (%s) VALUES (%s)"
	update_sql = fmt.Sprintf(update_sql, table, cs[2:], vs[2:])

	switch tdb.(type) {
	case *sql.DB:
		res, err = tdb.(*sql.DB).Prepare(update_sql)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		defer res.Close()
	case *sql.Tx:
		res, err = tdb.(*sql.Tx).Prepare(update_sql)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		defer res.Close()
	default:
		fmt.Println("db type error")
	}
	// var sr sql.Result
	var count int64
	count = 0
	for _, j := range m {
		var values []interface{}
		for _, v := range columns {
			values = append(values, j[v])

		}
		_, err = res.Exec(values...)
		if err != nil {
			fmt.Println(err)
			return count, err
		}
		count += 1
	}
	return count, err
}

func (db *MyDB) bagin() (*Transaction, error) {
	b, err := db.Conn.Begin()
	return &Transaction{b, "", []string{}, "", "", "", "", "", ""}, err
}

func (db *MyDB) close() {
	db.Conn.Close()
}

func Create(driver, con string, max_idle, max_open int) *MyDB {
	conn, err := sql.Open(driver, con)
	if err != nil {
		fmt.Println(err)
	}
	if max_idle != 0 {
		conn.SetMaxIdleConns(max_idle)
	}
	if max_open != 0 {
		conn.SetMaxOpenConns(max_open)
	}
	return &MyDB{conn, driver, con}
}

func (db *MyDB) getConn() *MyDB {
	// conn, err := sql.Open(db.driver, db.co)
	// db.conn = conn
	return db
}

type DB struct {
	*MyDB
	tablename string
	param     []string
	columnstr string
	where     string
	pk        string
	orderby   string
	limit     string
	join      string
}

func CreateMyDB(db *MyDB) *DB {
	return &DB{db, "", []string{}, "", "", "", "", "", ""}
}

// gomysql
func (m *DB) FindAll() map[int]map[string]string {

	result := make(map[int]map[string]string)
	if m.Conn == nil {
		fmt.Printf("mysql not connect")
		return result
	}
	if len(m.param) == 0 {
		m.columnstr = "*"
	} else {
		if len(m.param) == 1 {
			m.columnstr = m.param[0]
		} else {
			m.columnstr = strings.Join(m.param, ",")
		}

	}

	query := fmt.Sprintf("Select %v from %v %v %v %v %v", m.columnstr, m.tablename, m.join, m.where, m.orderby, m.limit)
	fmt.Println(query)
	rows, err := m.Conn.Query(query)
	// fmt.Println(query)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				// fmt.Printf("SQL syntax errors ")
				fmt.Println(query)
				fmt.Println(err)
			}
		}()
		err = errors.New("select sql failure")
	}
	result = QueryResult(rows)
	return result
}

func (m *DB) FindOne() map[int]map[string]string {
	empty := make(map[int]map[string]string)
	if m.Conn != nil {
		data := m.Limit(1).FindAll()

		return data
	}
	fmt.Printf("mysql not connect\r\n")

	return empty
}

func (m *DB) Count() int {
	data := m.Fileds("COUNT(1) AS total").Limit(1).FindAll()

	count, _ := strconv.Atoi(data[1]["total"])
	return count
}

func (m *DB) Insert(param map[string]interface{}) (num int, err error) {
	if m.Conn == nil {

		return 0, errors.New("mysql not connect")
	}
	var keys []string
	var values []string
	if len(m.pk) != 0 {
		delete(param, m.pk)
	}
	for key, value := range param {
		keys = append(keys, key)
		switch value.(type) {
		case int, int64, int32:
			values = append(values, fmt.Sprintf("%v", value.(int)))
		case string:
			value = strings.Replace(value.(string), "\\", "\\\\", -1)
			value = strings.Replace(value.(string), "'", "\\'", -1)
			value = strings.Replace(value.(string), "\"", "\\\"", -1)
			values = append(values, value.(string))
		case float32, float64:
			values = append(values, strconv.FormatFloat(value.(float64), 'f', -1, 64))

		}

	}
	fileValue := "'" + strings.Join(values, "','") + "'"
	fileds := "`" + strings.Join(keys, "`,`") + "`"
	sql := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", m.tablename, fileds, fileValue)
	result, err := m.Conn.Exec(sql)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
				fmt.Println(sql)
				fmt.Println(err)
			}
		}()
		err = errors.New("inster sql failure")

		return 0, err
	}
	i, err := result.LastInsertId()
	s, _ := strconv.Atoi(strconv.FormatInt(i, 10))
	if err != nil {
		err = errors.New("insert failure")
	}

	return s, err

}

func (m *DB) Fileds(param ...string) *DB {
	m.param = param
	return m
}

func (m *DB) Update(param map[string]interface{}) (num int, err error) {
	if m.Conn == nil {

		return 0, errors.New("mysql not connect")
	}
	var setValue []string
	for key, value := range param {
		switch value.(type) {
		case int, int64, int32:
			set := fmt.Sprintf("`%v` = %v", key, value)
			setValue = append(setValue, set)
		case string:
			value = strings.Replace(value.(string), "\\", "\\\\", -1)
			value = strings.Replace(value.(string), "'", "\\'", -1)
			value = strings.Replace(value.(string), "\"", "\\\"", -1)
			set := fmt.Sprintf("`%v` = '%v'", key, value.(string))
			setValue = append(setValue, set)
		case float32, float64:
			set := fmt.Sprintf("`%v` = '%v'", key, strconv.FormatFloat(value.(float64), 'f', -1, 64))
			setValue = append(setValue, set)
		}

	}
	setData := strings.Join(setValue, ",")
	sql := fmt.Sprintf("UPDATE %v SET %v %v", m.tablename, setData, m.where)
	fmt.Println(sql)
	result, err := m.Conn.Exec(sql)
	// fmt.Println(sql)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
				fmt.Println(sql)
				fmt.Println(err)
			}
		}()
		err = errors.New("update sql failure")

		return 0, err
	}
	i, err := result.RowsAffected()
	if err != nil {
		fmt.Println(sql)
		fmt.Println(err)
		err = errors.New("update failure")

		return 0, err
	}
	s, _ := strconv.Atoi(strconv.FormatInt(i, 10))

	return s, err
}

func (m *DB) Delete(param string) (num int, err error) {
	if m.Conn == nil {

		return 0, errors.New("mysql not connect")
	}
	h := m.Where(param).FindOne()
	if len(h) == 0 {

		return 0, errors.New("no Value")
	}
	sql := fmt.Sprintf("DELETE FROM %v WHERE %v", m.tablename, param)
	fmt.Println(sql)
	result, err := m.Conn.Exec(sql)
	// fmt.Println(sql)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
				fmt.Println(sql)
				fmt.Println(err)
			}
		}()
		err = errors.New("delete sql failure")

		return 0, err
	}
	i, err := result.RowsAffected()
	s, _ := strconv.Atoi(strconv.FormatInt(i, 10))
	if i == 0 {
		err = errors.New("delete failure")
	}

	return s, err
}

func (m *DB) Query(sql string) interface{} {
	if m.Conn == nil {

		return errors.New("mysql not connect")
	}
	var query = strings.TrimSpace(sql)
	s, err := regexp.MatchString(`(?i)^select`, query)
	if err == nil && s == true {
		result, eee := m.Conn.Query(sql)
		if eee != nil {
			fmt.Println(eee)
			return eee
		}
		c := QueryResult(result)

		return c
	}
	exec, err := regexp.MatchString(`(?i)^(update|delete)`, query)
	if err == nil && exec == true {
		m_exec, err := m.Conn.Exec(query)
		if err != nil {
			fmt.Println(sql)
			fmt.Println(err)

			return err
		}
		num, _ := m_exec.RowsAffected()
		id := strconv.FormatInt(num, 10)

		return id
	}

	insert, err := regexp.MatchString(`(?i)^insert`, query)
	if err == nil && insert == true {
		m_exec, err := m.Conn.Exec(query)
		if err != nil {
			fmt.Println(sql)
			fmt.Println(err)

			return err
		}
		num, _ := m_exec.LastInsertId()
		id := strconv.FormatInt(num, 10)

		return id
	}
	result, _ := m.Conn.Exec(query)

	return result

}

func QueryResult(rows *sql.Rows) map[int]map[string]string {
	var result = make(map[int]map[string]string)
	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanargs := make([]interface{}, len(values))
	for i := range values {
		scanargs[i] = &values[i]
	}

	var n = 1
	for rows.Next() {
		result[n] = make(map[string]string)
		err := rows.Scan(scanargs...)

		if err != nil {
			fmt.Println(err)
		}

		for i, v := range values {
			result[n][columns[i]] = string(v)
		}
		n++
	}

	return result
}

func (m *DB) SetTable(tablename string) *DB {
	m.Clean()
	m.tablename = tablename
	return m
}

func (m *DB) Where(param string) *DB {
	m.where = fmt.Sprintf(" where %v", param)
	return m
}

func (m *DB) SetPk(pk string) *DB {
	m.pk = pk
	return m
}

func (m *DB) OrderBy(param string) *DB {
	m.orderby = fmt.Sprintf("ORDER BY %v", param)
	return m
}

func (m *DB) Limit(size ...int) *DB {
	var end int
	start := size[0]
	if len(size) > 1 {
		end = size[1]
		m.limit = fmt.Sprintf("Limit %d,%d", start, end)
		return m
	}
	m.limit = fmt.Sprintf("Limit %d", start)
	return m
}

func (m *DB) LeftJoin(table, condition string) *DB {
	m.join = fmt.Sprintf("LEFT JOIN %v ON %v", table, condition)
	return m
}

func (m *DB) RightJoin(table, condition string) *DB {
	m.join = fmt.Sprintf("RIGHT JOIN %v ON %v", table, condition)
	return m
}

func (m *DB) Join(table, condition string) *DB {
	m.join = fmt.Sprintf("INNER JOIN %v ON %v", table, condition)
	return m
}

func (m *DB) FullJoin(table, condition string) *DB {
	m.join = fmt.Sprintf("FULL JOIN %v ON %v", table, condition)
	return m
}

func (m *DB) Clean() {
	m.tablename = ""
	m.columnstr = ""
	m.where = ""
	m.pk = ""
	m.orderby = ""
	m.limit = ""
	m.join = ""
	m.param = []string{}
}

//the function will use friendly way to print the data
func Print(slice map[int]map[string]string) {
	for _, v := range slice {
		for key, value := range v {
			fmt.Println(key, value)
		}
		fmt.Println("---------------")
	}
}

func (m *DB) DbClose() {
	m.Conn.Close()
}

func (d *DB) IGet() *sql.DB {
	return d.Conn
}

func (d *DB) IExecute(s string, params ...interface{}) ([]map[string]string, error) {
	fmt.Println(s, params)
	return execute(d.Conn, s, params...)
}

func (d *DB) IExecuteXML(s string, params ...interface{}) ([]map[string]string, error) {
	fmt.Println(conf.XmlGet(s), params)
	return execute(d.Conn, conf.XmlGet(s), params...)
}

func (d *DB) IUpdate(s string, params ...interface{}) (int64, error) {
	return update(d.Conn, s, params...)
}

func (d *DB) IUpdateXML(s string, params ...interface{}) (int64, error) {
	return update(d.Conn, conf.XmlGet(s), params...)
}

func (d *DB) IDel(s string, params ...interface{}) (int64, error) {
	return del(d.Conn, s, params...)
}

func (d *DB) IDelXML(s string, params ...interface{}) (int64, error) {
	return del(d.Conn, conf.XmlGet(s), params...)
}

func (d *DB) IInsert(s string, params ...interface{}) (int64, error) {
	return insert(d.Conn, s, params...)
}

func (d *DB) IInsertXML(s string, params ...interface{}) (int64, error) {
	return insert(d.Conn, conf.XmlGet(s), params...)
}

func (d *DB) IInsertOne(table string, m map[string]interface{}) (int64, error) {
	return insertone(d.Conn, table, m)
}

func (d *DB) IInsertmanay(table string, m ...map[string]interface{}) (int64, error) {
	return insertmanay(d.Conn, table, m...)
}

func (db *DB) IBagin() (*Transaction, error) {
	i, err := db.bagin()
	TxList[i] = time.Now().Unix()
	return i, err
}
