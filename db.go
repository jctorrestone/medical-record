package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Patient struct {
	ID       int64
	Name     string
	Lastname string
	Gender   string
}

type Exam struct {
	ID      int64
	Hg      bool
	Gvc     bool
	Hepat   bool
	RxTorax bool
	TemTM   bool
	DimeroD bool
	Pcr     bool
	Ferri   bool
	TpTPT   bool
	Procal  bool
	Fibri   bool
}

type Record struct {
	ID         int64
	PatientID  int64
	PatientObj Patient
	Date       string
	Age        int64
	Weight     int64
	Height     int64
	Duration   int64
	ExamID     int64
	ExamObj    Exam
}

type Symptom struct {
	ID          int64
	Description string
}

type RecordSymptom struct {
	ID         int64
	RecordID   int64
	RecordObj  Record
	SymptomID  int64
	SymptomObj Symptom
}

type Disease struct {
	ID          int64
	Description string
}

type Idx struct {
	ID         int64
	RecordID   int64
	RecordObj  Record
	DiseaseID  int64
	DiseaseObj Symptom
}

type Medicine struct {
	ID    int64
	Name  string
	Brand string
	Type  string
	Unit  string
	Dose  int64
}

type Treatment struct {
	ID          int64
	RecordID    int64
	RecordObj   Record
	MedicineID  int64
	MedicineObj Medicine
	Quantity    int64
	Dosage      float64
	Frequency   int64
	Note        string
}

func connect() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "medical_records",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func getPatients() ([]Patient, error) {
	var patients []Patient
	rows, err := db.Query("SELECT * FROM patient ORDER BY last_name ASC")
	if err != nil {
		return nil, fmt.Errorf("getPatients: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var pat Patient
		if err := rows.Scan(&pat.ID, &pat.Name, &pat.Lastname, &pat.Gender); err != nil {
			return nil, fmt.Errorf("getPatients: %v", err)
		}
		patients = append(patients, pat)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getPatients: %v", err)
	}
	return patients, nil
}

func getPatientsByName(query string) ([]Patient, error) {
	var patients []Patient
	query = "%" + query + "%"
	rows, err := db.Query("SELECT * FROM patient WHERE name LIKE ? OR last_name LIKE ? ORDER BY last_name ASC", query, query)
	if err != nil {
		return nil, fmt.Errorf("getPatientsByName %q: %v", query, err)
	}
	defer rows.Close()
	for rows.Next() {
		var patient Patient
		if err := rows.Scan(&patient.ID, &patient.Name, &patient.Lastname, &patient.Gender); err != nil {
			return nil, fmt.Errorf("getPatientsByName %q: %v", query, err)
		}
		patients = append(patients, patient)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getPatientsByName %q: %v", query, err)
	}
	return patients, nil
}

func getPatientById(id int64) (Patient, error) {
	var patient Patient
	row := db.QueryRow("SELECT * FROM patient WHERE id = ?", id)
	if err := row.Scan(&patient.ID, &patient.Name, &patient.Lastname, &patient.Gender); err != nil {
		if err == sql.ErrNoRows {
			return patient, fmt.Errorf("getPatientById %d: no such patient", id)
		}
		return patient, fmt.Errorf("getPatientById %d: %v", id, err)
	}
	return patient, nil
}

func addPatient(pat Patient) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO patient (name, last_name, gender) VALUES (?, ?, ?)",
		pat.Name, pat.Lastname, pat.Gender)
	if err != nil {
		return 0, fmt.Errorf("addPatient: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addPatient: %v", err)
	}
	return id, nil
}

func addExam(exam Exam) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO exam (hg, gvc, hepat, rx_torax, tem_t_m, dimero_d, pcr, ferri, tp_tpt, procal, fibri) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		exam.Hg, exam.Gvc, exam.Hepat, exam.RxTorax, exam.TemTM, exam.DimeroD,
		exam.Pcr, exam.Ferri, exam.TpTPT, exam.Procal, exam.Fibri)
	if err != nil {
		return 0, fmt.Errorf("addExam: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addExam: %v", err)
	}
	return id, nil
}

