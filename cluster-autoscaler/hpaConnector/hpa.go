package hpaConnector

import (
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeclient "k8s.io/client-go/kubernetes"
)

func WatchHpa(clientSet kubeclient.Interface) {

	hpaWatch, _ := clientSet.AutoscalingV1().HorizontalPodAutoscalers("").Watch(metav1.ListOptions{})
	data := <-hpaWatch.ResultChan()

	hpa, _ := meta.NewAccessor().Namespace(data.Object)
	glog.V(1).Info(hpa);

}
