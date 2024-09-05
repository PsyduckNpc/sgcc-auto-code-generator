package ast

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
)

// SvCommExpr describes syntax for api
type SvCommExpr struct {
	SvCommToken Expr
	SvCommValue Expr
	DocExpr     []Expr
	CommentExpr Expr
}

// VisitSvcommLit implements from api.BaseApiParserVisitor
func (v *ApiVisitor) VisitSvcommLit(ctx *api.SvcommLitContext) any {
	svCommToken := v.newExprWithToken(ctx.GetSvcommToken())
	svCommValue := v.newExprWithToken(ctx.GetSvcommValue())
	return []SvCommExpr{{
		SvCommToken: svCommToken,
		SvCommValue: svCommValue,
		DocExpr:     v.getDoc(ctx),
		CommentExpr: v.getComment(ctx),
	}}
}

// Format provides a formatter for api command, now nothing to do
func (s *SvCommExpr) Format() error {
	// todo
	return nil
}

// Equal compares whether the element literals in two SyntaxExpr are equal
func (s *SvCommExpr) Equal(v any) bool {
	if v == nil {
		return false
	}

	svComm, ok := v.(*SvCommExpr)
	if !ok {
		return false
	}

	if !EqualDoc(s, svComm) {
		return false
	}

	return s.SvCommToken.Equal(svComm.SvCommToken) &&
		s.SvCommValue.Equal(svComm.SvCommValue)
}

// Doc returns the document of SyntaxExpr, like // some text
func (s *SvCommExpr) Doc() []Expr {
	return s.DocExpr
}

// Comment returns the comment of SyntaxExpr, like // some text
func (s *SvCommExpr) Comment() Expr {
	return s.CommentExpr
}
