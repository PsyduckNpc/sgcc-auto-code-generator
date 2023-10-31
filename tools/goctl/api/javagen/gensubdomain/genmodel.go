package gensubdomain

import (
	"bytes"
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	apiutil "github.com/zeromicro/go-zero/tools/goctl/api/util"
	"strings"
	"text/template"
)

//go:embed subdomaintpl/DTO.tpl
var dtoTemplate string

//go:embed subdomaintpl/ROOT.tpl
var rootTemplate string

func GenModel(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool) error {
	defer func() { ch <- true }()
	for _, t := range api.Types {
		if err := typeCreateWith(dir, t, svName, subSvName); err != nil {
			return err
		}
	}

	return nil
}

func typeCreateWith(dir string, t spec.Type, svName, subSvName string) error {
	defineStruct := t.(spec.DefineStruct)
	rawName := defineStruct.RawName

	var fields []javagenutil.Field
	var path string
	var javaFile string
	var toStringFlag bool
	var imports string

	if defineStruct.Type == "DTO" {
		path = dir + "/application/dto"
	} else if defineStruct.Type == "ROOT" {
		path = dir + "/domain/" + strings.ToLower(subSvName) + "context/beanclass"
	} else {
		return nil
	}

	javaFile = rawName + ".java"
	fp, created, err := apiutil.MaybeCreateFile(path, "", javaFile)
	if err != nil {
		return err
	} //创建出现错误
	if !created {
		return nil
	} //已经存在文件
	defer fp.Close()

	//填充字段
	for _, d := range defineStruct.Members {
		f := javagenutil.Field{
			Desc:                 javagenutil.TagParse(d.Tag)["desc"],
			FieldType:            javagenutil.GoTypeToJavaConverter(d.Type.Name()),
			HeadLowCaseFieldName: javagenutil.HeadLowCase(d.Name),
			HeadUpCaseFieldName:  javagenutil.HeadUpCase(d.Name),
		}
		toStringFlag = true
		fields = append(fields, f)
	}

	var tpl *template.Template
	if defineStruct.Type == "DTO" {
		tpl = template.Must(template.New("dtoTemplate").Parse(dtoTemplate)) //解析模板了
	} else if defineStruct.Type == "ROOT" {
		tpl = template.Must(template.New("rootTemplate").Parse(rootTemplate)) //解析模板了
	}

	var tmplBytes bytes.Buffer
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":            svName,
		"subSvName":         subSvName,
		"imports":           imports,
		"headLowCaseBOName": javagenutil.HeadLowCase(defineStruct.RawName),
		"desc":              defineStruct.Desc,
		"headUpCaseBOName":  javagenutil.HeadUpCase(defineStruct.RawName),
		"extend":            false,
		"extendName":        "",
		"implements":        false,
		"implementContent":  "",
		"fields":            fields,
		"toStringFlag":      toStringFlag,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化
	return err

}
