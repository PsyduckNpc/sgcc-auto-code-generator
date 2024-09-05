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

//go:embed fronttpl/QueryRemoteServiceImpl.tpl
var queryRemoteServiceImplTemplate string

func GenQueryRemoteServiceImpl(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool) error {

	defer func() { ch <- true }()
	var imports string
	importMap := make(map[string]string)

	javaFile := "Query" + javagenutil.HeadUpCase(subSvName) + "RemoteServiceImpl.java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/"+subSvName+"/remote/datatransfer/impl", "", javaFile)
	if err != nil {
		return err
	} //创建出现错误
	if !created {
		return nil
	} //已经存在文件
	defer fp.Close()

	var inters []javagenutil.Inter
	for _, g := range api.Service.Groups {
		for _, v := range g.Routes {
			if javagenutil.JudgeSubdomainFunc(v.Handler) {
				continue
			}
			reqType, _ := javagenutil.AssembleParamType(v.MiddleType, nil)
			inter := javagenutil.Inter{
				InterfaceName: v.Handler,
				ReqType:       reqType,
			}
			inters = append(inters, inter)
		}

	}

	interfaceName := front.AssembleQueryRemoteService(svName, subSvName, &importMap)
	for _, v := range importMap {
		imports += v
	}

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "queryRemoteServiceImplTemplate").Parse(queryRemoteServiceImplTemplate))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":            strings.ToLower(svName),
		"subSvName":         strings.ToLower(subSvName),
		"imports":           imports,
		"interfaceImplName": front.AssembleQueryRemoteServiceImpl(svName, subSvName, nil),
		"interfaceName":     interfaceName,
		"inters":            inters,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
