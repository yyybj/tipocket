// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package scheme

import (
	"log"
	chaosoperatorv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kubescheme "k8s.io/client-go/kubernetes/scheme"
	"github.com/ngaut/log"
)

// Scheme gathers the schemes of native resources and custom resources used by tipocket
// in favor of the generic controller-runtime/client
var Scheme = runtime.NewScheme()

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(kubescheme.AddToScheme(Scheme))
	log.Printf("nemesis scheme init")
	utilruntime.Must(chaosoperatorv1alpha1.AddToScheme(Scheme))
	//v1.AddToGroupVersion(Scheme, schema.GroupVersion{Group: "chaos-mesh.org", Version: "v1alpha1"})

	//GroupVersion = schema.GroupVersion{Group: "chaos-mesh.org", Version: "v1alpha1"}
	//
	//// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	//SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
	//
	//// AddToScheme adds the types in this group-version to the given scheme.
	//AddToScheme = SchemeBuilder.AddToScheme
}
