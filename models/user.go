package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
	mysql_samp "wxcloudrun-golang/databases"
)


var _userDao *UserDao

type UserDao struct {
	gorm *gorm.DB
}

type User struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	WxInfo	  string	`gorm:"column:wxinfo" json:"wxinfo"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}



func GetUserDao() (*UserDao, error) {
	if _userDao == nil {
		db, err := mysql_samp.GetGorm()
		if err != nil {
			return nil, err
		}
		_userDao = &UserDao{
			gorm: db,
		}
		return _userDao, nil
	}else {
		err := _userDao.gorm.DB().Ping()
		if err != nil {
			//重连
			logrus.Debugf("ping error and reconnect:%s",err)
			_userDao.gorm.Close()
			db, err := mysql_samp.GetGorm()
			if err != nil {
				return nil, err
			}
			_userDao = &UserDao{
				gorm: db,
			}
			return _userDao, nil
		}
	}

	return _userDao, nil
}


func (m *UserDao) GetUser(phone *string) ([]User, error) {

	var (
		models []User
		err    error
	)

	sql := fmt.Sprintf(`
select
	id,phone,wxinfo,createdAt,updatedAt
from ops_logs
      where 1=1 %s
order by id asc
`, phone)

	err = m.gorm.Raw(sql).Scan(&models).Error
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return models, nil
}

func (m *UserDao) CreateUser(user *User)  error {

	err := m.gorm.Save(user).Error
	if err != nil {
		return err
	}
	return nil

}
