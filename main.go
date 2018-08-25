package main

import (
	"github.com/eriklupander/ocpvizz/service"
	"log"
	"sync"
	"github.com/eriklupander/ocpvizz/comms"
	"github.com/spf13/viper"
)

func main() {
	viper.Set("server.url", "https://192.168.99.100:8443")
	viper.Set("project.name", "local-test-env")
	//dcs, _ := service.GetDeploymentConfigurations("https://192.168.99.100:8443", "local-test-env")
	//logrus.Infof("Number of deployment configs: %v", len(dcs))
	//
	//services, _ := service.GetServices("https://192.168.99.100:8443", "local-test-env")
	//logrus.Infof("Number of services: %v", len(services))
	//
	//pods, _ := service.GetPods("https://192.168.99.100:8443", "local-test-env")
	//logrus.Infof("Number of pods: %v", len(pods))

	service.SetEventServer(comms.NewEventServer())

	go service.PublishTasks()
	log.Println("Initialized publishTasks")

	//go service.PublishServices(dockerClient)
	//log.Println("Initialized publishServices")
	//
	//go service.PublishNodes(dockerClient)
	//log.Println("Initialized publishNodes")

	// Block...
	log.Println("Waiting at block...")

	wg := sync.WaitGroup{} // Use a WaitGroup to block main() exit
	wg.Add(1)
	wg.Wait()
}
