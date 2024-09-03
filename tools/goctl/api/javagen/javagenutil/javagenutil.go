package javagenutil

import (
	"bufio"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"strings"
)

func GoTypeToJavaConverter(goType string) string {
	switch goType {
	case "string":
		return "String"
	case "data.Date":
		return "Date"
	case "int8", "int32", "int":
		return "int"
	case "int64":
		return "Long"
	case "byte":
		return "byte"
	case "bool":
		return "boolean"
	case "float32":
		return "float"
	case "float64":
		return "double"
	default:
		return ""
	}
}

func JavaTypeToGoConverter(javaType string) string {
	switch javaType {
	case "String":
		return "string"
	case "Date":
		return "date.Date"
	case "Integer", "int":
		return "int"
	case "Long", "long":
		return "int64"
	case "Byte", "char", "byte":
		return "byte"
	case "Boolean", "boolean":
		return "bool"
	case "Short":
		return "int32"
	case "Float", "float":
		return "float32"
	case "Double", "double":
		return "float64"
	default:
		return ""
	}
}

func HeadUpCase(str string) string {
	if len(str) > 0 && "a" <= str[:1] && str[:1] <= "z" {
		bs := []byte(str)
		bs[0] -= 32
		return string(bs)
	}
	if len(str) > 0 && "A" <= str[:1] && str[:1] <= "Z" {
		return str
	}
	return ""
}

func HeadLowCase(str string) string {
	if len(str) > 0 && "A" <= str[:1] && str[:1] <= "Z" {
		bs := []byte(str)
		bs[0] += 32
		return string(bs)
	}
	if len(str) > 0 && "a" <= str[:1] && str[:1] <= "z" {
		return str
	}
	return ""
}

func TagParse(tag string) (tarMap map[string]string) {
	if len(tag) == 0 {
		return tarMap
	}

	tarMap = make(map[string]string)
	for _, des := range strings.Split(strings.Trim(tag, "`"), ",") {
		t := strings.Split(des, ":")
		if len(t) == 2 {
			tarMap[strings.Trim(strings.TrimSpace(t[0]), "\"")] = strings.Trim(strings.TrimSpace(t[1]), "\"")
		} else if len(t) == 1 {
			tarMap[strings.Trim(strings.TrimSpace(t[0]), "\"")] = ""
		} else if len(t) > 2 {
			panic(errors.New("tag data error"))
		}
	}

	return tarMap
}

func GetKeyFromTag(tag, key string) string {
	tarMap := TagParse(tag)
	return tarMap[key]
}

func JudgeSubdomainFunc(funcName string) bool {
	return strings.HasPrefix(funcName, "create") ||
		strings.HasPrefix(funcName, "update") ||
		strings.HasPrefix(funcName, "delete")
}

func AssembleBoImport(projName, subProjName, comporent, boType string) string {
	if boType == "front" {
		str := "import " + strings.ToLower(projName) + "." + strings.ToLower(subProjName)
		if strings.HasSuffix(comporent, "AO") || strings.HasSuffix(comporent, "Ao") {
			return str + ".model.ao." + comporent + ";\n"
		}
		if strings.HasSuffix(comporent, "DTO") || strings.HasSuffix(comporent, "Dto") {
			return str + ".model.dto." + comporent + ";\n"
		}
		if strings.HasSuffix(comporent, "VO") || strings.HasSuffix(comporent, "Vo") {
			return str + ".model.vo." + comporent + ";\n"
		}
	}

	if boType == "subdomain" {
		str := "import " + strings.ToLower(projName)
		if strings.HasSuffix(comporent, "DTO") || strings.HasSuffix(comporent, "Dto") {
			return str + ".application.dto." + comporent + ";\n"
		}
		if strings.HasSuffix(comporent, "ROOT") || strings.HasSuffix(comporent, "Root") {
			return str + ".domain." + subProjName + "context.beanclass." + comporent + ";\n"
		}
	}

	//if strings.HasPrefix(comporent, "Query") && strings.HasSuffix(comporent, "RemoteService") {
	//	return str + ".remote.datatransfer." + comporent + ";\n"
	//}
	//if !strings.HasPrefix(comporent, "Query") && strings.HasSuffix(comporent, "RemoteService") {
	//	return str + ".remote.severtransfer." + comporent + ";\n"
	//}
	//if strings.HasPrefix(comporent, "Query") && !strings.HasSuffix(comporent, "RemoteService") && strings.HasSuffix(comporent, "Service") {
	//	return str + ".service.query." + comporent + ";\n"
	//}
	//if !strings.HasPrefix(comporent, "Query") && !strings.HasSuffix(comporent, "RemoteService") && strings.HasSuffix(comporent, "Service") {
	//	return str + ".remote.application." + comporent + ";\n"
	//}

	return ""
}

