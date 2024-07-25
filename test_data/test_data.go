package testdata

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"

	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/go-sql-driver/mysql"
)

// Engineer represents a row in the engineer table
type Engineer struct {
	FirstName string
	LastName  string
	Gender    int
	CountryID int64
	Title     string
}

func Insert() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		"mysql:3306",
		os.Getenv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	defer db.Close()

	numWorkers, _ := strconv.Atoi(os.Getenv("NUM_WORKERS")) 
	numRecords, _ := strconv.Atoi(os.Getenv("NUM_RECORDS"))

	jobs := make(chan Engineer, numRecords)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(db, jobs, &wg)
	}

	// Generate and send jobs
	for i := 0; i < numRecords; i++ {
		jobs <- generateEngineer()
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Printf("Successfully inserted %d records into the engineer table.\n", numRecords)
}

func generateEngineer() Engineer {
	gofakeit.Seed(0)
	return Engineer{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Gender:    rand.Intn(2),              // Assuming 0 for female, 1 for male
		CountryID: int64(rand.Intn(100) + 1), // Adjust based on your country table data
		Title:     gofakeit.JobTitle(),
	}
}

func insertEngineer(db *sql.DB, engineer Engineer) error {
	query := "INSERT INTO engineer (first_name, last_name, gender, country_id, title) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, engineer.FirstName, engineer.LastName, engineer.Gender, engineer.CountryID, engineer.Title)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}
	return nil
}

func worker(db *sql.DB, jobs <-chan Engineer, wg *sync.WaitGroup) {
	defer wg.Done()
	for engineer := range jobs {
		err := insertEngineer(db, engineer)
		if err != nil {
			log.Printf("Error inserting engineer: %s", err)
		}
	}
}
