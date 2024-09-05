package api

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Part 10
// The apiparser_parser.go file was split into multiple files because it
// was too large and caused a possible memory overflow during goctl installation.

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *ServiceNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceName() (localctx IServiceNameContext) {
	this := p
	_ = this

	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ApiParserParserRULE_serviceName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID {
		{
			p.SetState(356)
			p.Match(ApiParserParserID)
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ApiParserParserT__10 {
			{
				p.SetState(357)
				p.Match(ApiParserParserT__10)
			}

		}

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPathItem() []IPathItemContext
	PathItem(i int) IPathItemContext

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllPathItem() []IPathItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPathItemContext); ok {
			len++
		}
	}

	tst := make([]IPathItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPathItemContext); ok {
			tst[i] = t.(IPathItemContext)
			i++
		}
	}

	return tst
}

func (s *PathContext) PathItem(i int) IPathItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathItemContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Path() (localctx IPathContext) {
	this := p
	_ = this

	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ApiParserParserRULE_path)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(379)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ApiParserParserT__11 || _la == ApiParserParserT__12 {
			p.SetState(379)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case ApiParserParserT__11:
				{
					p.SetState(364)
					p.Match(ApiParserParserT__11)
				}

				{
					p.SetState(365)
					p.PathItem()
				}
				p.SetState(370)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == ApiParserParserT__10 {
					{
						p.SetState(366)
						p.Match(ApiParserParserT__10)
					}
					{
						p.SetState(367)
						p.PathItem()
					}

					p.SetState(372)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}

			case ApiParserParserT__12:
				{
					p.SetState(373)
					p.Match(ApiParserParserT__12)
				}

				{
					p.SetState(374)
					p.PathItem()
				}
				p.SetState(377)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == ApiParserParserT__10 {
					{
						p.SetState(375)
						p.Match(ApiParserParserT__10)
					}
					{
						p.SetState(376)
						p.PathItem()
					}

				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(381)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(383)
			p.Match(ApiParserParserT__11)
		}

	}

	return localctx
}

// IPathItemContext is an interface to support dynamic dispatch.
type IPathItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllLetterOrDigit() []antlr.TerminalNode
	LetterOrDigit(i int) antlr.TerminalNode

	// IsPathItemContext differentiates from other interfaces.
	IsPathItemContext()
}

type PathItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathItemContext() *PathItemContext {
	var p = new(PathItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_pathItem
	return p
}

func (*PathItemContext) IsPathItemContext() {}

func NewPathItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathItemContext {
	var p = new(PathItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_pathItem

	return p
}

func (s *PathItemContext) GetParser() antlr.Parser { return s.parser }

func (s *PathItemContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *PathItemContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *PathItemContext) AllLetterOrDigit() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserLetterOrDigit)
}

func (s *PathItemContext) LetterOrDigit(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserLetterOrDigit, i)
}

func (s *PathItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPathItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) PathItem() (localctx IPathItemContext) {
	this := p
	_ = this

	localctx = NewPathItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ApiParserParserRULE_pathItem)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID || _la == ApiParserParserLetterOrDigit {
		{
			p.SetState(386)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ApiParserParserID || _la == ApiParserParserLetterOrDigit) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *ApiParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 20:
		var t *FieldContext = nil
		if localctx != nil {
			t = localctx.(*FieldContext)
		}
		return p.Field_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ApiParserParser) Field_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return isNormal(p)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
