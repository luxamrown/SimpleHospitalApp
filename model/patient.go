package model

type Patient struct {
	PatientId     string `db:"id"`
	PatientName   string `db:"patient_name"`
	EarlySymptoms string `db:"early_symptoms"`
	ClinicId      string `db:"clinic_id"`
	IsDone        bool   `db:"is_done"`
}

func NewPatient(patientId, patientName, earlySymptoms, clinicId string, isDone bool) Patient {
	return Patient{
		PatientId:     patientId,
		PatientName:   patientName,
		EarlySymptoms: earlySymptoms,
		ClinicId:      clinicId,
	}
}
