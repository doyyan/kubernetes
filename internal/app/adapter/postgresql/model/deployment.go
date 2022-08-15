package model

type Deployment struct {
	ID            int               `gorm:"primaryKey" json:"id"`
	Name          string            `gorm:"column:name;type:text" json:"name" binding:"required"`
	Kind          string            `gorm:"column:kind;type:text" json:"kind"`
	Image         string            `gorm:"column:image;type:text" json:"image" binding:"required"`
	ContainerPort int               `gorm:"column:containerport;type:int" json:"containerPort" binding:"required"`
	ContainerName string            `gorm:"column:containername;type:text" json:"containerName" binding:"required"`
	NameSpace     string            `gorm:"column:namespace;type:text" json:"namespace" binding:"required"`
	Labels        map[string]string `gorm:"column:labels;type:text" json:"labels"`
	Replicas      int               `gorm:"column:replicas;type:int" json:"replicas"`
	CreatedAt     int64             `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     int64             `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (d *Deployment) FillDefaults() {
	d.Kind = "Deployment"
}
