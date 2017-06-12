package out_test

import (
	"github.com/concourse/cf-resource/out"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var oneEnvironmentPair = map[string]interface{}{"ENV_ONE": "env_one"}

// json from config is unmarshalled in to map[string]interface{}
// keys are always strings, but values can be anything
var multipleEnvironmentPairs = map[string]interface{}{
	"ENV_ONE":   "env_one",
	"ENV_TWO":   2,
	"ENV_THREE": true,
}
var fiveEnvironmentPair = map[string]interface{}{"ENV_FIVE": "env_five"}

var _ = Describe("CfEnvironment from Empty", func() {
	Context("happy path", func() {
		var cfEnvironment *out.CfEnvironment

		BeforeEach(func() {
			cfEnvironment = out.NewCfEnvironment()
		})

		It("default command environment should ONLY contain CF_COLOR=true", func() {
			cfEnv := cfEnvironment.CommandEnvironment()
			Ω(cfEnv).Should(HaveLen(1))
			Ω(cfEnv).Should(ContainElement("CF_COLOR=true"))
		})

		It("added environment switch ends up in environment", func() {

			cfEnvironment.AddCommandEnvironmentVariable(oneEnvironmentPair)
			cfEnv := cfEnvironment.CommandEnvironment()

			Ω(cfEnv).Should(HaveLen(2))
			Ω(cfEnv).Should(ContainElement("ENV_ONE=env_one"))
		})

		It("multiple environment switches all end up in environment", func() {

			cfEnvironment.AddCommandEnvironmentVariable(multipleEnvironmentPairs)
			cfEnv := cfEnvironment.CommandEnvironment()

			Ω(cfEnv).Should(HaveLen(4))
			Ω(cfEnv).Should(ContainElement("ENV_ONE=env_one"))
			Ω(cfEnv).Should(ContainElement("ENV_TWO=2"))
			Ω(cfEnv).Should(ContainElement("ENV_THREE=true"))
		})

		It("multiple adds to environment retains all additions", func() {
			cfEnvironment.AddCommandEnvironmentVariable(multipleEnvironmentPairs)
			cfEnvironment.AddCommandEnvironmentVariable(fiveEnvironmentPair)
			cfEnv := cfEnvironment.CommandEnvironment()

			Ω(cfEnv).Should(HaveLen(5))
			Ω(cfEnv).Should(ContainElement("ENV_ONE=env_one"))
			Ω(cfEnv).Should(ContainElement("ENV_TWO=2"))
			Ω(cfEnv).Should(ContainElement("ENV_THREE=true"))

			Ω(cfEnv).Should(ContainElement("ENV_FIVE=env_five"))
		})

	})
})

var _ = Describe("CfEnvironment from OS", func() {
	Context("happy path", func() {
		var cfEnvironment *out.CfEnvironment
		env := os.Environ()
		baseExpectedEnvVariables := len(env) + 1

		BeforeEach(func() {
			cfEnvironment = out.NewCfEnvironmentFromOS()
		})

		It("default command environment should contain CF_COLOR=true", func() {
			cfEnv := cfEnvironment.CommandEnvironment()
			Ω(cfEnv).Should(HaveLen(baseExpectedEnvVariables))
			Ω(cfEnv).Should(ContainElement("CF_COLOR=true"))
		})
	})
})
