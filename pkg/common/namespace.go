package common

import (
	apiV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("common_namespace")

func newCoreV1Client() (*v1.CoreV1Client, error) {
	client, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		return nil, err
	}
	return v1.NewForConfig(client)
}

func getNamespace(namespace string) (*apiV1.Namespace, error) {
	coreClient, err := newCoreV1Client()
	if err != nil {
		return nil, err
	}

	namespaceClient := coreClient.Namespaces()
	getOptions := metav1.GetOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: apiV1.SchemeGroupVersion.String(),
		},
		ResourceVersion: "1",
	}

	ns, err := namespaceClient.Get(namespace, getOptions)
	if err != nil {
		return nil, err
	}
	return ns, err
}

func AddLabelToNamespace(namespace string, label map[string]string) error {
	coreClient, err := newCoreV1Client()
	if err != nil {
		return err
	}

	namespaceClient := coreClient.Namespaces()

	ns, err := getNamespace(namespace)
	if err != nil {
		return err
	}
	ns.SetLabels(label)
	return err
}
