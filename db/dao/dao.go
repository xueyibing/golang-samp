package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/models"
)

const tableName = "Counters"

// ClearCounter 清除Counter
func (imp *CounterInterfaceImp) ClearCounter(id int32) error {
	cli := db.Get()
	return cli.Table(tableName).Delete(&models.CounterModel{Id: id}).Error
}

// UpsertCounter 更新/写入counter
func (imp *CounterInterfaceImp) UpsertCounter(counter *models.CounterModel) error {
	cli := db.Get()
	return cli.Table(tableName).Save(counter).Error
}

// GetCounter 查询Counter
func (imp *CounterInterfaceImp) GetCounter(id int32) (*models.CounterModel, error) {
	var err error
	var counter = new(models.CounterModel)

	cli := db.Get()
	err = cli.Table(tableName).Where("id = ?", id).First(counter).Error

	return counter, err
}
