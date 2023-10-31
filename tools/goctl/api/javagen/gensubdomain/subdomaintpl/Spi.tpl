//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.front.controller;

import com.sgcc.basictech.frame.response.ResponseResult;
import com.wdk.domain.plugin.annotations.SpecificationIdentity;
import com.wdk.domain.plugin.config.enums.ServiceType;
{{.imports}}

@SpecificationIdentity(id = "{{.subdomainName}}", name = "{{.subdomainName}}", detail = "{{.subdomainName}}", version = "1.0.0", serviceType = ServiceType.API)
public interface {{.className}} {

    {{ range .inters}}
    ResponseResult<String> {{.InterfaceName}}({{ if .DTOType}}{{.DTOType}} {{.DTOName}}{{ end }});
    {{ end }}
}