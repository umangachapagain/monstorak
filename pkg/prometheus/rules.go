package prometheus

import (
	monv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringclient "github.com/coreos/prometheus-operator/pkg/client/versioned"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var ruleLog = logf.Log.WithName("common_prometheusRule")

func newMonitoringClient() (*monitoringclient.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		return nil, err
	}
	monitoringClient, err := monitoringclient.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return monitoringClient, err
}

// CreateOrUpdatePrometheusRule Creates or Updates prometheusRule object
func CreateOrUpdatePrometheusRule(p *monv1.PrometheusRule) error {
	mclient, err := newMonitoringClient()
	pclient := mclient.MonitoringV1().PrometheusRules(p.GetNamespace())
	oldRule, err := pclient.Get(p.GetName(), metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		_, err := pclient.Create(p)
		if err != nil {
			ruleLog.Error(err, "creating PrometheusRule object failed", "Prometheus Namespace: ", p.ObjectMeta.Namespace)
			return err
		}
		ruleLog.Info("PrometheusRule Created.", "Prometheus Namespace: ", p.ObjectMeta.Namespace)
		return err
	}
	if err != nil {
		ruleLog.Error(err, "retrieving PrometheusRule object failed", "Prometheus Namespace: ", p.ObjectMeta.Namespace)
		return err
	}

	p.ResourceVersion = oldRule.ResourceVersion

	_, err = pclient.Update(p)
	if err != nil {
		ruleLog.Error(err, "updating PrometheusRule object failed", "Prometheus Namespace: ", p.ObjectMeta.Namespace)
		return err
	}
	ruleLog.Info("PrometheusRule Updated.", "Prometheus Namespace: ", p.ObjectMeta.Namespace)
	return err
}