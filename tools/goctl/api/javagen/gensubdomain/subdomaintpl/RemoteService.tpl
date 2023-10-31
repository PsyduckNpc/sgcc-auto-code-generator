//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.remote.datatransfer;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.context.annotation.Primary;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
{{.imports}}

@Primary
@FeignClient(value = "{{.dataService}}", fallback = {{.classImplName}}.class)
public interface {{.interfaceName}} {

    {{ range .inters}}
    @RequestMapping(value = "{{.Route}}", method = RequestMethod.POST, produces = MediaType.APPLICATION_JSON_UTF8_VALUE)
    String {{.InterfaceName}}(String jsonString);
    {{ end }}
}
