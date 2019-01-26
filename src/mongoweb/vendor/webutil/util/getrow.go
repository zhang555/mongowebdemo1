package util

import (
	"io/ioutil"
	"strings"
	"log"
	"os"
)

func GetRow() int {
	dir := `C:\MyInstall\go\src\bank_party`
	dir = `C:\MyInstall\gosrc\src\bank_party\`

	strs := readDirReturnList222(dir)
	count := 0
	for _, value := range strs {
		data, _ := ioutil.ReadFile(value)
		if strings.HasSuffix(value, ".go") {
			num := getRowByBytes(data)
			count += num
		}
	}
	return count
	//DB.Create(&model.SystemContent{Title: "行数", Num: count})
}

func readDirReturnList222(dir string) []string {
	fileList := []string{}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range fileInfos {
		if value.IsDir() {
			if value.Name() != "vendor" {
				//log.Println("dir  ", value.Name())
				fileList2 := readDirReturnList222(dir + `\` + value.Name())
				//dirList = append(dirList, dir+`\`+value.Name())
				fileList = append(fileList, fileList2...)
			}
		} else {
			fileList = append(fileList, dir+`\`+value.Name())
		}
	}
	return fileList
}

func GetRowByDir(dir string) int {
	strs := readDirReturnList222(dir)
	count := 0
	for _, value := range strs {
		data, _ := ioutil.ReadFile(value)
		if strings.HasSuffix(value, ".go") {
			num := getRowByBytes(data)
			count += num
		}
	}
	return count
}

func WriteAllCode() {

	var (
		dir     string
		newFile *os.File
		err     error
		data    []byte
	)

	newFile, err = os.Create("test222.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//newFile, err = os.OpenFile("code.txt", os.O_RDWR, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer newFile.Close()

	dir = `C:\MyInstall\go\src\bank_party\`
	strs := readDirReturnList222(dir)

	for _, value := range strs {
		if strings.HasSuffix(value, ".go") {
			data, _ = ioutil.ReadFile(value)

			_, err = newFile.Write(data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	dir = `C:\MyInstall\page\vue-xfh\src\`

	strs = readDirReturnList333(dir)

	for _, value := range strs {
		if strings.HasSuffix(value, ".vue") || strings.HasSuffix(value, ".js") || strings.HasSuffix(value, ".css") {

			data, _ = ioutil.ReadFile(value)

			_, err = newFile.Write(data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	dir = `C:\MyInstall\page\vue-xfh-home\src\`

	strs = readDirReturnList333(dir)

	for _, value := range strs {
		if strings.HasSuffix(value, ".vue") || strings.HasSuffix(value, ".js") || strings.HasSuffix(value, ".css") {

			data, _ = ioutil.ReadFile(value)

			_, err = newFile.Write(data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return

}

func GetRow2() int {

	var dir string

	dir = `C:\MyInstall\page\vue-xfh\src\`
	dir = `C:\MyInstall\page\vue-xfh-home\src\`
	strs := readDirReturnList333(dir)
	count := 0
	for _, value := range strs {
		data, _ := ioutil.ReadFile(value)
		if strings.HasSuffix(value, ".vue") || strings.HasSuffix(value, ".js") || strings.HasSuffix(value, ".css") {
			num := getRowByBytes(data)
			count += num
		}
	}
	return count

}

func readDirReturnList333(dir string) []string {
	fileList := []string{}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range fileInfos {
		if value.IsDir() {
			if value.Name() != "assets" {
				//log.Println("dir  ", value.Name())
				log.Println("dir  ", dir+`\`+value.Name())
				fileList2 := readDirReturnList333(dir + `\` + value.Name())
				//dirList = append(dirList, dir+`\`+value.Name())
				fileList = append(fileList, fileList2...)
			}
		} else {
			fileList = append(fileList, dir+`\`+value.Name())
		}
	}
	return fileList
}

//传入字节数组， 看有多少行
func getRowByBytes(b []byte) int {
	return len(strings.Split(string(b), "\n"))
}
