package gensubdomain

import (
	"bytes"
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil/subdomain"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	apiutil "github.com/zeromicro/go-zero/tools/goctl/api/util"
	"strings"
	"text/template"
)

//go:embed subdomaintpl/DomainService.tpl
var domainServiceTemplate string

func GenDomainService(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool) error {

	defer func() { ch <- true }()
	var imports string

	javaFile := javagenutil.HeadUpCase(subSvName) + "DomainService.java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/domain/"+subSvName+"context/service", "", javaFile)
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
		//if !javagenutil.JudgeSubdomainFunc(v.Handler) {
		//	continue
		//}
		inter := javagenutil.Inter{}
		inter.InterfaceName = v.Handler
		reqType, param := javagenutil.AssembleParamType(v.MiddleType, &importMap)
		inter.ReqType = reqType
		inter.ReqParamName = param
		interfaces = append(interfaces, inter)
	}

	for _, v := range importMap {
		imports += v
	}

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "domainServiceTemplate").Parse(domainServiceTemplate))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":      strings.ToLower(svName),
		"subSvName":   strings.ToLower(subSvName),
		"imports":     imports,
		"serviceName": subdomain.AssembleDomainService(svName, subSvName, nil),
		"interfaces":  interfaces,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
