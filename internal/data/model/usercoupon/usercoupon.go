// Code generated by entc, DO NOT EDIT.

package usercoupon

import (
	"time"
)

const (
	// Label holds the string label denoting the usercoupon type in the database.
	Label = "user_coupon"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCouponID holds the string denoting the coupon_id field in the database.
	FieldCouponID = "coupon_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// Table holds the table name of the usercoupon in the database.
	Table = "user_coupon"
)

// Columns holds all SQL columns for usercoupon fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldUserID,
	FieldCouponID,
	FieldStatus,
	FieldOrderID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
