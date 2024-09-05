package api

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// Part 2
// The apiparser_parser.go file was split into multiple files because it
// was too large and caused a possible memory overflow during goctl installation.

func (s *ImportBlockValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportBlockValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportBlockValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportBlockValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ImportBlockValue() (localctx IImportBlockValueContext) {
	this := p
	_ = this

	localctx = NewImportBlockValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ApiParserParserRULE_importBlockValue)

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
	{
		p.SetState(120)
		p.ImportValue()
	}

	return localctx
}

// IImportValueContext is an interface to support dynamic dispatch.
type IImportValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsImportValueContext differentiates from other interfaces.
	IsImportValueContext()
}

type ImportValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportValueContext() *ImportValueContext {
	var p = new(ImportValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_importValue
	return p
}

func (*ImportValueContext) IsImportValueContext() {}

func NewImportValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportValueContext {
	var p = new(ImportValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importValue

	return p
}

func (s *ImportValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserSTRING, 0)
}

func (s *ImportValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ImportValue() (localctx IImportValueContext) {
	this := p
	_ = this

	localctx = NewImportValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ApiParserParserRULE_importValue)

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
	checkImportValue(p)
	{
		p.SetState(123)
		p.Match(ApiParserParserSTRING)
	}

	return localctx
}

// IInfoSpecContext is an interface to support dynamic dispatch.
type IInfoSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInfoToken returns the infoToken token.
	GetInfoToken() antlr.Token

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetInfoToken sets the infoToken token.
	SetInfoToken(antlr.Token)

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	AllKvLit() []IKvLitContext
	KvLit(i int) IKvLitContext

	// IsInfoSpecContext differentiates from other interfaces.
	IsInfoSpecContext()
}

type InfoSpecContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	infoToken antlr.Token
	lp        antlr.Token
	rp        antlr.Token
}

func NewEmptyInfoSpecContext() *InfoSpecContext {
	var p = new(InfoSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_infoSpec
	return p
}

func (*InfoSpecContext) IsInfoSpecContext() {}

func NewInfoSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfoSpecContext {
	var p = new(InfoSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_infoSpec

	return p
}

func (s *InfoSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *InfoSpecContext) GetInfoToken() antlr.Token { return s.infoToken }

func (s *InfoSpecContext) GetLp() antlr.Token { return s.lp }

func (s *InfoSpecContext) GetRp() antlr.Token { return s.rp }

func (s *InfoSpecContext) SetInfoToken(v antlr.Token) { s.infoToken = v }

func (s *InfoSpecContext) SetLp(v antlr.Token) { s.lp = v }

func (s *InfoSpecContext) SetRp(v antlr.Token) { s.rp = v }

func (s *InfoSpecContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *InfoSpecContext) AllKvLit() []IKvLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKvLitContext); ok {
			len++
		}
	}

	tst := make([]IKvLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKvLitContext); ok {
			tst[i] = t.(IKvLitContext)
			i++
		}
	}

	return tst
}

func (s *InfoSpecContext) KvLit(i int) IKvLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKvLitContext); ok {
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

	return t.(IKvLitContext)
}

func (s *InfoSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfoSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfoSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitInfoSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) InfoSpec() (localctx IInfoSpecContext) {
	this := p
	_ = this

	localctx = NewInfoSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ApiParserParserRULE_infoSpec)
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
	match(p, "info")
	{
		p.SetState(126)

		var _m = p.Match(ApiParserParserID)

		localctx.(*InfoSpecContext).infoToken = _m
	}
	{
		p.SetState(127)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*InfoSpecContext).lp = _m
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID {
		{
			p.SetState(128)
			p.KvLit()
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*InfoSpecContext).rp = _m
	}

	return localctx
}

// ISvcommLitContext is an interface to support dynamic dispatch.
type ISvcommLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSvcommToken returns the svcommToken token.
	GetSvcommToken() antlr.Token

	// GetSvcommValue returns the svcommValue token.
	GetSvcommValue() antlr.Token

	// SetSvcommToken sets the svcommToken token.
	SetSvcommToken(antlr.Token)

	// SetSvcommValue sets the svcommValue token.
	SetSvcommValue(antlr.Token)

	// Getter signatures
	AS() antlr.TerminalNode
	ID() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsSvcommLitContext differentiates from other interfaces.
	IsSvcommLitContext()
}

type SvcommLitContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	svcommToken antlr.Token
	svcommValue antlr.Token
}

func NewEmptySvcommLitContext() *SvcommLitContext {
	var p = new(SvcommLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_svcommLit
	return p
}

func (*SvcommLitContext) IsSvcommLitContext() {}

func NewSvcommLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SvcommLitContext {
	var p = new(SvcommLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_svcommLit

	return p
}

func (s *SvcommLitContext) GetParser() antlr.Parser { return s.parser }

func (s *SvcommLitContext) GetSvcommToken() antlr.Token { return s.svcommToken }

func (s *SvcommLitContext) GetSvcommValue() antlr.Token { return s.svcommValue }

func (s *SvcommLitContext) SetSvcommToken(v antlr.Token) { s.svcommToken = v }

func (s *SvcommLitContext) SetSvcommValue(v antlr.Token) { s.svcommValue = v }

func (s *SvcommLitContext) AS() antlr.TerminalNode {
	return s.GetToken(ApiParserParserAS, 0)
}

func (s *SvcommLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *SvcommLitContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserSTRING, 0)
}

func (s *SvcommLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SvcommLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SvcommLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitSvcommLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) SvcommLit() (localctx ISvcommLitContext) {
	this := p
	_ = this

	localctx = NewSvcommLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ApiParserParserRULE_svcommLit)

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
	{
		p.SetState(135)

		var _m = p.Match(ApiParserParserID)

		localctx.(*SvcommLitContext).svcommToken = _m
	}
	{
		p.SetState(136)
		p.Match(ApiParserParserAS)
	}
	{
		p.SetState(137)

		var _m = p.Match(ApiParserParserSTRING)

		localctx.(*SvcommLitContext).svcommValue = _m
	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeLit() ITypeLitContext
	TypeBlock() ITypeBlockContext

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) TypeLit() ITypeLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeLitContext)
}

func (s *TypeSpecContext) TypeBlock() ITypeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeBlockContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeSpec() (localctx ITypeSpecContext) {
	this := p
	_ = this

	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ApiParserParserRULE_typeSpec)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.TypeLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.TypeBlock()
		}

	}

	return localctx
}

// ITypeLitContext is an interface to support dynamic dispatch.
type ITypeLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeToken returns the typeToken token.
	GetTypeToken() antlr.Token

	// SetTypeToken sets the typeToken token.
	SetTypeToken(antlr.Token)

	// Getter signatures
	TypeLitBody() ITypeLitBodyContext
	ID() antlr.TerminalNode
	AtType() IAtTypeContext

	// IsTypeLitContext differentiates from other interfaces.
	IsTypeLitContext()
}
