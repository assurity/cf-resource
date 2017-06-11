package out

import "os"

type CfEnvironment struct {
	commandEnvironment []string
	env map[string]string
}

func NewCfEnvironment() *CfEnvironment {
	//commandEnvironment := make([]string,10)
	env := make(map[string]string)
	commandEnvironment := os.Environ()
	cfe := &CfEnvironment{commandEnvironment, env}
	cfe.addCommandEnvironmentVariable("CF_COLOR", "true")
	return cfe
}


func (cf *CfEnvironment) addCommandEnvironmentVariable(key, value string) {
	cf.commandEnvironment = append(cf.commandEnvironment, key+"="+value)
}

func (cf *CfEnvironment) CommandEnvironment() []string {
	return cf.commandEnvironment
}

func (cf *CfEnvironment) AddCommandEnvironmentVariable(switchMap map[string]string) {
	for k, v := range switchMap {
		cf.addCommandEnvironmentVariable(k, v)
	}
}