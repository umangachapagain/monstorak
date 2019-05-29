package manifests

import (
	"bytes"
	"io"

	monv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

var (
	rules = map[string]map[string]string{
		"ceph": map[string]string{
			"v14.2.1": "jsonnet/manifests/ceph-prometheus-rules.yaml",
		},
		"noobaa": map[string]string{
			"v1.0": "jsonnet/manifests/noobaa-prometheus-rules.yaml",
		},
	}
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
	r, err := f.NewPrometheusRule(MustAssetReader(rules[storageProvider][storageVersion]))
	if err != nil {
		return nil, err
	}

	r.Namespace = f.namespace

	return r, nil
}

func (f *Factory) NewPrometheusRule(manifest io.Reader) (*monv1.PrometheusRule, error) {
	p, err := NewPrometheusRule(manifest)
	if err != nil {
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
		return nil, err
	}

	return &p, nil
}
