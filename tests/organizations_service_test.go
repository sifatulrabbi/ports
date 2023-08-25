package tests

import (
	"testing"

	"github.com/sifatulrabbi/ports/services"
)

func getOrganizationsService(t *testing.T) *services.OrganizationsService {
	db := getTestDB(t)
	if s, err := services.NewOrganizationsService(db); err != nil {
		t.Error(err)
		return nil
	} else {
		return s
	}
}

func TestCreateOrganization(t *testing.T) {
	s := getOrganizationsService(t)
}

func TestUpdateOneOrganization(t *testing.T) {
	s := getOrganizationsService(t)
}

func TestDeleteOneOrganization(t *testing.T) {
	s := getOrganizationsService(t)
}

func TestGetOneOrganization(t *testing.T) {
	s := getOrganizationsService(t)
}

func TestGetManyOrganization(t *testing.T) {
	s := getOrganizationsService(t)
}
