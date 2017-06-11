package out_test

import (
	"github.com/concourse/cf-resource/out"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"strings"
	"fmt"
)

var _ = Describe("CfEnvironment", func() {
	Context("happy path", func() {
		var cfEnvironment *out.CfEnvironment
		env := os.Environ()
		baseExpectedEnvVariables := len(env) + 1
		oneEnvironmentPair := map[string]string{"ENV_ONE": "env_one"}
		multipleEnvironmentPairs := map[string]string{
			"ENV_ONE":   "env_one",
			"ENV_TWO":   "env_two",
			"ENV_THREE": "env_three",
		}
		fiveEnvironmentPair := map[string]string{"ENV_FIVE": "env_five"}

		BeforeEach(func() {
			cfEnvironment = out.NewCfEnvironment()
		})

		It("default command environment should contain CF_COLOR=true", func() {
			cfEnv := cfEnvironment.CommandEnvironment()
			Ω(cfEnv).Should(HaveLen(baseExpectedEnvVariables))
			Ω(cfEnv).Should(ContainElement("CF_COLOR=true"))
		})

		It("added environment switch ends up in environment", func() {

			cfEnvironment.AddCommandEnvironmentVariable(oneEnvironmentPair)
			cfEnv := cfEnvironment.CommandEnvironment()

			Ω(cfEnv).Should(HaveLen(baseExpectedEnvVariables + 1))
			Ω(cfEnv).Should(ContainElement("ENV_ONE=env_one"))
		})

		It("multiple environment switches all end up in environment", func() {

			cfEnvironment.AddCommandEnvironmentVariable(multipleEnvironmentPairs)
			cfEnv := cfEnvironment.CommandEnvironment()

			Ω(cfEnv).Should(HaveLen(baseExpectedEnvVariables + 3))
			Ω(cfEnv).Should(ContainElement("ENV_ONE=env_one"))
			Ω(cfEnv).Should(ContainElement("ENV_TWO=env_two"))
			Ω(cfEnv).Should(ContainElement("ENV_THREE=env_three"))
		})

		It("multiple adds to environment retains all additions", func() {
			cfEnvironment.AddCommandEnvironmentVariable(multipleEnvironmentPairs)
			cfEnvironment.AddCommandEnvironmentVariable(fiveEnvironmentPair)
			cfEnv := cfEnvironment.CommandEnvironment()

			Ω(cfEnv).Should(HaveLen(baseExpectedEnvVariables + 4))
			Ω(cfEnv).Should(ContainElement("ENV_ONE=env_one"))
			Ω(cfEnv).Should(ContainElement("ENV_TWO=env_two"))
			Ω(cfEnv).Should(ContainElement("ENV_THREE=env_three"))

			Ω(cfEnv).Should(ContainElement("ENV_FIVE=env_five"))
		})

	})
})

func foo() {
	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
		items := make(map[string]string)
		for _, item := range data {
			key, val := getkeyval(item)
			items[key] = val
		}
		return items
	}
	environment := getenvironment(os.Environ(), func(item string) (key, val string) {
		//splits := strings.Split(item, "=")
		splits := strings.SplitN(item, "=", 2)
		key = splits[0]
		val = splits[1]
		return
	})
	fmt.Println(environment["PATH"])
	fmt.Println("Hello!")
}
