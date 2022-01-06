package models

import "time"

type Pacient struct {
	Surname      string
	Name         string
	Patronymic   string
	Gender       string
	Enp          string
	Birthday     time.Time
	Snils        string
	PrikAuto     string
	PrikDate     time.Time
	City         string
	NasPunkt     string
	Street       string
	House        string
	Korp         string
	Kvart        string
	SnilsDoc     string
	Pid          string
	Placeofbirth string
	RegionName   string
	UchZav       string
}
