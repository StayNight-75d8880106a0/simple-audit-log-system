package repository

import (
	"simple-audit-log-system/app/database"
	"simple-audit-log-system/app/model"

	"gorm.io/gorm"
)

type Criminal_Repository struct {
	DB *gorm.DB
}

func NewCriminalRepositoryRegistry() *Criminal_Repository {
	dbCluster := database.GetInstanceDbCluster()
	return &Criminal_Repository{
		DB: dbCluster.DB,
	}
}

func (repo *Criminal_Repository) Create(criminal *model.CriminalFugitives) error {

	errCreate := repo.DB.Table("criminal_fugitives").Create(&criminal).Error

	return errCreate

}

func (repo *Criminal_Repository) GetAll() ([]model.CriminalFugitives, error) {

	var criminal []model.CriminalFugitives

	errGet := repo.DB.Table("criminal_fugitives").Where("deleted_at IS null").Find(&criminal).Error

	return criminal, errGet

}

func (repo *Criminal_Repository) GetById(ID string) (*model.CriminalFugitives, error) {

	var criminal *model.CriminalFugitives

	errGet := repo.DB.Table("criminal_fugitives").Where("id = ? AND deleted_at IS null", ID).First(&criminal).Error

	return criminal, errGet

}

func (repo *Criminal_Repository) Update(ID string, criminal *model.CriminalFugitives) error {

	errUpdate := repo.DB.Table("criminal_fugitives").Where("id = ?", ID).Updates(&criminal).Error

	return errUpdate

}

func (repo *Criminal_Repository) Delete(ID string) error {

	var criminal *model.CriminalFugitives

	errDelete := repo.DB.Table("criminal_fugitives").Where("id = ?", ID).Delete(&criminal).Error

	return errDelete

}

func (repo *Criminal_Repository) CreateAudit(audit *model.AuditLogs) error {

	errCreate := repo.DB.Table("audit_logs").Create(&audit).Error

	return errCreate

}

func (repo *Criminal_Repository) GetAudit() ([]model.AuditLogs, error) {

	var audit []model.AuditLogs

	errGet := repo.DB.Table("audit_logs").Find(&audit).Error

	return audit, errGet

}
