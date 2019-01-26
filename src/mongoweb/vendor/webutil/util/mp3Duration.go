package util

import (
	"os/exec"
	"bytes"
	"strings"
	"strconv"
)

func MP3Info(path string) string {

	var (
		str    string
		stderr bytes.Buffer
	)

	cmd := exec.Command("ffmpeg", "-i", path)

	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		str = stderr.String()
		return str
	}
	return ""
}

//得到 分 秒
func MP3Duration1(path string) (a int, b int) {

	var (
		str   string
		str1  string
		str2  string
		index int
		l     int
	)

	str = MP3Info(path)

	l = len(str)

	index = strings.Index(str, "Duration:")

	if index == -1 {
		return 0, 0
	}

	if index+18 > l-1 {
		return 0, 0
	}

	str1 = str[index+13 : index+15]
	str2 = str[index+16 : index+18]

	num1, err := strconv.Atoi(str1)
	if err != nil {
		return 0, 0
	}

	num2, err := strconv.Atoi(str2)
	if err != nil {
		return 0, 0
	}

	return num1, num2

	return 0, 0
}

//总秒数
func MP3Duration3(path string) (int) {

	var (
		str   string
		str1  string
		str2  string
		index int
		l     int
	)

	str = MP3Info(path)

	l = len(str)

	index = strings.Index(str, "Duration:")

	if index == -1 {
		return 0
	}

	if index+18 > l-1 {
		return 0
	}

	str1 = str[index+13 : index+15]
	str2 = str[index+16 : index+18]

	num1, err := strconv.Atoi(str1)
	if err != nil {
		return 0
	}

	num2, err := strconv.Atoi(str2)
	if err != nil {
		return 0
	}

	return num1*60 + num2

}

func MP3Duration(path string) (duration int, str string) {

	var (
		//str   string
		str1  string
		str2  string
		index int
		l     int
	)

	str = MP3Info(path)

	l = len(str)

	index = strings.Index(str, "Duration:")

	if index == -1 {
		return
	}

	if index+18 > l-1 {
		return
	}

	str1 = str[index+13 : index+15]
	str2 = str[index+16 : index+18]

	num1, err := strconv.Atoi(str1)
	if err != nil {
		return
	}

	num2, err := strconv.Atoi(str2)
	if err != nil {
		return
	}

	duration = num1*60 + num2

	return

}
