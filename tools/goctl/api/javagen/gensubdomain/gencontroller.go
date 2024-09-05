package gensubdomain

import (
	"bytes"
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil/front"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil/subdomain"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	apiutil "github.com/zeromicro/go-zero/tools/goctl/api/util"
	"strings"
	"text/template"
)

//go:embed subdomaintpl/Controller.tpl
var controllerTemplate string

func GenController(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool, constants *map[string]javagenutil.Logic) error {

	defer func() { ch <- true }()

	javaFile := subdomain.AssembleController(svName, subSvName, nil) + ".java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/front/controller/impl", "", javaFile)
	if err != nil {
		return err
	} //创建出现错误
	if !created {
		return nil
	} //已经存在文件
	defer fp.Close()

	var imports string
	var serviceGroupDesc string
	var injects []javagenutil.Inject
	var logics []javagenutil.Logic
	importMap := make(map[string]string)

	injects = append(injects, javagenutil.Inject{
		InjectServiceName:   front.AssembleService(svName, subSvName, &importMap),
		VariableServiceName: javagenutil.HeadLowCase(front.AssembleService(svName, subSvName, nil)),
	})

	injects = append(injects, javagenutil.Inject{
		InjectServiceName:   front.AssembleQueryService(svName, subSvName, &importMap),
		VariableServiceName: javagenutil.HeadLowCase(front.AssembleQueryService(svName, subSvName, nil)),
	})

	for _, g := range api.Service.Groups {
		serviceDesc := strings.TrimSuffix(g.GetAnnotation("desc"), "/")
		if serviceDesc != "" {
			serviceGroupDesc += serviceDesc + ","
		}

		for _, v := range g.Routes {

			middleType, middleName := javagenutil.AssembleParamType(v.MiddleType, &importMap)

			group := g.GetAnnotation("group")
			if !strings.HasPrefix(group, "/") {
				group = "/" + group
			}
			group = strings.TrimSuffix(group, "/")

			if !strings.HasPrefix(v.Path, "/") {
				v.Path = "/" + v.Path
			}

			logics = append(logics, javagenutil.Logic{
				InterfaceName: v.Handler,
				DTOType:       middleType,
				DTOName:       middleName,
				Desc:          v.AtDesc,
				Route:         group + v.Path,
				MethodType:    strings.ToUpper(v.Method),
				ReqListFlag: func() bool {
					if v.RequestType == nil {
						return false
					}
					return strings.HasPrefix(v.RequestType.Name(), "[]")
				}(),
				RespListFlag: func() bool {
					if v.ResponseType == nil {
						return false
					}
					return strings.HasPrefix(v.ResponseType.Name(), "[]")
				}(),
				VariableServiceName: func() string {
					if javagenutil.JudgeSubdomainFunc(v.Handler) {
						return javagenutil.HeadLowCase(front.AssembleService(svName, subSvName, &importMap))
					} else {
						return javagenutil.HeadLowCase(front.AssembleQueryService(svName, subSvName, &importMap))
					}
				}(),
			})
		}

	}

	serviceGroupDesc = strings.TrimSuffix(serviceGroupDesc, ",")
	spiName := subdomain.AssembleController(svName, subSvName, &importMap)

	importMap["Resource"] = javagenutil.Resource
	for _, v := range logics {
		if v.ReqListFlag && importMap["ArrayList"] == "" {
			importMap["ArrayList"] = javagenutil.ArrayList
		}
		if v.ReqType != "" {
			importMap["RequestBody"] = javagenutil.RequestBody
		}
	}

	for _, v := range importMap {
		imports += v
	}

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "controllerTemplate").Parse(controllerTemplate))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":           strings.ToLower(svName),
		"subSvName":        strings.ToLower(subSvName),
		"spiName":          spiName,
		"imports":          imports,
		"serviceGroupDesc": serviceGroupDesc,
		"controllerName":   front.AssembleController(svName, subSvName, nil),
		"injects":          injects,
		"logics":           logics,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
