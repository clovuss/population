package models

import (
	"context"
	"fmt"

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
func (p *PacientDB) GetByUch(snilsdoc []string, params []string) ([]*Pacient, error) {
	pcts := make([]*Pacient, 0)
	if len(snilsdoc) == 1 {
		stmt := `SELECT pid, enp, surname, name, patronymic, birthday,
       gender, snils, city, naspunkt, street, house, korp, kvart FROm main WHERE snilsdoc=$1 order by surname ;`
		rows, err := p.DB.Query(context.Background(), stmt, snilsdoc[0])
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			pct := &Pacient{}
			err := rows.Scan(&pct.Pid, &pct.Enp, &pct.Surname, &pct.Name, &pct.Patronymic, &pct.Birthday, &pct.Gender, &pct.Snils, &pct.City, &pct.NasPunkt, &pct.Street, &pct.House,
				&pct.Korp, &pct.Kvart)
			if err != nil {
				fmt.Println(err)
			}
			pcts = append(pcts, pct)
		}
		if rows.Err() != nil {
			return nil, err
		}
	} else {

		stmt := `SELECT pid, enp, surname, name, patronymic, birthday,
       gender, snils, city, naspunkt, street, house, korp, kvart FROm main WHERE snilsdoc=$1 OR snilsdoc=$2 order by surname;`

		rows, err := p.DB.Query(context.Background(), stmt, snilsdoc[0], snilsdoc[1])
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			pct := &Pacient{}
			err := rows.Scan(&pct.Pid, &pct.Enp, &pct.Surname, &pct.Name, &pct.Patronymic, &pct.Birthday, &pct.Gender, &pct.Snils, &pct.City, &pct.NasPunkt, &pct.Street, &pct.House,
				&pct.Korp, &pct.Kvart)
			if err != nil {
				fmt.Println(err)
			}
			pcts = append(pcts, pct)
		}
		if rows.Err() != nil {
			return nil, err
		}
	}

	return pcts, nil
}
func (p *PacientDB) InnsertAll(prikrep preparedata.PRIKREP) error {
	stmt := `INSERT INTO main 
		 (pid, enp, surname, name, patronymic, birthday, gender, snils, placebirth, 
		rnname, city, naspunkt, street, house, korp, kvart, snilsdoc, prikrepdate, prikreptype) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19);`

	_, err := p.DB.Exec(context.Background(), stmt, prikrep.Pid, prikrep.ENP, prikrep.FAM, prikrep.IM, prikrep.OT, prikrep.BIRTHDAY, prikrep.GENDER,
		prikrep.SNILS, prikrep.PLACEOFBIRTH, prikrep.RNNAME, prikrep.CITY, prikrep.NP, prikrep.UL, prikrep.DOM, prikrep.KOR, prikrep.KV, prikrep.SSD,
		prikrep.LPUDT, prikrep.LPUAUTO)
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
