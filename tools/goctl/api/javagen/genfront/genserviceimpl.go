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

//go:embed fronttpl/ServiceImpl.tpl
var serviceTemplateImpl string

func GenServiceImpl(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool) error {

	defer func() { ch <- true }()

	javaFile := front.AssembleServiceImpl(svName, subSvName, nil) + ".java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/"+subSvName+"/service/application/impl", "", javaFile)
	if err != nil {
		return err
	} //创建出现错误
	if !created {
		return nil
	} //已经存在文件
	defer fp.Close()

	var imports string
	var injects []javagenutil.Inject
	var logics []javagenutil.Logic
	importMap := make(map[string]string)

	injects = append(injects, javagenutil.Inject{
		InjectServiceName:   front.AssembleRemoteService(svName, subSvName, &importMap),
		VariableServiceName: javagenutil.HeadLowCase(front.AssembleRemoteService(svName, subSvName, nil)),
	})

	for _, v := range api.Service.Routes() {
		if !javagenutil.JudgeSubdomainFunc(v.Handler) {
			continue
		}
		var fields []javagenutil.Field
		respType, _ := javagenutil.AssembleParamType(v.ResponseType, &importMap)
		reqType, reqName := javagenutil.AssembleParamType(v.RequestType, &importMap)

		voStruct := javagenutil.GetBo(api, "VO", nil)
		suffix := v.ResponseType != nil && strings.HasSuffix(v.ResponseType.Name(), voStruct.RawName)
		if suffix {
			importMap[voStruct.RawName] = voStruct.ImportPath
		}
		for _, v := range voStruct.Members {
			if code := javagenutil.GetKeyFromTag(v.Tag, "code"); code != "" {
				field := javagenutil.Field{
					Desc:                 javagenutil.GetKeyFromTag(v.Tag, "desc"),
					FieldType:            v.Type.Name(),
					HeadLowCaseFieldName: javagenutil.HeadLowCase(v.Name),
					HeadUpCaseFieldName:  javagenutil.HeadUpCase(v.Name),
					StandardCode:         code,
				}
				fields = append(fields, field)
			}
		}

		logics = append(logics, javagenutil.Logic{
			RespType:      respType,
			InterfaceName: v.Handler,
			ReqType:       reqType,
			AOName:        reqName,
			Desc:          v.AtDesc,
			ReqListFlag: func() bool {
				if v.RequestType == nil {
					return false
				}
				return strings.HasPrefix(v.RequestType.Name(), "[]")
			}(),
			DTOType:             javagenutil.HeadUpCase(javagenutil.GetBo(api, "DTO", &importMap).RawName),
			DTOName:             javagenutil.HeadLowCase(javagenutil.GetBo(api, "DTO", &importMap).RawName),
			VariableServiceName: javagenutil.HeadLowCase(front.AssembleRemoteService(svName, subSvName, &importMap)),
			RespListFlag: func() bool {
				if v.ResponseType == nil {
					return false
				}
				return strings.HasPrefix(v.ResponseType.Name(), "[]")
			}(),
			VOType:        javagenutil.HeadUpCase(javagenutil.GetBo(api, "VO", &importMap).RawName),
			CodeConverter: len(fields) > 0 && suffix,
			Fields:        fields,
		})
	}
	serviceName := front.AssembleQueryService(svName, subSvName, &importMap)

	importMap["Resource"] = javagenutil.Resource
	for _, v := range logics {
		if v.ReqListFlag && importMap["ArrayList"] == "" {
			importMap["ArrayList"] = javagenutil.ArrayList
		}
		if v.ReqType != "" && importMap["BeanCopyUtils"] == "" {
			importMap["BeanCopyUtils"] = javagenutil.BeanCopyUtils
			importMap["BusinessObjectConvertUtil"] = javagenutil.BusinessObjectConvertUtil
			importMap["JSONObject"] = javagenutil.JSONObject
		}

	}

	for _, v := range importMap {
		imports += v
	}

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "queryServiceTemplateImpl").Parse(queryServiceTemplateImpl))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":          strings.ToLower(svName),
		"subSvName":       strings.ToLower(subSvName),
		"imports":         imports,
		"serviceImplName": front.AssembleServiceImpl(svName, subSvName, nil),
		"serviceName":     serviceName,
		"injects":         injects,
		"logics":          logics,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
