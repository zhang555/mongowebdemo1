package util

import (
	"strings"
	"time"
	"path/filepath"
	"os"
	"reflect"
	"unicode"
)

//拿到后缀
func GetSuffix(str string, substr string) string {
	//strs := strings.Split(str, ".")
	//return strs[len(strs)-1]

	num := strings.LastIndex(str, substr)
	return str[num+1:]
}

//拿到后缀
func GetDotSuffix(str string) string {
	//strs := strings.Split(str, ".")
	//return strs[len(strs)-1]

	num := strings.LastIndex(str, ".")
	return str[num+1:]
}

//前缀
func GetDotPrefix(str string) string {
	num := strings.LastIndex(str, ".")
	return str[:num]
}

func GetPrefixByStr(str string, str2 string) string {
	num := strings.LastIndex(str, str2)
	return str[:num]
}

//
type StringArray []string

func (array StringArray) Contain(str string) bool {
	for _, value := range array {
		if str == value {
			return true
		}
	}
	return false
}
func (array StringArray) NotContain(str string) bool {
	return !array.Contain(str)
}

//
type IntArray []int

func (array IntArray) Contain(num int) bool {
	for _, value := range array {
		if num == value {
			return true
		}
	}
	return false
}
func (array IntArray) NotContain(num int) bool {
	return !array.Contain(num)
}

//
func IsIntArrayContainNum(array []int, num int) bool {
	for _, value := range array {
		if num == value {
			return true
		}
	}
	return false
}

//return
//1:图片
//2:视频
func JudgeFiletype(str string) int {
	var (
		image StringArray = []string{
			"png",
			"jpg",
			"jpeg",
		}
		video StringArray = []string{
			"avi",
			"avi",
		}
	)

	switch {
	case image.Contain(str):
		return 1
	case video.Contain(str):
		return 2
	default:
		return 0
	}

}

//
func IsImageType(str string) bool {
	var (
		image StringArray = []string{
			"png",
			"jpg",
			"jpeg",
			//"gif",
			//"bmp",
		}
	)
	return image.Contain(str)
}

//
func IsVideoType(str string) bool {
	var (
		image StringArray = []string{
			"mp4",
			//"avi",
			//"rmvb",
			//"wmv",
		}
	)
	return image.Contain(str)
}

func IsAudioType(str string) bool {
	var (
		image StringArray = []string{
			"mp3",
			//"wav",
		}
	)
	return image.Contain(str)
}

func IsDocumentType(str string) bool {
	var (
		file StringArray = []string{
			"pptx",
			"ppt",
			"doc",
			"docx",
			"xls",
			"xlsx",
		}
	)
	return file.Contain(str)
}

func IsTwtReportDocumentType(str string) bool {
	var (
		file StringArray = []string{
			"txt",
			"doc",
			"docx",
		}
	)
	return file.Contain(str)
}

func IsAppType(str string) bool {
	var (
		image StringArray = []string{
			"apk",
		}
	)
	return image.Contain(str)
}

func IsExcelType(str string) bool {
	var (
		image StringArray = []string{
			"xls",
			"xlsx",
		}
	)
	return image.Contain(str)
}

//得到年月日的字符串
func GetTimeString() string {
	return time.Now().Format("2006-01-02")
}

// 创建文件夹
func CreateDateDir(basePath string, folderName string) error {
	//folderName := time.Now().Format("2006-01-02")
	folderPath := filepath.Join(basePath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		err := os.Mkdir(folderPath, 0777)
		if err != nil {
			return err
		}

		// 再修改权限
		err = os.Chmod(folderPath, 0777)
		if err != nil {
			return err
		}
		return nil
	}
	return nil

	//return folderPath
}

//
func GetBeanName(bean interface{}) string {

	beantype := reflect.TypeOf(bean)
	beanstr := beantype.String()
	stringarr := strings.Split(beanstr, ".")

	controllerurl := stringarr[1]

	url := string(unicode.ToLower(rune(controllerurl[0]))) + controllerurl[1:]

	return url
}
