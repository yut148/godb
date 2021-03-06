package dbreflect

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

var scannableStructs map[string]bool

func init() {
	scannableStructs = make(map[string]bool)
	// time.Time is scannable by default
	RegisterScannableStruct(time.Time{})
	// Sql nullable type
	RegisterScannableStruct(sql.NullBool{})
	RegisterScannableStruct(sql.NullFloat64{})
	RegisterScannableStruct(sql.NullInt64{})
	RegisterScannableStruct(sql.NullString{})
}

// RegisterScannableStruct registers a struct (through an instance or pointer)
// as being scannable.
// The registered structs will not be considered as a sub structs in mappings.
func RegisterScannableStruct(instance interface{}) error {
	instanceType := reflect.TypeOf(instance)
	if instanceType.Kind() == reflect.Ptr {
		instanceType = instanceType.Elem()
	}
	if instanceType.Kind() != reflect.Struct {
		return fmt.Errorf("The given type is not a struct : %T", instance)
	}
	scannableStructs[instanceType.Name()] = true
	return nil
}

// isStructScannable return true if the struct is scannable.
func isStructScannable(typeName string) bool {
	_, isPresent := scannableStructs[typeName]
	return isPresent
}
