package tasks

import (
	monv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/monstorak/monstorak/pkg/manifests"
	"github.com/monstorak/monstorak/pkg/prometheus"
)

func GetPrometheusRule(storageNamespace, storageProvider, storageVersion string) (*monv1.PrometheusRule, error) {
	f := manifests.NewFactory(storageNamespace)
	prometheusK8sRules, err := f.PrometheusK8sRules(storageProvider, storageVersion)
	if err != nil {
		return nil, err
	}
	return prometheusK8sRules, nil
}

func DeployPrometheusRule(storageNamespace string, prometheusRule *monv1.PrometheusRule) error {
	prometheusRule.Namespace = storageNamespace
	err := prometheus.CreateOrUpdatePrometheusRule(prometheusRule)
	return err
}
