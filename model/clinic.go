package model

import (
	"time"
)

type Clinic struct {
	ClinicId     string    `db:"id"`
	ClinicName   string    `db:"clinic_name"`
	DoctorName   string    `db:"doctor_name"`
	StartTime    time.Time `db:"start_time"`
	EndTime      time.Time `db:"end_time"`
	PatientQueue int       `db:"patient_queue"`
}

func NewClinic(clinicId, clinicName, doctorName string, startTime, endTime time.Time, patientQueue int) Clinic {
	return Clinic{
		ClinicId:     clinicId,
		ClinicName:   clinicId,
		DoctorName:   doctorName,
		StartTime:    startTime,
		EndTime:      endTime,
		PatientQueue: patientQueue,
	}
}
