package config

//Package 文件最上面一行的包名称：package xxx
type Package struct {
	Name   string //包名
	Prefix string //包名前缀，无前缀设置为""
	Suffix string //包名后缀，无后缀设置为""
}

// ---------------------------------------------------- 包配置 ----------------------------------------------------------

// SetPackage 设置包(package)信息
// 	设置生成的model包名 SetPackage("包名","前缀","后缀")
func (ts *StructHandler) SetPackage(name, prefix, suffix string) *StructHandler {
	ts.Package.Name = name
	if prefix != "" {
		ts.Package.Prefix = prefix
	}
	if suffix != "" {
		ts.Package.Suffix = suffix
	}
	return ts
}
