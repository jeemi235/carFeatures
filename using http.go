package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	_ "github.com/lib/pq"
// )

// var db *sql.DB

// type Cardetails struct {
// 	Details_id int    `json:details_id`
// 	Name       string `json:name`
// 	Code       string `json:code`
// 	Color      string `json:color`
// }

// func getdetails(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("select * from car_details")
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	details := []Cardetails{}
// 	for rows.Next() {
// 		var detail Cardetails
// 		if err := rows.Scan(&detail.Details_id, &detail.Name, &detail.Code, &detail.Color); err != nil {
// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			return
// 		}
// 		details = append(details, detail)
// 	}
// 	if err := rows.Err(); err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(details)
// 	fmt.Println("successfully got data from car_details\n")
// }

// func postdetails(w http.ResponseWriter, r *http.Request) {
// 	var detail Cardetails
// 	err := json.NewDecoder(r.Body).Decode((&detail))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, err := db.Exec(
// 		"INSERT INTO car_details (details_id, name,code,color) VALUES ($1, $2,$3,$4)", detail.Details_id, detail.Name, detail.Code, detail.Color); err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(detail)
// 	w.Write([]byte("Data inserted successfully"))
// 	fmt.Println("successfully inserted data in car_details\n")
// }

// func putdetails(w http.ResponseWriter, r *http.Request) {
// 	var detail Cardetails
// 	err := json.NewDecoder(r.Body).Decode((&detail))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, err = db.Exec(
// 		"UPDATE car_details set color = $1 WHERE details_id = $2", detail.Color, detail.Details_id); err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(detail)
// 	w.Write([]byte("Data updated successfully"))
// 	fmt.Println("successfully updated data in car_details\n")
// }

// func deletedetails(w http.ResponseWriter, r *http.Request) {
// 	var detail Cardetails
// 	err := json.NewDecoder(r.Body).Decode((&detail))

// 	details_id := r.URL.Query().Get("details_id")
// 	if _, err = db.Exec(
// 		"delete from car_details where details_id=$1", details_id); err != nil {
// 		log.Fatal(err)
// 	}
// 	w.Write([]byte("Data deleted successfully"))
// 	fmt.Println("successfully deleted data in car_details\n")
// }

// func main() {
// 	var err error
// 	db, err = sql.Open("postgres", "postgresql://max:roach@localhost:26257/car?sslmode=require")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	http.HandleFunc("/get", getdetails)
// 	http.HandleFunc("/post", postdetails)
// 	http.HandleFunc("/put", putdetails)
// 	http.HandleFunc("/delete", deletedetails)
// 	http.ListenAndServe(":5050", nil)

// 	fmt.Println("success")
// }
