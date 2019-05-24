package tasks

import (
	monv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	alertsv1alpha1 "github.com/monstorak/monstorak/pkg/apis/alerts/v1alpha1"
	"github.com/monstorak/monstorak/pkg/manifests"
	"github.com/monstorak/monstorak/pkg/prometheus"
)

func GetPrometheusRule() (*monv1.PrometheusRule, error) {
	f := manifests.NewFactory(instance.Namespace)
	prometheusK8sRules, err := f.PrometheusK8sRules()
	if err != nil {
		return nil, err
	}
	return prometheusK8sRules, nil
}

func DeployPrometheusRule(instance *alertsv1alpha1.StorageAlert, prometheusRule *monv1.PrometheusRule) error {
	prometheusRule.Namespace = instance.Spec.StorageAlert.PrometheusNamespace
	err := prometheus.CreateOrUpdatePrometheusRule(prometheusRule)
	return err
}
