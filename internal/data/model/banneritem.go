// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"mall-go/internal/data/model/banner"
	"mall-go/internal/data/model/banneritem"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// BannerItem is the model entity for the BannerItem schema.
type BannerItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// Img holds the value of the "img" field.
	Img string `json:"img,omitempty"`
	// Keyword holds the value of the "keyword" field.
	Keyword string `json:"keyword,omitempty"`
	// Type holds the value of the "type" field.
	Type int `json:"type,omitempty"`
	// BannerID holds the value of the "banner_id" field.
	BannerID int64 `json:"banner_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BannerItemQuery when eager-loading is set.
	Edges BannerItemEdges `json:"edges"`
}

// BannerItemEdges holds the relations/edges for other nodes in the graph.
type BannerItemEdges struct {
	// Banner holds the value of the banner edge.
	Banner *Banner `json:"banner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BannerOrErr returns the Banner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BannerItemEdges) BannerOrErr() (*Banner, error) {
	if e.loadedTypes[0] {
		if e.Banner == nil {
			// The edge banner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: banner.Label}
		}
		return e.Banner, nil
	}
	return nil, &NotLoadedError{edge: "banner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BannerItem) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case banneritem.FieldID, banneritem.FieldType, banneritem.FieldBannerID:
			values[i] = new(sql.NullInt64)
		case banneritem.FieldImg, banneritem.FieldKeyword, banneritem.FieldName:
			values[i] = new(sql.NullString)
		case banneritem.FieldCreateTime, banneritem.FieldUpdateTime, banneritem.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BannerItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BannerItem fields.
func (bi *BannerItem) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case banneritem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bi.ID = int64(value.Int64)
		case banneritem.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				bi.CreateTime = value.Time
			}
		case banneritem.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				bi.UpdateTime = value.Time
			}
		case banneritem.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				bi.DeleteTime = value.Time
			}
		case banneritem.FieldImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field img", values[i])
			} else if value.Valid {
				bi.Img = value.String
			}
		case banneritem.FieldKeyword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keyword", values[i])
			} else if value.Valid {
				bi.Keyword = value.String
			}
		case banneritem.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				bi.Type = int(value.Int64)
			}
		case banneritem.FieldBannerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field banner_id", values[i])
			} else if value.Valid {
				bi.BannerID = value.Int64
			}
		case banneritem.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				bi.Name = value.String
			}
		}
	}
	return nil
}

// QueryBanner queries the "banner" edge of the BannerItem entity.
func (bi *BannerItem) QueryBanner() *BannerQuery {
	return (&BannerItemClient{config: bi.config}).QueryBanner(bi)
}

// Update returns a builder for updating this BannerItem.
// Note that you need to call BannerItem.Unwrap() before calling this method if this BannerItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (bi *BannerItem) Update() *BannerItemUpdateOne {
	return (&BannerItemClient{config: bi.config}).UpdateOne(bi)
}

// Unwrap unwraps the BannerItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bi *BannerItem) Unwrap() *BannerItem {
	tx, ok := bi.config.driver.(*txDriver)
	if !ok {
		panic("model: BannerItem is not a transactional entity")
	}
	bi.config.driver = tx.drv
	return bi
}

// String implements the fmt.Stringer.
func (bi *BannerItem) String() string {
	var builder strings.Builder
	builder.WriteString("BannerItem(")
	builder.WriteString(fmt.Sprintf("id=%v", bi.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(bi.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(bi.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", delete_time=")
	builder.WriteString(bi.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", img=")
	builder.WriteString(bi.Img)
	builder.WriteString(", keyword=")
	builder.WriteString(bi.Keyword)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", bi.Type))
	builder.WriteString(", banner_id=")
	builder.WriteString(fmt.Sprintf("%v", bi.BannerID))
	builder.WriteString(", name=")
	builder.WriteString(bi.Name)
	builder.WriteByte(')')
	return builder.String()
}

// BannerItems is a parsable slice of BannerItem.
type BannerItems []*BannerItem

func (bi BannerItems) config(cfg config) {
	for _i := range bi {
		bi[_i].config = cfg
	}
}