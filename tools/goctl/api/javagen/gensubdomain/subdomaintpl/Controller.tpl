//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.front.controller.impl;

import com.sgcc.basictech.frame.response.ResponseResult;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import org.springframework.http.MediaType;
import io.swagger.annotations.ApiParam;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import javax.annotation.Resource;

@Api(tags = "{{.serviceGroupDesc}}")
@RestController
public class {{.controllerName}} implements {{.spiName}} {

    {{ range .injects}}
    @Resource
    private {{.InjectServiceName}} {{.VariableServiceName}};
    {{ end }}

    {{ range .logics}}
    @Override
    @ApiOperation("{{.Desc}}")
    @RequestMapping(value = "{{.Route}}", method = RequestMethod.POST, produces = MediaType.APPLICATION_JSON_UTF8_VALUE)
    public ResponseResult<String> {{.InterfaceName}}({{ if .DTOType}}
            @ApiParam(name = "{{.DTOType}}", value = "{{.DTOName}}", required = true)
            @RequestBody {{.DTOType}} {{.DTOName}}{{ end }}) {
        return {{.VariableServiceName}}.{{.InterfaceName}}({{ if .DTOType}}{{.DTOName}}{{ end }});
    }
    {{ end }}
}
