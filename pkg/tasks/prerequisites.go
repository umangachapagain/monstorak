package tasks

import (
	alertsv1alpha1 "github.com/monstorak/monstorak/pkg/apis/alerts/v1alpha1"
	"github.com/monstorak/monstorak/pkg/common"
	"github.com/monstorak/monstorak/pkg/prometheus"
)

// Prerequisites needed for Monitoring to work with Storage
func Prerequisites(instance *alertsv1alpha1.StorageAlert) error {
	namespace := instance.Spec.StorageAlert.PrometheusNamespace
	// Check if Namespace exists
	_, err := common.GetNamespace(namespace)
	if err != nil {
		return err
	}
	// Check if Namespace has required labels
	label := make(map[string]string)
	label["openshift.io/cluster-monitoring"] = "true"
	err = common.NamespaceHasLabels(namespace, label)
	if err != nil {
		return err
	}
	// Check if ServiceMonitor exists
	svcMonitorLabel := make(map[string]string)
	svcMonitorLabel["app"] = "rook-ceph-mgr"
	svcMonitorLabel["rook_cluster"] = "openshift-storage"
	// Create a Service Monitor object
	sm := prometheus.CreateServiceMonitor("rook-ceph-mgr", "http-metrics", "openshift-storage", svcMonitorLabel)
	_, err = prometheus.ServiceMonitorExists(sm)
	if err != nil {
		return err
	}
	// TODO:Check if Role and RoleBinding exists

	return nil
}
