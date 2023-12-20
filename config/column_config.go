package config

import (
	"fmt"
	"gotable/common"
	"os"
	"strings"
)

// StructColumn 结构体内字段配置
type StructColumn struct {
	Order            string //排序方式 字典顺序/数据库字段建立顺序
	NameCaseType     string //字段名命名规则 CAMEL_CASE骆驼命名/FIRST_UPPER首字母大写
	ColumnNameSuffix string //生成模型行后缀
	ColumnNamePrefix string //生成模型行前缀

	ModelOrmTagType string   //生成orm结构体标签类型
	TagType         bool     //是否生成 type 标签
	IsNull          bool     //是否生成 NOT NULL 标签
	Primary         bool     //是否生成主键标签中的 PRIMARY KEY
	OtherTag        []string //其他标签

	//数据库对应行信息
	Columns         []column
	MaxLenFieldType int
	MaxLenFieldTag  int
	MaxLenFieldName int
}

// column 列信息
type column struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	Type          string `gorm:"column:DATA_TYPE"`      //字段数据类型
	ColumnType    string `gorm:"column:COLUMN_TYPE"`    //字段类型(全)
	Nullable      string `gorm:"column:IS_NULLABLE"`    //是否可 null
	Default       string `gorm:"column:COLUMN_DEFAULT"` //默认值
	TableName     string `gorm:"column:TABLE_NAME"`     //表名
	ColumnComment string `gorm:"column:COLUMN_COMMENT"` //字段注释
	ColumnKey     string `gorm:"column:COLUMN_KEY"`     //索引

	FieldContent Field `gorm:"-"`
}

// Field 字段信息
type Field struct {
	Name    string //生成的字段名
	Type    string //生成的字段类型
	Tag     string //生成的字段tag
	Comment string //字段注释
}

// SetStructOrmTag 设置所生成对应的orm 标记类型  默认为 `gorm:"column:xxx"`
//	modelOrmTagType: 	"gorm"
//	tagType: 			是否生成 type 标签
//	primary: 			是否生成 PRIMARY KEY 标签
//	isNull: 			是否生成 NOT NULL 标签
func (ts *StructHandler) SetStructOrmTag(modelOrmTagType string, tagType, primary, isNull bool) *StructHandler {
	ts.StructColumn.ModelOrmTagType = modelOrmTagType
	//仅 gorm 生效
	ts.StructColumn.TagType = tagType
	ts.StructColumn.IsNull = isNull
	ts.StructColumn.Primary = primary
	return ts
}

// SetOtherTag 扩展标签 SetOtherTag("json","column"...)
func (ts *StructHandler) SetOtherTag(otherTag ...string) *StructHandler {
	if otherTag != nil {
		//添加其他的标签 如json ==> `json:"xxx"`
		ts.StructColumn.OtherTag = otherTag
	}
	return ts
}

// SetStructColumnName 设置行信息，结构体的属性配置
//	SetStructColumnName("命名规则","排序规则","前缀","后缀")
func (ts *StructHandler) SetStructColumnName(columnNameType, columnOrder, prefix, suffix string) *StructHandler {
	ts.StructColumn.NameCaseType = columnNameType
	ts.StructColumn.ColumnNameSuffix = suffix
	ts.StructColumn.ColumnNamePrefix = prefix
	ts.StructColumn.Order = columnOrder
	return ts
}

