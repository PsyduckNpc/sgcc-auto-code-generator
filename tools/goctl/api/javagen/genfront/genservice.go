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

//go:embed fronttpl/Service.tpl
var serviceTemplate string

func GenService(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool) error {

	defer func() { ch <- true }()
	var imports string

	javaFile := javagenutil.HeadUpCase(subSvName) + "Service.java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/"+subSvName+"/service/application", "", javaFile)
	if err != nil {
		return err
	} //创建出现错误
	if !created {
		return nil
	} //已经存在文件
	defer fp.Close()

	var interfaces []javagenutil.Inter
	importMap := make(map[string]string)
	for _, v := range api.Service.Routes() {
		if !javagenutil.JudgeSubdomainFunc(v.Handler) {
			continue
		}
		inter := javagenutil.Inter{}
		respType, _ := javagenutil.AssembleParamType(v.ResponseType, &importMap)
		inter.RespType = respType
		inter.InterfaceName = v.Handler
		reqType, param := javagenutil.AssembleParamType(v.RequestType, &importMap)
		inter.ReqType = reqType
		inter.ReqParamName = param
		interfaces = append(interfaces, inter)
	}

	for _, v := range importMap {
		imports += v
	}

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "serviceTemplate").Parse(serviceTemplate))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":      strings.ToLower(svName),
		"subSvName":   strings.ToLower(subSvName),
		"imports":     imports,
		"serviceName": front.AssembleService(svName, subSvName, nil),
		"interfaces":  interfaces,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
