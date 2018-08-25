package ocpclient

import (
	"github.com/eriklupander/ocpvizz/model"
	"fmt"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/json"
	"os"
)

var client = &http.Client{}
var authToken string

func init() {
	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Transport = transport

	authToken = os.Getenv("OCP_AUTH_TOKEN")
}

func GetDeploymentConfigurations(serverUrl string, projectName string) ([]model.DCItem, error) {

	url := fmt.Sprintf("%s/apis/apps.openshift.io/v1/namespaces/%s/deploymentconfigs", serverUrl, projectName)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer " + authToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Call to AP returned error: %v\n", err.Error())
		return nil, err
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		ocpDcResponse := &model.OcpDc{}
		err := json.Unmarshal(body, &ocpDcResponse)
		if err != nil {
			fmt.Errorf("Error parsing returned DCs JSON: %v\n", err.Error())
			return nil, err
		} else {
			return ocpDcResponse.Items, nil
		}
	} else {
		return nil, fmt.Errorf("HTTP status was %v", resp.StatusCode)
	}
}

func GetServices(serverUrl string, projectName string) ([]model.ServiceItem, error) {

	url := fmt.Sprintf("%s/api/v1/namespaces/%s/services", serverUrl, projectName)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer 4n-vyDhi2bCjXDqGRdjV-4Xe0xkLTFVvcvsy2MKK-s8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Call to AP returned error: %v\n", err.Error())
		return nil, err
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		ocpServiceResponse := &model.OcpService{}
		err := json.Unmarshal(body, &ocpServiceResponse)
		if err != nil {
			fmt.Errorf("Error parsing returned Pods JSON: %v\n", err.Error())
			return nil, err
		} else {
			return ocpServiceResponse.Items, nil
		}
	} else {
		return nil, fmt.Errorf("HTTP status was %v", resp.StatusCode)
	}
}


func GetPods(serverUrl string, projectName string) ([]model.PodItem, error) {

	url := fmt.Sprintf("%s/api/v1/namespaces/%s/pods", serverUrl, projectName)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer 4n-vyDhi2bCjXDqGRdjV-4Xe0xkLTFVvcvsy2MKK-s8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Call to AP returned error: %v\n", err.Error())
		return nil, err
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		ocpPodResponse := &model.OcpPod{}
		err := json.Unmarshal(body, &ocpPodResponse)
		if err != nil {
			fmt.Errorf("Error parsing returned Pods JSON: %v\n", err.Error())
			return nil, err
		} else {
			return ocpPodResponse.Items, nil
		}
	} else {
		return nil, fmt.Errorf("HTTP status was %v", resp.StatusCode)
	}
}
