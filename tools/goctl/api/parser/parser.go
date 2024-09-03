package parser

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/env"
	apiParser "github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/parser"
)

type parser struct {
	ast  *ast.Api
	spec *spec.ApiSpec
}

// Parse parses the api file
func Parse(filename string) (*spec.ApiSpec, error) {
	if env.UseExperimental() {
		return apiParser.Parse(filename, "")
	}

	astParser := ast.NewParser(ast.WithParserPrefix(filepath.Base(filename)), ast.WithParserDebug())
	parsedApi, err := astParser.Parse(filename)
	if err != nil {
		return nil, err
	}

	apiSpec := new(spec.ApiSpec)
	p := parser{ast: parsedApi, spec: apiSpec}
	err = p.convert2Spec()
	if err != nil {
		return nil, err
	}

	return apiSpec, nil
}

func parseContent(content string, skipCheckTypeDeclaration bool, filename ...string) (*spec.ApiSpec, error) {
	var astParser *ast.Parser
	if skipCheckTypeDeclaration {
		astParser = ast.NewParser(ast.WithParserSkipCheckTypeDeclaration())
	} else {
		astParser = ast.NewParser()
	}
	parsedApi, err := astParser.ParseContent(content, filename...)
	if err != nil {
		return nil, err
	}

	apiSpec := new(spec.ApiSpec)
	p := parser{ast: parsedApi, spec: apiSpec}
	err = p.convert2Spec()
	if err != nil {
		return nil, err
	}

	return apiSpec, nil
}

// ParseContent parses the api content
func ParseContent(content string, filename ...string) (*spec.ApiSpec, error) {
	return parseContent(content, false, filename...)
}

// ParseContentWithParserSkipCheckTypeDeclaration parses the api content with skip check type declaration
func ParseContentWithParserSkipCheckTypeDeclaration(content string, filename ...string) (*spec.ApiSpec, error) {
	return parseContent(content, true, filename...)
}

func (p parser) convert2Spec() error {
	p.fillInfo()
	p.fillSyntax()
	p.fillImport()
	p.fillSvComm()
	err := p.fillTypes()
	if err != nil {
		return err
	}

	return p.fillService()
}

func (p parser) fillInfo() {
	properties := make(map[string]string)
	if p.ast.Info != nil {
		for _, kv := range p.ast.Info.Kvs {
			properties[kv.Key.Text()] = kv.Value.Text()
		}
	}
	p.spec.Info.Properties = properties
}

func (p parser) fillSyntax() {
	if p.ast.Syntax != nil {
		p.spec.Syntax = spec.ApiSyntax{
			Version: p.ast.Syntax.Version.Text(),
			Doc:     p.stringExprs(p.ast.Syntax.DocExpr),
			Comment: p.stringExprs([]ast.Expr{p.ast.Syntax.CommentExpr}),
		}
	}
}

func (p parser) fillImport() {
	if len(p.ast.Import) > 0 {
		for _, item := range p.ast.Import {
			p.spec.Imports = append(p.spec.Imports, spec.Import{
				Value:   item.Value.Text(),
				Doc:     p.stringExprs(item.DocExpr),
				Comment: p.stringExprs([]ast.Expr{item.CommentExpr}),
			})
		}
	}
}

func (p parser) fillSvComm() {
	for _, item := range p.ast.SvComm {
		p.spec.SvComm = append(p.spec.SvComm, spec.SvComm{
			Key:   item.SvCommToken.Text(),
			Value: strings.Trim(item.SvCommValue.Text(), "\""),
			Doc:   p.stringExprs(item.DocExpr),
		})
	}
}

