//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.{{.subSvName}}.model.dto;

import {{.svName}}.infrastructure.util.CodeToName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
{{.imports}}

@ApiModel(value = "{{.headLowCaseBOName}}", Description = "{{.desc}}")
public class {{.headUpCaseBOName}} {{ if .extend }}extends {{$.extendName}} {{ end }} {{ if .implements }} implements {{$.implementContent}}  {{ end }} {


    {{ range .fields}}
    @ApiModelProperty("{{.Desc}}")
    private {{.FieldType}} {{.HeadLowCaseFieldName}};
    {{ end }}
    
    {{ range .fields}}
    public {{.FieldType}} get{{.HeadUpCaseFieldName}}() { 
        return {{.HeadLowCaseFieldName}};
    }
    public void set{{.HeadUpCaseFieldName}}({{.FieldType}} {{.HeadLowCaseFieldName}}) {
        this.{{.HeadLowCaseFieldName}} = {{.HeadLowCaseFieldName}};
    }
    {{ end }}
    
    {{ if .toStringFlag}}
    @Override
    public String toString() {
        return "{{.headUpCaseBOName}}{" +
        {{ range $index, $element := .fields}}"{{if ne $index 0}}, {{ end }}{{$element.HeadLowCaseFieldName}}='" + {{$element.HeadLowCaseFieldName}} + '\'' +
        {{ end }}'}';
    }
    {{ end }}
}
