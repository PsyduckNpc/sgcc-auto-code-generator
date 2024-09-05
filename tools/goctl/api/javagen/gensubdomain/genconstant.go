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

//go:embed subdomaintpl/Constant.tpl
var constantTemplate string

func GenConstant(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool, constants *map[string]javagenutil.Logic) error {

	defer func() { ch <- true }()
	var imports string
	var injects []javagenutil.Logic

	javaFile := subdomain.AssembleConstant(svName, subSvName, nil) + ".java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/constants/"+subSvName+"constants/", "", javaFile)
	if err != nil {
		return err
	} //创建出现错误
	if !created {
		return nil
	} //已经存在文件
	defer fp.Close()

	for _, v := range *constants {
		injects = append(injects, v)
	}

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "constantTemplate").Parse(constantTemplate))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":       strings.ToLower(svName),
		"subSvName":    strings.ToLower(subSvName),
		"constantName": subdomain.AssembleConstant(svName, subSvName, nil),
		"imports":      imports,
		"constants":    injects,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
