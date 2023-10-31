package genfront

import (
	"bytes"
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	apiutil "github.com/zeromicro/go-zero/tools/goctl/api/util"
	"text/template"
)

//go:embed fronttpl/AO.tpl
var aoTemplate string

//go:embed fronttpl/DTO.tpl
var dtoTemplate string

//go:embed fronttpl/VO.tpl
var voTemplate string

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

	if defineStruct.Type == "AO" {
		path = dir + "/" + subSvName + "/model/ao"
	} else if defineStruct.Type == "DTO" {
		path = dir + "/" + subSvName + "/model/dto"
	} else if defineStruct.Type == "VO" {
		path = dir + "/" + subSvName + "/model/vo"
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
	if defineStruct.Type == "AO" {
		tpl = template.Must(template.New("aoTemplate").Parse(aoTemplate)) //解析模板了
	} else if defineStruct.Type == "DTO" {
		tpl = template.Must(template.New("dtoTemplate").Parse(dtoTemplate)) //解析模板了
	} else if defineStruct.Type == "VO" {
		tpl = template.Must(template.New("voTemplate").Parse(voTemplate)) //解析模板了
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
