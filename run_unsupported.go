//go:build !linux && !darwin && !freebsd && !windows

package buildah

import (
	"errors"

	nettypes "github.com/containers/common/libnetwork/types"
	"github.com/containers/storage"
)

func setChildProcess() error {
	return errors.New("child process function not supported on non-linux systems")
}

func runUsingRuntimeMain() {}

func (b *Builder) Run(command []string, options RunOptions) error {
	return errors.New("run function not supported on non-linux systems")
}

func DefaultNamespaceOptions() (NamespaceOptions, error) {
	return NamespaceOptions{}, errors.New("namespace options function not supported on non-linux systems")
}

// getNetworkInterface creates the network interface
func getNetworkInterface(store storage.Store, cniConfDir, cniPluginPath string) (nettypes.ContainerNetwork, error) {
	return nil, errors.New("function not supported on non-linux systems")
}
