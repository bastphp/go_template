package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logRequest "kuke_logger/common/request"
	"kuke_logger/common/response"
)

type LogApi struct {
}

func (l *LogApi) CreateLonger(c *gin.Context) {
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//filePath := "./log.json"
	//file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 777)
	//defer file.Close()
	//write := bufio.NewWriter(file)
	//_, err := write.WriteString(string(data) + "\n")
	//fmt.Println(string(data))
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//write.Flush()
	var log []logRequest.LogRequest
	_ = c.ShouldBindJSON(&log)
	fmt.Print(len(log))
	_ = logService.CreateLog(log)
	response.Result(1, "{}", "success", c)
}
