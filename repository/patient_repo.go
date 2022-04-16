package repository

import (
	"enigmacamp.com/rumah_sakit/model"
	"github.com/jmoiron/sqlx"
)

type PatientRepo interface {
	InsertPatient(newPatient model.Patient) (model.Patient, error)
	PatientStatusSetDone(patientId string) (model.Patient, error)
	GetClinicIdByPatient(patientId string) string
}

type patientRepoImpl struct {
	patientDb *sqlx.DB
}

func (p *patientRepoImpl) InsertPatient(newPatient model.Patient) (model.Patient, error) {

	_, err := p.patientDb.Exec("INSERT INTO patient_list(id, patient_name, early_symptoms, clinic_id, is_done) VALUES($1, $2, $3, $4, $5)", newPatient.PatientId, newPatient.PatientName, newPatient.EarlySymptoms, newPatient.ClinicId, newPatient.IsDone)
	if err != nil {
		return model.Patient{}, err
	}
	return newPatient, nil
}

func (p *patientRepoImpl) GetClinicIdByPatient(patientId string) string {
	var clinicId string
	err := p.patientDb.Get(&clinicId, "SELECT clinic_id FROM patient_list WHERE id = $1", patientId)
	if err != nil {
		return ""
	}
	return clinicId
}

func (p *patientRepoImpl) PatientStatusSetDone(patientId string) (model.Patient, error) {
	selectedPatient := model.Patient{}
	_, err := p.patientDb.Exec("UPDATE patient_list SET is_done = true WHERE id=$1", patientId)
	if err != nil {
		return model.Patient{}, err
	}
	err = p.patientDb.Get(&selectedPatient, "SELECT * FROM patient_list WHERE id=$1", patientId)
	if err != nil {
		return model.Patient{}, err
	}
	return selectedPatient, nil
}

func NewPatientRepoImpl(patientDb *sqlx.DB) PatientRepo {
	patientRepo := patientRepoImpl{
		patientDb: patientDb,
	}
	return &patientRepo
}
