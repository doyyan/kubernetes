package model

import (
	"math/rand"
	"time"
)

type Deployment struct {
	ID        int               `gorm:"primaryKey" json:"id"`
	Name      string            `gorm:"column:name;type:text" json:"name" binding:"required"`
	Kind      string            `gorm:"column:kind;type:text" json:"kind"`
	NameSpace string            `gorm:"column:namespace;type:text" json:"namespace" binding:"required"`
	Labels    map[string]string `gorm:"column:labels;type:text" json:"labels"`
	Replicas  int               `gorm:"column:replicas;type:text" json:"replicas"`
	CreatedAt int64             `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64             `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (d *Deployment) FillDefaults() {
	d.Kind = "Deployment"
	if d.ID == 0 {
		rand.Seed(time.Now().UnixNano())
		min := 10
		max := 30
		d.ID = rand.Intn(max-min+1) + min
	}
}
