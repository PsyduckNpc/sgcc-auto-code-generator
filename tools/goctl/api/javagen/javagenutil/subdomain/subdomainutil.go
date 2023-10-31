package subdomain

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil"
	"strings"
)

func AssembleConstant(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble remote service error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "Constant"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + ".infrastructure.constant." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleController(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble remote service error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "Controller"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + ".front.controller.impl." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleDomainService(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble remote service error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "DomainService"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + ".domain." + strings.ToLower(subSvName) + "context.service." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleDomainServiceImpl(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble remote service error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "DomainServiceImpl"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + ".domain." + strings.ToLower(subSvName) + "context.service.impl" + serviceName + ";\n"
	}
	return serviceName
}

func AssembleRemoteService(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble remote service error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "RemoteService"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + ".remote.datatransfer." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleRemoteServiceImpl(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble remote service error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "RemoteServiceImpl"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + ".remote.datatransfer.impl." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleSpi(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble remote service error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "Spi"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + ".front.controller." + serviceName + ";\n"
	}
	return serviceName
}

func GetOperate(interfaceName string) (operate string) {
	if strings.HasPrefix(interfaceName, "create") || strings.HasPrefix(interfaceName, "Create") {
		operate = "create"
	}
	if strings.HasPrefix(interfaceName, "update") || strings.HasPrefix(interfaceName, "Update") {
		operate = "update"
	}
	if strings.HasPrefix(interfaceName, "delete") || strings.HasPrefix(interfaceName, "Delete") {
		operate = "delete"
	}
	if strings.HasPrefix(interfaceName, "select") || strings.HasPrefix(interfaceName, "Select") {
		operate = "select"
	}
	return
}
