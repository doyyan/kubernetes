package domain

type Deployment struct {
	namespace string
	name      string
	kind      string
	labels    []string
	replicas  int
	ready     int
	current   int
	available int
	status    string
}