func addRecord(record Record) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO record (patient_id, rdate, age, weight, height, duration, exam_id) VALUES (?, ?, ?, ?, ?, ?, ?)",
		record.PatientID, record.Date, record.Age, record.Weight, record.Height, record.Duration, record.ExamID)
	if err != nil {
		return 0, fmt.Errorf("addRecord: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addRecord: %v", err)
	}
	return id, nil
}

func getRecords() ([]Record, error) {
	var records []Record
	rows, err := db.Query("SELECT * FROM record INNER JOIN patient ON record.patient_id=patient.id ORDER BY rdate DESC LIMIT 30")
	if err != nil {
		return nil, fmt.Errorf("getRecords: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var rec Record
		if err := rows.Scan(&rec.ID, &rec.PatientID, &rec.Date, &rec.Age, &rec.Weight, &rec.Height, &rec.Duration, &rec.ExamID, &rec.PatientObj.ID, &rec.PatientObj.Name, &rec.PatientObj.Lastname, &rec.PatientObj.Gender); err != nil {
			return nil, fmt.Errorf("getRecords: %v", err)
		}
		records = append(records, rec)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getRecords: %v", err)
	}
	return records, nil
}

func getRecordsByPatient(query string) ([]Record, error) {
	var records []Record
	query = "%" + query + "%"
	rows, err := db.Query("SELECT * FROM record INNER JOIN patient ON record.patient_id=patient.id WHERE patient.name LIKE ? OR patient.last_name LIKE ? ORDER BY rdate DESC LIMIT 30", query, query)
	if err != nil {
		return nil, fmt.Errorf("getRecordsByPatient %q: %v", query, err)
	}
	defer rows.Close()
	for rows.Next() {
		var rec Record
		if err := rows.Scan(&rec.ID, &rec.PatientID, &rec.Date, &rec.Age, &rec.Weight, &rec.Height, &rec.Duration, &rec.ExamID, &rec.PatientObj.ID, &rec.PatientObj.Name, &rec.PatientObj.Lastname, &rec.PatientObj.Gender); err != nil {
			return nil, fmt.Errorf("getRecordsByPatient %q: %v", query, err)
		}
		records = append(records, rec)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getRecordsByPatient %q: %v", query, err)
	}
	return records, nil
}

func addSymptom(symptom Symptom) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO symptom (description) VALUES (?)",
		symptom.Description)
	if err != nil {
		return 0, fmt.Errorf("addSymptom: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addSymptom: %v", err)
	}
	return id, nil
}

func getSymptomById(id int64) (Symptom, error) {
	var symptom Symptom
	row := db.QueryRow("SELECT * FROM symptom WHERE id = ?", id)
	if err := row.Scan(&symptom.ID, &symptom.Description); err != nil {
		if err == sql.ErrNoRows {
			return symptom, fmt.Errorf("getSymptomById %d: no such symptom", id)
		}
		return symptom, fmt.Errorf("getSymptomById %d: %v", id, err)
	}
	return symptom, nil
}

func getSymptomsByDesc(query string) ([]Symptom, error) {
	var symptoms []Symptom
	query = "%" + query + "%"
	rows, err := db.Query("SELECT * FROM symptom WHERE description LIKE ?", query)
	if err != nil {
		return nil, fmt.Errorf("getSymptomsByDesc %q: %v", query, err)
	}
	defer rows.Close()
	for rows.Next() {
		var symptom Symptom
		if err := rows.Scan(&symptom.ID, &symptom.Description); err != nil {
			return nil, fmt.Errorf("getSymptomsByDesc %q: %v", query, err)
		}
		symptoms = append(symptoms, symptom)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getSymptomsByDesc %q: %v", query, err)
	}
	return symptoms, nil
}

func addRecordSymptom(recSymptom RecordSymptom) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO record_symptom (record_id, symptom_id) VALUES (?, ?)",
		recSymptom.RecordID, recSymptom.SymptomID)
	if err != nil {
		return 0, fmt.Errorf("addRecordSymptom: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addRecordSymptom: %v", err)
	}
	return id, nil
}

