package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Content struct {
	ID          string
	Title       string
	Description string
}


func SelectContentAll() []Content {
	var data []Content
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Sukses!")
	}

	rows, err := db.Query("select ID, Title, Description from content ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID          string
			Title       string
			Description string
		)
		err := rows.Scan(&ID, &Title, &Description)
		if err != nil {
			log.Fatal(err)
		}
		var tempData Content
		tempData.ID = ID
		tempData.Title = Title
		tempData.Description = Description
		data = append(data, tempData)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func SelectContentById(idContent string) Content {
	var data Content
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Sukses!")
	}

	rows, err := db.Query("select ID, Title, Description from content where ID=?",idContent)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID          string
			Title       string
			Description string
		)
		err := rows.Scan(&ID, &Title, &Description)
		if err != nil {
			log.Fatal(err)
		}
		data.ID = ID
		data.Title = Title
		data.Description = Description
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func AddContent(ID string, Title string, Description string) bool {

	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Sukses!")
	}

	_, err = db.Exec("INSERT INTO content(ID,Title,Description) value(?,?,?) ", ID, Title, Description)

	if err != nil {
		return false
	}
	return true
}

func UpdateContent(ID string, Title string, Description string) bool {

	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Sukses!")
	}

	_, err = db.Exec("UPDATE content SET Title=?,Description=? WHERE ID =?", Title, Description,ID)

	if err != nil {
		return false
	}
	return true
}

func DeleteContent(ID string) bool {

	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Sukses!")
	}

	_, err = db.Exec("DELETE FROM content WHERE ID =?",ID)

	if err != nil {
		return false
	}
	return true
}