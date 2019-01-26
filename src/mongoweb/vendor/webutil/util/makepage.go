package util

//根据当前页，总记录数，每页记录个数，页码个数
// 计算出 总页数，显示的页码，
func MakePage(pageNum, count, pageSize, pageNumCount int) map[string]interface{} {

	if pageSize == 0 {
		page := map[string]interface{}{
			"pageNum":  250,
			"pageSize": 250,
			"count":    250,
		}
		return page
	}

	var (
		minPage  int
		showPage []int
		pageAll  int
	)

	if count%pageSize > 0 {
		pageAll = count/pageSize + 1
	} else {
		pageAll = count / pageSize
	}

	if pageAll < pageNumCount {
		pageNumCount = pageAll
	}

	minPage = pageNum - (pageNumCount-1)/2

	if minPage >= (pageAll - pageNumCount + 1) {
		minPage = (pageAll - pageNumCount + 1)
	}
	if minPage < 1 {
		minPage = 1
	}

	for i := 0; i < pageNumCount; i++ {
		showPage = append(showPage, minPage)
		minPage++
	}

	page := map[string]interface{}{
		"pageNum":  pageNum,
		"pageSize": pageSize,
		"count":    count,

		//"pageNumCount": pageNumCount,
		//"pageShow":     showPage,
		//"pageAll":      pageAll,
	}

	return page
}

//type PageRequest struct {
//	Status       int //0: 正常   -1：异常 ，
//	PageNum      int
//	PageSize     int
//	PageNumCount int
//}

/*
处理 url中 带的三个参数， 字符串转为数字，返回结构体。
*/
//func HandlePageRequest(c *gin.Context) PageRequest {
//
//	pageNum := c.DefaultQuery("pageNum", PageNumString)
//	pageSize := c.DefaultQuery("pageSize", PageSizeString)
//	pageNumCount := c.DefaultQuery("pageNumCount", PageNumCountString)
//
//	num, err := strconv.Atoi(pageNum)
//	if err != nil {
//		return PageRequest{Status: -1}
//	}
//
//	size, err := strconv.Atoi(pageSize)
//	if err != nil {
//		return PageRequest{Status: -1}
//	}
//
//	numCount, err := strconv.Atoi(pageNumCount)
//	if err != nil {
//		return PageRequest{Status: -1}
//	}
//
//	return PageRequest{
//		Status:       0,
//		PageNum:      num,
//		PageSize:     size,
//		PageNumCount: numCount,
//	}
//
//}
