// Code generated by ent, DO NOT EDIT.

package banneritem

import (
	"mall-go/app/app/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// Img applies equality check predicate on the "img" field. It's identical to ImgEQ.
func Img(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImg), v))
	})
}

// Keyword applies equality check predicate on the "keyword" field. It's identical to KeywordEQ.
func Keyword(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKeyword), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// BannerID applies equality check predicate on the "banner_id" field. It's identical to BannerIDEQ.
func BannerID(v int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBannerID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
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
func DeleteTimeNotIn(vs ...time.Time) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
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
func DeleteTimeGT(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeleteTime)))
	})
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeleteTime)))
	})
}

// ImgEQ applies the EQ predicate on the "img" field.
func ImgEQ(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImg), v))
	})
}

// ImgNEQ applies the NEQ predicate on the "img" field.
func ImgNEQ(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImg), v))
	})
}

// ImgIn applies the In predicate on the "img" field.
func ImgIn(vs ...string) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImg), v...))
	})
}

// ImgNotIn applies the NotIn predicate on the "img" field.
func ImgNotIn(vs ...string) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImg), v...))
	})
}

// ImgGT applies the GT predicate on the "img" field.
func ImgGT(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImg), v))
	})
}

// ImgGTE applies the GTE predicate on the "img" field.
func ImgGTE(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImg), v))
	})
}

// ImgLT applies the LT predicate on the "img" field.
func ImgLT(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImg), v))
	})
}

// ImgLTE applies the LTE predicate on the "img" field.
func ImgLTE(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImg), v))
	})
}

// ImgContains applies the Contains predicate on the "img" field.
func ImgContains(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImg), v))
	})
}

// ImgHasPrefix applies the HasPrefix predicate on the "img" field.
func ImgHasPrefix(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImg), v))
	})
}

// ImgHasSuffix applies the HasSuffix predicate on the "img" field.
func ImgHasSuffix(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImg), v))
	})
}

// ImgEqualFold applies the EqualFold predicate on the "img" field.
func ImgEqualFold(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImg), v))
	})
}

// ImgContainsFold applies the ContainsFold predicate on the "img" field.
func ImgContainsFold(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImg), v))
	})
}

// KeywordEQ applies the EQ predicate on the "keyword" field.
func KeywordEQ(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKeyword), v))
	})
}

// KeywordNEQ applies the NEQ predicate on the "keyword" field.
func KeywordNEQ(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKeyword), v))
	})
}

// KeywordIn applies the In predicate on the "keyword" field.
func KeywordIn(vs ...string) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKeyword), v...))
	})
}

// KeywordNotIn applies the NotIn predicate on the "keyword" field.
func KeywordNotIn(vs ...string) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKeyword), v...))
	})
}

// KeywordGT applies the GT predicate on the "keyword" field.
func KeywordGT(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKeyword), v))
	})
}

// KeywordGTE applies the GTE predicate on the "keyword" field.
func KeywordGTE(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKeyword), v))
	})
}

// KeywordLT applies the LT predicate on the "keyword" field.
func KeywordLT(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKeyword), v))
	})
}

// KeywordLTE applies the LTE predicate on the "keyword" field.
func KeywordLTE(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKeyword), v))
	})
}

// KeywordContains applies the Contains predicate on the "keyword" field.
func KeywordContains(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKeyword), v))
	})
}

// KeywordHasPrefix applies the HasPrefix predicate on the "keyword" field.
func KeywordHasPrefix(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKeyword), v))
	})
}

// KeywordHasSuffix applies the HasSuffix predicate on the "keyword" field.
func KeywordHasSuffix(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKeyword), v))
	})
}

// KeywordEqualFold applies the EqualFold predicate on the "keyword" field.
func KeywordEqualFold(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKeyword), v))
	})
}

// KeywordContainsFold applies the ContainsFold predicate on the "keyword" field.
func KeywordContainsFold(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKeyword), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// BannerIDEQ applies the EQ predicate on the "banner_id" field.
func BannerIDEQ(v int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBannerID), v))
	})
}

// BannerIDNEQ applies the NEQ predicate on the "banner_id" field.
func BannerIDNEQ(v int64) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBannerID), v))
	})
}

// BannerIDIn applies the In predicate on the "banner_id" field.
func BannerIDIn(vs ...int64) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBannerID), v...))
	})
}

// BannerIDNotIn applies the NotIn predicate on the "banner_id" field.
func BannerIDNotIn(vs ...int64) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBannerID), v...))
	})
}

// BannerIDIsNil applies the IsNil predicate on the "banner_id" field.
func BannerIDIsNil() predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBannerID)))
	})
}

// BannerIDNotNil applies the NotNil predicate on the "banner_id" field.
func BannerIDNotNil() predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBannerID)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.BannerItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BannerItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// HasBanner applies the HasEdge predicate on the "banner" edge.
func HasBanner() predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BannerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BannerTable, BannerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBannerWith applies the HasEdge predicate on the "banner" edge with a given conditions (other predicates).
func HasBannerWith(preds ...predicate.Banner) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BannerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BannerTable, BannerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BannerItem) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BannerItem) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
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
func Not(p predicate.BannerItem) predicate.BannerItem {
	return predicate.BannerItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
