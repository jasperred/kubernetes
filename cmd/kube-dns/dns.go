/*
Copyright 2014 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/spf13/pflag"
	"k8s.io/kubernetes/cmd/kube-dns/app"
	"k8s.io/kubernetes/cmd/kube-dns/app/options"
	"k8s.io/kubernetes/pkg/util"
	"k8s.io/kubernetes/pkg/util/flag"
	"k8s.io/kubernetes/pkg/version/verflag"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config := options.NewKubeDNSConfig()
	config.AddFlags(pflag.CommandLine)

	flag.InitFlags()
	util.InitLogs()
	defer util.FlushLogs()

	verflag.PrintAndExitIfRequested()
	server := app.NewKubeDNSServerDefault(config)
	server.Run()
}
