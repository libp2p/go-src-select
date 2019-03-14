package sourceselect

import (
	"net"

	"github.com/vishvananda/netlink"
)

type LinuxSourceSelect struct {
}

func New() *LinuxSourceSelect {
	return &LinuxSourceSelect{}
}

func (s *LinuxSourceSelect) Select(destination net.IP) (net.IP, error) {
	routes, err := netlink.RouteGet(destination)
	if err != nil {
		return nil, err
	}

	for _, route := range routes {
		if route.Src != nil && len(route.Src) != 0 {
			return route.Src, nil
		}
	}

	return nil, ErrNotFound
}
