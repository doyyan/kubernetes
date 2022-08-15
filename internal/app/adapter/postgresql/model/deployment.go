package model

type Deployment struct {
	ID            int               `gorm:"primaryKey" json:"id"`
	Name          string            `gorm:"column:name;type:text" json:"name"`
	Kind          string            `gorm:"column:kind;type:text" json:"kind"`
	Image         string            `gorm:"column:image;type:text" json:"image"`
	ContainerPort int               `gorm:"column:containerport;type:int" json:"containerPort"`
	ContainerName string            `gorm:"column:containername;type:text" json:"containerName"`
	NameSpace     string            `gorm:"column:namespace;type:text" json:"namespace"`
	Labels        map[string]string `json:"labels"`
	LabelKeys     []string          `gorm:"column:labelkeys;type:[]text"`
	LabelValues   []string          `gorm:"column:labelvalues;type:[]text"`
	Replicas      int               `gorm:"column:replicas;type:int" json:"replicas"`
	CreatedAt     int64             `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     int64             `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (d *Deployment) FillDefaults() {
	d.Kind = "Deployment"
}
