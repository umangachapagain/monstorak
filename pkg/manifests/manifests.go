package manifests

import (
	"bytes"
	"io"

	monv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	manLog = logf.Log.WithName("manifests_manifests")
	rules  = map[string]map[string]string{
		"ceph": map[string]string{
			"v14.2.1": "jsonnet/manifests/ceph-prometheus-rules.yaml",
		},
		"noobaa": map[string]string{
			"v1.0": "jsonnet/manifests/noobaa-prometheus-rules.yaml",
		},
	}
)

const (
	ManifestDoesNotExist  string = "Manifest for requested Storage Provider/Version may not exist"
	PrometheusObjectError string = "Prometheus object could not be created"
	ManifestParseError    string = "Manifest could not be parsed"
)

func MustAssetReader(asset string) io.Reader {
	return bytes.NewReader(MustAsset(asset))
}

type Factory struct {
	namespace string
}

func NewFactory(namespace string) *Factory {
	return &Factory{
		namespace: namespace,
	}
}

func (f *Factory) PrometheusK8sRules(storageProvider, storageVersion string) (*monv1.PrometheusRule, error) {
	manLog.WithValues("Storage Provider", storageProvider, "Storage Version", storageVersion)
	r, err := f.NewPrometheusRule(MustAssetReader(rules[storageProvider][storageVersion]))
	if err != nil {
		manLog.Error(err, ManifestDoesNotExist)
		return nil, err
	}
	return r, nil
}

func (f *Factory) NewPrometheusRule(manifest io.Reader) (*monv1.PrometheusRule, error) {
	p, err := NewPrometheusRule(manifest)
	if err != nil {
		manLog.Error(err, PrometheusObjectError)
		return nil, err
	}

	if p.GetNamespace() == "" {
		p.SetNamespace(f.namespace)
	}

	return p, nil
}

func NewPrometheusRule(manifest io.Reader) (*monv1.PrometheusRule, error) {
	p := monv1.PrometheusRule{}
	err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&p)
	if err != nil {
		manLog.Error(err, ManifestParseError)
		return nil, err
	}

	return &p, nil
}
