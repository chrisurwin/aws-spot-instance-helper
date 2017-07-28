package awshelpers

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

const (
	awsMetaData = "http://169.254.169.254"
)

// GetAWSInfoBool - A function to check a http response code against an AWS meta-data object query
func GetAWSInfoBool(path string, code int) (bool, error) {

	resp, err := http.Get(awsMetaData + path)
	if err != nil {
		log.Error("GetAWSInfo: Failed to connect to AWS Metadata " + err.Error())
		if resp != nil {
			defer resp.Body.Close()
		}
		return false, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode != code {
		return false, err
	}
	return true, err

}
