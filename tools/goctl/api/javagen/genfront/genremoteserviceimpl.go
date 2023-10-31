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

//go:embed fronttpl/RemoteServiceImpl.tpl
var remoteServiceImplTemplate string

func GenRemoteServiceImpl(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool) error {

	defer func() { ch <- true }()
	var imports string
	importMap := make(map[string]string)

	javaFile := javagenutil.HeadUpCase(subSvName) + "RemoteServiceImpl.java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/"+subSvName+"/remote/servertransfer/impl", "", javaFile)
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
		for _, v := range g.Routes {
			if !javagenutil.JudgeSubdomainFunc(v.Handler) {
				continue
			}
			middleType, middleTypeParam := javagenutil.AssembleParamType(v.MiddleType, &importMap)
			inter := javagenutil.Inter{
				InterfaceName: v.Handler,
				ReqType:       middleType,
				DTOName:       middleTypeParam,
			}
			inters = append(inters, inter)
		}

	}

	interfaceImplName := front.AssembleRemoteService(svName, subSvName, &importMap)
	for _, v := range importMap {
		imports += v
	}

	imports += javagenutil.GetBo(api, "DTO", nil).ImportPath

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "remoteServiceTemplateImpl").Parse(remoteServiceImplTemplate))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":            strings.ToLower(svName),
		"subSvName":         strings.ToLower(subSvName),
		"imports":           imports,
		"interfaceName":     front.AssembleRemoteServiceImpl(svName, subSvName, nil),
		"interfaceImplName": interfaceImplName,
		"inters":            inters,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
