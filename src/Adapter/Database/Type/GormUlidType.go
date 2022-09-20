package Type

import (
	"context"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm/schema"
	"reflect"
)

type GormUlid ulid.ULID

func (id *GormUlid) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	switch value := dbValue.(type) {
	case GormUlid:
		*id = value
	case string:
		*id = GormUlid(ulid.MustParse(value))
	default:
		errorMessage := fmt.Sprintf("Unsupported data while parsing GormUlidType: %s", dbValue)
		return errors.New(errorMessage)
	}
	return nil
}

func (id GormUlid) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return ulid.ULID(id).String(), nil
}
