//go:build linux && cgo && !agent
// +build linux,cgo,!agent

package state

import (
	"context"
	"net/http"
	"net/url"

	"github.com/lxc/lxd/lxd/db"
	"github.com/lxc/lxd/lxd/endpoints"
	"github.com/lxc/lxd/lxd/events"
	"github.com/lxc/lxd/lxd/firewall"
	"github.com/lxc/lxd/lxd/maas"
	"github.com/lxc/lxd/lxd/sys"
	"github.com/lxc/lxd/shared"
)

// State is a gateway to the two main stateful components of LXD, the database
// and the operating system. It's typically used by model entities such as
// containers, volumes, etc. in order to perform changes.
type State struct {
	// Context
	Context context.Context

	// Databases
	Node    *db.Node
	Cluster *db.Cluster

	// MAAS server
	MAAS *maas.Controller

	// OS access
	OS    *sys.OS
	Proxy func(req *http.Request) (*url.URL, error)

	// LXD server
	Endpoints *endpoints.Endpoints

	// Event server
	DevlxdEvents *events.Server
	Events       *events.Server

	// Firewall instance
	Firewall firewall.Firewall

	// Server certificate
	ServerCert             func() *shared.CertInfo
	UpdateCertificateCache func()
}
