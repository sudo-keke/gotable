# gotable ⚡



### 🔨 生成数据库对应的go模型文件


---

点击此处打开测试文件  `-->`   [✅ main.go](./main/main.go)

🌱 使用方法： 复制以下代码片段即可
    

### ⏩ 单表生成（示例）
```go
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
		SetStructNameConfig(StructName, common.CamelCase).
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
```

---

### ⏩ 多表生成（示例）

```go

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
```


---

🚥 有问题可以提 issues
