package front

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/javagen/javagenutil"
)

func AssembleRemoteService(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble remote service error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "RemoteService"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + "." + subSvName + ".remote.severtransfer." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleRemoteServiceImpl(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble remote service implement error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "RemoteServiceImpl"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + "." + subSvName + ".remote.severtransfer.impl." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleQueryRemoteService(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble query remote service error, subSvName is %s", subSvName))
	}
	serviceName := "Query" + javagenutil.HeadUpCase(subSvName) + "RemoteService"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + "." + subSvName + ".remote.datatransfer." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleQueryRemoteServiceImpl(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble query remote service error, subSvName is %s", subSvName))
	}
	serviceName := "Query" + javagenutil.HeadUpCase(subSvName) + "RemoteServiceImpl"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + "." + subSvName + ".remote.datatransfer.impl." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleService(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble service error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "Service"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + "." + subSvName + ".service.application." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleServiceImpl(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble service implement error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "ServiceImpl"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + "." + subSvName + ".service.application.impl." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleQueryService(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble query service error, subSvName is %s", subSvName))
	}
	serviceName := "Query" + javagenutil.HeadUpCase(subSvName) + "Service"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + "." + subSvName + ".service.query." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleQueryServiceImpl(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble query service implement error, subSvName is %s", subSvName))
	}
	serviceName := "Query" + javagenutil.HeadUpCase(subSvName) + "ServiceImpl"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + "." + subSvName + ".service.query.impl." + serviceName + ";\n"
	}
	return serviceName
}

func AssembleController(svName, subSvName string, importMap *map[string]string) string {
	if subSvName == "" {
		panic(fmt.Sprintf("assemble controller error, subSvName is %s", subSvName))
	}
	serviceName := javagenutil.HeadUpCase(subSvName) + "Controller"
	if importMap != nil && (*importMap)[serviceName] == "" {
		(*importMap)[serviceName] = "import " + svName + "." + subSvName + ".controller." + serviceName + ";\n"
	}
	return serviceName
}
