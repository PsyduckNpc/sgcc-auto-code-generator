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

//go:embed subdomaintpl/RemoteService.tpl
var remoteServiceTemplate string

func GenRemoteService(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool) error {

	defer func() { ch <- true }()
	var imports string
	importMap := make(map[string]string)

	javaFile := javagenutil.HeadUpCase(subSvName) + "RemoteService.java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/remote/datatransfer", "", javaFile)
	if err != nil {
		return err
	} //创建出现错误
	if !created {
		return nil
	} //已经存在文件
	defer fp.Close()

	//var importTypes []spec.Type
	var inters []javagenutil.Inter
	for _, g := range api.Service.Groups {

		group := g.GetAnnotation("group")
		if !strings.HasPrefix(group, "/") {
			group = "/" + group
		}

		for _, v := range g.Routes {
			//if javagenutil.JudgeSubdomainFunc(v.Handler) {
			//	continue
			//}
			reqType, reqName := javagenutil.AssembleParamType(v.MiddleType, &importMap)

			inter := javagenutil.Inter{
				InterfaceName: v.Handler,
				ReqType:       reqType,
				DTOName:       reqName,
				Route:         group + v.Path,
			}
			inters = append(inters, inter)
		}
	}

	serviceImpl := front.AssembleRemoteServiceImpl(svName, subSvName, &importMap)
	javagenutil.GetBo(api, "DTO", &importMap)
	for _, v := range importMap {
		imports += v
	}

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "remoteServiceTemplate").Parse(remoteServiceTemplate))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":           strings.ToLower(svName),
		"subSvName":        strings.ToLower(subSvName),
		"interfaceName":    subdomain.AssembleRemoteService(svName, subSvName, nil),
		"imports":          imports,
		"dataService":      "emss-acap-data",
		"classImplName":    serviceImpl,
		"inters":           inters,
		"subdomainService": "emss-scp-" + subSvName + "-subdomain",
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