// SetColumns 字段设置
func (ts *StructHandler) SetColumns() {
	if ts.TableName == "" {
		fmt.Printf("\x1b[%dm------------->ERROR: [%s] \x1b[0m\n", 32, "请先调用SetTableName设置要生成结构的数据库表哦")
		os.Exit(2)
	}

	//查出所有字段信息
	db := ts.DB
	var cols []column
	qr := db.Table("information_schema.COLUMNS").
		Select("COLUMN_NAME, DATA_TYPE, COLUMN_TYPE, IS_NULLABLE, COLUMN_DEFAULT, TABLE_NAME, COLUMN_COMMENT, COLUMN_KEY").
		Where("table_schema = DATABASE()").
		Where("TABLE_NAME", ts.TableName)

	//设置字段排序
	switch ts.StructColumn.Order {
	case common.FieldOrderFieldName:
		qr.Order(common.FieldOrderFieldName).
			Find(&cols)
	case common.FieldOrderFollowDB:
		qr.Order(common.FieldOrderFollowDB).
			Find(&cols)
	case "":
		qr.Order("COLUMN_NAME").
			Find(&cols)
	default:
		qr.Order(ts.StructColumn.Order).
			Find(&cols)
	}
	if len(cols) < 1 {
		fmt.Printf("\x1b[%dm------------->ERROR: [%s] \x1b[0m\n", 32, "此表不存在或者数据库连接不正确请检查哦")
		os.Exit(1)
	}
	ts.StructColumn.MaxLenFieldName = 0
	ts.StructColumn.MaxLenFieldTag = 0
	ts.StructColumn.MaxLenFieldType = 0

	//组装要生成的字段信息
	var tsColumn []column
	for _, col := range cols {
		switch ts.TimeType {
		case common.TimeTypeString:
			switch col.Type {
			case "date", "datetime", "timestamp", "time":
				col.Type = fmt.Sprintf("%s_string", col.Type)
			}
		}
		var tag string
		switch ts.StructColumn.ModelOrmTagType {
		case common.ORM:
			tag = fmt.Sprintf("`orm:\"%s\"", col.ColumnName)
		case common.GORM:
			tag = fmt.Sprintf("`gorm:\"column:%s;", col.ColumnName)

			//生成 type
			if ts.StructColumn.TagType {
				tag += fmt.Sprintf(" type:%s;", col.ColumnType)
			}

			//生成 PRIMARY KEY
			if ts.StructColumn.Primary && col.ColumnKey == "PRI" {
				tag += fmt.Sprintf(" %s;", common.PRI)
			}

			//生成 not null
			if ts.StructColumn.IsNull && col.Nullable == "NO" {
				//不能为null
				//`gorm:"column:sn;type:BIGINT(19) UNSIGNED; PRIMARY_KEY; NOT NULL"`
				tag += fmt.Sprintf(" NOT NULL;")
			}

			tag += "\""
		default:
			tag = fmt.Sprintf("`gorm:\"column:%s\"", col.ColumnName)
		}
		for _, v := range ts.StructColumn.OtherTag {
			if v != "" {
				tag += fmt.Sprintf(" %s:\"%s\"", v, col.ColumnName)
			}
		}

		tag += "`"
		fieldName := fmt.Sprintf("%s%s%s",
			ts.StructColumn.ColumnNamePrefix,
			col.ColumnName,
			ts.StructColumn.ColumnNameSuffix,
		)
		fieldName = ts.GenerateChangeChara(fieldName, ts.StructColumn.NameCaseType)

		if len(fieldName) > ts.StructColumn.MaxLenFieldName {
			ts.StructColumn.MaxLenFieldName = len(fieldName)
		}
		if len(common.SqlTypeToGoType[col.Type]) > ts.StructColumn.MaxLenFieldType {
			ts.StructColumn.MaxLenFieldType = len(common.SqlTypeToGoType[col.Type])
		}
		if len(tag) > ts.StructColumn.MaxLenFieldTag {
			ts.StructColumn.MaxLenFieldTag = len(tag)
		}

		col.FieldContent = Field{
			Name:    fieldName,
			Type:    common.SqlTypeToGoType[col.Type],
			Tag:     tag,
			Comment: fmt.Sprintf("// %s", col.ColumnComment),
		}

		// col.ColunmContent = fmt.Sprintf("%s %s %s//是否可空：%s %s\n",
		// 	fieldName,
		// 	sqlTypeToGoType[col.Type],
		// 	tag,
		// 	col.Nullable,
		// 	col.ColumnComment)
		tsColumn = append(tsColumn, col)

		// ts.columns = append(ts.columns, col)
	}

	ts.StructColumn.Columns = tsColumn

}

//GenerateChangeChara 设置命名规则
func (ts *StructHandler) GenerateChangeChara(str string, Type string) string {

	var text string
	//不开启字段转为骆驼写法则仅仅将首字母大写
	switch Type {

	case common.CamelCase:
		for _, p := range strings.Split(str, "_") {
			// 考虑不规范字段定义,只保证不报错，不矫正字段 比如 ID_、NAME__AAA
			if p == "" {
				continue
			}
			text += strings.ToUpper(p[0:1]) + p[1:]
		}

	case common.FirstUpper:
		text += strings.ToUpper(str[0:1]) + strings.ToLower(str[1:])

	case common.LowAfterCamelCase:
		// 先全转小写
		lower := strings.ToLower(str)

		// 驼峰
		for _, p := range strings.Split(lower, "_") {
			// 考虑不规范字段定义,只保证不报错，不矫正字段 比如 ID_、NAME__AAA
			if p == "" {
				continue
			}
			text += strings.ToUpper(p[0:1]) + p[1:]
		}

	default:
		for _, p := range strings.Split(str, "_") {
			// 考虑不规范字段定义,只保证不报错，不矫正字段 比如 ID_、NAME__AAA
			if p == "" {
				continue
			}
			text += strings.ToUpper(p[0:1]) + p[1:]
		}
	}
	return text
}
