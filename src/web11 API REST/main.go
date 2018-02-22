package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// API REST que devolvera formato JSON y que será completamente independiente el backend con el frontend

// Note Modelo para la BD (que en realidad será un Map)
type Note struct {
	Title       string    `json:"title"` // Notacion para json: Cuando se codifique a json poga title en vez de Title
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

var noteStore = make(map[string]Note) // Simulamos una BD
var id int                            // Simula indice de la BD

func main() {
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening #8080..")
	log.Fatal(server.ListenAndServe())
}

// GetNoteHandler Handler para el GET ../api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note

	for _, value := range noteStore {
		notes = append(notes, value)
	}

	j, err := json.Marshal(notes) // Nos formatea a JSON

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") // Headers de HTTP
	w.WriteHeader(http.StatusOK)                       // 200
	w.Write(j)
}

// PostNoteHandler Handler para el POST
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note) // Decodifica un JSON y lo rellena en la estructura Note

	if err != nil {
		log.Println(err)
	}

	note.CreatedAt = time.Now()

	id++
	idString := strconv.Itoa(id)
	noteStore[idString] = note // Añadimos al map (BD)

	j, err := json.Marshal(note) // Nos formatea a JSON
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") // Headers de HTTP
	w.WriteHeader(http.StatusCreated)                  // 201
	w.Write(j)
}

// PutNoteHandler Handler para el PUT
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)   // Nos devuelve un map con las variables de la peticion en nuestro caso el "id" --> /api/notes/{id}
	idLeido := vars["id"] // Si hubiera mas nos quedamos solo con el "id"

	var noteUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		log.Println(err)
	}

	if note, ok := noteStore[idLeido]; ok { // Si existe el ID
		noteUpdate.CreatedAt = note.CreatedAt // A los datos que nos mando el usuario le agregamos la fecha de cuando se creó
		delete(noteStore, idLeido)            // Borramos la nota antigua
		noteStore[idLeido] = noteUpdate
	} else {
		log.Println("No se ha encontrado el ID: %s", idLeido)
	}

	w.WriteHeader(http.StatusNoContent) // 204
}

// DeleteNoteHandler Handler para el DELETE
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)   // Nos devuelve un map con las variables de la peticion en nuestro caso el "id" --> /api/notes/{id}
	idLeido := vars["id"] // Si hubiera mas nos quedamos solo con el "id"

	if _, ok := noteStore[idLeido]; ok { // Si existe el ID
		delete(noteStore, idLeido) // Borramos la nota
	} else {
		log.Println("No se ha encontrado el ID: %s", idLeido)
	}

	w.WriteHeader(http.StatusNoContent) // 204
}
