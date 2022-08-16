package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONMap map[string]string
type Deployment struct {
	ID            int               `gorm:"primaryKey" json:"id"`
	Name          string            `gorm:"column:name;type:text" json:"name"`
	Kind          string            `gorm:"column:kind;type:text" json:"kind"`
	Image         string            `gorm:"column:image;type:text" json:"image"`
	ContainerPort int               `gorm:"column:containerport;type:int" json:"containerPort"`
	ContainerName string            `gorm:"column:containername;type:text" json:"containerName"`
	NameSpace     string            `gorm:"column:namespace;type:text" json:"namespace"`
	LabelsDB      JSONMap           `gorm:"column:labels;type:jsonb;default:'[]';not null"`
	Labels        map[string]string `json:"labels" gorm:"-"`
	Replicas      int               `gorm:"column:replicas;type:int" json:"replicas"`
	CreatedAt     int64             `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     int64             `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (d *Deployment) FillDefaults() {
	d.Kind = "Deployment"
}

// Value return json value, implement driver.Valuer interface
func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (m *JSONMap) Scan(value interface{}) error {
	ba, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	*m = make(map[string]string)
	if err := json.Unmarshal(ba, m); err != nil {
		return err
	}
	return nil
}
