// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package recharge_order

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table gf_recharge_order.
type Entity struct {
    Id         uint    `orm:"id,primary" json:"id"`         // 主键ID     
    Orderid    string  `orm:"orderid"    json:"orderid"`    // 订单ID     
    UserId     uint    `orm:"user_id"    json:"user_id"`    // 会员ID     
    Amount     float64 `orm:"amount"     json:"amount"`     // 订单金额   
    Payamount  float64 `orm:"payamount"  json:"payamount"`  // 支付金额   
    Paytype    string  `orm:"paytype"    json:"paytype"`    // 支付类型   
    Paytime    int     `orm:"paytime"    json:"paytime"`    // 支付时间   
    Ip         string  `orm:"ip"         json:"ip"`         // IP地址     
    Useragent  string  `orm:"useragent"  json:"useragent"`  // UserAgent  
    Memo       string  `orm:"memo"       json:"memo"`       // 备注       
    Createtime int     `orm:"createtime" json:"createtime"` // 添加时间   
    Updatetime int     `orm:"updatetime" json:"updatetime"` // 更新时间   
    Status     string  `orm:"status"     json:"status"`     // 状态       
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