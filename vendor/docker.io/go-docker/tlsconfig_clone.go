// +build go1.8

package docker // import "docker.io/go-docker"

import "crypto/tls"

// tlsConfigClone returns a clone of tls.Config. This function is provided for
// compatibility for go1.7 that doesn't include this method in stdlib.
func tlsConfigClone(c *tls.Config) *tls.Config {
	return c.Clone()
}
