package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type StructHandler struct {
	DB *gorm.DB //数据库连接

	TableName  string   //要生成model的数据库表名
	TableNames []string //要生成model的数据库表名
	SavePath   string   //保存model文件的位置
	TimeType   string   //时间类型对应go类型 string/time.Time

	Package      Package      //模型文件包名配置
	StructName   StructName   //结构体模型名配置
	StructColumn StructColumn //结构体内容配置
}

// -------------------------------------------------- 连接数据库(3种方式) ---------------------------------------------------------

// SetConnectByDSN 通过 DSN 链接数据库
// 	root:123456@(127.0.0.1:3306)/mysqlDB?charset=utf8mb4&parseTime=True&loc=Local
func (ts *StructHandler) SetConnectByDSN(dsn string) *StructHandler {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("\x1b[%dm------------->ERROR: [%s] \x1b[0m\n", 32, err)
		os.Exit(4)
	}
	ts.DB = db
	return ts
}

// SetConnectByParam 配置数据库链接
// 	User:Password@(Host:Port)/DB?config
// 	用户:密码@(IP:端口)/数据库?其他配置
func (ts *StructHandler) SetConnectByParam(User string, Password string, Host string, Port string, DB string, config string) *StructHandler {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", User, Password, Host, Port, DB, config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("\x1b[%dm------------->ERROR: [%s] \x1b[0m\n", 32, err)
		os.Exit(3)
	}
	ts.DB = db
	return ts
}

// SetConnectByDB 直接设置DB，自己写连接数据库的方法，连接成功后把DB传进来
func (ts *StructHandler) SetConnectByDB(db *gorm.DB) *StructHandler {
	ts.DB = db
	return ts
}

// ---------------------------------------------------------------------------------------------------------------------

// SetSavePath 设置生成模型的保存位置
func (ts *StructHandler) SetSavePath(savePath string) *StructHandler {
	ts.SavePath = savePath
	return ts
}

// SetTableName 设置要生成哪张数据库表的结构
func (ts *StructHandler) SetTableName(tableName string) *StructHandler {
	ts.TableName = tableName
	return ts
}

// SetTableNames 设置要生成哪些数据库表的结构(多表)
func (ts *StructHandler) SetTableNames(tableName []string) *StructHandler {
	ts.TableNames = tableName
	return ts
}

// SetTimeType 设置数据库中的时间类型对应 struct 中的什么 time.Time/string
func (ts *StructHandler) SetTimeType(timeType string) *StructHandler {
	ts.TimeType = timeType
	return ts
}
