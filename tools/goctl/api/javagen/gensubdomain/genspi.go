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

//go:embed subdomaintpl/Spi.tpl
var spiTemplate string

func GenSpi(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool) error {

	defer func() { ch <- true }()
	var imports string
	importMap := make(map[string]string)

	javaFile := subdomain.AssembleSpi(svName, subSvName, nil) + ".java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/front/controller", "", javaFile)
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
				DTOType:       reqType,
				DTOName:       reqName,
				Route:         group + v.Path,
			}
			inters = append(inters, inter)
		}
	}

	javagenutil.GetBo(api, "DTO", &importMap)
	for _, v := range importMap {
		imports += v
	}

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "spiTemplate").Parse(spiTemplate))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":        strings.ToLower(svName),
		"subSvName":     strings.ToLower(subSvName),
		"subdomainName": "emss-scp-" + strings.ToLower(subSvName) + "subdomin",
		"className":     subdomain.AssembleSpi(svName, subSvName, nil),
		"imports":       imports,
		"inters":        inters,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
