package models

import (
	"context"

	"github.com/clovuss/population/preparedata"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PacientDB struct {
	DB *pgxpool.Pool
}

func (p *PacientDB) GetByPid(pid string) (*Pacient, error) {
	stmt := "SELECT pid, enp, surname, name, patronymic, birthday, gender, snils, city, naspunkt, street, house, korp, kvart FROm main WHERE pid=$1;"
	pct := &Pacient{}
	row := p.DB.QueryRow(context.Background(), stmt, pid)
	err := row.Scan(&pct.Pid, &pct.Enp, &pct.Surname, &pct.Name, &pct.Patronymic, &pct.Birthday, &pct.Gender, &pct.Snils, &pct.City, &pct.NasPunkt, &pct.Street, &pct.House,
		&pct.Korp, &pct.Kvart)
	if err != nil {
		return nil, err
	}
	return pct, nil
}

func (p *PacientDB) InnsertAll(prikrep preparedata.PRIKREP) error {
	stmt := "INSERT INTO main" +
		" (pid, enp, surname, name, patronymic, birthday, gender, snils, placebirth, " +
		"rnname, city, naspunkt, street, house, korp, kvart) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16 );"

	_, err := p.DB.Exec(context.Background(), stmt, prikrep.Pid, prikrep.ENP, prikrep.FAM, prikrep.IM, prikrep.OT, prikrep.BIRTHDAY, prikrep.GENDER,
		prikrep.SNILS, prikrep.PLACEOFBIRTH, prikrep.RNNAME, prikrep.CITY, prikrep.NP, prikrep.UL, prikrep.DOM, prikrep.KOR, prikrep.KV)
	if err != nil {
		return err
	}
	return nil

}

func (p *PacientDB) InnsertOne(prikrep interface{}) error {
	stmt := "INSERT INTO main (pid) VALUES($1);"
	_, err := p.DB.Exec(context.Background(), stmt, prikrep)
	if err != nil {
		return err
	}
	return nil

}
