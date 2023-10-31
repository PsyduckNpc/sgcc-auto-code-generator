//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.domain.{{.subSvName}}context.service;

import com.sgcc.basictech.frame.response.ResponseResult;
import com.sgcc.gridmgtf.subdomain.application.dto.CustDemandAppDTO;

public interface {{.serviceName}} {

    {{ range .interfaces}}
    ResponseResult<String> {{.InterfaceName}}({{ if .ReqType}}{{.ReqType}} {{.ReqParamName}}{{ end }});
    {{ end }}
}