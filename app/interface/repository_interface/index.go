package repository_interface

import "simple-audit-log-system/app/model"

type Repository_Interface interface {
	Create(criminal *model.CriminalFugitives) error
	GetAll() ([]model.CriminalFugitives, error)
	Update(ID string, criminal *model.CriminalFugitives) error
	Delete(ID string) error
	CreateAudit(audit *model.AuditLogs) error
	GetAudit() ([]model.AuditLogs, error)
	GetById(ID string) (*model.CriminalFugitives, error)
}
