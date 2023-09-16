package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ OrganizationsCRUD = &OrganizationsService{}

func NewOrganizationsService(db *gorm.DB) (*OrganizationsService, error) {
	if err := db.AutoMigrate(&Organization{}); err != nil {
		return nil, err
	}
	service := newServiceWithDB[Organization, OrganizationPayload, OrganizationFilter](db, "OrganizationsService")
	orgService := &OrganizationsService{ServiceWithDB: service}
	return orgService, nil
}

type OrganizationsCRUD interface {
	crudService[Organization, OrganizationPayload, OrganizationFilter]
	AddMember(f OrganizationFilter, memberId string) (*Organization, error)
	RemoveMember(f OrganizationFilter, memberId string) (*Organization, error)
	AddAdmin(f OrganizationFilter, adminId string) (*Organization, error)
	RemoveAdmin(f OrganizationFilter, adminId string) (*Organization, error)
}

type OrganizationsService struct {
	*ServiceWithDB[Organization, OrganizationPayload, OrganizationFilter]
}

type Organization struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;not null" json:"id"`
	Name        string    `gorm:"type:text;not null" json:"name"`
	Email       string    `gorm:"type:text;not null" json:"email"`
	Description string    `gorm:"type:text;not null" json:"description"`
	AdminIDs    UUIDArray `gorm:"type:uuid[];not nul" json:"admin_ids"`
	MemberIDs   UUIDArray `gorm:"type:uuid[];not null" json:"member_ids"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (org *Organization) String() string {
	return fmt.Sprintf(
		"Organization <%s, %s, %s, %s, %s, %s, %s>\n",
		org.ID, org.Name, org.Email, org.Description, org.MemberIDs, org.CreatedAt, org.UpdatedAt,
	)
}

type OrganizationFilter struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
}

type OrganizationPayload struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Description string   `json:"description"`
	AdminIDs    []string `json:"admin_ids"`
	MemberIDs   []string `json:"member_ids"`
}

func (s *OrganizationsService) GetOne(f OrganizationFilter) (*Organization, error) {
	org := Organization{ID: f.ID}
	if err := s.db.First(&org, f.ID).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (s *OrganizationsService) GetMany(f OrganizationFilter) (*[]Organization, error) {
	orgs := []Organization{}
	if err := s.db.Find(&orgs).Error; err != nil {
		return nil, err
	}
	return &orgs, nil
}

func (s *OrganizationsService) CreateOne(p OrganizationPayload) (*Organization, error) {
	memberIds := UUIDArray{}
	if err := memberIds.ParseStringArr(&p.MemberIDs); err != nil {
		return nil, err
	}
	org := Organization{
		ID:          uuid.New(),
		Name:        p.Name,
		Email:       p.Email,
		Description: p.Description,
		MemberIDs:   memberIds,
		AdminIDs:    memberIds,
	}
	if org.Name == "" {
		return nil, errors.New("'name' is required")
	} else if org.Email == "" {
		return nil, errors.New("'email' is required")
	} else if org.Description == "" {
		return nil, errors.New("'description' is required")
	} else if org.MemberIDs == nil {
		return nil, errors.New("'member_ids' is required")
	}
	if err := s.db.Create(&org).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (s *OrganizationsService) UpdateOne(f OrganizationFilter, p OrganizationPayload) (*Organization, error) {
	org := Organization{ID: f.ID}
	updateData := map[string]any{}
	if p.Name != "" {
		updateData["name"] = p.Name
	}
	if p.Email != "" {
		updateData["email"] = p.Email
	}
	if p.Description != "" {
		updateData["description"] = p.Description
	}
	if len(p.MemberIDs) > 0 {
		memberIds := UUIDArray{}
		if err := memberIds.ParseStringArr(&p.MemberIDs); err != nil {
			return nil, err
		}
		updateData["member_ids"] = memberIds
	}
	if len(p.AdminIDs) > 0 {
		adminIds := UUIDArray{}
		if err := adminIds.ParseStringArr(&p.AdminIDs); err != nil {
			return nil, err
		}
		updateData["admin_ids"] = adminIds
	}
	if err := s.db.Model(&org).Where("id = ?", f.ID).Updates(updateData).Error; err != nil {
		return nil, err
	}
	if err := s.db.First(&org, f.ID).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (s *OrganizationsService) DeleteOne(f OrganizationFilter) error {
	err := s.db.Delete(&Organization{ID: f.ID}, f.ID).Error
	return err
}

func (s *OrganizationsService) AddMember(f OrganizationFilter, memberId string) (*Organization, error) {
	// instantiate users service
	usersService, err := NewUsersService(s.db)
	if err != nil {
		return nil, err
	}
	mid, err := uuid.Parse(memberId)
	if err != nil {
		return nil, err
	}
	if _, err = usersService.GetOne(UserFilter{ID: mid}); err != nil {
		return nil, err
	}
	org, err := s.GetOne(f)
	if err != nil {
		return nil, err
	}
	stringIds := []string{memberId}
	for _, id := range org.MemberIDs {
		if id.String() == memberId {
			return org, nil
		}
		stringIds = append(stringIds, id.String())
	}
	org, err = s.UpdateOne(f, OrganizationPayload{
		MemberIDs: stringIds,
	})
	return org, err
}

func (s *OrganizationsService) RemoveMember(f OrganizationFilter, memberId string) (*Organization, error) {
	org, err := s.GetOne(f)
	if err != nil {
		return nil, err
	}
	if _, err := uuid.Parse(memberId); err != nil {
		return nil, err
	}
	newIdsList := []string{}
	for _, id := range org.MemberIDs {
		if id.String() != memberId {
			newIdsList = append(newIdsList, id.String())
		}
	}
	if org, err = s.UpdateOne(f, OrganizationPayload{MemberIDs: newIdsList}); err != nil {
		return nil, err
	} else {
		return org, nil
	}
}

func (s *OrganizationsService) AddAdmin(f OrganizationFilter, adminId string) (*Organization, error) {
	// instantiate users service
	usersService, err := NewUsersService(s.db)
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(adminId)
	if err != nil {
		return nil, err
	}
	if _, err = usersService.GetOne(UserFilter{ID: id}); err != nil {
		return nil, err
	}
	org, err := s.GetOne(f)
	if err != nil {
		return nil, err
	}
	stringIds := []string{adminId}
	for _, id := range org.AdminIDs {
		if id.String() == adminId {
			return org, nil
		}
		stringIds = append(stringIds, id.String())
	}
	org, err = s.UpdateOne(f, OrganizationPayload{
		AdminIDs: stringIds,
	})
	return org, err
}

func (s *OrganizationsService) RemoveAdmin(f OrganizationFilter, adminId string) (*Organization, error) {
	org, err := s.GetOne(f)
	if err != nil {
		return nil, err
	}
	if _, err := uuid.Parse(adminId); err != nil {
		return nil, err
	}
	newIdsList := []string{}
	for _, id := range org.AdminIDs {
		if id.String() != adminId {
			newIdsList = append(newIdsList, id.String())
		}
	}
	if org, err = s.UpdateOne(f, OrganizationPayload{AdminIDs: newIdsList}); err != nil {
		return nil, err
	} else {
		return org, nil
	}
}
