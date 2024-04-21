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

type Record struct {
	ID         int64
	PatientID  int64
	PatientOBj Patient
	Date       string
	Age        int64
	Weight     int64
	Height     int64
	Duration   int64
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

func addPatient(pat Patient) (int64, error) {
	result, err := db.Exec("INSERT INTO patient (name, last_name, gender) VALUES (?, ?, ?)", pat.Name, pat.Lastname, pat.Gender)
	if err != nil {
		return 0, fmt.Errorf("addPatient: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addPatient: %v", err)
	}
	return id, nil
}

func getRecords() ([]Record, error) {
	var records []Record
	rows, err := db.Query("SELECT * FROM record INNER JOIN patient ON record.patient_id=patient.id ORDER BY rdate DESC LIMIT 20")
	if err != nil {
		return nil, fmt.Errorf("getRecords: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var rec Record
		if err := rows.Scan(&rec.ID, &rec.PatientID, &rec.Date, &rec.Age, &rec.Weight, &rec.Height, &rec.Duration, &rec.PatientOBj.ID, &rec.PatientOBj.Name, &rec.PatientOBj.Lastname, &rec.PatientOBj.Gender); err != nil {
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
	rows, err := db.Query("SELECT * FROM record INNER JOIN patient ON record.patient_id=patient.id WHERE patient.name LIKE ? OR patient.last_name LIKE ?", query, query)
	if err != nil {
		return nil, fmt.Errorf("getRecordsByPatient %q: %v", query, err)
	}
	defer rows.Close()
	for rows.Next() {
		var rec Record
		if err := rows.Scan(&rec.ID, &rec.PatientID, &rec.Date, &rec.Age, &rec.Weight, &rec.Height, &rec.Duration, &rec.PatientOBj.ID, &rec.PatientOBj.Name, &rec.PatientOBj.Lastname, &rec.PatientOBj.Gender); err != nil {
			return nil, fmt.Errorf("getRecordsByPatient %q: %v", query, err)
		}
		records = append(records, rec)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getRecordsByPatient %q: %v", query, err)
	}
	return records, nil
}
