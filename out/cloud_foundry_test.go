package out_test

import (
	"github.com/concourse/cf-resource/out"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("CloudFoundry", func() {
	Context("happy path", func() {
		var cf *out.CloudFoundry
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
			cf = out.NewCloudFoundry()
		})

		It("default command environment should contain CF_COLOR=true", func() {
			cfEnv := cf.CommandEnvironment()
			Ω(cfEnv).Should(HaveLen(baseExpectedEnvVariables))
			Ω(cfEnv).Should(ContainElement("CF_COLOR=true"))
		})

		It("added environment switch ends up in environment", func() {

			cf.AddCommandEnvironmentVariable(oneEnvironmentPair)
			cfEnv := cf.CommandEnvironment()

			Ω(cfEnv).Should(HaveLen(baseExpectedEnvVariables + 1))
			Ω(cfEnv).Should(ContainElement("ENV_ONE=env_one"))
		})

		It("multiple environment switches all end up in environment", func() {

			cf.AddCommandEnvironmentVariable(multipleEnvironmentPairs)
			cfEnv := cf.CommandEnvironment()

			Ω(cfEnv).Should(HaveLen(baseExpectedEnvVariables + 3))
			Ω(cfEnv).Should(ContainElement("ENV_ONE=env_one"))
			Ω(cfEnv).Should(ContainElement("ENV_TWO=env_two"))
			Ω(cfEnv).Should(ContainElement("ENV_THREE=env_three"))
		})

		It("multiple adds to environment retains all additions", func() {
			cf.AddCommandEnvironmentVariable(multipleEnvironmentPairs)
			cf.AddCommandEnvironmentVariable(fiveEnvironmentPair)
			cfEnv := cf.CommandEnvironment()

			Ω(cfEnv).Should(HaveLen(baseExpectedEnvVariables + 4))
			Ω(cfEnv).Should(ContainElement("ENV_ONE=env_one"))
			Ω(cfEnv).Should(ContainElement("ENV_TWO=env_two"))
			Ω(cfEnv).Should(ContainElement("ENV_THREE=env_three"))

			Ω(cfEnv).Should(ContainElement("ENV_FIVE=env_five"))
		})

	})
})
