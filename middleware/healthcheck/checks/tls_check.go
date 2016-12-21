package checks

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/mtojek/greenwall/middleware/monitoring"
)

const (
	tlsCheckName = "tls_check"
	warnDays     = 30

	errExpiringShortly = "%s: ** '%s' (S/N %X) expires in %d hours! **"
	errExpiringSoon    = "%s: '%s' (S/N %X) expires in roughly %d days."
	errSunsetAlg       = "%s: '%s' (S/N %X) expires after the sunset date for its signature algorithm '%s'."
)

type sigAlgSunset struct {
	name      string    // Human readable name of signature algorithm
	sunsetsAt time.Time // Time the algorithm will be sunset
}

// sunsetSigAlgs is an algorithm to string mapping for signature algorithms
// which have been or are being deprecated.  See the following links to learn
// more about SHA1's inclusion on this list.
//
// - https://technet.microsoft.com/en-us/library/security/2880823.aspx
// - http://googleonlinesecurity.blogspot.com/2014/09/gradually-sunsetting-sha-1.html
var sunsetSigAlgs = map[x509.SignatureAlgorithm]sigAlgSunset{
	x509.MD2WithRSA: {
		name:      "MD2 with RSA",
		sunsetsAt: time.Now(),
	},
	x509.MD5WithRSA: {
		name:      "MD5 with RSA",
		sunsetsAt: time.Now(),
	},
	x509.SHA1WithRSA: {
		name:      "SHA1 with RSA",
		sunsetsAt: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	x509.DSAWithSHA1: {
		name:      "DSA with SHA1",
		sunsetsAt: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	x509.ECDSAWithSHA1: {
		name:      "ECDSA with SHA1",
		sunsetsAt: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
	},
}

// TLSCheck is a TLS certificate health check.
// The implementation bases on project - github.com/timewasted/go-check-certs
type TLSCheck struct {
	monitoringConfiguration *monitoring.Configuration
	nodeConfig              *monitoring.Node
}

// Initialize method initializes the check instance.
func (t *TLSCheck) Initialize(monitoringConfiguration *monitoring.Configuration, nodeConfig *monitoring.Node) {
	t.monitoringConfiguration = monitoringConfiguration
	t.nodeConfig = nodeConfig
}

// Run method executes the check. This is invoked periodically.
func (t *TLSCheck) Run() CheckResult {
	conn, err := tls.Dial("tcp", t.nodeConfig.Endpoint, nil)
	if err != nil {
		log.Println(err)
		return CheckResult{
			Status:  StatusDanger,
			Message: err.Error(),
		}
	}

	err = conn.Close()
	if err != nil {
		log.Println(err)
		return CheckResult{
			Status:  StatusDanger,
			Message: err.Error(),
		}
	}

	timeNow := time.Now()
	checkedCerts := make(map[string]struct{})
	for _, chain := range conn.ConnectionState().VerifiedChains {
		for certNum, cert := range chain {
			if _, checked := checkedCerts[string(cert.Signature)]; checked {
				continue
			}
			checkedCerts[string(cert.Signature)] = struct{}{}

			// Check the expiration.
			if timeNow.AddDate(0, 0, warnDays).After(cert.NotAfter) {
				expiresIn := int64(cert.NotAfter.Sub(timeNow).Hours())
				if expiresIn <= 48 {
					return CheckResult{
						Status:  StatusDanger,
						Message: fmt.Sprintf(errExpiringShortly, t.nodeConfig.Endpoint, cert.Subject.CommonName, cert.SerialNumber, expiresIn),
					}
				}
				return CheckResult{
					Status:  StatusDanger,
					Message: fmt.Sprintf(errExpiringSoon, t.nodeConfig.Endpoint, cert.Subject.CommonName, cert.SerialNumber, expiresIn/24),
				}
			}

			// Check the signature algorithm, ignoring the root certificate.
			if alg, exists := sunsetSigAlgs[cert.SignatureAlgorithm]; exists && certNum != len(chain)-1 {
				if cert.NotAfter.Equal(alg.sunsetsAt) || cert.NotAfter.After(alg.sunsetsAt) {
					return CheckResult{
						Status:  StatusDanger,
						Message: fmt.Sprintf(errSunsetAlg, t.nodeConfig.Endpoint, cert.Subject.CommonName, cert.SerialNumber, alg.name),
					}
				}
			}
		}
	}

	return CheckResult{
		Message: MessageOK,
		Status:  StatusSuccess,
	}
}

func init() {
	registerType(tlsCheckName, reflect.TypeOf(TLSCheck{}))
}
