package model

import (
	"gorm.io/gorm"
)

type CPU struct {
	gorm.Model
	Microarchitectures string `json:"microarchitectures"`
	CoreCount          int    `json:"corecount"`
	Packages           string `json:"packages"`
}

type Memory struct {
	gorm.Model
	Total     string `json:"total"`
	Used      string `json:"used"`
	Free      string `json:"free"`
	Shared    string `json:"shared"`
	Buff      string `json:"buff"`
	Available string `json:"available"`
}

func (table *CPU) TableName() string {
	return "CPU"
}

func (table *Memory) TableName() string {
	return "Memory"
}
