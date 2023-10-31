package javagenutil

type NoRepeatMap struct {
	NoRepeatMap map[string]any
}

func (m *NoRepeatMap) Add(key string, value any) {
	if (*m).NoRepeatMap[key] == nil || (*m).NoRepeatMap[key] == "" {
		(*m).NoRepeatMap[key] = value
	}
}

type Field struct {
	Desc                 string
	FieldType            string
	HeadLowCaseFieldName string
	HeadUpCaseFieldName  string
	StandardCode         string
}

type Inject struct {
	InjectServiceName   string
	VariableServiceName string
}

type Logic struct {
	RespType            string
	InterfaceName       string
	ReqType             string
	AOName              string
	Desc                string
	ReqListFlag         bool
	DTOType             string
	DTOName             string
	VariableServiceName string
	RespListFlag        bool
	VOType              string
	CodeConverter       bool
	Fields              []Field
	MethodType          string
	Route               string

	Operate       string
	RootType      string
	PrimaryKeySQL string
	PrimaryKey    string
}

type Inter struct {
	RespType      string
	InterfaceName string
	ReqType       string
	ReqParamName  string
	Route         string
	DTOType       string
	DTOName       string
}
