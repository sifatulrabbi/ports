package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sifatulrabbi/ports/services"
)

const (
	mockOrgId = "599e1754-6e42-4555-bc27-8a9eeb972d4c"
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
	p := services.OrganizationPayload{
		Name:        "Sifatul's Org",
		Email:       "sifatul@sifatul.com",
		Description: "Testing organization",
		MemberIDs:   []string{"da525320-43e0-478a-a3c2-b5424a6f8fa5"},
	}
	if org, err := s.CreateOne(p); err != nil {
		t.Error(err)
	} else {
		t.Log(org.String())
	}
}

func TestUpdateOneOrganization(t *testing.T) {
	s := getOrganizationsService(t)
	id, err := uuid.Parse(mockOrgId)
	if err != nil {
		t.Error(err)
		return
	}
	p := services.OrganizationPayload{
		Name: "Sifatul's Organization",
	}
	f := services.OrganizationFilter{
		ID: id,
	}
	if org, err := s.UpdateOne(f, p); err != nil {
		t.Error(err)
	} else {
		t.Log(org.String())
	}
}

func TestDeleteOneOrganization(t *testing.T) {
	s := getOrganizationsService(t)
	id, err := uuid.Parse(mockOrgId)
	if err != nil {
		t.Error(err)
		return
	}
	f := services.OrganizationFilter{ID: id}
	if err := s.DeleteOne(f); err != nil {
		t.Error(err)
	}
}

func TestGetOneOrganization(t *testing.T) {
	s := getOrganizationsService(t)
	id, err := uuid.Parse(mockOrgId)
	if err != nil {
		t.Error(err)
		return
	}
	f := services.OrganizationFilter{ID: id}
	if org, err := s.GetOne(f); err != nil {
		t.Error(err)
	} else {
		t.Log(org.String())
	}
}

func TestGetManyOrganization(t *testing.T) {
	s := getOrganizationsService(t)
	if orgs, err := s.GetMany(services.OrganizationFilter{}); err != nil {
		t.Error(err)
	} else {
		t.Logf("organizations found: %v", len(*orgs))
	}
}
