package out

import (
	"strings"
	"os"
	"fmt"
)

type CfEnvironment struct {
	env map[string]string
}

func NewCfEnvironment() *CfEnvironment {
	env := make(map[string]string)
	env["CF_COLOR"]="true"

	cfe := &CfEnvironment{env}

	return cfe
}

func NewCfEnvironmentFromOS() *CfEnvironment {
	cfe := NewCfEnvironment()

	environment := getenvironment(os.Environ(), splitKeyValueString)
	cfe.AddCommandEnvironmentVariable(environment)

	return cfe
}

func getenvironment(data []string, getkeyval func(item string) (key, val string)) map[string]interface{} {
	items := make(map[string]interface{})
	for _, item := range data {
		key, val := getkeyval(item)
		items[key] = val
	}
	return items
}

func splitKeyValueString(item string)(key, val string) {
	splits := strings.SplitN(item, "=", 2)
	key = splits[0]
	val = splits[1]
	return
}


func (cfe *CfEnvironment) addCommandEnvironmentVariable(key, value string) {
	cfe.env[key] = value
}

func (cfe *CfEnvironment) CommandEnvironment() []string {

	var commandEnvironment []string

	for k, v := range cfe.env {
		commandEnvironment = append(commandEnvironment, k+"="+v)
	}
	return commandEnvironment
}

func (cfe *CfEnvironment) AddCommandEnvironmentVariable(switchMap map[string]interface{}) {
	for k, v := range switchMap {
		vString := fmt.Sprintf("%v", v)
		cfe.env[k] = vString
	}
}
