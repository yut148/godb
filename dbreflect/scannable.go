package dbreflect

import (
	"fmt"
	"reflect"
	"time"
)

var scannableStructs map[string]bool

func init() {
	scannableStructs = make(map[string]bool)
	// time.Time is scannable by default
	RegisterScannableStruct(time.Time{})
}

// Register a struct (through an instance or pointer) as being scannable.
// The registered structs will not be considered as sub structs in mappings.
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

// isStructScannable return true if the struct is scannable (not a sub mappin)
func isStructScannable(typeName string) bool {
	_, isPresent := scannableStructs[typeName]
	return isPresent
}