package registrationservice

import (
	"strings"

	"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1"
	"github.com/codeready-toolchain/host-operator/pkg/configuration"
	apply "github.com/codeready-toolchain/toolchain-common/pkg/client"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateOrUpdateResources(client client.Client, s *runtime.Scheme, namespace string, confg *configuration.Registry) error {
	envs := map[string]string{}
	for key, value := range confg.GetAllRegistrationServiceParameters() {
		envs[strings.TrimPrefix(key, configuration.RegServiceEnvPrefix+"_")] = value
	}

	regService := &v1alpha1.RegistrationService{
		ObjectMeta: v1.ObjectMeta{
			Namespace: namespace,
			Name:      "registration-service",
		},
		Spec: v1alpha1.RegistrationServiceSpec{
			EnvironmentVariables: envs,
		}}
	applyClient := apply.NewApplyClient(client, s)
	_, err := applyClient.CreateOrUpdateObject(regService, false, nil)
	return err
}
