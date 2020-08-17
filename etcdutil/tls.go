package etcdutil

import (
	"crypto/tls"

	"code.cloudfoundry.org/tlsconfig"
)

func GetTlsFromFiles(cafile, certfile, keyfile string) (*tls.Config, error) {
	return tlsconfig.Build(
		tlsconfig.WithIdentityFromFile(certfile, keyfile)).
		Client(tlsconfig.WithAuthorityFromFile(cafile))
}
