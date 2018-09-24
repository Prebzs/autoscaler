package hpaConnector

import (
	"fmt"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
)

func watchHpa() {

	response, err := http.Get("/apis/autoscaling/v1/watch/namespaces/{namespace}/horizontalpodautoscalers/{name}?includeUninitialized=true")
	if err != nil {
		fmt.Print("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
