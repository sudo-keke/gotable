package main

import (
	"gotable/common"
	"gotable/config"
)

const (
	user     = "root"
	password = "root"
	host     = "localhost"
	port     = "3306"
	db       = "rent"
)

// TableName 要生成的表名
var TableName = "citys"

// TableNames 多表生成的表名
var TableNames = []string{"citys", "steps"}

// StructName struct的名称
var StructName = "City"

// SavePackage 文件最上面一行的包名称：package xxx
var SavePackage = "model"

// SavePath 存在哪个包下，此处加不加[.go]后缀、包名后加不加[/] 都可以
var SavePath = "./" + SavePackage

func main() {
	Multiple()
	Single()
}

// Single 单表生成
func Single() {
	handler := config.NewTblToStructHandler()
	handler.
		//设置数据库dsn连接地址
		SetConnectByDSN(user+":"+password+"@("+host+":"+port+")/"+db+"?charset=utf8mb4&parseTime=True&loc=Local").
		//生成哪张数据库表的结构
		SetTableName(TableName).
		//文件最上面一行的包名称 SetPackage("包名","前缀","后缀")
		SetPackage(SavePackage, "", "").
		//保存到什么位置
		SetSavePath(SavePath).
		//单表可以指定生成的struct的名称和格式(单表生成时使用)
		SetStructNameConfig(common.CamelCase, StructName).
		//ORM标签信息, SetStructOrmTag("gorm","是否生成type标签","是否生成PRIMARY KEY标签","是否生成NOT NULL标签")
		SetStructOrmTag(common.GORM, true, true, true).
		//扩展标签信息
		SetOtherTag(common.JSON).
		//时间类型要生成什么格式
		SetTimeType(common.TimeTypeTime).
		//设置struct中字段的格式
		SetStructColumnName(common.CamelCase, common.FieldOrderFollowDB, "", "")

	//生成
	config.GenerateTblStruct(handler)

}

// Multiple 多表生成
func Multiple() {
	handler := config.NewTblToStructHandler()
	handler.
		//设置数据库dsn连接地址
		SetConnectByDSN(user+":"+password+"@("+host+":"+port+")/"+db+"?charset=utf8mb4&parseTime=True&loc=Local").
		//生成哪张数据库表的结构
		SetTableNames(TableNames).
		//文件最上面一行的包名称 SetPackage("包名","前缀","后缀")
		SetPackage(SavePackage, "", "").
		//保存到什么位置
		SetSavePath(SavePath).
		//多表生成不可指定struct的名称，但可以设置格式(多表生成时使用)
		SetManyStructNameConfig(common.CamelCase, "", "").
		//ORM标签信息, SetStructOrmTag("gorm","是否生成type标签","是否生成PRIMARY KEY标签","是否生成NOT NULL标签")
		SetStructOrmTag(common.GORM, true, true, true).
		//扩展标签信息
		SetOtherTag(common.JSON).
		//时间类型要生成什么格式
		SetTimeType(common.TimeTypeTime).
		//设置struct中字段的格式
		SetStructColumnName(common.CamelCase, common.FieldOrderFollowDB, "", "")

	//生成
	config.GenerateAllTblStruct(handler)
}
