// Code generated by "enumer -type=Access -yaml"; DO NOT EDIT.

//
package main

import (
	"fmt"
)

const (
	_AccessName_0 = "NoAccess"
	_AccessName_1 = "Guest"
	_AccessName_2 = "Reporter"
	_AccessName_3 = "Developer"
	_AccessName_4 = "Maintainer"
	_AccessName_5 = "Owner"
	_AccessName_6 = "Admin"
)

var (
	_AccessIndex_0 = [...]uint8{0, 8}
	_AccessIndex_1 = [...]uint8{0, 5}
	_AccessIndex_2 = [...]uint8{0, 8}
	_AccessIndex_3 = [...]uint8{0, 9}
	_AccessIndex_4 = [...]uint8{0, 10}
	_AccessIndex_5 = [...]uint8{0, 5}
	_AccessIndex_6 = [...]uint8{0, 5}
)

func (i Access) String() string {
	switch {
	case i == 0:
		return _AccessName_0
	case i == 10:
		return _AccessName_1
	case i == 20:
		return _AccessName_2
	case i == 30:
		return _AccessName_3
	case i == 40:
		return _AccessName_4
	case i == 50:
		return _AccessName_5
	case i == 60:
		return _AccessName_6
	default:
		return fmt.Sprintf("Access(%d)", i)
	}
}

var _AccessValues = []Access{0, 10, 20, 30, 40, 50, 60}

var _AccessNameToValueMap = map[string]Access{
	_AccessName_0[0:8]:  0,
	_AccessName_1[0:5]:  10,
	_AccessName_2[0:8]:  20,
	_AccessName_3[0:9]:  30,
	_AccessName_4[0:10]: 40,
	_AccessName_5[0:5]:  50,
	_AccessName_6[0:5]:  60,
}

// AccessString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AccessString(s string) (Access, error) {
	if val, ok := _AccessNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Access values", s)
}

// AccessValues returns all values of the enum
func AccessValues() []Access {
	return _AccessValues
}

// IsAAccess returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Access) IsAAccess() bool {
	for _, v := range _AccessValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalYAML implements a YAML Marshaler for Access
func (i Access) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Access
func (i *Access) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = AccessString(s)
	return err
}
