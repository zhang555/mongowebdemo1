package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"unicode"
	"fmt"
	"os"
)

type Columns struct {
	TableName      string `gorm:"column:TABLE_NAME" json:"tableName"`
	ColumnName     string `gorm:"column:COLUMN_NAME" `
	DataType       string `gorm:"column:DATA_TYPE" `
	Ordinal        int    `gorm:"column:ORDINAL_POSITION" `
	COLUMN_COMMENT string `gorm:"column:COLUMN_COMMENT" `
}

type BeanColumn struct {
	Name           string
	Type           string
	Ordinal        int
	COLUMN_COMMENT string
}

type DbParam struct {
	IP       string
	Password string
	DbName   string
}

func GenerateBean(param DbParam, tableName string) {
	var (
		err  error
		path string
		db   *gorm.DB
		ip   string
	)

	if param.IP == "localhost" {
		ip = ""
	} else {
		ip = "tcp(" + param.IP + ":3306)"
	}

	//param.DbName = "information_schema"

	path = "root:" + param.Password + "@" + ip + "/" + "information_schema" + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open("mysql", path)
	if err != nil {
		panic("failed to connect database")
	}
	//db.LogMode(true)
	db.SingularTable(true)
	defer db.Close()

	var column []Columns
	//var str string
	//str = "system_log"

	//tableName = "%" + tableName + "%"

	db.Where("TABLE_NAME like ? and TABLE_SCHEMA = ? ", tableName, param.DbName).Find(&column)

	var beans map[string][]BeanColumn = make(map[string][]BeanColumn)

	for _, value := range column {
		beanc := BeanColumn{
			Name:           value.ColumnName,
			Type:           value.DataType,
			Ordinal:        value.Ordinal,
			COLUMN_COMMENT: value.COLUMN_COMMENT,
		}
		beans[value.TableName] = append(beans[value.TableName], beanc)
	}

	//log.Println(beans)

	for key, value := range beans {

		fmt.Println("type " + makestr2(key) + " struct {")
		for _, v2 := range value {
			if v2.Type == "timestamp" {
				continue
			}

			var dataType string
			var name string
			var jsonname string
			jsonname = makestr(v2.Name)
			name = string(unicode.ToUpper(rune(jsonname[0]))) + jsonname[1:]
			if name == "Id" {
				name = "ID"
			}
			if strings.Contains(v2.Type, "char") || strings.Contains(v2.Type, "text") {
				dataType = "string"
			} else if strings.Contains(v2.Type, "int") {
				dataType = "int"
			} else {
				dataType = v2.Type
			}
			fmt.Println(name, dataType, "`json:\""+jsonname+"\"`", ` //`, v2.COLUMN_COMMENT)
		}
		fmt.Println("ModelTime")
		fmt.Println("}")
		//break
	}

}

func f1() {
	var (
		err      error
		path     string
		db       *gorm.DB
		ip       string
		DbParam1 DbParam = DbParam{
			IP: "39.107.80.58",
			//IP:       "localhost",
			Password: os.Getenv("dbpassword"),
			//DbName:   "bank_party2",
			DbName: "information_schema",
		}
	)

	if DbParam1.IP == "localhost" {
		ip = ""
	} else {
		ip = "tcp(" + DbParam1.IP + ":3306)"
	}

	//path = "root:88888888@/bank_party2?charset=utf8&parseTime=True&loc=Local"
	//path = "root:" + password + "@/information_schema?charset=utf8&parseTime=True&loc=Local"
	//path = "root:" + password + "@tcp(193.112.40.84:3306)/information_schema?charset=utf8&parseTime=True&loc=Local"

	path = "root:" + DbParam1.Password + "@" + ip + "/" + DbParam1.DbName + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open("mysql", path)
	if err != nil {
		panic("failed to connect database")
	}
	//db.LogMode(true)
	db.SingularTable(true)
	defer db.Close()

	var column []Columns
	var str string
	str = "system_log"

	db.Where("TABLE_NAME like ? ", str).Find(&column)

	var beans map[string][]BeanColumn = make(map[string][]BeanColumn)

	for _, value := range column {
		beanc := BeanColumn{
			Name:           value.ColumnName,
			Type:           value.DataType,
			Ordinal:        value.Ordinal,
			COLUMN_COMMENT: value.COLUMN_COMMENT,
		}
		beans[value.TableName] = append(beans[value.TableName], beanc)
	}

	//log.Println(beans)

	for key, value := range beans {

		fmt.Println("type " + makestr2(key) + " struct {")
		for _, v2 := range value {
			if v2.Type == "timestamp" {
				continue
			}

			var dataType string
			var name string
			var jsonname string
			jsonname = makestr(v2.Name)
			name = string(unicode.ToUpper(rune(jsonname[0]))) + jsonname[1:]
			if name == "Id" {
				name = "ID"
			}
			if strings.Contains(v2.Type, "char") || strings.Contains(v2.Type, "text") {
				dataType = "string"
			} else if strings.Contains(v2.Type, "int") {
				dataType = "int"
			} else {
				dataType = v2.Type
			}
			fmt.Println(name, dataType, "`json:\""+jsonname+"\"`", ` //`, v2.COLUMN_COMMENT)
		}
		fmt.Println("ModelTime")
		fmt.Println("}")
		//break
	}

}

//去掉下划线，  首字母不大写
func makestr(str string) string {
	var (
		strs []string
		str2 string
	)

	strs = strings.Split(str, "_")

	for index, value := range strs {
		if index == 0 {
			str2 += value
		} else {
			str2 += string(unicode.ToUpper(rune(value[0]))) + value[1:]
		}
	}
	return str2
}

//去掉下划线，  首字母大写
func makestr2(str string) string {
	var (
		strs []string
		str2 string
	)
	strs = strings.Split(str, "_")
	for _, value := range strs {
		str2 += string(unicode.ToUpper(rune(value[0]))) + value[1:]
	}
	return str2
}
