// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.PrometheusSpec":     schema_pkg_apis_alerts_v1alpha1_PrometheusSpec(ref),
		"github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.StorageAlert":       schema_pkg_apis_alerts_v1alpha1_StorageAlert(ref),
		"github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.StorageAlertSpec":   schema_pkg_apis_alerts_v1alpha1_StorageAlertSpec(ref),
		"github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.StorageAlertStatus": schema_pkg_apis_alerts_v1alpha1_StorageAlertStatus(ref),
		"github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.StorageSpec":        schema_pkg_apis_alerts_v1alpha1_StorageSpec(ref),
	}
}

func schema_pkg_apis_alerts_v1alpha1_PrometheusSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PrometheusSpec defines the prometheus to be used",
				Properties: map[string]spec.Schema{
					"label": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_alerts_v1alpha1_StorageAlert(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageAlert is the Schema for the storagealerts API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.StorageAlertSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.StorageAlertStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.StorageAlertSpec", "github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.StorageAlertStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_alerts_v1alpha1_StorageAlertSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageAlertSpec defines the desired state of StorageAlert",
				Properties: map[string]spec.Schema{
					"storage": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.StorageSpec"),
									},
								},
							},
						},
					},
					"prometheus": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.PrometheusSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.PrometheusSpec", "github.com/monstorak/monstorak-operator/pkg/apis/alerts/v1alpha1.StorageSpec"},
	}
}

func schema_pkg_apis_alerts_v1alpha1_StorageAlertStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageAlertStatus defines the observed state of StorageAlert",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_alerts_v1alpha1_StorageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageSpec defines the storages to be monitored",
				Properties: map[string]spec.Schema{
					"provider": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"serviceMonitor": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
