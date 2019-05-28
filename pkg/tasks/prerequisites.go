package tasks

import (
	"github.com/monstorak/monstorak/pkg/common"
	"github.com/monstorak/monstorak/pkg/prometheus"
)

// Prerequisites needed for Monitoring to work with Storage
func Prerequisites(storageNamespace, serviceMonitor string) error {
	// Check if Namespace exists
	_, err := common.GetNamespace(storageNamespace)
	if err != nil {
		return err
	}
	// Check if Namespace has required labels
	label := make(map[string]string)
	label["openshift.io/cluster-monitoring"] = "true"
	err = common.NamespaceHasLabels(storageNamespace, label)
	if err != nil {
		return err
	}
	svcMonitorLabel := make(map[string]string)
	svcMonitorLabel["app"] = serviceMonitor
	// Create a ServiceMonitor object
	sm := prometheus.CreateServiceMonitor(serviceMonitor, "http-metrics", storageNamespace, svcMonitorLabel)
	// Check if ServiceMonitor exists
	_, err = prometheus.ServiceMonitorExists(sm)
	if err != nil {
		return err
	}
	// TODO:Check if Role and RoleBinding exists

	return nil
}
