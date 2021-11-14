package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"onprem-k8s.lab0-apps.net/k8s-set-configuration/library"
	"os"
	"path"
	"strings"
)

func getFullPathFromConfigValue(name string) string {
	var homeDir, _ = os.UserHomeDir()
	var kubePath = path.Join(homeDir, ".kube")
	if strings.HasPrefix(name,"!") {
		return strings.ReplaceAll(name,"!",kubePath)
	} else {
		return name
	}
}

func main()  {
	var homeDir, _ = os.UserHomeDir()
	var configFile = fmt.Sprintf("%s/.kube/k8s-set-config.txt", homeDir)
	fmt.Println("User Home Directory ................... :" +homeDir)
	fmt.Println("K8S Selector Config File .............. :"+configFile)
	configs, err := library.ReadConfig(configFile)

	if err != nil {
		log.Fatal("Error Reading Config ... " + err.Error())
		return
	}

	if len(configs) == 0 {
		log.Println("No Configs Found")
		return
	}

	keys := make([]string, 0, len(configs))
	for k := range configs {
		keys = append(keys, k)
	}

	prompt := promptui.Select{
		Label: "Select Kubernetes Configuration",
		Items: keys,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	sourceKubeFile := getFullPathFromConfigValue(configs[result])

	fmt.Printf("You choose %q with file %q\n", result, sourceKubeFile)
}