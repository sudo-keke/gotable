package config

import (
	"fmt"
	"gotable/common"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// NewTblToStructHandler 初始化生成器
func NewTblToStructHandler() *StructHandler {
	return &StructHandler{
		Package: Package{
			Name:   "model",
			Suffix: "",
			Prefix: "",
		},
		StructName: StructName{
			Prefix:       "",
			Suffix:       "",
			NameCaseType: common.CamelCase,
		},
		StructColumn: StructColumn{
			ModelOrmTagType:  common.GORM,
			NameCaseType:     common.CamelCase,
			Order:            common.FieldOrderFollowDB,
			ColumnNamePrefix: "",
			ColumnNameSuffix: "",
		},

		SavePath: "./model",
	}
}

// GenerateTblStruct 单表生成
func GenerateTblStruct(handler *StructHandler) {
	fmt.Println("\r")
	handler.SetColumns()
	// 设置包信息
	packageName := fmt.Sprintf("package %s%s%s",
		handler.Package.Prefix,
		handler.Package.Name,
		handler.Package.Suffix,
	)
	var timeTypeCount int64
	handler.DB.Table("INFORMATION_SCHEMA.COLUMNS").
		Where("TABLE_SCHEMA=database()").
		Where("TABLE_NAME", handler.TableName).
		Where("DATA_TYPE in ?", []string{
			"date", "datetime", "timestamp", "time",
			"DATE", "DATETIME", "TIMESTAMP", "TIME",
		}).Count(&timeTypeCount)

	packageImport := ""
	if handler.TimeType != common.TimeTypeString && timeTypeCount > 0 {
		packageImport = "\nimport \"time\"\n"
	}

	var tableComment string
	// select * from INFORMATION_SCHEMA.TABLES where TABLE_SCHEMA=database()
	handler.DB.Table("INFORMATION_SCHEMA.TABLES").
		Select("TABLE_COMMENT").
		Where("TABLE_SCHEMA=database()").
		Where("TABLE_NAME", handler.TableName).Find(&tableComment)

	if handler.StructName.Name == "" {
		handler.StructName.Name = fmt.Sprintf("%s%s%s",
			handler.StructName.Prefix,
			handler.TableName,
			handler.StructName.Suffix,
		)
	} else {
		handler.StructName.Name = fmt.Sprintf("%s%s%s",
			handler.StructName.Prefix,
			handler.StructName.Name,
			handler.StructName.Suffix,
		)
	}

	structName := handler.GenerateChangeChara(handler.StructName.Name, handler.StructName.NameCaseType)
	tableComment = fmt.Sprintf("// %s %s\n", structName, tableComment)

	structContent := fmt.Sprintf("type %s struct {\n", structName)

	for _, v := range handler.StructColumn.Columns {
		match := fmt.Sprint("\t%-",
			handler.StructColumn.MaxLenFieldName, "s %-",
			handler.StructColumn.MaxLenFieldType, "s %-",
			handler.StructColumn.MaxLenFieldTag, "s %s\n")
		structContent += fmt.Sprintf(match, v.FieldContent.Name, v.FieldContent.Type, v.FieldContent.Tag, v.FieldContent.Comment)
	}
	structContent += "}\n"

	funcTableName := fmt.Sprintf("func (*%s) TableName() string {\n", structName) +
		fmt.Sprintf("\treturn \"%s\"\n", handler.TableName) +
		"}\n"

	fileContent := fmt.Sprintf("%s\n%s\n%s%s\n%s", packageName, packageImport, tableComment, structContent, funcTableName)

	//处理保存路径
	savePathDealWith(handler)

	filePath := fmt.Sprint(handler.SavePath)
	paths, _ := filepath.Split(handler.SavePath)

	err := os.MkdirAll(paths, os.ModePerm)
	if err != nil {
		fmt.Printf("\x1b[%dm------------->ERROR: 创建文件夹 [%s] 出错 \x1b[0m\n", 32, paths)
		os.Exit(5)
	}
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("\x1b[%dm------------->Table: %s 生成失败\x1b[0m\n", 31, handler.TableName)
		os.Exit(6)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("\x1b[%dm------------->ERROR: 关闭文件出错 \x1b[0m\n", 32)
			os.Exit(7)
		}
	}(f)

	_, err = f.WriteString(fileContent)
	if err != nil {
		fmt.Printf("\x1b[%dm------------->ERROR: 写入内容出错 \x1b[0m\n", 32)
		os.Exit(11)
	} else {
		fmt.Printf("\x1b[%dm------------->Table: [%s] 生成成功\x1b[0m\n", 32, handler.TableName)
		os.Exit(8)
	}

	return
}

