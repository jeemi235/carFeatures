package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Cardetails struct {
	Id    int    `json:id`
	Name  string `json:name`
	Code  string `json:code`
	Color string `json:color`
}

type Featuredetails struct {
	Id   int    `json:id`
	Name string `json:name`
}

type Relationdetails struct {
	Id         int `json:id`
	Car_id     int `json:car_id`
	Feature_id int `json:feature_id`
}

type CarFeatures struct {
	Id       int              `json:id`
	Name     string           `json:name`
	Code     string           `json:code`
	Color    string           `json:color`
	Features []Featuredetails `json:"features"`
}

func main() {
	var err error
	db, err = sql.Open("postgres", "postgresql://max:roach@localhost:26257/cardb?sslmode=require")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/cars", getcardetails).Methods("GET")                 //0
	router.HandleFunc("/cars", addcar).Methods("POST")                       //1
	router.HandleFunc("/cars/{id}", updatecar).Methods("PUT")                //2
	router.HandleFunc("/features", addfeature).Methods("POST")               //3
	router.HandleFunc("/features", getcarbyfeatures).Methods("GET")          //4
	router.HandleFunc("/carsbycolors", getcarbycolors).Methods("GET")        //5
	router.HandleFunc("/carwithfeatures", getcarwithfeatures).Methods("GET") //6
	router.HandleFunc("/cars/{name}", searchcar).Methods("GET")              //7

	log.Println("Listening ...")
	http.ListenAndServe(":5050", router)
	fmt.Println("success")
}

func getcardetails(resp http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from car")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	details := []Cardetails{}
	for rows.Next() {
		var detail Cardetails
		if err := rows.Scan(&detail.Id, &detail.Name, &detail.Code, &detail.Color); err != nil {
			log.Fatal(err)
		}
		details = append(details, detail)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(resp).Encode(details)

	fmt.Print("successfully got all cars detail\n")
}

func addcar(resp http.ResponseWriter, r *http.Request) {
	var detail Cardetails
	err := json.NewDecoder(r.Body).Decode((&detail))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		"INSERT INTO car (name,code,color) VALUES ($1,$2,$3)", detail.Name, detail.Code, detail.Color); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(resp).Encode(detail)
	resp.Write([]byte("Data inserted successfully"))
	fmt.Print("successfully inserted data in car details\n")
}

func updatecar(resp http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	temp := vars["id"]
	//id := r.URL.Query().Get("id")
	var detail Cardetails
	err := json.NewDecoder(r.Body).Decode((&detail))
	if err != nil {
		log.Fatal(err)
	}

	if _, err = db.Exec(
		"UPDATE car set color = $1 WHERE id = $2", detail.Color, temp); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(resp).Encode(detail)
	resp.Write([]byte("Data updated successfully"))
	fmt.Print("successfully updated data in car details\n")
}

func addfeature(resp http.ResponseWriter, r *http.Request) {
	var relation Relationdetails
	err := json.NewDecoder(r.Body).Decode((&relation))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		"INSERT INTO relation (car_id,feature_id) VALUES ($1,$2)", relation.Car_id, relation.Feature_id); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(resp).Encode(relation)
	resp.Write([]byte("Data inserted successfully"))
	fmt.Print("successfully inserted data in relation details\n")
}

func getcarbyfeatures(resp http.ResponseWriter, r *http.Request) {
	values := []string{}
	values = r.URL.Query()["values"]
	color := r.URL.Query().Get("color")
	rows, err := db.Query("select car.id ,car.name ,car.code ,car.color,features.id ,features.name from car left join relation on relation.car_id = car.id left join features on features.id = relation.feature_id where features.name=any($1) and car.color=$2", pq.Array(values), color)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cars := make(map[int]CarFeatures)
	for rows.Next() {
		var carID int
		var carName string
		var carCode string
		var carColor string
		var featureID int
		var featureName string

		if err := rows.Scan(&carID, &carName, &carCode, &carColor, &featureID, &featureName); err != nil {
			log.Fatal(err)
		}
		carFeatures, ok := cars[carID]
		if !ok {
			// Create a new CarResponse object for this car
			carFeatures = CarFeatures{
				Id:    carID,
				Name:  carName,
				Code:  carCode,
				Color: carColor,
			}
		}
		carFeatures.Features = append(carFeatures.Features, Featuredetails{Id: featureID, Name: featureName})
		cars[carID] = carFeatures
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(resp).Encode(cars)
	fmt.Print("successfully got car details by features\n")
}

func getcarbycolors(resp http.ResponseWriter, r *http.Request) {
	colors := []string{}
	colors = r.URL.Query()["colors"]
	rows, err := db.Query("select * from car where color=any($1)", pq.Array(colors))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	details := []Cardetails{}
	for rows.Next() {
		var detail Cardetails
		if err := rows.Scan(&detail.Id, &detail.Name, &detail.Code, &detail.Color); err != nil {
			log.Fatal(err)
		}
		details = append(details, detail)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(resp).Encode(details)

	fmt.Print("successfully got cars detail by colors\n")
}

func getcarwithfeatures(resp http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select car.id ,car.name ,car.code ,car.color,features.id ,features.name from car left join relation on relation.car_id = car.id left join features on features.id = relation.feature_id where features.id is not null and features.name is not null order by car_id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cars := make(map[int]CarFeatures)
	for rows.Next() {
		var carID int
		var carName string
		var carCode string
		var carColor string
		var featureID int
		var featureName string

		if err := rows.Scan(&carID, &carName, &carCode, &carColor, &featureID, &featureName); err != nil {
			log.Fatal(err)
		}
		carFeatures, ok := cars[carID]
		if !ok {
			// Create a new CarResponse object for this car
			carFeatures = CarFeatures{
				Id:    carID,
				Name:  carName,
				Code:  carCode,
				Color: carColor,
			}
		}
		carFeatures.Features = append(carFeatures.Features, Featuredetails{Id: featureID, Name: featureName})
		cars[carID] = carFeatures
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(resp).Encode(cars)

	fmt.Print("successfully got car details along with it's features\n")
}

func searchcar(resp http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	temp := vars["name"]
	//temp := r.URL.Query().Get("param1")
	rows, err := db.Query("select * from car where name=$1 ", temp)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	details := []Cardetails{}
	for rows.Next() {
		var detail Cardetails
		if err := rows.Scan(&detail.Id, &detail.Name, &detail.Code, &detail.Color); err != nil {
			log.Fatal(err)
		}
		details = append(details, detail)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(resp).Encode(details)

	fmt.Print("successfully got car details by name\n")
}
