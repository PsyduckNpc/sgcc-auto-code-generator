//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.{{.subSvName}}.remote.severtransfer;

import com.nbf.broker.annotations.SpecificationIdentityClient;
import com.nbf.broker.annotations.SpiClientMethodMapping;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.context.annotation.Primary;
import org.springframework.http.MediaType;
{{.imports}}

@Primary
@FeignClient(name = "{{.subdomainService}}", fallback = {{.serviceImpl}}.class)
@SpecificationIdentityClient(spi = "{{.subdomainService}}")
public interface {{.interfaceName}} {

    {{range .inters }} 
    @RequestMapping(value = "{{.Route}}", method = RequestMethod.POST, produces = MediaType.APPLICATION_JSON_UTF8_VALUE)
    @SpiClientMethodMapping(spi = "{{$.subdomainService}}", path = "/inner{{.Route}}")
    String {{.InterfaceName}}({{ if .ReqType}}{{.ReqType}} {{.DTOName}}{{ end }});
    {{ end }}
}
