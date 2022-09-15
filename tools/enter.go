package tools

import "kuke_logger/tools/es"

type ToolsGroup struct {
	EsGroup es.EsToolGroup
}

var ToolsGroupApp = new(ToolsGroup)
