//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.{{.subSvName}}.remote.severtransfer.impl;

import {{.svName}}.infrastructure.util.FallbackMessage;
import org.springframework.stereotype.Component;
import java.util.List;
{{.imports}}

@Component
public class {{.className}} implements {{.implementName}} {

    {{range .inters }} 
    @Override
    public String {{.InterfaceName}}({{ if .ReqType}}{{.ReqType}} {{.DTOName}}{{ end }}) {
        return FallbackMessage.fallbackTimeOut();
    }
    {{ end }}
}
