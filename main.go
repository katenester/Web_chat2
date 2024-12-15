package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// Подключение к базе данных SQLite
	db, err := sql.Open("sqlite3", "main.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Получение списка таблиц
	tablesQuery := `
		SELECT name 
		FROM sqlite_master 
		WHERE type='table'
	`
	rows, err := db.Query(tablesQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Таблицы в базе данных:")
	var tableName string
	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tableName)

		// Вывод полей таблицы
		printTableFields(db, tableName)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func printTableFields(db *sql.DB, tableName string) {
	fmt.Printf("Поля в таблице %s:\n", tableName)

	// Выполнение PRAGMA-запроса для получения информации о полях
	query := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		cid        int
		fieldName  string
		fieldType  string
		notnull    int
		dflt_value sql.NullString
		pk         int
	)

	for rows.Next() {
		err := rows.Scan(&cid, &fieldName, &fieldType, &notnull, &dflt_value, &pk)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  - %s (%s)\n", fieldName, fieldType)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
