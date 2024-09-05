//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.{{.subSvName}}.remote.datatransfer.impl;

import {{.svName}}.infrastructure.util.FallbackMessage;
import org.springframework.stereotype.Component;
{{.imports}}

@Component
public class {{.interfaceImplName}} implements {{.interfaceName}} {
    
    {{range .inters }} 
    @Override
    public String {{.InterfaceName}}({{ if .ReqType}}String jsonString{{ end }}) {
        return FallbackMessage.fallbackTimeOut();
    }
    {{ end }}
}
