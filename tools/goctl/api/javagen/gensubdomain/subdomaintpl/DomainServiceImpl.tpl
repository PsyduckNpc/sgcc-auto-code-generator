//Code generator by goctl.

package {{.svName}}.domain.{{.allLowSubSvName}}context.service.impl;

import org.springframework.util.StringUtils;
import com.sgcc.basictech.frame.response.ResponseConstant;
import com.sgcc.basictech.frame.response.ResponseResult;
import com.sgcc.basictech.integration.redis.api.IRedisService;
import com.sgcc.gridmgtf.subdomain.infrastructure.util.AssembleUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import java.util.Objects;
{{.imports}}

@Service
public class {{.className}} implements {{.interfaceName}} {

    private static final Logger logger = LoggerFactory.getLogger({{.className}}.class);

    {{ range .injects}}
    @Resource
    private {{.InjectServiceName}} {{.VariableServiceName}};
    {{ end }}

    @Autowired
    private IRedisService redisService;

    {{ range .logics}}
    @Override
    public ResponseResult<String> {{.InterfaceName}}({{ if .DTOType}}{{.DTOType}} {{.DTOName}}{{ end }}) {
        {{ if .DTOType}}
        logger.info("{{.Desc}}-数据中台微服务收到参数：{}", JSONObject.toJSONString({{.DTOName}}));
        {{ if .ReqListFlag}}
        ArrayList<{{.RootType}}> roots = new ArrayList<>();
        BeanCopyUtils.copyList({{.DTOName}}, roots, {{.RootType}}.class);

        for ({{.RootType}} root : roots) {
            {{  if eq .Operate "create"}}
            String id = redisService.getID({{$.headUpCaseSubSvName}}.{{.PrimaryKeySQL}}, null);
            if (!Objects.isNull(id) && !"".equals(id) && StringUtils.isEmpty(root.get{{.PrimaryKey}}())) {
                root.set{{.PrimaryKey}}(id);
            }
            {{ else }}
            if (StringUtils.isEmpty(root.get{{.PrimaryKey}}())) {
                throw new BusinessException("主键标识为空", ResponseConstant.FAIL);
            }
            {{ end }}
            root.{{.Operate}}();
        }
        HashMap<String, Object> map = new HashMap<>(1);
        map.put({{.RootType}}.class.getSimpleName(), roots);
        String rootReq = JsonUtil.convertMapToJson(map);
        {{ else }}
        {{.RootType}} root = new {{.RootType}}();
        BeanCopyUtils.copy({{.DTOName}}, root);

        {{  if eq .Operate "create"}}
        String id = redisService.getID({{$.headUpCaseSubSvName}}.{{.PrimaryKeySQL}}, null);
        if (!Objects.isNull(id) && !"".equals(id) && StringUtils.isEmpty(root.get{{.PrimaryKey}}())) {
            root.set{{.PrimaryKey}}(id);
        }
        {{ else }}
        if (StringUtils.isEmpty(root.get{{.PrimaryKey}}())) {
            throw new BusinessException("主键标识为空", ResponseConstant.FAIL);
        }
        {{ end }}

        root.{{.Operate}}();
        String rootReq = BusinessObjectConvertUtil.convertBoToPoJson(root);
        {{ end }}
        logger.info("{{.Desc}}-调用数据微服务参数：{}",rootReq);{{ end }}
        String result = {{.VariableServiceName}}.{{.InterfaceName}}(rootReq);
        logger.info("{{.Desc}}-数据微服务返回参数：{}", result);
        ResponseResult<String> response = AssembleUtil.convertResultString(result);

        return Objects.isNull(response) || !ResponseConstant.SUCCESS.equals(response.getCode())
                ? ResponseResult.fail(ResponseConstant.FAIL, "{{.Desc}}失败")
                : response;
    }
    {{ end }}
}