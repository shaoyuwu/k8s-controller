package utils

import (
	"bytes"
	"fmt"
	"text/template"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/yaml"

	"kubebuiler-demo/api/v1beta1"
)

func parseTemplate(templateName string, app *v1beta1.App) []byte {
	tmpl, err := template.ParseFiles("controllers/template/" + templateName + ".yml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	b := new(bytes.Buffer)
	tmpl.Execute(b, app)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return b.Bytes()
}

func NewDeployment(app *v1beta1.App) *appv1.Deployment {
	d := &appv1.Deployment{}
	err := yaml.Unmarshal(parseTemplate("deployment", app), d)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return d
}

func NewIngress(app *v1beta1.App) *netv1.Ingress {
	i := &netv1.Ingress{}
	err := yaml.Unmarshal(parseTemplate("ingress", app), i)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return i
}

func NewService(app *v1beta1.App) *corev1.Service {
	s := &corev1.Service{}
	err := yaml.Unmarshal(parseTemplate("service", app), s)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return s
}
