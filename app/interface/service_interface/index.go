package service_interface

import (
	"simple-audit-log-system/app/model"
	"simple-audit-log-system/app/request"
)

type Service_Interface interface {
	Create(criminal *request.Criminal_Request) (*model.CriminalFugitives, error)
	GetAll() ([]model.CriminalFugitives, error)
	Update(ID string, criminal *request.Criminal_Request) (*model.CriminalFugitives, error)
	Delete(ID string) error
	GetAllAudit() ([]model.AuditLogs, error)
}
