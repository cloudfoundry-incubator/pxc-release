package bosh_wait_scripts_test

import (
	"os"
	"testing"

	helpers "specs/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBoshWaitScripts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PXC Acceptance Tests -- BoshWaitScripts Suite")
}

var _ = BeforeSuite(func() {
	requiredEnvs := []string{
		"BOSH_ENVIRONMENT",
		"BOSH_CA_CERT",
		"BOSH_CLIENT",
		"BOSH_CLIENT_SECRET",
		"BOSH_DEPLOYMENT",
		"CREDHUB_SERVER",
		"CREDHUB_CLIENT",
		"CREDHUB_SECRET",
	}
	helpers.CheckForRequiredEnvVars(requiredEnvs)

	helpers.SetupBoshDeployment()
	if os.Getenv("BOSH_ALL_PROXY") != "" {
		helpers.SetupSocks5Proxy()
	}
})