func GetBo(api *spec.ApiSpec, boSuffix string, importMap *map[string]string) *spec.DefineStruct {
	for _, t := range api.Types {
		defineStruct := t.(spec.DefineStruct)
		if defineStruct.Type == boSuffix {
			if importMap != nil && (*importMap)[defineStruct.RawName] == "" {
				(*importMap)[defineStruct.RawName] = defineStruct.ImportPath
			}
			return &defineStruct
		}
	}
	panic("未找到对应Bo")
}

func ComputeImportPath(api *spec.ApiSpec, packetName, serviceType string) (imports string) {

	var importTypes []spec.Type
	for _, v := range api.Service.Routes() {
		if JudgeSubdomainFunc(v.Handler) {
			if v.RequestType != nil && !JudgeElemInSlice(importTypes, v.RequestType) {
				imports += AssembleBoImport(packetName, packetName, v.RequestTypeName(), serviceType)
				importTypes = append(importTypes, v.RequestType)
			}
			if v.ResponseType != nil && !JudgeElemInSlice(importTypes, v.ResponseType) {
				imports += AssembleBoImport(packetName, packetName, v.ResponseTypeName(), serviceType)
				importTypes = append(importTypes, v.ResponseType)
			}
		}
	}
	return
}

func JudgeElemInSlice(anys []spec.Type, elem spec.Type) bool {
	for _, v := range anys {
		if cmp.Equal(elem, v) {
			return true
		}
	}
	return false
}

func FindSvComm(key string, api *spec.ApiSpec) (*spec.SvComm, error) {
	for _, item := range api.SvComm {
		if item.Key == key {
			return &item, nil
		}
	}
	return nil, errors.New("No matches found " + key + " in []api.SvComm")
}

func AssembleParamType(t spec.Type, importMap *map[string]string) (string, string) {
	if t == nil {
		return "", ""
	}
	switch t.(type) {
	case spec.DefineStruct:
		defineStruct := t.(spec.DefineStruct)
		if importMap != nil && (*importMap)[defineStruct.RawName] == "" {
			(*importMap)[defineStruct.RawName] = defineStruct.ImportPath
		}
		return defineStruct.RawName, HeadLowCase(defineStruct.RawName)
	case spec.PrimitiveType:
		return GoTypeToJavaConverter(t.(spec.PrimitiveType).RawName), "param"
	case spec.ArrayType:
		if importMap != nil && (*importMap)["List"] == "" {
			(*importMap)["List"] = List
		}
		arrayType := t.(spec.ArrayType)
		paramType, param := AssembleParamType(arrayType.Value, importMap)
		return fmt.Sprintf("List<%s>", paramType), param + "List"
	case spec.MapType:
		if importMap != nil && (*importMap)["Map"] == "" {
			(*importMap)["Map"] = Map
		}
		paramType, param := AssembleParamType(t.(spec.MapType).Value, importMap)
		return fmt.Sprintf("Map<%s>", paramType), param + "Map"
	case spec.InterfaceType:
		return t.Name(), HeadLowCase(t.Name())
	case spec.PointerType:
		pointerType := t.(spec.PointerType)
		return AssembleParamType(pointerType.Type, importMap)
	default:
		panic("Unrecognized type")
	}

}

func FormatSource(source string) string {
	var builder strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(source))
	preIsBreakLine := false
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" && preIsBreakLine {
			continue
		}
		preIsBreakLine = text == ""
		builder.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return builder.String()
}

func ParamToSQLConverter(fields []string) []string {
	if fields == nil || len(fields) == 0 {
		return make([]string, 0)
	}
	var t []string
	for _, f := range fields {
		tf := ""
		bytes := []byte(f)
		for i, b := range bytes {
			if b >= 'A' && b <= 'Z' && i != 0 && (bytes[i-1] < 'A' || bytes[i-1] > 'Z') {
				tf += "_" + string(b)
			} else if b >= 'a' && b <= 'z' {
				tf += string(b - 32)
			} else {
				tf += string(b)
			}
		}
		t = append(t, tf)
	}
	return t
}

func SQLToParamConverter(sqls []string) []string {
	if sqls == nil || len(sqls) == 0 {
		return make([]string, 0)
	}
	t := make([]string, len(sqls))
	for _, s := range sqls {
		s = strings.ToLower(s)
		bytes := []byte(s)
		i := 0
		for {
			i = strings.IndexByte(s, '_')
			if i < 0 {
				break
			}
			i++
			if len(s) >= i && bytes[i] >= 'a' && bytes[i] <= 'z' {
				bytes[i] -= 32
			}
		}
		ts := string(bytes)
		ts = strings.ReplaceAll(ts, "_", "")
		t = append(t, ts)
	}
	return t
}
