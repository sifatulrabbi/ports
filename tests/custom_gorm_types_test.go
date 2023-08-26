package tests

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/sifatulrabbi/ports/services"
)

func TestUUIDArrayType(t *testing.T) {
	id := services.UUIDArray{uuid.New(), uuid.New(), uuid.New()}

	// returned value test
	v, _ := id.Value()
	t.Log(v)
	str, ok := v.(string)
	if !ok {
		t.Error("returned value is not a string")
		return
	}
	if !strings.ContainsAny(str, "\"{ , }\"") {
		t.Error("returned string isn't properly formatted")
		return
	}

	// Scan value test
	newIds := services.UUIDArray{}
	if err := newIds.Scan(str); err != nil {
		t.Error(err)
		return
	}
	l := len(newIds)
	if l > 0 {
		t.Log(newIds[0].String())
	}
}
