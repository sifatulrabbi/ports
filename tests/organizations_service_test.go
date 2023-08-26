package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sifatulrabbi/ports/services"
)

const (
	mockOrgId  = "0f841993-41fa-4f55-b536-b4f057eaf454"
	mockUserId = "da525320-43e0-478a-a3c2-b5424a6f8fa5"
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
		MemberIDs:   []string{},
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

func TestAddMember(t *testing.T) {
	s := getOrganizationsService(t)
	id, err := uuid.Parse(mockOrgId)
	if err != nil {
		t.Error(err)
		return
	}
	org, err := s.AddMember(services.OrganizationFilter{ID: id}, mockUserId)
	if err != nil {
		t.Error(err)
	}
	idFound := false
	for _, mid := range org.MemberIDs {
		if mid.String() == mockOrgId {
			idFound = false
		}
	}
	if !idFound {
		t.Error("member id is not added")
		return
	}
	t.Log(org.String())
}

func TestRemoveMember(t *testing.T) {

}

func TestAddAdmin(t *testing.T) {
	s := getOrganizationsService(t)
	id, err := uuid.Parse(mockOrgId)
	if err != nil {
		return
	}
	org, err := s.AddAdmin(services.OrganizationFilter{ID: id}, mockUserId)
	if err != nil {
		t.Error(err)
	}
	idFound := false
	for _, aid := range org.AdminIDs {
		if aid.String() == mockOrgId {
			idFound = false
		}
	}
	if !idFound {
		t.Error("admin id is not added")
		return
	}
	t.Log(org.String())
}
