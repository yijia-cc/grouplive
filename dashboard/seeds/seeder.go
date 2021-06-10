package main

import (
	"database/sql"
	"fmt"
	"github.com/yijia-cc/grouplive/dashboard/db"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yijia-cc/grouplive/dashboard/config"
)

func main() {
	cfg := config.FromEnv()
	db.Conn(cfg)
	defer db.Close()

	log.Printf("Connected to DB %s successfully!\n", cfg.Db["dash"].DBName)
	var err error

	err = createEventCategoryTable(db.DashDB)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = createEventTypeTable(db.DashDB)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = createEventTable(db.DashDB)
	if err != nil {
		log.Fatalf("Error %s when creating event table", err)
	}

	err = createMediaTable(db.DashDB)
	if err != nil {
		log.Fatalf("Error %s when creating media table", err)
	}

	err = createUserReactionTable(db.DashDB)
	if err != nil {
		log.Fatalf("Error %s when creating media table", err)
	}

	log.Printf("Connection to DB <%s> is still alive!\n", cfg.Db["dash"].DBName)
}

func createMediaTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS media (
				id INT PRIMARY KEY AUTO_INCREMENT,  
				event_id INT NOT NULL, 
				media_name VARCHAR(255) NOT NULL,
				media_url VARCHAR(255) NOT NULL,
    			created_at DATETIME,
    			updated_at DATETIME,
    			active BOOL DEFAULT TRUE, 
    			CONSTRAINT media_fk_event 
    				FOREIGN KEY (event_id) REFERENCES event(id)
			)`

	_, err := db.Exec(sql)
	return err
}


func createEventCategoryTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS category (
				id TINYINT PRIMARY KEY AUTO_INCREMENT,  
				name VARCHAR(100) UNIQUE NOT NULL
			)`

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	// add data to the table
	err = addCategory(db)
	if err != nil {
		return err
	}

	return nil
}

func createEventTypeTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS type (
				id TINYINT PRIMARY KEY AUTO_INCREMENT,  
				category_id TINYINT NOT NULL, 
				name VARCHAR(200) UNIQUE NOT NULL,
    			CONSTRAINT subcategory_fk_category 
    				FOREIGN KEY (category_id) REFERENCES category(id)
			)`

	_, err := db.Exec(sql)
	if err != nil {
		return fmt.Errorf("error when creating sub_category table: %s", err)
	}

	// add data to the table
	err = addType(db)
	if err != nil {
		return fmt.Errorf("error when inserting into sub_category table: %s", err)
	}

	return err
}


func createEventTable(db *sql.DB) error {
	// Policy may have a start_time, but no end_time
	sql := `CREATE TABLE IF NOT EXISTS event (
				id INT PRIMARY KEY AUTO_INCREMENT,  
				type_id TINYINT NOT NULL, 
				username VARCHAR(50) NOT NULL,
				title VARCHAR(255) NOT NULL,
    			description TEXT,
    			start_time DATETIME NOT NULL,
    			end_time DATETIME,  
    			rsvp_required BOOL DEFAULT FALSE,
    			created_at DATETIME,
    			updated_at DATETIME,
    			active BOOL DEFAULT TRUE,
    			CONSTRAINT event_fk_type
    				FOREIGN KEY (type_id) REFERENCES type(id)
			)`

	_, err := db.Exec(sql)
	return err
}


func createUserReactionTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS reaction (
				id INT PRIMARY KEY AUTO_INCREMENT,
				event_id INT NOT NULL,  
				username VARCHAR(50) NOT NULL, 
    			attend BOOL,
    			created_at DATETIME,
    			updated_at DATETIME,
    			active bool DEFAULT TRUE,
    			CONSTRAINT reaction_fk_event 
    				FOREIGN KEY (event_id) REFERENCES event(id)
			)`

	_, err := db.Exec(sql)
	return err
}


func addCategory(db *sql.DB) error {
	sql := "INSERT INTO category (name) VALUES (?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec("Policy")
	if err != nil {
		return err
	}

	stmt.Exec("Announcement")
	if err != nil {
		return err
	}

	stmt.Exec("Alert")
	if err != nil {
		return err
	}

	stmt.Exec("Event")
	if err != nil {
		return err
	}

	return nil
}

func addType(db *sql.DB) error {
	sql := "INSERT INTO type (category_id, name) VALUES (?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 1. Policy
	_, err = stmt.Exec(1, "Condo Policy")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(1, "Leasing Agreement")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(1, "Tenant Rules")
	if err != nil {
		return err
	}

	// 2. Announcement
	_, err = stmt.Exec(2, "Found & Lost")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(2, "Fire Alarm Testing")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(2, "Elevator Reservation")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(2, "Gym Maintenance")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(2, "Swimming Pool Maintenance")
	if err != nil {
		return err
	}


	// 3. Alert
	_, err = stmt.Exec(3, "Heat Warning")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(3, "Storm Alert")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(3, "Hurricane Alert")
	if err != nil {
		return err
	}

	// 4. Event
	_, err = stmt.Exec(4, "BBQ")
	if err != nil {
		return err
	}

	stmt.Exec(4, "Music Night")
	if err != nil {
		return err
	}

	stmt.Exec(4, "Dog Party")
	if err != nil {
		return err
	}

	stmt.Exec(4, "Movie Night")
	if err != nil {
		return err
	}

	stmt.Exec(4, "Coffee Gathering")
	if err != nil {
		return err
	}

	stmt.Exec(4, "Cookie Night")
	if err != nil {
		return err
	}

	stmt.Exec(4, "Halloween Night")
	if err != nil {
		return err
	}

	stmt.Exec(4, "Kids Gathering")
	if err != nil {
		return err
	}

	stmt.Exec(4, "Bicycle Race")
	if err != nil {
		return err
	}

	return nil
}


