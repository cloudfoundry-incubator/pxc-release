package tls_test

import (
	"crypto/tls"
	"database/sql"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"

	helpers "github.com/cloudfoundry/pxc-release/specs/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTls(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tls Suite")
}

var (
	mysqlConn *sql.DB
)

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

	Expect(mysql.RegisterTLSConfig("deprecated-tls11", &tls.Config{
		MaxVersion: tls.VersionTLS11,
		InsecureSkipVerify: true,
	})).To(Succeed())

	mysqlUsername := "root"
	mysqlPassword, err := helpers.GetMySQLAdminPassword()
	Expect(err).NotTo(HaveOccurred())
	firstProxy, err := helpers.FirstProxyHost(helpers.BoshDeployment)
	Expect(err).NotTo(HaveOccurred())
	mysqlConn = helpers.DbConnWithUser(mysqlUsername, mysqlPassword, firstProxy)
})
