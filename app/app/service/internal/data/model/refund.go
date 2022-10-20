// Code generated by ent, DO NOT EDIT.

package model

import (
	"fmt"
	"mall-go/app/app/service/internal/data/model/refund"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Refund is the model entity for the Refund schema.
type Refund struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// RefundNo holds the value of the "refund_no" field.
	RefundNo string `json:"refund_no,omitempty"`
	// 支付平台流水号
	TransactionID string `json:"transaction_id,omitempty"`
	// user表外键
	UserID int64 `json:"user_id,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason string `json:"reason,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int64 `json:"order_id,omitempty"`
	// OrderSubID holds the value of the "order_sub_id" field.
	OrderSubID int64 `json:"order_sub_id,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Refund) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case refund.FieldID, refund.FieldUserID, refund.FieldOrderID, refund.FieldOrderSubID, refund.FieldStatus:
			values[i] = new(sql.NullInt64)
		case refund.FieldRefundNo, refund.FieldTransactionID, refund.FieldReason:
			values[i] = new(sql.NullString)
		case refund.FieldCreateTime, refund.FieldUpdateTime, refund.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Refund", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Refund fields.
func (r *Refund) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case refund.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int64(value.Int64)
		case refund.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				r.CreateTime = value.Time
			}
		case refund.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				r.UpdateTime = value.Time
			}
		case refund.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				r.DeleteTime = value.Time
			}
		case refund.FieldRefundNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refund_no", values[i])
			} else if value.Valid {
				r.RefundNo = value.String
			}
		case refund.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				r.TransactionID = value.String
			}
		case refund.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				r.UserID = value.Int64
			}
		case refund.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				r.Reason = value.String
			}
		case refund.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				r.OrderID = value.Int64
			}
		case refund.FieldOrderSubID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_sub_id", values[i])
			} else if value.Valid {
				r.OrderSubID = value.Int64
			}
		case refund.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Refund.
// Note that you need to call Refund.Unwrap() before calling this method if this Refund
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Refund) Update() *RefundUpdateOne {
	return (&RefundClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Refund entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Refund) Unwrap() *Refund {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("model: Refund is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Refund) String() string {
	var builder strings.Builder
	builder.WriteString("Refund(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("create_time=")
	builder.WriteString(r.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(r.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_time=")
	builder.WriteString(r.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("refund_no=")
	builder.WriteString(r.RefundNo)
	builder.WriteString(", ")
	builder.WriteString("transaction_id=")
	builder.WriteString(r.TransactionID)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", r.UserID))
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(r.Reason)
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", r.OrderID))
	builder.WriteString(", ")
	builder.WriteString("order_sub_id=")
	builder.WriteString(fmt.Sprintf("%v", r.OrderSubID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Refunds is a parsable slice of Refund.
type Refunds []*Refund

func (r Refunds) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
