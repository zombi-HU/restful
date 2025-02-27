// Copyright 2021-2023 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

package restful

import (
	"crypto/tls"
	"net/http"
)

// TLSClientCert adds client certs to server, enabling mutual TLS (mTLS).
// If path is a directory then scans for files recursively. If path is not set then defaults to /etc.
// File name should match *.crt or *.pem.
func (s *Server) TLSClientCert(path string) *Server {
	if s.server.TLSConfig == nil {
		s.server.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}
	s.server.TLSConfig.ClientCAs = NewCertPool(path)
	s.server.TLSConfig.ClientAuth = tls.RequireAndVerifyClientCert
	return s
}

// TLSServerCert sets server cert + key.
func (s *Server) TLSServerCert(certFile, keyFile string) *Server {
	s.certFile = certFile
	s.keyFile = keyFile
	return s
}

// ListenAndServeTLS acts like standard http.ListenAndServeTLS().
// Logs, except for automatically served LivenessProbePath and HealthCheckPath.
func ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error {
	return NewServer().Addr(addr).Handler(handler).TLSServerCert(certFile, keyFile).ListenAndServe()
}

// ListenAndServeMTLS acts like standard http.ListenAndServeTLS(). Just authenticates client.
// Parameter clientCerts is a PEM cert file or a directory of PEM cert files case insensitively matching *.pem or *.crt.
// Logs, except for automatically served LivenessProbePath and HealthCheckPath.
func ListenAndServeMTLS(addr, certFile, keyFile, clientCerts string, handler http.Handler) error {
	return NewServer().Addr(addr).Handler(handler).TLSServerCert(certFile, keyFile).TLSClientCert(clientCerts).ListenAndServe()
}