func (p parser) fillTypes() error {

	var boType = ""
	for i, comm := range p.spec.SvComm {
		if p.spec.SvComm[i].Key == "svname" {
			value := strings.TrimSpace(comm.Value)
			boType = value[strings.LastIndex(strings.TrimSpace(comm.Value), ".")+1:]
		}

	}
	if !(strings.ToUpper(boType) == "FRONT" || strings.ToUpper(boType) == "SUBDOMAIN") {
		return fmt.Errorf("\"svname\" must end in front or subdomain.")
	}
	for _, item := range p.ast.Type {
		switch v := item.(type) {
		case *ast.TypeStruct:
			var members []spec.Member
			for _, item := range v.Fields {
				members = append(members, p.fieldToMember(item))
			}
			p.spec.Types = append(p.spec.Types, spec.DefineStruct{
				RawName: strings.Trim(strings.TrimSpace(v.Name.Text()), "/"),
				Members: members,
				Docs:    p.stringExprs(v.Doc()),
				Group: func() string {
					if strings.HasSuffix(v.Name.Text(), "DTO") || strings.HasSuffix(v.Name.Text(), "Dto") {
						return v.Name.Text()[:len(v.Name.Text())-3]
					}
					if strings.HasSuffix(v.Name.Text(), "AO") || strings.HasSuffix(v.Name.Text(), "Ao") || strings.HasSuffix(v.Name.Text(), "VO") || strings.HasSuffix(v.Name.Text(), "Vo") {
						return v.Name.Text()[:len(v.Name.Text())-2]
					}
					if strings.HasSuffix(v.Name.Text(), "ROOT") || strings.HasSuffix(v.Name.Text(), "Root") {
						return v.Name.Text()[:len(v.Name.Text())-4]
					}
					panic("Types other than AO, VO, DTO, ROOT are not supported")
				}(),
				Type: func() string {
					if strings.HasSuffix(v.Name.Text(), "DTO") || strings.HasSuffix(v.Name.Text(), "Dto") {
						return "DTO"
					}
					if strings.HasSuffix(v.Name.Text(), "AO") || strings.HasSuffix(v.Name.Text(), "Ao") {
						return "AO"
					}
					if strings.HasSuffix(v.Name.Text(), "VO") || strings.HasSuffix(v.Name.Text(), "Vo") {
						return "VO"
					}
					if strings.HasSuffix(v.Name.Text(), "ROOT") || strings.HasSuffix(v.Name.Text(), "Root") {
						return "ROOT"
					}
					panic("Types other than AO, VO, DTO, ROOT are not supported")
				}(),
				ImportPath: func() string {
					svName, err := javagenutil.FindSvComm("svname", p.spec)
					if err != nil {
						panic(err)
					}
					subSvComm, err := javagenutil.FindSvComm("subsvname", p.spec)
					if err != nil {
						panic(err)
					}
					return javagenutil.AssembleBoImport(svName.Value, subSvComm.Value, v.Name.Text(), strings.ToLower(boType))
				}(),
				Desc: func() string {
					if v.AtType.LineDoc != nil {
						return strings.Trim(v.AtType.LineDoc.Text(), "\"")
					}
					for _, item := range v.AtType.Kv {
						if item.Key.Text() == "name" {
							return strings.Trim(item.Value.Text(), "\"")
						}

					}
					return ""
				}(),
			})
		default:
			return fmt.Errorf("unknown type %+v", v)
		}
	}

	var types []spec.Type
	for _, item := range p.spec.Types {
		switch v := (item).(type) {
		case spec.DefineStruct:
			var members []spec.Member
			for _, member := range v.Members {
				switch v := member.Type.(type) {
				case spec.DefineStruct:
					tp, err := p.findDefinedType(v.RawName)
					if err != nil {
						return err
					}

					member.Type = *tp
				}
				members = append(members, member)
			}
			v.Members = members
			types = append(types, v)
		default:
			return fmt.Errorf("unknown type %+v", v)
		}
	}
	p.spec.Types = types

	return nil
}

func (p parser) findDefinedType(name string) (*spec.Type, error) {
	for _, item := range p.spec.Types {
		if _, ok := item.(spec.DefineStruct); ok {
			if item.Name() == name {
				return &item, nil
			}
		}
		if _, ok := item.(spec.ArrayType); ok {
			if item.Name() == name {
				return &item, nil
			}
		}
	}

	return nil, fmt.Errorf("type %s not defined", name)
}

func (p parser) fieldToMember(field *ast.TypeField) spec.Member {
	name := ""
	tag := ""
	if !field.IsAnonymous {
		name = field.Name.Text()
		//if field.Tag == nil { //为什么必须要要填写Tag？
		//	panic(fmt.Sprintf("error: line %d:%d field %s has no tag",
		//		field.Name.Line(), field.Name.Column(), field.Name.Text()))
		//}

		if field.Tag != nil {
			tag = field.Tag.Text()
		}
	}
	return spec.Member{
		Name:     name,
		Type:     p.astTypeToSpec(field.DataType),
		Tag:      tag,
		Comment:  p.commentExprs(field.Comment()),
		Docs:     p.stringExprs(field.Doc()),
		IsInline: field.IsAnonymous,
	}
}

func (p parser) astTypeToSpec(in ast.DataType) spec.Type {
	switch v := (in).(type) {
	case *ast.Literal:
		raw := v.Literal.Text()
		if api.IsBasicType(raw) {
			return spec.PrimitiveType{
				RawName: raw,
			}
		}

		return spec.DefineStruct{
			RawName: raw,
		}
	case *ast.Interface:
		return spec.InterfaceType{RawName: v.Literal.Text()}
	case *ast.Map:
		return spec.MapType{RawName: v.MapExpr.Text(), Key: v.Key.Text(), Value: p.astTypeToSpec(v.Value)}
	case *ast.Array:
		return spec.ArrayType{RawName: v.ArrayExpr.Text(), Value: p.astTypeToSpec(v.Literal)}
	case *ast.Pointer:
		raw := v.Name.Text()
		if api.IsBasicType(raw) {
			return spec.PointerType{RawName: v.PointerExpr.Text(), Type: spec.PrimitiveType{RawName: raw}}
		}

		return spec.PointerType{RawName: v.PointerExpr.Text(), Type: spec.DefineStruct{RawName: raw}}
	}

	panic(fmt.Sprintf("unspported type %+v", in))
}

