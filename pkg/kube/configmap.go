// Copyright 2019 RedHat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kube

import (
	"fmt"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateConfigMap returns a configmap that can be created with the oc client
func GenerateConfigMap(namespace string, cmName string, pdServiceID, pdIntegrationID, pdEscalationPolicyID string, limitedSupport, serviceOrchestrationEnabled bool, serviceOrchestrationRuleApplied, alertGroupingType string, alertGroupingTimeout uint) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cmName,
			Namespace: namespace,
		},
		Data: map[string]string{
			"SERVICE_ID":                         pdServiceID,
			"INTEGRATION_ID":                     pdIntegrationID,
			"ESCALATION_POLICY_ID":               pdEscalationPolicyID,
			"LIMITED_SUPPORT":                    strconv.FormatBool(limitedSupport),
			"SERVICE_ORCHESTRATION_ENABLED":      strconv.FormatBool(serviceOrchestrationEnabled),
			"SERVICE_ORCHESTRATION_RULE_APPLIED": serviceOrchestrationRuleApplied,
			"ALERT_GROUPING_TYPE":                alertGroupingType,
			"ALERT_GROUPING_TIMEOUT":             fmt.Sprintf("%d", alertGroupingTimeout),
		},
	}
}
