package javagen

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/genfront"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/gensubdomain"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var (
	// VarStringDir describes a directory.
	VarStringDir string
	// VarStringAPI describes an API.
	VarStringAPI string

	//生成的类型，front或者subdomain
	VarServerType string
)

// JavaCommand generates java code command entrance.
func JavaCommand(_ *cobra.Command, _ []string) error {
	apiFile := VarStringAPI
	//apiFile := "../../../../attachment/greet2.api"
	dir := VarStringDir
	//dir := "C:/Users/Psydu/Desktop/GoctlGen"
	if len(apiFile) == 0 {
		return errors.New("missing -api")
	}
	if len(dir) == 0 {
		return errors.New("missing -dir")
	}

	api, err := parser.Parse(apiFile) //解析.api文件
	if err != nil {
		return err
	}

	if err := api.Validate(); err != nil {
		return err
	}

	api.Service = api.Service.JoinPrefix()
	//packetName := strings.TrimSuffix(api.Service.Name, "-api")
	comm, err := javagenutil.FindSvComm("svname", api)
	if err != nil {
		panic(err)
	}
	svName := comm.Value
	comm, err = javagenutil.FindSvComm("subsvname", api)
	if err != nil {
		panic(err)
	}
	subSvName := comm.Value

	logx.Must(pathx.MkdirIfNotExist(dir)) //检查生成目录是否存在，不存在则创建
	//logx.Must(genPacket(dir, packetName, api))     //包生成
	//logx.Must(genComponents(dir, packetName, api)) //生成组件

	FillMissBo(api)
	//importPathPadding(api, packetName

	//VarStringType = "SUBDOMAIN"
	if strings.Contains(strings.ToUpper(VarServerType), "FRONT") {
		genFront(dir, svName, subSvName, api)
	} else if strings.Contains(strings.ToUpper(VarServerType), "SUBDOMAIN") {
		genSubdomain(dir, svName, subSvName, api)
	} else {
		return errors.New("type is nil or not support")
	}

	fmt.Println(aurora.Green("Done."))
	return nil
}

func genFront(dir, svName, subSvName string, api *spec.ApiSpec) {
	size := 10
	ch := make(chan bool, size)

	go logx.Must(genfront.GenModel(dir, svName, subSvName, api, ch))
	go logx.Must(genfront.GenQueryRemoteService(dir, svName, subSvName, api, ch))
	go logx.Must(genfront.GenQueryRemoteServiceImpl(dir, svName, subSvName, api, ch))
	go logx.Must(genfront.GenRemoteService(dir, svName, subSvName, api, ch))
	go logx.Must(genfront.GenRemoteServiceImpl(dir, svName, subSvName, api, ch))
	go logx.Must(genfront.GenQueryService(dir, svName, subSvName, api, ch))
	go logx.Must(genfront.GenQueryServiceImpl(dir, svName, subSvName, api, ch))
	go logx.Must(genfront.GenService(dir, svName, subSvName, api, ch))
	go logx.Must(genfront.GenServiceImpl(dir, svName, subSvName, api, ch))
	go logx.Must(genfront.GenController(dir, svName, subSvName, api, ch))

	for i := 0; i < size; i++ {
		<-ch
	}
}

func genSubdomain(dir, svName, subSvName string, api *spec.ApiSpec) {
	size := 7
	ch := make(chan bool, size)
	constants := make(map[string]javagenutil.Logic)

	go logx.Must(gensubdomain.GenModel(dir, svName, subSvName, api, ch, &constants))
	go logx.Must(gensubdomain.GenRemoteService(dir, svName, subSvName, api, ch, &constants))
	go logx.Must(gensubdomain.GenRemoteServiceImpl(dir, svName, subSvName, api, ch, &constants))
	go logx.Must(gensubdomain.GenDomainService(dir, svName, subSvName, api, ch, &constants))
	go logx.Must(gensubdomain.GenDomainServiceImpl(dir, svName, subSvName, api, ch, &constants))
	go logx.Must(gensubdomain.GenSpi(dir, svName, subSvName, api, ch, &constants))
	go logx.Must(gensubdomain.GenController(dir, svName, subSvName, api, ch, &constants))

	for i := 0; i < size; i++ {
		<-ch
	}

	ch = make(chan bool, 1)
	go logx.Must(gensubdomain.GenConstant(dir, svName, subSvName, api, ch, &constants))
	<-ch
}

