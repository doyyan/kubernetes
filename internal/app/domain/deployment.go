package domain

// Deployment the domain object for Kubernetes Deployment payloads that has all the
// columns needed to process a k8s deployment
type Deployment struct {
	NameSpace     string
	Name          string
	Kind          string
	Image         string
	ContainerName string
	ContainerPort int
	Labels        map[string]string
	Replicas      int
	Ready         int
	Current       int
	Available     int
	Status        string
}
