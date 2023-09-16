package services

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

func (v *UUIDArray) Scan(value interface{}) error {
	bytes, ok := value.(string)
	if !ok {
		return errors.New(fmt.Sprint("failed to scan UUIDArray, value:", value))
	}
	str := string(bytes)
	str = regexp.MustCompile("{|}|'|'").ReplaceAllString(str, "")
	ids := []uuid.UUID{}
	for _, strId := range strings.Split(str, ",") {
		if strId == "" {
			continue
		}
		id, err := uuid.Parse(strId)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}
	*v = ids
	return nil
}

func (v UUIDArray) Value() (driver.Value, error) {
	if len(v) == 0 {
		return "{}", nil
	}
	strArr := []string{}
	for _, id := range v {
		if str := id.String(); str != "" {
			strArr = append(strArr, fmt.Sprintf("\"%s\"", id.String()))
		}
	}
	joinedStr := fmt.Sprintf("{%s}", strings.Join(strArr, ","))
	return joinedStr, nil
}

func (v *UUIDArray) ParseStringArr(stringIds *[]string) error {
	for _, strId := range *stringIds {
		if strId == "" {
			return errors.New("invalid id in the array")
		}
		id, err := uuid.Parse(strId)
		if err != nil {
			return err
		}
		*v = append(*v, id)
	}
	return nil
}

func (v *UUIDArray) GetStringArr() []string {
	strArr := []string{}
	for _, id := range *v {
		strArr = append(strArr, id.String())
	}
	return strArr
}
