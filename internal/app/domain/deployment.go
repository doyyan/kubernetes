package domain

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
