package services

import "lvl2/develop/dev11/internal/app/models"

type Service struct {
	repo Repository
}

type Repository interface {
	AddEvent(event models.Event) error
	UpdateEvent(oldEventName string, event models.Event) error
	DeleteEvent(event models.Event)
	GetEventsForDay(date string) ([]models.Event, error)
	GetEventsForWeek(date string) ([]models.Event, error)
	GetEventsForMonth() []models.Event
}

func NewStore(repo Repository) *Service {
	service := Service{
		repo: repo,
	}

	return &service
}

func (service *Service) AddEvent(event models.Event) error {
	err := service.repo.AddEvent(event)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) UpdateEvent(oldEventName string, event models.Event) error {
	err := service.repo.UpdateEvent(oldEventName, event)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) DeleteEvent(event models.Event) {
	service.repo.DeleteEvent(event)
}

func (service *Service) GetEventsForDay(date string) ([]models.Event, error) {
	events, err := service.repo.GetEventsForDay(date)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (service *Service) GetEventsForWeek(date string) ([]models.Event, error) {
	events, err := service.repo.GetEventsForWeek(date)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (service *Service) GetEventsForMonth() []models.Event {
	events := service.repo.GetEventsForMonth()

	return events
}
