package models

import "time"

type Patient struct {
	ID              uint      `gorm:"primaryKey"`
	Name            string    `json:"name"`
	Gender          string    `json:"gender"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	Address         string    `json:"address"`
	PhoneNumber     string    `json:"phone_number"`
	Email           string    `json:"email"`
	SocialSecurity  string    `json:"social_security"`
	MedicalRecord   string    `json:"medical_record"`
	InsuranceNumber string    `json:"insurance_number"`
	MedicalReports  string    `json:"medical_reports"`
	Forms           string    `json:"forms"`
	UpdatedBy       string    `json:"updated_by"`
	UpdatedAt       time.Time
}

type MedicalHistory struct {
	PatientID           uint   `gorm:"primaryKey"`
	PatientName         string `json:"patient_name"`
	Allergies           string `json:"allergies"`
	Medications         string `json:"medications"`
	PastMedicalHistory  string `json:"past_medical_history"`
	FamilyHistory       string `json:"family_history"`
	ImmunizationHistory string `json:"immunization_history"`
	DietaryPreferences  string `json:"dietary_preferences"`
}

type Appointments struct {
	PatientID          uint      `gorm:"primaryKey"`
	PatientName        string    `json:"patient_name"`
	AppointmentTime    time.Time `json:"appointment_time"`
	VisitNotes         string    `json:"visit_notes"`
	ConsultationReport string    `json:"consultation_report"`
}