func (p parser) stringExprs(docs []ast.Expr) []string {
	var result []string
	for _, item := range docs {
		if item == nil {
			continue
		}

		result = append(result, item.Text())
	}
	return result
}

func (p parser) commentExprs(comment ast.Expr) string {
	if comment == nil {
		return ""
	}

	return comment.Text()
}

func (p parser) fillService() error {
	var groups []spec.Group
	for _, item := range p.ast.Service {
		var group spec.Group
		p.fillAtServer(item, &group)

		for _, astRoute := range item.ServiceApi.ServiceRoute {
			route := spec.Route{
				AtServerAnnotation: spec.Annotation{},
				Method:             astRoute.Route.Method.Text(),
				Path:               astRoute.Route.Path.Text(),
				Doc:                p.stringExprs(astRoute.Route.DocExpr),
				Comment:            p.stringExprs([]ast.Expr{astRoute.Route.CommentExpr}),
			}
			if astRoute.AtHandler != nil {
				route.Handler = astRoute.AtHandler.Name.Text()
				if astRoute.AtDesc != nil {
					route.AtDesc = astRoute.AtDesc.Name.Text()
				}
				route.HandlerDoc = append(route.HandlerDoc, p.stringExprs(astRoute.AtHandler.DocExpr)...)
				route.HandlerComment = append(route.HandlerComment, p.stringExprs([]ast.Expr{astRoute.AtHandler.CommentExpr})...)
			}

			err := p.fillRouteAtServer(astRoute, &route)
			if err != nil {
				return err
			}

			if astRoute.Route.Req != nil {
				route.RequestType = p.astTypeToSpec(astRoute.Route.Req.Name)
			}
			if astRoute.Route.Reply != nil {
				route.ResponseType = p.astTypeToSpec(astRoute.Route.Reply.Name)
			}
			if astRoute.AtDoc != nil {
				properties := make(map[string]string)
				for _, kv := range astRoute.AtDoc.Kv {
					properties[kv.Key.Text()] = kv.Value.Text()
				}
				route.AtDoc.Properties = properties
				if astRoute.AtDoc.LineDoc != nil {
					route.AtDoc.Text = astRoute.AtDoc.LineDoc.Text()
				}
			}

			err = p.fillRouteType(&route)
			if err != nil {
				return err
			}

			group.Routes = append(group.Routes, route)
			name := item.ServiceApi.Name.Text()
			// todo 不支持不同名称的service块，后续待更改
			if len(p.spec.Service.Name) > 0 && p.spec.Service.Name != name {
				return fmt.Errorf("multiple service names defined %s and %s",
					name, p.spec.Service.Name)
			}

			p.spec.Service.Name = name
		}
		groups = append(groups, group)
	}
	p.spec.Service.Groups = groups

	return nil
}

func (p parser) fillRouteAtServer(astRoute *ast.ServiceRoute, route *spec.Route) error {
	if astRoute.AtServer != nil {
		properties := make(map[string]string)
		for _, kv := range astRoute.AtServer.Kv {
			properties[kv.Key.Text()] = kv.Value.Text()
		}
		route.AtServerAnnotation.Properties = properties
		if len(route.Handler) == 0 {
			route.Handler = properties["handler"]
		}
		if len(route.Handler) == 0 {
			return fmt.Errorf("missing handler annotation for %q", route.Path)
		}

		for _, char := range route.Handler {
			if !unicode.IsDigit(char) && !unicode.IsLetter(char) {
				return fmt.Errorf("route [%s] handler [%s] invalid, handler name should only contains letter or digit",
					route.Path, route.Handler)
			}
		}
	}
	return nil
}

func (p parser) fillAtServer(item *ast.Service, group *spec.Group) {
	if item.AtServer != nil {
		properties := make(map[string]string)
		for _, kv := range item.AtServer.Kv {
			properties[kv.Key.Text()] = kv.Value.Text()
		}
		group.Annotation.Properties = properties
	}
}

func (p parser) fillRouteType(route *spec.Route) error {
	if route.RequestType != nil {
		switch route.RequestType.(type) {
		case spec.DefineStruct:
			tp, err := p.findDefinedType(route.RequestType.Name())
			if err != nil {
				return err
			}

			route.RequestType = *tp
		case spec.ArrayType:
			tp, err := p.findDefinedType(strings.TrimPrefix(route.RequestType.Name(), "[]"))
			if err != nil {
				return err
			}
			arrayType := route.RequestType.(spec.ArrayType)
			arrayType.Value = *tp
			route.RequestType = arrayType
		}

	}

	if route.ResponseType != nil {
		switch route.ResponseType.(type) {
		case spec.DefineStruct:
			tp, err := p.findDefinedType(route.ResponseType.Name())
			if err != nil {
				return err
			}

			route.ResponseType = *tp
		case spec.ArrayType:
			tp, err := p.findDefinedType(strings.TrimPrefix(route.ResponseType.Name(), "[]"))
			if err != nil {
				return err
			}
			arrayType := route.ResponseType.(spec.ArrayType)
			arrayType.Value = *tp
			route.ResponseType = arrayType
		}
	}

	return nil
}
