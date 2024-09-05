//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.remote.datatransfer.impl;

import com.sgcc.gridmgtf.subdomain.infrastructure.util.FallbackMessage;
import org.springframework.stereotype.Component;
{{.imports}}

@Component
public class {{.className}} implements {{.implementName}} {

    {{ range .inters}}
    public String {{.InterfaceName}}(String jsonString) {
        return FallbackMessage.fallbackTimeOut();
    }
    {{ end }}
}