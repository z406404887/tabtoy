package main

import (
	"flag"
	"fmt"
	"github.com/0990/tabtoy/v2"
	"github.com/0990/tabtoy/v2/i18n"
	"github.com/0990/tabtoy/v2/printer"
	"github.com/davyxu/golog"
	"os"
)

var log = golog.New("main")

const (
	Version_v2 = "2.9.1"
	Version_v3 = "3.0.0"
)

var (
	// 显示版本号
	paramVersion = flag.Bool("version", false, "Show version")

	// 工作模式
	paramMode = flag.String("mode", "", "v2")

	// 并发导出,提高导出速度, 输出日志会混乱
	paramPara = flag.Bool("para", false, "parallel export by your cpu count")

	paramLanguage = flag.String("lan", "en_us", "set output language")

)

var (
	paramProtoImport       = flag.String("protoimport", "", "proto import header")
	paramProtoOutputIgnoreFile       = flag.String("protooutputignorefile", "", "protooutputignorefile")
	paramPackageName       = flag.String("package", "", "override the package name in table @Types")
	paramCombineStructName = flag.String("combinename", "Config", "combine struct name, code struct name")
	paramProtoOut          = flag.String("proto_out", "", "output protobuf define (*.proto)")
	paramPbtOut            = flag.String("pbt_out", "", "output proto text format (*.pbt)")
	paramLuaOut            = flag.String("lua_out", "", "output lua code (*.lua)")
	paramJsonOut           = flag.String("json_out", "", "output json format (*.json)")
	paramCSharpOut         = flag.String("csharp_out", "", "output c# class and deserialize code (*.cs)")
	paramGoOut             = flag.String("go_out", "", "output golang code (*.go)")
	paramBinaryOut         = flag.String("binary_out", "", "output binary format(*.bin)")
	paramTypeOut           = flag.String("type_out", "", "output table types(*.json)")
	paramCppOut            = flag.String("cpp_out", "", "output c++ format (*.cpp)")
)

// 特殊文件格式参数
var (
	paramProtoVersion = flag.Int("protover", 3, "output .proto file version, 2 or 3")

	paramLuaEnumIntValue = flag.Bool("luaenumintvalue", false, "use int type in lua enum value")
	paramLuaTabHeader    = flag.String("luatabheader", "", "output string to lua tab header")

	paramGenCSharpBinarySerializeCode = flag.Bool("cs_gensercode", true, "generate c# binary serialize code, default is true")
)

func main() {

	flag.Parse()

	// 版本
	if *paramVersion {
		fmt.Printf("%s, %s", Version_v2, Version_v3)
		return
	}

	switch *paramMode {
	//case "v3":
	//	V3Entry()
	case "exportorv2", "v2":
		V2Entry()
	//case "v2tov3":
	//	V2ToV3Entry()
	default:
		fmt.Println("--mode not specify")
		os.Exit(1)
	}

}

func V2Entry() {
	g := printer.NewGlobals()

	if *paramLanguage != "" {
		if !i18n.SetLanguage(*paramLanguage) {
			log.Infof("language not support: %s", *paramLanguage)
		}
	}

	g.Version = Version_v2

	for _, v := range flag.Args() {
		g.InputFileList = append(g.InputFileList, v)
	}

	g.ParaMode = *paramPara
	g.CombineStructName = *paramCombineStructName
	g.ProtoVersion = *paramProtoVersion
	g.LuaEnumIntValue = *paramLuaEnumIntValue
	g.LuaTabHeader = *paramLuaTabHeader
	g.GenCSSerailizeCode = *paramGenCSharpBinarySerializeCode
	g.PackageName = *paramPackageName
	g.ProtoImport = *paramProtoImport
	g.Protooutputignorefile  = *paramProtoOutputIgnoreFile

	if *paramProtoOut != "" {
		g.AddOutputType("proto", *paramProtoOut)
	}

	if *paramPbtOut != "" {
		g.AddOutputType("pbt", *paramPbtOut)
	}

	if *paramJsonOut != "" {
		g.AddOutputType("json", *paramJsonOut)
	}

	if *paramLuaOut != "" {
		g.AddOutputType("lua", *paramLuaOut)
	}

	if *paramCSharpOut != "" {
		g.AddOutputType("cs", *paramCSharpOut)
	}

	if *paramGoOut != "" {
		g.AddOutputType("go", *paramGoOut)
	}

	if *paramCppOut != "" {
		g.AddOutputType("cpp", *paramCppOut)
	}

	if *paramBinaryOut != "" {
		g.AddOutputType("bin", *paramBinaryOut)
	}

	if *paramTypeOut != "" {
		g.AddOutputType("type", *paramTypeOut)
	}

	if !v2.Run(g) {
		os.Exit(1)
	}
}
