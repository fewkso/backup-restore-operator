/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	v1 "github.com/mrajashree/backup/pkg/apis/some.api.group/v1"
	clientset "github.com/mrajashree/backup/pkg/generated/clientset/versioned/typed/some.api.group/v1"
	informers "github.com/mrajashree/backup/pkg/generated/informers/externalversions/some.api.group/v1"
	"github.com/rancher/wrangler/pkg/generic"
)

type Interface interface {
	Foo() FooController
}

func New(controllerManager *generic.ControllerManager, client clientset.SomeV1Interface,
	informers informers.Interface) Interface {
	return &version{
		controllerManager: controllerManager,
		client:            client,
		informers:         informers,
	}
}

type version struct {
	controllerManager *generic.ControllerManager
	informers         informers.Interface
	client            clientset.SomeV1Interface
}

func (c *version) Foo() FooController {
	return NewFooController(v1.SchemeGroupVersion.WithKind("Foo"), c.controllerManager, c.client, c.informers.Foos())
}