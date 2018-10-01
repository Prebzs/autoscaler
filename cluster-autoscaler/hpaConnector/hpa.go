package hpaConnector

import (
	"encoding/json"
	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeclient "k8s.io/client-go/kubernetes"
	"strings"
)

func WatchHpa(clientSet kubeclient.Interface) {

	hpaWatch, _ := clientSet.AutoscalingV1().HorizontalPodAutoscalers("").Watch(metav1.ListOptions{})
	data := <-hpaWatch.ResultChan()
	dataJSON, _ :=json.Marshal(data)
	dataString := string(dataJSON);

	// Seach for DesiredReplicas String
	DesiredReplicasIndex := strings.Index(dataString, "desiredReplicas")
	DesiredReplicasIndex=DesiredReplicasIndex+17
	DesiredReplicaString1:=dataString[DesiredReplicasIndex:DesiredReplicasIndex+3]+""

	//Search for } in String
	DesiredReplicasKommaIndex := strings.Index(DesiredReplicaString1, "}")
	DesiredReplicaString := DesiredReplicaString1[0:DesiredReplicasKommaIndex]

	// Log for Debugging
	glog.V(1).Info(DesiredReplicaString)


}

