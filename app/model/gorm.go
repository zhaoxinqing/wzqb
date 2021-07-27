package model

import (
	"gorm.io/gorm"
)

// DB ...
var DB *gorm.DB

// Create ...
func Create(value interface{}) error {
	return DB.Create(value).Error
}

// Save ...
func Save(value interface{}) error {
	return DB.Save(value).Error
}

// Updates ...
func Updates(where interface{}, value interface{}) error {
	return DB.Model(where).Updates(value).Error
}

// DeleteByModel ...
func DeleteByModel(model interface{}) (count int64, err error) {
	db := DB.Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByWhere ...
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := DB.Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByID ...
func DeleteByID(model interface{}, id int64) (count int64, err error) {
	db := DB.Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByIDS ...
func DeleteByIDS(model interface{}, ids []int64) (count int64, err error) {
	db := DB.Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// FirstByID ...
func FirstByID(out interface{}, id int64) (notFound bool, err error) {
	err = DB.First(out, id).Error
	if err != nil {
		notFound = true
		err = gorm.ErrRecordNotFound
	}
	return
}

// First ...
func First(where interface{}, out interface{}) (notFound bool, err error) {
	err = DB.Where(where).First(out).Error
	if err != nil {
		notFound = true
		err = gorm.ErrRecordNotFound
	}
	return
}

// Find ...
func Find(where interface{}, out interface{}, orders ...string) error {
	db := DB.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Scan ...
func Scan(model, where interface{}, out interface{}) (notFound bool, err error) {
	err = DB.Model(model).Where(where).Scan(out).Error
	if err != nil {
		// notFound = gorm.IsRecordNotFoundError(err)
		notFound = true
		err = gorm.ErrRecordNotFound
	}
	return
}

// ScanList ...
func ScanList(model, where interface{}, out interface{}, orders ...string) error {
	db := DB.Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Scan(out).Error
}

// ScanListByIds ...
func ScanListByIds(model interface{}, ids []int64, out interface{}) error {
	db := DB.Model(model).Where("id in (?)", ids)
	return db.Scan(out).Error
}

type PageWhereOrder struct {
	Order string
	Where string
	Value []interface{}
}

// GetPage ...
func GetPage(model, where interface{}, out interface{}, pageIndex, pageSize int64, totalCount *int64, whereOrder ...PageWhereOrder) error {
	db := DB.Model(model).Where(where)
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Order != "" {
				db = db.Order(wo.Order)
			}
			if wo.Where != "" {
				db = db.Where(wo.Where, wo.Value...)
			}
		}
	}
	err := db.Count(totalCount).Error
	if err != nil {
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	// size, err := convert.ToIntE(pageSize)
	// index, err := convert.ToIntE(pageIndex)
	// return db.Offset((index - 1) * size).Limit(size).Find(out).Error
	return err
}

// PluckList ...
func PluckList(model, where interface{}, out interface{}, fieldName string) error {
	return DB.Model(model).Where(where).Pluck(fieldName, out).Error
}

// GetPage ...
func GetPageBySql(name string, where string, out interface{}, pageIndex, pageSize int64, totalCount *int64, whereOrder ...PageWhereOrder) error {
	db := DB.Table(name).Where(where)
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Order != "" {
				db = db.Order(wo.Order)
			}
			if wo.Where != "" {
				db = db.Where(wo.Where, wo.Value...)
			}
		}
	}
	err := db.Count(totalCount).Error
	if err != nil {
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	// size, err := convert.ToIntE(pageSize)
	// index, err := convert.ToIntE(pageIndex)
	// return db.Offset((index - 1) * size).Limit(size).Find(out).Error
	return err
}
