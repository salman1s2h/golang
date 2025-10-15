package main

import (
	"demo/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"math/rand"
)

var (
	ChContact chan *models.Contact
	ChErr     chan error
)

func init() {
	ChContact = make(chan *models.Contact, 10)
	ChErr = make(chan error, 10)
}

func init() {
	println("init is called-2")
}

func init() {
	println("init is called-3")
}

func init() {
	println("init is called-4")
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/health", Health)

	http.HandleFunc("/contact", CreateContact)

	log.Println("Server is up and running on port:", 8084)
	go SaveContacts("contacts.dat")
	go ProcessError()
	if err := http.ListenAndServe("0.0.0.0:8084", nil); err != nil {
		log.Println(err.Error())
		runtime.Goexit()
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("Oops.. something went wrong"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contact := new(models.Contact)
	err = json.Unmarshal(bytes, contact)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("Oops.. something went wrong"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	contact.Id = rand.Intn(10000)
	contact.Lastmodified = time.Now().Unix()
	contact.Status = "active"

	// err = contact.Validate()

	// if err != nil {
	// log.Println(err.Error())
	// w.Write([]byte(err.Error()))
	// w.WriteHeader(http.StatusBadRequest)
	// return
	// }

	// err = SaveContactToFile("contacts.dat", contact)
	// if err != nil {
	// log.Println(err.Error())
	// w.Write([]byte("Oops.. something went wrong"))
	// w.WriteHeader(http.StatusBadRequest)
	// return
	// }
	ChContact <- contact

	w.Write([]byte("contact successfully stored"))
	w.WriteHeader(http.StatusCreated)

}

func SaveContactToFile(filename string, contact *models.Contact) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, err := contact.ToBytes()
	if err != nil {
		return err
	}
	bytes = append(bytes, []byte("\n")...)
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil

}

func SaveContacts(filename string) {
	for contact := range ChContact {
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			ChErr <- err
		}
		defer f.Close()
		bytes, err := contact.ToBytes()
		if err != nil {
			ChErr <- err
		}
		err = contact.Validate()
		if err != nil {
			ChErr <- err
		}
		bytes = append(bytes, []byte("\n")...)
		_, err = f.Write(bytes)
		if err != nil {
			ChErr <- err
		}

		//SaveContactToFile(filename, c)
	}
}

func ProcessError() {
	for e := range ChErr {
		log.Println(e.Error())
		println(e.Error())
	}
}
