// Code generated by "enumer -type=DataField -json"; DO NOT EDIT.

package enum

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _DataFieldName = "NameAgeRollNoAddress"

var _DataFieldIndex = [...]uint8{0, 4, 7, 13, 20}

const _DataFieldLowerName = "nameagerollnoaddress"

func (i DataField) String() string {
	if i < 0 || i >= DataField(len(_DataFieldIndex)-1) {
		return fmt.Sprintf("DataField(%d)", i)
	}
	return _DataFieldName[_DataFieldIndex[i]:_DataFieldIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _DataFieldNoOp() {
	var x [1]struct{}
	_ = x[Name-(0)]
	_ = x[Age-(1)]
	_ = x[RollNo-(2)]
	_ = x[Address-(3)]
}

var _DataFieldValues = []DataField{Name, Age, RollNo, Address}

var _DataFieldNameToValueMap = map[string]DataField{
	_DataFieldName[0:4]:        Name,
	_DataFieldLowerName[0:4]:   Name,
	_DataFieldName[4:7]:        Age,
	_DataFieldLowerName[4:7]:   Age,
	_DataFieldName[7:13]:       RollNo,
	_DataFieldLowerName[7:13]:  RollNo,
	_DataFieldName[13:20]:      Address,
	_DataFieldLowerName[13:20]: Address,
}

var _DataFieldNames = []string{
	_DataFieldName[0:4],
	_DataFieldName[4:7],
	_DataFieldName[7:13],
	_DataFieldName[13:20],
}

// DataFieldString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DataFieldString(s string) (DataField, error) {
	if val, ok := _DataFieldNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _DataFieldNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DataField values", s)
}

// DataFieldValues returns all values of the enum
func DataFieldValues() []DataField {
	return _DataFieldValues
}

// DataFieldStrings returns a slice of all String values of the enum
func DataFieldStrings() []string {
	strs := make([]string, len(_DataFieldNames))
	copy(strs, _DataFieldNames)
	return strs
}

// IsADataField returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DataField) IsADataField() bool {
	for _, v := range _DataFieldValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for DataField
func (i DataField) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for DataField
func (i *DataField) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("DataField should be a string, got %s", data)
	}

	var err error
	*i, err = DataFieldString(s)
	return err
}
