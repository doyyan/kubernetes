package valueObject

type Deployment struct {
	Namespace string
	Name      string
	Labels    map[string]string
	Replicas  int
}
