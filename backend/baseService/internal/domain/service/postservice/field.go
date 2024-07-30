package postservice

import (
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type StringField interface {
	Eq(value string) field.Expr
	Neq(value string) field.Expr
	Gt(value string) field.Expr
	Gte(value string) field.Expr
	Lt(value string) field.Expr
	Lte(value string) field.Expr
	In(values ...string) field.Expr
	NotIn(values ...string) field.Expr
	Between(min, max string) field.Expr
	NotBetween(min, max string) field.Expr
	Like(value string) field.Expr
	NotLike(value string) field.Expr
}

type TemplateSearchField struct {
	Field     StringField
	FieldName string
	Operator  api.SearchOperator
}

func NewTemplateSearchFiled(field StringField, fieldName string, operator api.SearchOperator) *TemplateSearchField {
	return &TemplateSearchField{
		Field:     field,
		FieldName: fieldName,
		Operator:  operator,
	}
}

func (t *TemplateSearchField) ToGormCondition(value string) (gen.Condition, error) {
	switch t.Operator {
	case api.SearchOperator_EQ:
		return t.Field.Eq(value), nil
	case api.SearchOperator_NE:
		return t.Field.Neq(value), nil
	case api.SearchOperator_GT:
		return t.Field.Gt(value), nil
	case api.SearchOperator_GE:
		return t.Field.Gte(value), nil
	case api.SearchOperator_LT:
		return t.Field.Lt(value), nil
	case api.SearchOperator_LE:
		return t.Field.Lte(value), nil
	case api.SearchOperator_LIKE:
		return t.Field.Like("%" + value + "%"), nil
	case api.SearchOperator_IN:
		return t.Field.In(value), nil
	case api.SearchOperator_NOT_IN:
		return t.Field.NotIn(value), nil
	case api.SearchOperator_BETWEEN:
		return t.Field.Between(value, value), nil
	default:
		return nil, fmt.Errorf("unsupported operator %v for %s", t.Operator, t.FieldName)
	}
}
