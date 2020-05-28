/*
 * @Author: lin.zhenhui
 * @Date: 2020-04-06 07:51:44
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-06 13:52:41
 */

 package model

 import (
	"time"

	"github.com/jinzhu/gorm"

	"api.ys1994-vod/database"
)

 // Base base modelStruct
 type Base struct {
   DeletedAt    *time.Time      `sql:"index" json:"-"`
   Filter       *interface{}    `json:"-"`
   Model        *interface{}    `json:"-"`
 }

 // Super init Base
 func (b *Base) Super(model interface{}, filter interface{}) {
   b.Filter = &filter
   b.Model = &model
 }

 // FirstOrCreate (*gorm.DB).FirstOrCreate
 func (b *Base) FirstOrCreate() {
   database.DB.Where(*b.Filter).FirstOrCreate(*b.Filter)
 }

 // Create (*gorm.DB).Create
 func (b *Base) Create() error {
   return database.DB.Create(*b.Filter).Error
 }

 // Delete (*gorm.DB).Delete
 func (b *Base) Delete(ids []uint) error {
   return database.DB.Where("id IN (?)", ids).Delete(*b.Model).Error
 }

 // Detail (*gorm.DB).First
 func (b *Base) Detail() {
   database.DB.Where(*b.Filter).First(*b.Model)
 }

 // Update (*gorm.DB).Updates; Filter.ID != 0
 func (b *Base) Update() {
   database.DB.Model(*b.Filter).Updates(*b.Filter)
 }

 // Increment Increment
 func (b *Base) Increment(attribute string) error {
   return database.DB.Model(*b.Filter).Update(attribute, gorm.Expr(attribute + " + ?", 1)).Error
 }

 // Related (*gorm.DB).Related
 func (b *Base) Related(value interface{}, foreignKeys ...string) error {
   return database.DB.Model(*b.Model).Related(value, foreignKeys...).Error
 }
