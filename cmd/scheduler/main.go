/*
Copyright 2022 The Arbiter Authors.

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
	"os"

	"k8s.io/component-base/cli"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	"github.com/kube-arbiter/arbiter/pkg/extend"

	_ "github.com/kube-arbiter/arbiter/pkg/apis/scheme"
)

func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(extend.Name, extend.New),
	)

	code := cli.Run(command)
	os.Exit(code)
}
