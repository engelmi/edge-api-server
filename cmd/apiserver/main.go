/* SPDX-License-Identifier: LGPL-2.1-or-later */

package main

import (
	"os"

	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/cli"
)

func main() {
	stopCh := genericapiserver.SetupSignalHandler()
	options := NewEdgeServerOptions(os.Stdout, os.Stderr)
	cmd := NewCommandStartEdgeServer(options, stopCh)
	code := cli.Run(cmd)
	os.Exit(code)
}
