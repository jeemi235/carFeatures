package main

// import (
// 	"errors"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got / request\n")
// 	io.WriteString(w, "This is my website!\n")

// }

// func main() {
// 	http.HandleFunc("/", getRoot)
// 	err := http.ListenAndServe(":3333", nil)
// 	if errors.Is(err, http.ErrServerClosed) {
// 		fmt.Printf("server closed\n")
// 	} else if err != nil {
// 		fmt.Printf("error starting server: %s\n", err)
// 		os.Exit(1)
// 	}
// }

// create both the tables for details and features
	// if _, err := db.Exec(
	// 	"CREATE TABLE IF NOT EXISTS car_features(features_id int primary key,engine string,sppedometer string,gear string)"); err != nil {
	// 	log.Fatal(err)
	// }
	// if _, err := db.Exec(
	// 	"CREATE TABLE IF NOT EXISTS car_details(id int primary key,name string,code varchar(5),color string,feature int references car_features(features_id))"); err != nil {
	// 	log.Fatal(err)
	// }

	//insert
	// if _, err := db.Exec(
	// 	"INSERT INTO car_details (id, name) VALUES (14, 'BMW')"); err != nil {
	// 	log.Fatal(err)
	// }

	//select
	// rows, err := db.Query("select * from car_details where feature=1")
	// for rows.Next() {
	// 	var id, feature int
	// 	var name, code, color string
	// 	if err := rows.Scan(&id, &name, &code, &color, &feature); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%d %s %s %s %d\n", id, name, code, color, feature)
	// }

	//update
	// 	if _, err := db.Exec(
	// 		"update car_details set color='blue' where id=7"); err != nil {
	// 		log.Fatal(err)
	// 	}