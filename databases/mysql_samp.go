package mysql_samp

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	addr       string
	user     string
	password string
	db       string
)

type MyDb struct {
	db *sql.DB
}

func Init(mAddr string, mUser string, mPassword string, mDb string) error {
	addr = mAddr
	user = mUser
	password = mPassword
	db = mDb

	_, err := GetConn()
	if err != nil {
		return err
	}
	return nil
}

func GetConn() (*MyDb, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", user, password, addr, db)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	//验证连接
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &MyDb{db: db}, nil
}

func GetGorm() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", user, password, addr, db)
	g, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	sqlDB := g.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	return g, err
}

func (m *MyDb) Query(SQL string) ([]map[string]string, error) {
	rows, err := m.db.Query(SQL)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i, _ := range values {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	ret := make([]map[string]string, 0) //创建返回值：不定长的map类型切片
	for rows.Next() {
		err := rows.Scan(values...)  //开始读行，Scan函数只接受指针变量
		m := make(map[string]string) //用于存放1列的 [键/值] 对
		if err != nil {
			logrus.Error(err)
		}
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{})) //读出raw数据，类型为byte
			b, _ := raw_value.([]byte)
			v := string(b) //将raw数据转换成字符串
			m[colName] = v //colName是键，v是值
		}
		ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
	}

	defer rows.Close()
	return ret, nil

}
