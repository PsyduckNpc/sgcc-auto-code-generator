//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.{{.subSvName}}.controller;

import com.sgcc.basictech.frame.response.ResponseResult;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import io.swagger.annotations.ApiParam;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
{{.imports}}

@Api(tags = "{{.serviceGroupDesc}}")
@RestControlle
public class {{.controllerName}} {

    private static final Logger logger = LoggerFactory.getLogger({{.controllerName}}.class);

    {{ range .injects}}
    @Resource
    private {{.InjectServiceName}} {{.VariableServiceName}};
    {{ end }}

    {{ range .logics }}
    @ApiOperation("{{.Desc}}")
    @RequestMapping(value = "{{.Route}}", method = RequestMethod.{{.MethodType}}, produces = MediaType.APPLICATION_JSON_UTF8_VALUE)
    public ResponseResult<{{ if .RespType}}{{.RespType}}{{ else }}String{{ end }}> {{.InterfaceName}}({{ if .ReqType}}
            @ApiParam(name = "{{.ReqType}}", value = "{{.AOName}}", required = true) 
            @RequestBody {{.ReqType}} {{.AOName}}{{ end }}) {
        return {{.VariableServiceName}}.{{.InterfaceName}}({{ if .ReqType}}{{.AOName}}{{ end }});
    }
    {{ end }}
}
