package controllers

import (
	"encoding/json"
	"lvl2/develop/dev11/internal/app/models"

	"log"
	"net/http"
)

const success = "success"

type Controller struct {
	service Service
}

type Service interface {
	AddEvent(event models.Event) error
	UpdateEvent(oldEventName string, event models.Event) error
	DeleteEvent(event models.Event)
	GetEventsForDay(date string) ([]models.Event, error)
	GetEventsForWeek(date string) ([]models.Event, error)
	GetEventsForMonth() []models.Event
}

func NewService(service Service) *Controller {
	controller := Controller{
		service: service,
	}

	return &controller
}

func (controller *Controller) CreateEvent(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var event models.Event
	err := decoder.Decode(&event)
	if err != nil {
		http.Error(writer, "can't read body", http.StatusBadRequest)
		log.Fatal(err)
	}

	err = controller.service.AddEvent(event)
	if err != nil {
		result := models.Result{Result: err.Error()}
		resultJsonBytes, _ := json.Marshal(result)
		_, err = writer.Write(resultJsonBytes)
		writer.WriteHeader(http.StatusBadRequest)

	}

	result := models.Result{Result: success}
	resultJsonBytes, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.Write(resultJsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	writer.WriteHeader(http.StatusOK)
}

func (controller *Controller) UpdateEvent(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var event models.Event
	err := decoder.Decode(&event)
	if err != nil {
		http.Error(writer, "can't read body", http.StatusBadRequest)
		log.Fatal(err)
	}
	var name models.Name
	err = decoder.Decode(&name)
	if err != nil {
		http.Error(writer, "can't read body", http.StatusBadRequest)
		log.Fatal(err)
	}

	err = controller.service.UpdateEvent(name.OldNameEvent, event)
	if err != nil {
		result := models.Result{Result: err.Error()}
		resultJsonBytes, _ := json.Marshal(result)
		_, err = writer.Write(resultJsonBytes)
		writer.WriteHeader(http.StatusBadRequest)
	}

	result := models.Result{Result: success}
	resultJsonBytes, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.Write(resultJsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	writer.WriteHeader(http.StatusOK)
}

func (controller *Controller) DeleteEvent(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var event models.Event
	err := decoder.Decode(&event)
	if err != nil {
		http.Error(writer, "can't read body", http.StatusBadRequest)
		log.Fatal(err)
	}

	controller.service.DeleteEvent(event)

	result := models.Result{Result: success}
	resultJsonBytes, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.Write(resultJsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	writer.WriteHeader(http.StatusOK)

}

func (controller *Controller) EventsForDay(writer http.ResponseWriter, request *http.Request) {
	date := request.URL.Query().Get("date")

	events, err := controller.service.GetEventsForDay(date)
	if err != nil {
		result := models.Result{Result: err.Error()}
		resultJsonBytes, _ := json.Marshal(result)
		_, err = writer.Write(resultJsonBytes)
		writer.WriteHeader(http.StatusBadRequest)
	}

	resultJsonBytes, err := json.Marshal(events)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.Write(resultJsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	writer.WriteHeader(http.StatusOK)
}

func (controller *Controller) EventsForWeek(writer http.ResponseWriter, request *http.Request) {

	date := request.URL.Query().Get("date")

	events, err := controller.service.GetEventsForWeek(date)
	if err != nil {
		result := models.Result{Result: err.Error()}
		resultJsonBytes, _ := json.Marshal(result)
		_, err = writer.Write(resultJsonBytes)
		writer.WriteHeader(http.StatusBadRequest)
	}

	resultJsonBytes, err := json.Marshal(events)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.Write(resultJsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	writer.WriteHeader(http.StatusOK)

}

func (controller *Controller) EventsForMonth(writer http.ResponseWriter, request *http.Request) {

	events := controller.service.GetEventsForMonth()

	resultJsonBytes, err := json.Marshal(events)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.Write(resultJsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	writer.WriteHeader(http.StatusOK)
}
