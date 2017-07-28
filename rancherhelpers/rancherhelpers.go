package rancherhelpers

import (
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/rancher/go-rancher/v2"
)

const (
	rancherMetaData = "http://169.254.169.250"
)

//GetRancherMetadata - Function to query local rancher metadata
func GetRancherMetadata(path string) (string, error) {

	resp, err := http.Get(rancherMetaData + path)
	if err != nil {
		log.Warn("invalid path passed " + path)
		if resp != nil {
			defer resp.Body.Close()
		}
		return "", err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn("GetRancherMetadata: Received invalid response " + path)
		return "", err
	}
	return string(body), err
}

//EvacuateHost - Function to evacuate a Rancher host
func EvacuateHost(hostName string, c *client.RancherClient) (bool, error) {

	//Get a list of Hosts
	hosts, err := c.Host.List(nil)
	if err != nil {
		log.Error("EvacuateHost: Error getting host list")
		return false, err
	}
	for _, h := range hosts.Data {
		if h.Hostname == hostName {
			_, err := c.Host.ActionEvacuate(&h)
			if err != nil {
				return false, err
			}
		}
	}
	return true, err
}