// GenerateAllTblStruct 多表生成
func GenerateAllTblStruct(handler *StructHandler) {
	for _, tableName := range handler.TableNames {
		fmt.Println("\r")
		handler.TableName = tableName
		handler.SetColumns()
		// 设置包信息
		packageName := fmt.Sprintf("package %s%s%s",
			handler.Package.Prefix,
			handler.Package.Name,
			handler.Package.Suffix,
		)
		var timeTypeCount int64
		handler.DB.Table("INFORMATION_SCHEMA.COLUMNS").
			Where("TABLE_SCHEMA=database()").
			Where("TABLE_NAME", tableName).
			Where("DATA_TYPE in ?", []string{
				"date", "datetime", "timestamp", "time",
				"DATE", "DATETIME", "TIMESTAMP", "TIME",
			}).Count(&timeTypeCount)

		packageImport := ""
		if handler.TimeType != common.TimeTypeString && timeTypeCount > 0 {
			packageImport = "\nimport \"time\"\n"
		}

		var tableComment string
		// select * from INFORMATION_SCHEMA.TABLES where TABLE_SCHEMA=database()
		handler.DB.Table("INFORMATION_SCHEMA.TABLES").
			Select("TABLE_COMMENT").
			Where("TABLE_SCHEMA=database()").
			Where("TABLE_NAME", tableName).Find(&tableComment)

		handler.StructName.Name = fmt.Sprintf("%s%s%s",
			handler.StructName.Prefix,
			tableName,
			handler.StructName.Suffix,
		)

		structName := handler.GenerateChangeChara(handler.StructName.Name, handler.StructName.NameCaseType)

		tableComment = fmt.Sprintf("// %s %s\n", structName, tableComment)

		structContent := fmt.Sprintf("type %s struct {\n", structName)

		for _, v := range handler.StructColumn.Columns {
			match := fmt.Sprint("\t%-",
				handler.StructColumn.MaxLenFieldName, "s %-",
				handler.StructColumn.MaxLenFieldType, "s %-",
				handler.StructColumn.MaxLenFieldTag, "s %s\n")
			structContent += fmt.Sprintf(match, v.FieldContent.Name, v.FieldContent.Type, v.FieldContent.Tag, v.FieldContent.Comment)
		}
		structContent += "}\n"

		funcTableName := fmt.Sprintf("func (*%s) TableName() string {\n", structName) +
			fmt.Sprintf("\treturn \"%s\"\n", tableName) +
			"}\n"

		fileContent := fmt.Sprintf("%s\n%s\n%s%s\n%s", packageName, packageImport, tableComment, structContent, funcTableName)

		//处理保存路径
		savePathDealWith(handler)

		filePath := fmt.Sprint(handler.SavePath)
		paths, _ := filepath.Split(handler.SavePath)

		err := os.MkdirAll(paths, os.ModePerm)
		if err != nil {
			fmt.Printf("\x1b[%dm------------->Table: 创建文件夹 [%s] 出错 \x1b[0m\n", 32, paths)
			os.Exit(9)
		}
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("\x1b[%dm->table: %s 生成失败\x1b[0m\n", 31, tableName)
			os.Exit(10)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				fmt.Printf("\x1b[%dm------------->ERROR: 关闭文件出错 \x1b[0m\n", 32)
				os.Exit(7)
			}
		}(f)

		_, err = f.WriteString(fileContent)
		if err != nil {
			fmt.Printf("\x1b[%dm------------->ERROR: 写入内容出错 \x1b[0m\n", 32)
			os.Exit(11)
		} else {
			fmt.Printf("\x1b[%dm------------->Table: [%s] 生成成功\x1b[0m\n", 32, tableName)
		}
	}
	return
}

// savePathDealWith 处理保存位置
func savePathDealWith(handler *StructHandler) *StructHandler {

	//如果文件名和表名不一致，将表名设置为文件名
	baseName := path.Base(handler.SavePath)
	fileName := handler.TableName + ".go"

	if baseName != fileName {
		lastChar := handler.SavePath[len(handler.SavePath)-1 : len(handler.SavePath)]
		if lastChar == "/" {
			//最后一个字符是"/",直接添加表名.go
			//e.g. ./model/
			handler.SavePath = handler.SavePath + fileName
		} else {
			reg := regexp.MustCompile(".{3}$")
			//最后三位是".go",并且与表名不同，替换原来的文件名
			//e.g.  ./model/XXX.go
			if reg.FindString(handler.SavePath) == ".go" {
				handler.SavePath = strings.ReplaceAll(handler.SavePath, baseName, fileName)
			} else {
				//路径结尾没有文件夹分隔符"/"
				//e.g.  ./model
				handler.SavePath = handler.SavePath + "/" + fileName
			}
			//handler.SavePath = handler.SavePath[0:strings.LastIndex(handler.SavePath, "/")+1] + fileName
		}

	}
	fmt.Printf("文件生成位置：%s\n", handler.SavePath)
	return handler
}
