package kubernetes

import (
	"context"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/joyrex2001/kubedock/internal/container"
)

// StartContainer will start given container object in kubernetes.
func (in *instance) StartContainer(tainr container.Container) error {
	log.Printf("starting container %s (%s)", tainr.GetName(), tainr.GetID())

	name := tainr.GetKubernetesName()
	matchlabels := map[string]string{
		"app":  name,
		"tier": "kubedock",
	}

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: in.namespace,
			Labels:    tainr.GetLabels(),
		},

		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: matchlabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: matchlabels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   tainr.GetImage(),
						Name:    tainr.GetKubernetesName(),
						Command: tainr.GetCmd(),
						Env:     tainr.GetEnvVar(),
						Ports:   tainr.GetContainerPorts(),
					}},
				},
			},
		},
	}

	log.Printf("%#v", dep)

	_, err := in.cli.AppsV1().Deployments(in.namespace).Create(context.TODO(), dep, metav1.CreateOptions{})
	return err
}
