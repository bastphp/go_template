package log

import (
	"github.com/gin-gonic/gin"
	logRequest "kuke_logger/common/request"
	"kuke_logger/common/response"
)

type LogApi struct {
}

func (l *LogApi) CreateLonger(c *gin.Context) {
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf(string(data))
	//filePath := "./log.json"
	//file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 066)
	//defer file.Close()
	//write := bufio.NewWriter(file)
	//write.WriteString(string(data) + "/n")
	var log []logRequest.LogRequest
	_ = c.ShouldBindJSON(&log)
	//fmt.Print(len(log))
	_ = logService.CreateLog(log)
	response.Result(1, "{}", "success", c)
}
