package config

// StructName struct的名称配置
type StructName struct {
	Name         string //struct 名称
	Prefix       string //struct 名称前缀
	Suffix       string //struct 名称后缀
	NameCaseType string //struct 名称命名规则 CAMEL_CASE骆驼命名/FIRST_UPPER首字母大写
}

// --------------------------------------------------- 结构体配置 --------------------------------------------------------

// SetManyStructNameConfig 结构名称配置，多表生成时使用
// SetManyStructNameConfig("命名规则","前缀","后缀")
// 不设置此项，默认生成的结构名类型为CamelCase写法，无前后缀
func (ts *StructHandler) SetManyStructNameConfig(nameCaseType, prefix, suffix string) *StructHandler {
	ts.StructName.NameCaseType = nameCaseType
	ts.StructName.Prefix = prefix
	ts.StructName.Suffix = suffix
	return ts
}

// SetStructNameConfig 结构名称配置，已简化，无前后缀,单表生成时使用
// SetStructNameConfig("结构名称","命名规则")
// 不设置此项，默认生成的结构名类型为CamelCase写法
func (ts *StructHandler) SetStructNameConfig(nameCaseType, structName string) *StructHandler {
	ts.StructName.Name = structName
	ts.StructName.NameCaseType = nameCaseType
	ts.StructName.Prefix = ""
	ts.StructName.Suffix = ""
	return ts
}

// --------------------------------------------------- 结构体处理 --------------------------------------------------------
