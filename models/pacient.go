package models

import "time"

type Pacient struct {
	Pid          string
	Enp          string
	Surname      string
	Name         string
	Patronymic   string
	Birthday     time.Time
	Gender       string
	Snils        string
	Placeofbirth string
	RegionName   string
	City         string
	NasPunkt     string
	Street       string
	House        string
	Korp         string
	Kvart        string
	PrikAuto     string
	PrikDate     time.Time
	SnilsDoc     string
}
