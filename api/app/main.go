package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	Persons []Person
)

// entity
type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Gener string `json:"gender"`
	Age   int    `json:"age"`
}

type RespondOK struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

// repository
func SavePersonRepo(person *Person) (*Person, error) {
	Persons = append(Persons, *person)
	return person, nil
}

func GetPersonsRepo() ([]Person, error) {
	return Persons, nil
}

// usecase
func SavePersonUsecase(person *Person) (*Person, error) {
	data, err := SavePersonRepo(person)
	return data, err
}

func GetPersonUsecase() ([]Person, error) {
	data, err := GetPersonsRepo()
	return data, err
}

func RespondWithJSON(w http.ResponseWriter, result interface{}, code int) {
	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, string(response))
}

// delivery
func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "hello %s", name)
}

func SavePerson(w http.ResponseWriter, r *http.Request) {
	functionName := "SavePerson"
	if r.Method == http.MethodPost {
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := SavePersonRepo(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		success := RespondOK{
			Data:    res,
			Message: http.StatusText(http.StatusCreated),
			Status:  http.StatusCreated,
		}

		log.WithFields(log.Fields{"status": success.Status, "method": http.MethodPost, "name": functionName}).Info()
		RespondWithJSON(w, success, success.Status)
	}
}

func GetPersons(w http.ResponseWriter, r *http.Request) {
	functionName := "GetPersons"
	if r.Method == http.MethodGet {
		res, err := GetPersonsRepo()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		success := RespondOK{
			Data:    res,
			Message: http.StatusText(http.StatusOK),
			Status:  http.StatusOK,
		}

		log.WithFields(log.Fields{"status": success.Status, "method": http.MethodPost, "name": functionName}).Info()
		RespondWithJSON(w, success, success.Status)
	}
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/save", SavePerson)
	http.HandleFunc("/getperson", GetPersons)

	fmt.Println("api start on http://loalhost:8081/")
	http.ListenAndServe(":8081", nil)
}
