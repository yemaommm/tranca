package obj

import (
	"IginServer/conf"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// import "log"

type Transaction struct {
	tx        *sql.Tx
	tablename string
	param     []string
	columnstr string
	where     string
	pk        string
	orderby   string
	limit     string
	join      string
}

func (d *Transaction) IGet() *sql.Tx {
	return d.tx
}

func (d *Transaction) IExecute(s string, params ...interface{}) ([]map[string]string, error) {
	return execute(d.tx, s, params...)
}

func (d *Transaction) IExecuteXML(s string, params ...interface{}) ([]map[string]string, error) {
	return execute(d.tx, conf.XmlGet(s), params...)
}

func (d *Transaction) IUpdate(s string, params ...interface{}) (int64, error) {
	return update(d.tx, s, params...)
}

func (d *Transaction) IUpdateXML(s string, params ...interface{}) (int64, error) {
	return update(d.tx, conf.XmlGet(s), params...)
}

func (d *Transaction) IDel(s string, params ...interface{}) (int64, error) {
	return del(d.tx, s, params...)
}

func (d *Transaction) IDelXML(s string, params ...interface{}) (int64, error) {
	return del(d.tx, conf.XmlGet(s), params...)
}

func (d *Transaction) IInsert(s string, params ...interface{}) (int64, error) {
	return insert(d.tx, s, params...)
}

func (d *Transaction) IInsertXML(s string, params ...interface{}) (int64, error) {
	return insert(d.tx, conf.XmlGet(s), params...)
}

func (d *Transaction) IInsertOne(table string, m map[string]interface{}) (int64, error) {
	return insertone(d.tx, table, m)
}

func (d *Transaction) IInsertmanay(table string, m ...map[string]interface{}) (int64, error) {
	return insertmanay(d.tx, table, m...)
}

func (d *Transaction) Commit() (err error) {
	err = d.tx.Commit()
	delete(TxList, d)
	return
}

func (d *Transaction) Rollback() (err error) {
	err = d.tx.Rollback()
	delete(TxList, d)
	return
}

// gomysql
func (m *Transaction) FindAll() map[int]map[string]string {

	result := make(map[int]map[string]string)
	if m.tx == nil {
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
	rows, err := m.tx.Query(query)
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

func (m *Transaction) FindOne() map[int]map[string]string {
	empty := make(map[int]map[string]string)
	if m.tx != nil {
		data := m.Limit(1).FindAll()

		return data
	}
	fmt.Printf("mysql not connect\r\n")

	return empty
}

func (m *Transaction) Insert(param map[string]interface{}) (num int, err error) {
	if m.tx == nil {

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
			values = append(values, value.(string))
		case float32, float64:
			values = append(values, strconv.FormatFloat(value.(float64), 'f', -1, 64))

		}

	}
	fileValue := "'" + strings.Join(values, "','") + "'"
	fileds := "`" + strings.Join(keys, "`,`") + "`"
	sql := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", m.tablename, fileds, fileValue)
	fmt.Println(sql)
	result, err := m.tx.Exec(sql)
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
		fmt.Println(sql)
		fmt.Println(err)
	}

	return s, err

}

func (m *Transaction) Fileds(param ...string) *Transaction {
	m.param = param
	return m
}

func (m *Transaction) Update(param map[string]interface{}) (num int, err error) {
	if m.tx == nil {

		return 0, errors.New("mysql not connect")
	}
	var setValue []string
	for key, value := range param {
		switch value.(type) {
		case int, int64, int32:
			set := fmt.Sprintf("`%v` = %v", key, value)
			setValue = append(setValue, set)
		case string:
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
	result, err := m.tx.Exec(sql)
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
		err = errors.New("update failure")

		return 0, err
	}
	s, _ := strconv.Atoi(strconv.FormatInt(i, 10))

	return s, err
}

func (m *Transaction) Delete(param string) (num int, err error) {
	if m.tx == nil {

		return 0, errors.New("mysql not connect")
	}
	h := m.Where(param).FindOne()
	if len(h) == 0 {

		return 0, errors.New("no Value")
	}
	sql := fmt.Sprintf("DELETE FROM %v WHERE %v", m.tablename, param)
	fmt.Println(sql)
	result, err := m.tx.Exec(sql)
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

func (m *Transaction) Query(sql string) interface{} {
	if m.tx == nil {

		return errors.New("mysql not connect")
	}
	var query = strings.TrimSpace(sql)
	s, err := regexp.MatchString(`(?i)^select`, query)
	if err == nil && s == true {
		result, _ := m.tx.Query(sql)
		c := QueryResult(result)

		return c
	}
	exec, err := regexp.MatchString(`(?i)^(update|delete)`, query)
	if err == nil && exec == true {
		m_exec, err := m.tx.Exec(query)
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
		m_exec, err := m.tx.Exec(query)
		if err != nil {
			fmt.Println(sql)
			fmt.Println(err)

			return err
		}
		num, _ := m_exec.LastInsertId()
		id := strconv.FormatInt(num, 10)

		return id
	}
	result, _ := m.tx.Exec(query)

	return result

}

func (m *Transaction) Count() int {
	data := m.Fileds("COUNT(1) AS total").Limit(1).FindAll()

	count, _ := strconv.Atoi(data[1]["total"])
	return count
}

func (m *Transaction) SetTable(tablename string) *Transaction {
	m.Clean()
	m.tablename = tablename
	return m
}

func (m *Transaction) Where(param string) *Transaction {
	m.where = fmt.Sprintf(" where %v", param)
	return m
}

func (m *Transaction) SetPk(pk string) *Transaction {
	m.pk = pk
	return m
}

func (m *Transaction) OrderBy(param string) *Transaction {
	m.orderby = fmt.Sprintf("ORDER BY %v", param)
	return m
}

func (m *Transaction) Limit(size ...int) *Transaction {
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

func (m *Transaction) LeftJoin(table, condition string) *Transaction {
	m.join = fmt.Sprintf("LEFT JOIN %v ON %v", table, condition)
	return m
}

func (m *Transaction) RightJoin(table, condition string) *Transaction {
	m.join = fmt.Sprintf("RIGHT JOIN %v ON %v", table, condition)
	return m
}

func (m *Transaction) Join(table, condition string) *Transaction {
	m.join = fmt.Sprintf("INNER JOIN %v ON %v", table, condition)
	return m
}

func (m *Transaction) FullJoin(table, condition string) *Transaction {
	m.join = fmt.Sprintf("FULL JOIN %v ON %v", table, condition)
	return m
}

func (m *Transaction) Clean() {
	m.tablename = ""
	m.columnstr = ""
	m.where = ""
	m.pk = ""
	m.orderby = ""
	m.limit = ""
	m.join = ""
	m.param = []string{}
}
