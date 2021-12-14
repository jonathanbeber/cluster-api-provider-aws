/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha1

import (
	"sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/api/bootstrap/v1beta1"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts the v1alpha1 AWSIAMConfiguration receiver to a v1beta1 AWSIAMConfiguration.
func (r *AWSIAMConfiguration) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.AWSIAMConfiguration)
	if err := Convert_v1alpha1_AWSIAMConfiguration_To_v1beta1_AWSIAMConfiguration(r, dst, nil); err != nil {
		return err
	}

	restored := &v1beta1.AWSIAMConfiguration{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	if restored.Spec.StackTags != nil {
		dst.Spec.StackTags = restored.Spec.StackTags
	}

	return nil
}

// ConvertFrom converts the v1beta1 AWSIAMConfiguration receiver to a v1alpha1 AWSIAMConfiguration.
func (r *AWSIAMConfiguration) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.AWSIAMConfiguration)

	if err := Convert_v1beta1_AWSIAMConfiguration_To_v1alpha1_AWSIAMConfiguration(src, r, nil); err != nil {
		return err
	}

	if err := utilconversion.MarshalData(src, r); err != nil {
		return err
	}

	return nil
}
