package service

import (
	"encoding/json"
	"simple-audit-log-system/app/helper"
	"simple-audit-log-system/app/interface/repository_interface"
	"simple-audit-log-system/app/model"
	"simple-audit-log-system/app/request"
)

type Criminal_Service struct {
	repository repository_interface.Repository_Interface
}

func NewCriminalServiceRegistry(criminal_repository repository_interface.Repository_Interface) *Criminal_Service {
	return &Criminal_Service{
		repository: criminal_repository,
	}
}

func (s *Criminal_Service) Create(criminal *request.Criminal_Request) (*model.CriminalFugitives, error) {

	if criminal.Full_Name == nil || *criminal.Full_Name == "" {
		return nil, helper.NewBadRequest("Full Name Cannot Be Empty")
	}

	if criminal.Danger_Level == nil {
		return nil, helper.NewBadRequest("Danger Level Cannot Be Empty")
	}

	if criminal.Last_Known_Location == nil || *criminal.Last_Known_Location == "" {
		return nil, helper.NewBadRequest("Last Location Cannot Be Empty")
	}

	if criminal.Is_Captured == nil {
		return nil, helper.NewBadRequest("Captured Cannot Be Empty")
	}

	if criminal.Reward_Amount == nil {
		return nil, helper.NewBadRequest("Reward Amount Cannot Be Empty")
	}

	criminal_input := &model.CriminalFugitives{
		Full_Name:           criminal.Full_Name,
		Crime_Description:   criminal.Crime_Description,
		Danger_Level:        criminal.Danger_Level,
		Last_Known_Location: criminal.Last_Known_Location,
		Is_Captured:         criminal.Is_Captured,
		Reward_Amount:       criminal.Reward_Amount,
	}

	errCreateCriminal := s.repository.Create(criminal_input)

	if errCreateCriminal != nil {
		return nil, helper.NewInternalServerError("Error During Create Criminal : " + errCreateCriminal.Error())
	}

	NewDataTypes, _ := json.Marshal(criminal_input)

	audit_input := &model.AuditLogs{
		Entity_Name: "Criminal_Fugitives",
		Entity_ID:   criminal_input.ID,
		Action:      "Create",
		Old_Data:    nil,
		New_Data:    NewDataTypes,
	}

	errCreateAudit := s.repository.CreateAudit(audit_input)

	if errCreateAudit != nil {
		return nil, helper.NewInternalServerError("Error During Create Audit : " + errCreateAudit.Error())
	}

	return criminal_input, nil
}

func (s *Criminal_Service) GetAll() ([]model.CriminalFugitives, error) {

	criminal, errGet := s.repository.GetAll()

	if errGet != nil {
		return nil, helper.NewInternalServerError("Error During Get Criminal : " + errGet.Error())
	}

	return criminal, nil

}

func (s *Criminal_Service) Update(ID string, criminal *request.Criminal_Request) (*model.CriminalFugitives, error) {

	getCriminal, errGet := s.repository.GetById(ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("Error During Get Criminal : " + errGet.Error())
	}

	if criminal.Full_Name == nil || *criminal.Full_Name == "" {
		return nil, helper.NewBadRequest("Full Name Cannot Be Empty")
	}

	if criminal.Danger_Level == nil {
		return nil, helper.NewBadRequest("Danger Level Cannot Be Empty")
	}

	if criminal.Last_Known_Location == nil || *criminal.Last_Known_Location == "" {
		return nil, helper.NewBadRequest("Last Location Cannot Be Empty")
	}

	if criminal.Is_Captured == nil {
		return nil, helper.NewBadRequest("Captured Cannot Be Empty")
	}

	if criminal.Reward_Amount == nil {
		return nil, helper.NewBadRequest("Reward Amount Cannot Be Empty")
	}

	getCriminal.Full_Name = criminal.Full_Name
	getCriminal.Crime_Description = criminal.Crime_Description
	getCriminal.Danger_Level = criminal.Danger_Level
	getCriminal.Last_Known_Location = criminal.Last_Known_Location
	getCriminal.Is_Captured = criminal.Is_Captured
	getCriminal.Reward_Amount = criminal.Reward_Amount

	errUpdate := s.repository.Update(ID, getCriminal)

	if errUpdate != nil {
		return nil, helper.NewInternalServerError("Error During Update Criminal : " + errUpdate.Error())
	}

	NewDataTypes, _ := json.Marshal(getCriminal)
	OldDataTypes, _ := json.Marshal(getCriminal)

	audit_input := &model.AuditLogs{
		Entity_Name: "Criminal_Fugitives",
		Entity_ID:   getCriminal.ID,
		Action:      "Update",
		Old_Data:    OldDataTypes,
		New_Data:    NewDataTypes,
	}

	errCreateAudit := s.repository.CreateAudit(audit_input)

	if errCreateAudit != nil {
		return nil, helper.NewInternalServerError("Error During Create Audit : " + errCreateAudit.Error())
	}

	return getCriminal, nil

}

func (s *Criminal_Service) Delete(ID string) error {

	getCriminal, errGet := s.repository.GetById(ID)

	if errGet != nil {
		return helper.NewInternalServerError("Error During Get Criminal : " + errGet.Error())
	}

	errDelete := s.repository.Delete(ID)

	if errDelete != nil {
		return helper.NewInternalServerError("Error During Delete Criminal : " + errDelete.Error())
	}

	OldDataTypes, _ := json.Marshal(getCriminal)

	audit_input := &model.AuditLogs{
		Entity_Name: "Criminal_Fugitives",
		Entity_ID:   getCriminal.ID,
		Action:      "Delete",
		Old_Data:    OldDataTypes,
		New_Data:    nil,
	}

	errCreateAudit := s.repository.CreateAudit(audit_input)

	if errCreateAudit != nil {
		return helper.NewInternalServerError("Error During Create Audit : " + errCreateAudit.Error())
	}

	return nil
}

func (s *Criminal_Service) GetAllAudit() ([]model.AuditLogs, error) {

	audit, errGet := s.repository.GetAudit()

	if errGet != nil {
		return nil, helper.NewInternalServerError("Error During Get Criminal : " + errGet.Error())
	}

	return audit, nil

}
