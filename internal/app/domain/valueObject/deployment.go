package valueObject

type Deployment struct {
	Namespace     string
	Name          string
	Labels        map[string]string
	Image         string
	ContainerPort int
	ContainerName string
	Replicas      int
}
