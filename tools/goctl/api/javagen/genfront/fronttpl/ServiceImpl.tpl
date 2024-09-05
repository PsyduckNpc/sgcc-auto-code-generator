//Code generator by goctl.

package {{.svName}}.{{.subSvName}}.service.application.impl;

import com.sgcc.acapdevice.front.sd.infrastructure.util.AssembleUtil;
import com.sgcc.basictech.frame.response.ResponseConstant;
import com.sgcc.basictech.frame.response.ResponseResult;
import com.sgcc.basictech.util.BusinessObjectConvertUtil;
import com.sgcc.gridmgtf.front.sd.infrastructure.util.AssembleUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import java.util.Objects;
{{.imports}}

@Service
public class {{.serviceImplName}} implements {{.serviceName}} {

    private static final Logger logger = LoggerFactory.getLogger({{.serviceImplName}}.class);
    
    {{ range .injects}}
    @Resource
    private {{.InjectServiceName}} {{.VariableServiceName}};
    {{ end }}

    {{ range .logics}}
    @Override
    public ResponseResult<{{ if .RespType}}{{.RespType}}{{ else }}String{{ end }}> {{.InterfaceName}}({{.ReqType}} {{.AOName}}) {
        {{ if .ReqType }}
        logger.info("{{.Desc}}前台微服务收到参数：{}", JSONObject.toJSONString({{.AOName}}));
        {{ if .ReqListFlag}}
        ArrayList<{{.DTOType}}> {{.DTOName}} = new ArrayList<>();
        BeanCopyUtils.copyList({{.AOName}}, {{.DTOName}}, {{.DTOType}}.class);
        {{ else }}
        {{.DTOType}} {{.DTOName}} = new {{.DTOType}}();
        BeanCopyUtils.copy({{.AOName}}, {{.DTOName}});
        {{ end }}String param = BusinessObjectConvertUtil.convertObjToJson({{.DTOName}});
        logger.info("{{.Desc}}-调用下游数据微服务参数：{}", param);
        String result = {{.VariableServiceName}}.{{.InterfaceName}}(param);{{ else }}
        logger.info("{{.Desc}}前台微服务开始执行");
        String result = {{.VariableServiceName}}.{{.InterfaceName}}();{{ end }}
        logger.info("{{.Desc}}收到数据中心返回的信息：{}", result);
        ResponseResult<{{ if .RespType}}{{.RespType}}{{ else }}String{{ end }}> responseResult = AssembleUtil.{{ if .RespType}}{{ if .RespListFlag}}convertJsonToResponseList{{ else }}converterJsonToResponse{{ end }}{{ else }}converterJsonToString{{ end }}(result{{ if .RespType}}, {{.VOType}}.class{{ end }});
        
        if (Objects.isNull(responseResult)) {
            return ResponseResult.fail(ResponseConstant.FAIL, "{{.Desc}}执行失败");
        }
        if (!ResponseConstant.SUCCESS.equals(responseResult.getCode)) {
            return responseResult;
        }
        
        {{if .CodeConverter}}
        {{.RespType}}data = listResponseResult.getData();
        for ({{.VOType}} datum : data) {
            {{ range $i, $v := .Fields }}datum.set{{$v.HeadUpCaseFieldName}}(StCodeUtil.getCodeName("{{$v.StandardCode}}", datum.get{{$v.HeadUpCaseFieldName}}));{{ end }}
        }
        responseResult.setData(data);
        {{end}}
        return responseResult;
    }
    {{ end }}
}
