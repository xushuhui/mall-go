// Code generated by entc, DO NOT EDIT.

package orderdetail

import (
	"mall-go/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// PayWay applies equality check predicate on the "pay_way" field. It's identical to PayWayEQ.
func PayWay(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPayWay), v))
	})
}

// ClientType applies equality check predicate on the "client_type" field. It's identical to ClientTypeEQ.
func ClientType(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientType), v))
	})
}

// ShipNo applies equality check predicate on the "ship_no" field. It's identical to ShipNoEQ.
func ShipNo(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipNo), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeleteTime)))
	})
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeleteTime)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// PayWayEQ applies the EQ predicate on the "pay_way" field.
func PayWayEQ(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPayWay), v))
	})
}

// PayWayNEQ applies the NEQ predicate on the "pay_way" field.
func PayWayNEQ(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPayWay), v))
	})
}

// PayWayIn applies the In predicate on the "pay_way" field.
func PayWayIn(vs ...int) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPayWay), v...))
	})
}

// PayWayNotIn applies the NotIn predicate on the "pay_way" field.
func PayWayNotIn(vs ...int) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPayWay), v...))
	})
}

// PayWayGT applies the GT predicate on the "pay_way" field.
func PayWayGT(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPayWay), v))
	})
}

// PayWayGTE applies the GTE predicate on the "pay_way" field.
func PayWayGTE(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPayWay), v))
	})
}

// PayWayLT applies the LT predicate on the "pay_way" field.
func PayWayLT(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPayWay), v))
	})
}

// PayWayLTE applies the LTE predicate on the "pay_way" field.
func PayWayLTE(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPayWay), v))
	})
}

// ClientTypeEQ applies the EQ predicate on the "client_type" field.
func ClientTypeEQ(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientType), v))
	})
}

// ClientTypeNEQ applies the NEQ predicate on the "client_type" field.
func ClientTypeNEQ(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientType), v))
	})
}

// ClientTypeIn applies the In predicate on the "client_type" field.
func ClientTypeIn(vs ...int) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldClientType), v...))
	})
}

// ClientTypeNotIn applies the NotIn predicate on the "client_type" field.
func ClientTypeNotIn(vs ...int) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldClientType), v...))
	})
}

// ClientTypeGT applies the GT predicate on the "client_type" field.
func ClientTypeGT(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientType), v))
	})
}

// ClientTypeGTE applies the GTE predicate on the "client_type" field.
func ClientTypeGTE(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientType), v))
	})
}

// ClientTypeLT applies the LT predicate on the "client_type" field.
func ClientTypeLT(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientType), v))
	})
}

// ClientTypeLTE applies the LTE predicate on the "client_type" field.
func ClientTypeLTE(v int) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientType), v))
	})
}

// ShipNoEQ applies the EQ predicate on the "ship_no" field.
func ShipNoEQ(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipNo), v))
	})
}

// ShipNoNEQ applies the NEQ predicate on the "ship_no" field.
func ShipNoNEQ(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShipNo), v))
	})
}

// ShipNoIn applies the In predicate on the "ship_no" field.
func ShipNoIn(vs ...string) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldShipNo), v...))
	})
}

// ShipNoNotIn applies the NotIn predicate on the "ship_no" field.
func ShipNoNotIn(vs ...string) predicate.OrderDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldShipNo), v...))
	})
}

// ShipNoGT applies the GT predicate on the "ship_no" field.
func ShipNoGT(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShipNo), v))
	})
}

// ShipNoGTE applies the GTE predicate on the "ship_no" field.
func ShipNoGTE(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShipNo), v))
	})
}

// ShipNoLT applies the LT predicate on the "ship_no" field.
func ShipNoLT(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShipNo), v))
	})
}

// ShipNoLTE applies the LTE predicate on the "ship_no" field.
func ShipNoLTE(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShipNo), v))
	})
}

// ShipNoContains applies the Contains predicate on the "ship_no" field.
func ShipNoContains(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldShipNo), v))
	})
}

// ShipNoHasPrefix applies the HasPrefix predicate on the "ship_no" field.
func ShipNoHasPrefix(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldShipNo), v))
	})
}

// ShipNoHasSuffix applies the HasSuffix predicate on the "ship_no" field.
func ShipNoHasSuffix(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldShipNo), v))
	})
}

// ShipNoEqualFold applies the EqualFold predicate on the "ship_no" field.
func ShipNoEqualFold(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldShipNo), v))
	})
}

// ShipNoContainsFold applies the ContainsFold predicate on the "ship_no" field.
func ShipNoContainsFold(v string) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldShipNo), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderDetail) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderDetail) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderDetail) predicate.OrderDetail {
	return predicate.OrderDetail(func(s *sql.Selector) {
		p(s.Not())
	})
}
