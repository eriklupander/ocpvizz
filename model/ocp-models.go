package model

import "time"

type DCItem struct {
	Metadata struct {
		Name              string    `json:"name"`
		Namespace         string    `json:"namespace"`
		SelfLink          string    `json:"selfLink"`
		UID               string    `json:"uid"`
		ResourceVersion   string    `json:"resourceVersion"`
		Generation        int       `json:"generation"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Labels struct {
			App string `json:"app"`
		} `json:"labels"`
		Annotations struct {
			OpenshiftIoGeneratedBy string `json:"openshift.io/generated-by"`
		} `json:"annotations"`
	} `json:"metadata"`
	Spec struct {
		Strategy struct {
			Type string `json:"type"`
			RollingParams struct {
				UpdatePeriodSeconds int    `json:"updatePeriodSeconds"`
				IntervalSeconds     int    `json:"intervalSeconds"`
				TimeoutSeconds      int    `json:"timeoutSeconds"`
				MaxUnavailable      string `json:"maxUnavailable"`
				MaxSurge            string `json:"maxSurge"`
			} `json:"rollingParams"`
			Resources struct {
			} `json:"resources"`
			ActiveDeadlineSeconds int `json:"activeDeadlineSeconds"`
		} `json:"strategy"`
		Triggers []struct {
			Type string `json:"type"`
			ImageChangeParams struct {
				Automatic      bool     `json:"automatic"`
				ContainerNames []string `json:"containerNames"`
				From struct {
					Kind      string `json:"kind"`
					Namespace string `json:"namespace"`
					Name      string `json:"name"`
				} `json:"from"`
				LastTriggeredImage string `json:"lastTriggeredImage"`
			} `json:"imageChangeParams,omitempty"`
		} `json:"triggers"`
		Replicas             int  `json:"replicas"`
		RevisionHistoryLimit int  `json:"revisionHistoryLimit"`
		Test                 bool `json:"test"`
		Selector struct {
			App              string `json:"app"`
			Deploymentconfig string `json:"deploymentconfig"`
		} `json:"selector"`
		Template struct {
			Metadata struct {
				CreationTimestamp interface{} `json:"creationTimestamp"`
				Labels struct {
					App              string `json:"app"`
					Deploymentconfig string `json:"deploymentconfig"`
				} `json:"labels"`
				Annotations struct {
					OpenshiftIoGeneratedBy string `json:"openshift.io/generated-by"`
				} `json:"annotations"`
			} `json:"metadata"`
			Spec struct {
				Containers []struct {
					Name  string `json:"name"`
					Image string `json:"image"`
					Ports []struct {
						ContainerPort int    `json:"containerPort"`
						Protocol      string `json:"protocol"`
					} `json:"ports"`
					Resources struct {
					} `json:"resources"`
					TerminationMessagePath   string `json:"terminationMessagePath"`
					TerminationMessagePolicy string `json:"terminationMessagePolicy"`
					ImagePullPolicy          string `json:"imagePullPolicy"`
				} `json:"containers"`
				RestartPolicy                 string `json:"restartPolicy"`
				TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
				DNSPolicy                     string `json:"dnsPolicy"`
				SecurityContext struct {
				} `json:"securityContext"`
				SchedulerName string `json:"schedulerName"`
			} `json:"spec"`
		} `json:"template"`
	} `json:"spec"`
	Status struct {
		LatestVersion       int `json:"latestVersion"`
		ObservedGeneration  int `json:"observedGeneration"`
		Replicas            int `json:"replicas"`
		UpdatedReplicas     int `json:"updatedReplicas"`
		AvailableReplicas   int `json:"availableReplicas"`
		UnavailableReplicas int `json:"unavailableReplicas"`
		Details struct {
			Message string `json:"message"`
			Causes []struct {
				Type string `json:"type"`
			} `json:"causes"`
		} `json:"details"`
		Conditions []struct {
			Type               string    `json:"type"`
			Status             string    `json:"status"`
			LastUpdateTime     time.Time `json:"lastUpdateTime"`
			LastTransitionTime time.Time `json:"lastTransitionTime"`
			Reason             string    `json:"reason,omitempty"`
			Message            string    `json:"message"`
		} `json:"conditions"`
		ReadyReplicas int `json:"readyReplicas"`
	} `json:"status"`
}

type OcpDc struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata struct {
		SelfLink        string `json:"selfLink"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []DCItem `json:"items"`
}

type ServiceItem struct {
	Metadata struct {
		Name              string    `json:"name"`
		Namespace         string    `json:"namespace"`
		SelfLink          string    `json:"selfLink"`
		UID               string    `json:"uid"`
		ResourceVersion   string    `json:"resourceVersion"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Labels struct {
			App string `json:"app"`
		} `json:"labels"`
		Annotations struct {
			OpenshiftIoGeneratedBy string `json:"openshift.io/generated-by"`
		} `json:"annotations"`
	} `json:"metadata"`
	Spec struct {
		Ports []struct {
			Name       string `json:"name"`
			Protocol   string `json:"protocol"`
			Port       int    `json:"port"`
			TargetPort int    `json:"targetPort"`
		} `json:"ports"`
		Selector struct {
			App              string `json:"app"`
			Deploymentconfig string `json:"deploymentconfig"`
		} `json:"selector"`
		ClusterIP       string `json:"clusterIP"`
		Type            string `json:"type"`
		SessionAffinity string `json:"sessionAffinity"`
	} `json:"spec"`
	Status struct {
		LoadBalancer struct {
		} `json:"loadBalancer"`
	} `json:"status"`
}

type OcpService struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata struct {
		SelfLink        string `json:"selfLink"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []ServiceItem `json:"items"`
}

type PodItem struct {
	Metadata struct {
		Name              string    `json:"name"`
		GenerateName      string    `json:"generateName"`
		Namespace         string    `json:"namespace"`
		SelfLink          string    `json:"selfLink"`
		UID               string    `json:"uid"`
		ResourceVersion   string    `json:"resourceVersion"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Labels struct {
			App              string `json:"app"`
			Deployment       string `json:"deployment"`
			Deploymentconfig string `json:"deploymentconfig"`
		} `json:"labels"`
		//Annotations struct {
		//	OpenshiftIoDeploymentConfigLatestVersion string `json:"openshift.io/deployment-config.latest-version"`
		//	OpenshiftIoDeploymentConfigName          string `json:"openshift.io/deployment-config.name"`
		//	OpenshiftIoDeploymentName                string `json:"openshift.io/deployment.name"`
		//	OpenshiftIoGeneratedBy                   string `json:"openshift.io/generated-by"`
		//	OpenshiftIoScc                           string `json:"openshift.io/scc"`
		//} `json:"annotations"`
		OwnerReferences []struct {
			APIVersion         string `json:"apiVersion"`
			Kind               string `json:"kind"`
			Name               string `json:"name"`
			UID                string `json:"uid"`
			Controller         bool   `json:"controller"`
			BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
		} `json:"ownerReferences"`
	} `json:"metadata"`
	Spec struct {
		Volumes []struct {
			Name string `json:"name"`
			Secret struct {
				SecretName  string `json:"secretName"`
				DefaultMode int    `json:"defaultMode"`
			} `json:"secret"`
		} `json:"volumes"`
		Containers []struct {
			Name  string `json:"name"`
			Image string `json:"image"`
			Ports []struct {
				ContainerPort int    `json:"containerPort"`
				Protocol      string `json:"protocol"`
			} `json:"ports"`
			Resources struct {
			} `json:"resources"`
			VolumeMounts []struct {
				Name      string `json:"name"`
				ReadOnly  bool   `json:"readOnly"`
				MountPath string `json:"mountPath"`
			} `json:"volumeMounts"`
			TerminationMessagePath   string `json:"terminationMessagePath"`
			TerminationMessagePolicy string `json:"terminationMessagePolicy"`
			ImagePullPolicy          string `json:"imagePullPolicy"`
			SecurityContext struct {
				Capabilities struct {
					Drop []string `json:"drop"`
				} `json:"capabilities"`
				RunAsUser int `json:"runAsUser"`
			} `json:"securityContext"`
		} `json:"containers"`
		RestartPolicy                 string `json:"restartPolicy"`
		TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
		DNSPolicy                     string `json:"dnsPolicy"`
		ServiceAccountName            string `json:"serviceAccountName"`
		ServiceAccount                string `json:"serviceAccount"`
		NodeName                      string `json:"nodeName"`
		SecurityContext struct {
			SeLinuxOptions struct {
				Level string `json:"level"`
			} `json:"seLinuxOptions"`
			FsGroup int `json:"fsGroup"`
		} `json:"securityContext"`
		ImagePullSecrets []struct {
			Name string `json:"name"`
		} `json:"imagePullSecrets"`
		SchedulerName string `json:"schedulerName"`
	} `json:"spec"`
	Status struct {
		Phase string `json:"phase"`
		Conditions []struct {
			Type               string      `json:"type"`
			Status             string      `json:"status"`
			LastProbeTime      interface{} `json:"lastProbeTime"`
			LastTransitionTime time.Time   `json:"lastTransitionTime"`
		} `json:"conditions"`
		HostIP    string    `json:"hostIP"`
		PodIP     string    `json:"podIP"`
		StartTime time.Time `json:"startTime"`
		ContainerStatuses []struct {
			Name string `json:"name"`
			State struct {
				Running struct {
					StartedAt time.Time `json:"startedAt"`
				} `json:"running"`
			} `json:"state"`
			LastState struct {
			} `json:"lastState"`
			Ready        bool   `json:"ready"`
			RestartCount int    `json:"restartCount"`
			Image        string `json:"image"`
			ImageID      string `json:"imageID"`
			ContainerID  string `json:"containerID"`
		} `json:"containerStatuses"`
		QosClass string `json:"qosClass"`
	} `json:"status"`
}

type OcpPod struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata struct {
		SelfLink        string `json:"selfLink"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []PodItem `json:"items"`
}
