package repository

import (
	"github.com/jmoiron/sqlx"
)

type ClinicRepo interface {
	RemoveQueue(idClinic string) error
	AddQueue(idClinic string) error
}

type clinicRepoImpl struct {
	clinicDb *sqlx.DB
}

func (c *clinicRepoImpl) RemoveQueue(idClinic string) error {
	_, err := c.clinicDb.Exec("UPDATE clinic_list SET patient_queue = patient_queue - 1 WHERE id = $1", idClinic)
	if err != nil {
		return err
	}
	return nil
}

func (c *clinicRepoImpl) AddQueue(idClinic string) error {
	_, err := c.clinicDb.Exec("UPDATE clinic_list SET patient_queue = patient_queue + 1 WHERE id = $1", idClinic)
	if err != nil {
		return err
	}
	return nil
}

func NewClinicRepoImpl(clinicDb *sqlx.DB) ClinicRepo {
	clinicRepo := clinicRepoImpl{
		clinicDb: clinicDb,
	}
	return &clinicRepo
}