func FillMissBo(api *spec.ApiSpec) {
	//如果只添加AO或DTO，就按照已有的生成缺少的AO或者DTO
	var aoFlag bool
	var dtoFlag bool
	var aoType spec.DefineStruct
	var dtoType spec.DefineStruct

	for _, v := range api.Types {
		defineStruct := v.(spec.DefineStruct)
		rawName := defineStruct.RawName
		if strings.HasSuffix(rawName, "DTO") || strings.HasSuffix(rawName, "Dto") {
			dtoFlag = true
			dtoType = v.(spec.DefineStruct)
		}
		if strings.HasSuffix(rawName, "AO") || strings.HasSuffix(rawName, "Ao") {
			aoFlag = true
			aoType = v.(spec.DefineStruct)
		}
	}
	//有AO
	if (dtoFlag || aoFlag) && dtoFlag == false {
		s := aoType.RawName[:len(aoType.RawName)-2]
		aoType.ImportPath = strings.TrimSuffix(aoType.ImportPath, "ao."+aoType.RawName+";\n")
		aoType.RawName = s + "DTO"
		aoType.Group = s
		aoType.Type = "DTO"
		aoType.ImportPath += "dto." + aoType.RawName + ";\n"
		api.Types = append(api.Types, aoType)

		for ig, g := range api.Service.Groups {
			for iv, v := range g.Routes {
				switch v.RequestType.(type) {
				case nil:
					continue
				case spec.DefineStruct:
					api.Service.Groups[ig].Routes[iv].MiddleType = aoType
				case spec.ArrayType:
					arrayType := spec.ArrayType{
						RawName: "[]" + aoType.RawName,
						Value:   aoType,
					}
					api.Service.Groups[ig].Routes[iv].MiddleType = arrayType
				default:
					panic("not supported RequestType type")
				}

			}
		}
	}

	//有DTO
	if (aoFlag || dtoFlag) && aoFlag == false {
		s := dtoType.RawName[:len(dtoType.RawName)-3]
		dtoType.ImportPath = strings.TrimSuffix(dtoType.ImportPath, "dto."+dtoType.RawName+";\n")
		dtoType.RawName = s + "AO"
		dtoType.Group = s
		dtoType.Type = "AO"
		dtoType.ImportPath += "ao." + dtoType.RawName + ";\n"
		api.Types = append(api.Types, dtoType)

		for ig, g := range api.Service.Groups {
			for iv, v := range g.Routes {
				switch v.RequestType.(type) {
				case nil:
					continue
				case spec.DefineStruct:
					api.Service.Groups[ig].Routes[iv].MiddleType = v.RequestType
					api.Service.Groups[ig].Routes[iv].RequestType = dtoType
				case spec.ArrayType:
					arrayType := spec.ArrayType{
						RawName: "[]" + aoType.RawName,
						Value:   aoType,
					}
					api.Service.Groups[ig].Routes[iv].MiddleType = v.RequestType
					api.Service.Groups[ig].Routes[iv].RequestType = arrayType
				default:
					panic("not supported RequestType type")
				}

			}
		}
	}
}

func importPathPadding(api *spec.ApiSpec, packageName string) {
	for i, v := range api.Types {
		defineStruct := v.(spec.DefineStruct)
		defineStruct.ImportPath = javagenutil.AssembleBoImport(packageName, packageName, defineStruct.RawName, "front")
		api.Types[i] = defineStruct
		if api.ImportPathMap == nil {
			api.ImportPathMap = make(map[string]string)
		}
		api.ImportPathMap[defineStruct.RawName] = defineStruct.ImportPath
	}
}
