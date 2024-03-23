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
	InsuranceNumber string    `json:"insurance_number"`
	UpdatedBy       string    `json:"updated_by"`
	UpdatedAt       time.Time
	MedicalHistory  []MedicalHistory `gorm:"foreignKey:PatientID"`
	MedicalReports  []MedicalReports `gorm:"foreignKey:PatientID"`
	MedicalForms    []MedicalForms   `gorm:"foreignKey:PatientID"`
	Appointments    []Appointments   `gorm:"foreignKey:PatientID"`
}

type MedicalHistory struct {
	PatientID           uint   `gorm:"primaryKey"`
	Allergies           string `json:"allergies"`
	Medications         string `json:"medications"`
	PastMedicalHistory  string `json:"past_medical_history""`
	FamilyHistory       string `json:"family_history"`
	ImmunizationHistory string `json:"immunization_history"`
	DietaryPreferences  string `json:"dietary_preferences"`
}

type MedicalReports struct {
	PatientID uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	Content   string    `json:"content"`
}

type MedicalForms struct {
	PatientID uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	Content   string    `json:"content"`
}

type Appointments struct {
	PatientID          uint      `gorm:"primaryKey"`
	AppointmentTime    time.Time `json:"appointment_time"`
	VisitNotes         string    `json:"visit_notes"`
	ConsultationReport string    `json:"consultation_report"`
}
