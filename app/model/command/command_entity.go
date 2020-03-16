// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package command

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table gf_command.
type Entity struct {
    Id          uint   `orm:"id,primary"  json:"id"`          // ID        
    Type        string `orm:"type"        json:"type"`        // 类型      
    Params      string `orm:"params"      json:"params"`      // 参数      
    Command     string `orm:"command"     json:"command"`     // 命令      
    Content     string `orm:"content"     json:"content"`     // 返回结果  
    Executetime uint   `orm:"executetime" json:"executetime"` // 执行时间  
    Createtime  uint   `orm:"createtime"  json:"createtime"`  // 创建时间  
    Updatetime  uint   `orm:"updatetime"  json:"updatetime"`  // 更新时间  
    Status      string `orm:"status"      json:"status"`      // 状态      
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}