func addDisease(disease Disease) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO disease (description) VALUES (?)",
		disease.Description)
	if err != nil {
		return 0, fmt.Errorf("addDisease: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addDisease: %v", err)
	}
	return id, nil
}

func getDiseaseById(id int64) (Disease, error) {
	var disease Disease
	row := db.QueryRow("SELECT * FROM disease WHERE id = ?", id)
	if err := row.Scan(&disease.ID, &disease.Description); err != nil {
		if err == sql.ErrNoRows {
			return disease, fmt.Errorf("getDiseaseById %d: no such disease", id)
		}
		return disease, fmt.Errorf("getDiseaseById %d: %v", id, err)
	}
	return disease, nil
}

func getDiseasesByDesc(query string) ([]Disease, error) {
	var diseases []Disease
	query = "%" + query + "%"
	rows, err := db.Query("SELECT * FROM disease WHERE description LIKE ?", query)
	if err != nil {
		return nil, fmt.Errorf("getDiseasesByDesc %q: %v", query, err)
	}
	defer rows.Close()
	for rows.Next() {
		var disease Disease
		if err := rows.Scan(&disease.ID, &disease.Description); err != nil {
			return nil, fmt.Errorf("getDiseasesByDesc %q: %v", query, err)
		}
		diseases = append(diseases, disease)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getDiseasesByDesc %q: %v", query, err)
	}
	return diseases, nil
}

func addIdx(idx Idx) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO idx (record_id, disease_id) VALUES (?, ?)",
		idx.RecordID, idx.DiseaseID)
	if err != nil {
		return 0, fmt.Errorf("addIdx: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addIdx: %v", err)
	}
	return id, nil
}

func addMedicine(medicine Medicine) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO medicine (name, brand, type, unit, dose) VALUES (?, ?, ?, ?, ?)",
		medicine.Name, medicine.Brand, medicine.Type, medicine.Unit, medicine.Dose)
	if err != nil {
		return 0, fmt.Errorf("addMedicine: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addMedicine: %v", err)
	}
	return id, nil
}

func getMedicinesByDesc(query string) ([]Medicine, error) {
	var medicines []Medicine
	query = "%" + query + "%"
	rows, err := db.Query("SELECT * FROM medicine WHERE name LIKE ?", query)
	if err != nil {
		return nil, fmt.Errorf("getMedicinesByDesc %q: %v", query, err)
	}
	defer rows.Close()
	for rows.Next() {
		var medicine Medicine
		if err := rows.Scan(&medicine.ID, &medicine.Name, &medicine.Brand, &medicine.Type, &medicine.Unit, &medicine.Dose); err != nil {
			return nil, fmt.Errorf("getMedicinesByDesc %q: %v", query, err)
		}
		medicines = append(medicines, medicine)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getMedicinesByDesc %q: %v", query, err)
	}
	return medicines, nil
}

func getMedicineById(id int64) (Medicine, error) {
	var medicine Medicine
	row := db.QueryRow("SELECT * FROM medicine WHERE id = ?", id)
	if err := row.Scan(&medicine.ID, &medicine.Name, &medicine.Brand, &medicine.Type, &medicine.Unit, &medicine.Dose); err != nil {
		if err == sql.ErrNoRows {
			return medicine, fmt.Errorf("getMedicineById %d: no such medicine", id)
		}
		return medicine, fmt.Errorf("getMedicineById %d: %v", id, err)
	}
	return medicine, nil
}

func addTreatment(treatment Treatment) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO treatment (record_id, medicine_id, quantity, dosage, frequency, note) VALUES (?, ?, ?, ?, ?, ?)",
		treatment.RecordID, treatment.MedicineID, treatment.Quantity, treatment.Dosage, treatment.Frequency, treatment.Note)
	if err != nil {
		return 0, fmt.Errorf("addTreatment: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addTreatment: %v", err)
	}
	return id, nil
}
