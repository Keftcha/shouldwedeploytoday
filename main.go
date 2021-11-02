package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const HOST = "0.0.0.0"
const PORT = 5461

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/{weekday:[0-6]}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		weekday, err := strconv.ParseInt(vars["weekday"], 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		today := GetDay(time.Weekday(weekday))
		jsn, err := json.Marshal(today)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		fmt.Fprint(w, string(jsn))
	}).Methods(http.MethodGet)

	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./images/")))).Methods(http.MethodGet)

	fmt.Printf("Server started on %s:%d\n", HOST, PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", HOST, PORT), r))
}
