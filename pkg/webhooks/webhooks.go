/*
Copyright 2024 The Kubeflow Authors.

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

package webhooks

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"

	trainingoperator "github.com/kubeflow/training-operator/pkg/apis/kubeflow.org/v1"
	"github.com/kubeflow/training-operator/pkg/webhooks/pytorch"
)

type WebhookSetupFunc func(manager manager.Manager) error

var (
	SupportedSchemeWebhook = map[string]WebhookSetupFunc{
		trainingoperator.PyTorchJobKind: pytorch.SetupWebhook,
		trainingoperator.TFJobKind:      scaffold,
		trainingoperator.MXJobKind:      scaffold,
		trainingoperator.XGBoostJobKind: scaffold,
		trainingoperator.MPIJobKind:     scaffold,
		trainingoperator.PaddleJobKind:  scaffold,
	}
)

func scaffold(manager.Manager) error {
	return nil
}