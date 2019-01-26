package util

import "strconv"

//将数组 拼接成 sql 中  where in （） ， 需要的字符串。
//如果 数组为空 ， 返回 随便一个字符串 ， 查不到结果就行。
func AttachArrayWithComma(ids []int) string {
	if ids == nil || len(ids) == 0 {
		return "-1"
	}

	sql := ""
	for index, value := range ids {
		sql += strconv.Itoa(value)
		if index != len(ids)-1 {
			sql += ","
		}
	}
	return sql
}
