package gensubdomain

import (
	"bytes"
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil/front"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil/subdomain"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	apiutil "github.com/zeromicro/go-zero/tools/goctl/api/util"
	"strings"
	"text/template"
)

//go:embed subdomaintpl/DomainServiceImpl.tpl
var domainServiceImplTemplate string

func GenDomainServiceImpl(dir, svName, subSvName string, api *spec.ApiSpec, ch chan bool, constants *map[string]javagenutil.Logic) error {

	defer func() { ch <- true }()

	javaFile := subdomain.AssembleDomainServiceImpl(svName, subSvName, nil) + ".java"
	fp, created, err := apiutil.MaybeCreateFile(dir+"/domain/"+subSvName+"context/service/impl", "", javaFile)
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
		InjectServiceName:   subdomain.AssembleDomainService(svName, subSvName, &importMap),
		VariableServiceName: javagenutil.HeadLowCase(subdomain.AssembleDomainService(svName, subSvName, nil)),
	})

	for _, v := range api.Service.Routes() {
		reqType, reqName := javagenutil.AssembleParamType(v.MiddleType, &importMap)
		root := javagenutil.GetBo(api, "ROOT", &importMap)
		var primary spec.Member
		var id string
		for _, m := range root.Members {
			tarMap := javagenutil.TagParse(m.Tag)
			if _, ok := tarMap["primary"]; ok {
				primary = m
				id = tarMap["primary"]
			}
		}

		primaryKey := javagenutil.HeadUpCase(primary.Name)
		primaryKeySQL := func() string {
			if id != "" {
				return id
			} else {
				toSQLConverter := javagenutil.ParamToSQLConverter([]string{strings.TrimLeft(v.MiddleType.Name(), "[]"), primary.Name})
				return toSQLConverter[0] + "_" + toSQLConverter[1]
			}
		}()
		if constants != nil && (*constants)[root.Group].PrimaryKeySQL == "" {
			(*constants)[root.Group] = javagenutil.Logic{PrimaryKey: primaryKey, PrimaryKeySQL: primaryKeySQL, TableSQL: javagenutil.ParamToSQLConverter([]string{root.Group})[0]}
		}

		logics = append(logics, javagenutil.Logic{
			Operate:       subdomain.GetOperate(v.Handler),
			PrimaryKey:    primaryKey,
			PrimaryKeySQL: primaryKeySQL,
			InterfaceName: v.Handler,
			DTOType:       reqType,
			DTOName:       reqName,
			RootType:      javagenutil.GetBo(api, "ROOT", &importMap).RawName,
			Desc:          v.AtDesc,
			ReqListFlag: func() bool {
				if v.RequestType == nil {
					return false
				}
				return strings.HasPrefix(v.RequestType.Name(), "[]")
			}(),
			VariableServiceName: javagenutil.HeadLowCase(subdomain.AssembleDomainService(svName, subSvName, &importMap)),
		})
	}
	serviceName := front.AssembleQueryService(svName, subSvName, &importMap)

	importMap["Resource"] = javagenutil.Resource
	for _, v := range logics {
		if v.ReqListFlag && importMap["ArrayList"] == "" {
			importMap["ArrayList"] = javagenutil.ArrayList
		}
		if v.RootType != "" && importMap["BeanCopyUtils"] == "" {
			importMap["BeanCopyUtils"] = javagenutil.BeanCopyUtils
			importMap["BusinessObjectConvertUtil"] = javagenutil.BusinessObjectConvertUtil
			importMap["JSONObject"] = javagenutil.JSONObject
			importMap["Constant"] = "import " + svName + ".infrastructure.constant." + javagenutil.HeadUpCase(subSvName) + "Constant;\n"
		}
	}

	for _, v := range importMap {
		imports += v
	}

	var tmplBytes bytes.Buffer
	tpl := template.Must(template.New(dir + "domainServiceImplTemplate").Parse(domainServiceImplTemplate))
	err = tpl.Execute(&tmplBytes, map[string]any{ //渲染模板了
		"svName":              strings.ToLower(svName),
		"allLowSubSvName":     strings.ToLower(subSvName),
		"headUpCaseSubSvName": javagenutil.HeadUpCase(subSvName) + "Constant",
		"imports":             imports,
		"className":           front.AssembleServiceImpl(svName, subSvName, nil),
		"interfaceName":       serviceName,
		"injects":             injects,
		"logics":              logics,
	})
	if err != nil {
		return err
	}

	_, err = fp.WriteString(javagenutil.FormatSource(tmplBytes.String())) //格式化

	return err
}
