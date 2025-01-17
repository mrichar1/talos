// Code generated by "enumer -type=Family -linecomment -text"; DO NOT EDIT.

//
package nethelpers

import (
	"fmt"
)

const (
	_FamilyName_0 = "inet4"
	_FamilyName_1 = "inet6"
)

var (
	_FamilyIndex_0 = [...]uint8{0, 5}
	_FamilyIndex_1 = [...]uint8{0, 5}
)

func (i Family) String() string {
	switch {
	case i == 2:
		return _FamilyName_0
	case i == 10:
		return _FamilyName_1
	default:
		return fmt.Sprintf("Family(%d)", i)
	}
}

var _FamilyValues = []Family{2, 10}

var _FamilyNameToValueMap = map[string]Family{
	_FamilyName_0[0:5]: 2,
	_FamilyName_1[0:5]: 10,
}

// FamilyString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FamilyString(s string) (Family, error) {
	if val, ok := _FamilyNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Family values", s)
}

// FamilyValues returns all values of the enum
func FamilyValues() []Family {
	return _FamilyValues
}

// IsAFamily returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Family) IsAFamily() bool {
	for _, v := range _FamilyValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for Family
func (i Family) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Family
func (i *Family) UnmarshalText(text []byte) error {
	var err error
	*i, err = FamilyString(string(text))
	return err
}
