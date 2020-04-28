/*

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

package v1

import (
	"encoding/json"

	"github.com/newrelic/newrelic-client-go/pkg/alerts"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	IncidentPreference string               `json:"incident_preference,omitempty"`
	Name               string               `json:"name"`
	APIKey             string               `json:"api_key,omitempty"`
	APIKeySecret       NewRelicAPIKeySecret `json:"api_key_secret,omitempty"`
	Region             string               `json:"region"`
}

// PolicyStatus defines the observed state of Policy
type PolicyStatus struct {
	AppliedSpec *PolicySpec `json:"applied_spec"`
	PolicyID    int         `json:"policy_id"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the policies API
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PolicySpec   `json:"spec,omitempty"`
	Status PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policy
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}

func (in PolicySpec) APIPolicy() alerts.Policy {
	jsonString, _ := json.Marshal(in)
	var APIPolicy alerts.Policy
	json.Unmarshal(jsonString, &APIPolicy) //nolint

	//APICondition.PolicyID = spec.ExistingPolicyId

	return APIPolicy
}