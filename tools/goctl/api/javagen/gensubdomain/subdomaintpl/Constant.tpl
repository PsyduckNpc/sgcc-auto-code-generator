//Code generator by goctl. DO NOT EDIT.

package {{.svName}}.infrastructure.constant;

public class {{.constantName}} {

    {{ range .constants}}
    public static final String {{.PrimaryKeySQL}} = "{{TableSQL}}:{{.PrimaryKeySQL}}";
    {{ end }}
}