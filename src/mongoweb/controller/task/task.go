package task

func init() {
	go task1()

}

//备份数据库
//func updateTicker1() *time.Ticker {
//	nextTick := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
//	for nextTick.Before(time.Now()) {
//		nextTick = nextTick.Add(1 * time.Hour)
//	}
//	diff := nextTick.Sub(time.Now())
//	return time.NewTicker(diff)
//}
//
func task1() {
	//	ticker := updateTicker1()
	//	for {
	//		<-ticker.C
	//
	//		//BackupDatabase()
	//
	//		ticker = updateTicker1()
	//	}
}
