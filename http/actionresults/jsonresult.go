package actionresults

import "encoding/json"

type JSONResult struct {
	data any
}

func (j *JSONResult) Execute(ctx *ActionContext) error {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(ctx.ResponseWriter).Encode(j.data)
}

func NewJSONResult(data any) ActionResult {
	return &JSONResult{data: data}
}
