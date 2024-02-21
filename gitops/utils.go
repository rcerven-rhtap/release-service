/*
Copyright 2022.

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

package gitops

import (
	applicationapiv1alpha1 "github.com/redhat-appstudio/application-api/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// hasDeploymentFinished returns a boolean that is only true if the first passed object
// is a SnapshotEnvironmentBinding with the componentDeployment status missing or not true and the
// second passed object is a SnapshotEnvironmentBinding with the componentDeployment status True/False.
func hasDeploymentFinished(objectOld, objectNew client.Object) bool {
	var ok bool
	var oldBinding, newBinding *applicationapiv1alpha1.SnapshotEnvironmentBinding
	var oldCondition, newCondition *metav1.Condition

	if oldBinding, ok = objectOld.(*applicationapiv1alpha1.SnapshotEnvironmentBinding); !ok {
		return false
	}
	if newBinding, ok = objectNew.(*applicationapiv1alpha1.SnapshotEnvironmentBinding); !ok {
		return false
	}

	oldCondition = meta.FindStatusCondition(oldBinding.Status.ComponentDeploymentConditions,
		applicationapiv1alpha1.ComponentDeploymentConditionAllComponentsDeployed)

	newCondition = meta.FindStatusCondition(newBinding.Status.ComponentDeploymentConditions,
		applicationapiv1alpha1.ComponentDeploymentConditionAllComponentsDeployed)

	return (oldCondition == nil || oldCondition.Status != metav1.ConditionTrue) &&
		(newCondition != nil && newCondition.Status != metav1.ConditionUnknown)
}
