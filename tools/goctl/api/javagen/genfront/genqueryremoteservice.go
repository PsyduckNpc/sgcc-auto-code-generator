package genfront

import (
	"bytes"
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil/front"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	apiutil "github.com/zeromicro/go-zero/tools/goctl/api/util"
	"strings"
	"text/template"
)

//go:embed fronttpl/QueryRemoteService.tpl
var queryRemoteServiceTemplate string

func GenQueryRemoteService(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool) error {

	defer func() { ch <- true }()
	var imports string
	importMap := make(map[string]string)

	javaFile := "Query" + javagenutil.HeadUpCase(subSvName) + "RemoteService.java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/"+subSvName+"/remote/datatransfer", "", javaFile)
	if err != nil {
		return err
	} //创建出现错误
	if !created {
		return nil
	} //已经存在文件
	defer fp.Close()

	var inters []javagenutil.Inter
	for _, g := range api.Service.Groups {
		group := g.GetAnnotation("group")
		if !strings.HasPrefix(group, "/") {
			group = "/" + group
		}

		for _, v := range g.Routes {
			if javagenutil.JudgeSubdomainFunc(v.Handler) {
				continue
			}
			inter := javagenutil.Inter{
				InterfaceName: v.Handler,
				Route:         group + v.Path,
			}
			inters = append(inters, inter)
		}

	}

	interfaceImplName := front.AssembleQueryRemoteServiceImpl(svName, subSvName, &importMap)
	for _, v := range importMap {
		imports += v
	}

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "queryRemoteServiceTemplate").Parse(queryRemoteServiceTemplate))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":            strings.ToLower(svName),
		"subSvName":         strings.ToLower(subSvName),
		"imports":           imports,
		"dataService":       "emss-acap-data",
		"interfaceName":     front.AssembleQueryRemoteService(svName, subSvName, nil),
		"interfaceImplName": interfaceImplName,
		"inters":            inters,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
