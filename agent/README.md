# aws-spot-instance-helper
A simple service for rancher hosts running on spot instances

### Info
This service runs as a container under a Rancher cattle environment, it monitors the state of the host. If the host is a spot instance and becomes marked for termination then this service will automatically deactivate the host and evacuate the containers.

This is an alpine based image that deploys globally