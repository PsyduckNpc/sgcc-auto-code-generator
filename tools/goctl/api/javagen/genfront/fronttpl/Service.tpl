//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.{{.subSvName}}.service.application;

import com.sgcc.basictech.frame.response.ResponseResult;
{{ if .imports}}{{.imports}} {{ end }}

public interface {{.serviceName}} {


    {{range .interfaces }} 
    ResponseResult<{{ if .RespType }}{{.RespType}}{{ else }}String{{ end }}> {{.InterfaceName}}({{.ReqType}} {{.ReqParamName}});
    {{ end }}
}
