package common

// SqlTypeToGoType mysql数据类型<=>go数据类型
var SqlTypeToGoType = map[string]string{
	"int":                "int8",
	"integer":            "int8",
	"tinyint":            "int8",
	"smallint":           "int",
	"mediumint":          "int",
	"bigint":             "int64",
	"int unsigned":       "int8",
	"integer unsigned":   "int8",
	"tinyint unsigned":   "int8",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int",
	"bigint unsigned":    "int64",
	"bit":                "int64",
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"binary":             "string",
	"varbinary":          "string",
	"json":               "string",
	"date":               "time.Time",
	"datetime":           "time.Time",
	"timestamp":          "time.Time",
	"time":               "time.Time",
	"date_string":        "string", // time.Time
	"datetime_string":    "string", // time.Time
	"timestamp_string":   "string", // time.Time
	"time_string":        "string", // time.Time

	//canNullAble
	"point_int":                "int",
	"point_integer":            "int",
	"point_tinyint":            "int",
	"point_smallint":           "int",
	"point_mediumint":          "int",
	"point_bigint":             "int",
	"point_int unsigned":       "int",
	"point_integer unsigned":   "int",
	"point_tinyint unsigned":   "int",
	"point_smallint unsigned":  "int",
	"point_mediumint unsigned": "int",
	"point_bigint unsigned":    "int",
	"point_bit":                "int",
	"point_float":              "float64",
	"point_double":             "float64",
	"point_decimal":            "float64",
	"point_bool":               "bool",
	"point_enum":               "string",
	"point_set":                "string",
	"point_varchar":            "string",
	"point_char":               "string",
	"point_tinytext":           "string",
	"point_mediumtext":         "string",
	"point_text":               "string",
	"point_longtext":           "string",
	"point_blob":               "string",
	"point_tinyblob":           "string",
	"point_mediumblob":         "string",
	"point_longblob":           "string",
	"point_binary":             "string",
	"point_varbinary":          "string",
	"point_json":               "string",

	"point_date":      "time.Time", //
	"point_datetime":  "time.Time", //
	"point_timestamp": "time.Time", //
	"point_time":      "time.Time", //

	"point_date_string":      "string", // time.Time
	"point_datetime_string":  "string", // time.Time
	"point_timestamp_string": "string", // time.Time
	"point_time_string":      "string", // time.Time

}

const (
	CamelCase  = "CamelCase"   // CamelCase 驼峰
	FirstUpper = "First_upper" // FirstUpper 首字母大写

	ORM  = "orm"
	GORM = "gorm"
	JSON = "json"

	TimeTypeString = "string" // TimeTypeString 时间对应-string
	TimeTypeTime   = "time"   // TimeTypeTime 时间对应-time

	FieldOrderFieldName = "COLUMN_NAME"      //FieldOrderFieldName 字典顺序
	FieldOrderFollowDB  = "ORDINAL_POSITION" //FieldOrderFollowDB 数据库字段建立顺序

	PRI = "PRIMARY KEY" //主键索引
)
