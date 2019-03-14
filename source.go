package sourceselect

import (
	"errors"
	"net"
)

type SourceSelect interface {
	Select(net.IP) (net.IP, error)
}

var (
	ErrNotFound       = errors.New("Select Not Found")
	ErrNotImplemented = errors.New("not implemented")
	sourceSelect      = New()
)

func Select(destination net.IP) (net.IP, error) {
	if sourceSelect == nil {
		return nil, ErrNotImplemented
	}
	return sourceSelect.Select(destination)
}
