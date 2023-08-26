package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ OrganizationsCRUD = &OrganizationsService{}

type OrganizationsCRUD interface {
	crudService[Organization, OrganizationPayload, OrganizationFilter]
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
	if err := memberIds.NewFromStrings(&p.MemberIDs); err != nil {
		return nil, err
	}
	org := Organization{
		ID:          uuid.New(),
		Name:        p.Name,
		Email:       p.Email,
		Description: p.Description,
		MemberIDs:   memberIds,
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
	updateData := map[string]any{
		"name":        p.Name,
		"email":       p.Email,
		"description": p.Description,
	}
	if len(p.MemberIDs) > 0 {
		memberIds := UUIDArray{}
		if err := memberIds.NewFromStrings(&p.MemberIDs); err != nil {
			return nil, err
		}
		updateData["member_ids"] = memberIds
	}
	res := s.db.Model(&org).Where("id = ?", f.ID).Updates(updateData)
	if res.Error != nil {
		return nil, res.Error
	}
	if err := res.First(&org, f.ID).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (s *OrganizationsService) DeleteOne(f OrganizationFilter) error {
	err := s.db.Delete(&Organization{ID: f.ID}, f.ID).Error
	return err
}

func NewOrganizationsService(db *gorm.DB) (*OrganizationsService, error) {
	if err := db.AutoMigrate(&Organization{}); err != nil {
		return nil, err
	}
	service := newServiceWithDB[Organization, OrganizationPayload, OrganizationFilter](db, "OrganizationsService")
	orgService := &OrganizationsService{ServiceWithDB: service}
	return orgService, nil
}
