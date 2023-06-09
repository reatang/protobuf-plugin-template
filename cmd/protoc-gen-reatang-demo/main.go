package main

import (
	"os"

	"github.com/reatang/protobuf-plugin-template/internal/demo"
	"github.com/reatang/protobuf-plugin-template/pkg/generator"
	"github.com/reatang/protobuf-plugin-template/pkg/log"
	"github.com/reatang/protobuf-plugin-template/pkg/protoio"
)

func main() {
	// 解码参数
	req := protoio.RequestDecode(os.Stdin)
	params := protoio.ParamsMapDecode(req)
	err := log.ConfigureLogging(params)
	if err != nil {
		panic(err)
	}

	if v, ok := params["reqjson"]; ok && v == "true" {
		protoio.RequestJson(req)
	}

	// 注册分析器
	generator.AddAnalyse(demo.NewDemoAnalyseBuilder())

	// 启动生成器
	g := generator.NewProtoFileGenerator(params)

	// 执行生成器
	resp, err := g.Generate(demo.DemoAnalyse, req)

	// 返回 protoc
	protoio.ResponseEncode(resp)
}
