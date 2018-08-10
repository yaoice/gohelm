package main

import (
	"fmt"
	"log"
	"k8s.io/helm/pkg/helm"
)

func main() {
	var opt helm.Option
	opt = helm.Host("10.125.233.67:44134")
	c := helm.NewClient(opt)

	if _, err := c.GetVersion(); err != nil {
		log.Fatal("failed to connect to Tiller, are you sure it is installed?")
	}

	ops := []helm.ReleaseListOption{
		helm.ReleaseListNamespace("default"),
	}

	releaseList, err := c.ListReleases(ops...)
	if err != nil {
		log.Fatalf("did not expect error but got (%v)\n``", err)
	}

	for _, release := range releaseList.Releases {
		fmt.Println(release.Name)
	}
}
