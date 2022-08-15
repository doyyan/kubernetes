package domain

type Deployment struct {
	Namespace string
	Name      string
	Kind      string
	Labels    map[string]string
	Replicas  int
	Ready     int
	Current   int
	Available int
	Status    string
}
