/*
Copyright 2020 The Tekton Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	pipelinev1beta1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1beta1"
	gentype "k8s.io/client-go/gentype"
)

// fakeCustomRuns implements CustomRunInterface
type fakeCustomRuns struct {
	*gentype.FakeClientWithList[*v1beta1.CustomRun, *v1beta1.CustomRunList]
	Fake *FakeTektonV1beta1
}

func newFakeCustomRuns(fake *FakeTektonV1beta1, namespace string) pipelinev1beta1.CustomRunInterface {
	return &fakeCustomRuns{
		gentype.NewFakeClientWithList[*v1beta1.CustomRun, *v1beta1.CustomRunList](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("customruns"),
			v1beta1.SchemeGroupVersion.WithKind("CustomRun"),
			func() *v1beta1.CustomRun { return &v1beta1.CustomRun{} },
			func() *v1beta1.CustomRunList { return &v1beta1.CustomRunList{} },
			func(dst, src *v1beta1.CustomRunList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.CustomRunList) []*v1beta1.CustomRun { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.CustomRunList, items []*v1beta1.CustomRun) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
