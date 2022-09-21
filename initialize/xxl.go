package initialize

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	xxl_job_executor_gin "github.com/gin-middleware/xxl-job-executor"
	"github.com/xxl-job/xxl-job-executor-go"
	"kuke_logger/global"
	"kuke_logger/tools"
	"strconv"
)

func Xxl(r *gin.Engine) {
	xxlCfg := global.GVA_CONFIG.Xxl
	exec := xxl.NewExecutor(
		xxl.ServerAddr(xxlCfg.XxlJobAdminAddress),
		xxl.ExecutorPort(strconv.Itoa(global.GVA_CONFIG.System.Addr)),
		xxl.RegistryKey(xxlCfg.XxlJobAppName),
	)
	exec.Init()
	defer exec.Stop()
	xxl_job_executor_gin.XxlJobMux(r, exec)
	exec.RegTask("createMapping", func(cxt context.Context, param *xxl.RunReq) string {
		var esMapping = tools.ToolsGroupApp.EsGroup
		logType := [4]string{"error", "info", "waring", "notice"}
		for _, value := range logType {
			indexName := esMapping.GetIndexName(value, "+1")
			fmt.Println(indexName)
			b := esMapping.IndexExits(indexName)
			if !b {
				esMapping.InitMapping(indexName, "_doc", esMapping.LogMappingTmp())
			}
		}
		return "success"
	})
}
