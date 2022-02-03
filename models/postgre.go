package models

import (
	"context"
	"fmt"
	"github.com/clovuss/population/preparedata"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
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

func (p *PacientDB) GetByUch(params map[string][]string, snilsdoc []string) ([]*Pacient, error) {
	pcts := make([]*Pacient, 0)
	pctemp := &Pacient{}
	rawsqlparams := make([]string, 0, len(params)+8)
	rawsqlparams = append(rawsqlparams, "surname", "name", "patronymic", "gender")
	destField := make([]interface{}, 0, len(params)+8)
	destField = append(destField, &pctemp.Surname, &pctemp.Name, &pctemp.Patronymic, &pctemp.Gender)
	tableparams := " FROM main"

	for k, _ := range params {
		if k != "adress" {
			rawsqlparams = append(rawsqlparams, k)
		} else {
			rawsqlparams = append(rawsqlparams, "city", "naspunkt", "street", "house", "korp", "kvart")
		}
		switch k {
		case "enp":
			destField = append(destField, &pctemp.Enp)
		case "birthday":
			destField = append(destField, &pctemp.Birthday)
		case "snils":
			destField = append(destField, &pctemp.Snils)
		case "prikreptype":
			destField = append(destField, &pctemp.PrikAuto)
		case "prikrepdate":
			destField = append(destField, &pctemp.PrikDate)
		case "adress":
			destField = append(destField, &pctemp.City, &pctemp.NasPunkt, &pctemp.Street, &pctemp.House, &pctemp.Korp, &pctemp.Kvart)
		case "uch_zav":
			destField = append(destField, &pctemp.UchZav)
			tableparams += ", uch7"
		}
	}
	//fmt.Println(rawsqlparams)
	//fmt.Println(tableparams)
	queryString := fmt.Sprintf(`SELECT %v`, strings.Join(rawsqlparams, ", "))
	queryString += tableparams
	snilsparams := make([]interface{}, 0, 2)
	snilsargs := " WHERE snilsdoc=$1 OR snilsdoc=$2 order by surname;"
	snilsparams = append(snilsparams, snilsdoc[0])
	if len(snilsdoc) == 2 {
		snilsparams = append(snilsparams, snilsdoc[1])
	} else {
		snilsargs = strings.Replace(snilsargs, "OR snilsdoc=$2", "", 1)
	}
	queryString += snilsargs
	//fmt.Println(74, queryString)
	rows, err := p.DB.Query(context.Background(), queryString, snilsparams...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(destField...)
		if err != nil {
			fmt.Println(err)
		}
		pct := &Pacient{}
		*pct = *pctemp
		pcts = append(pcts, pct)

	}
	if rows.Err() != nil {
		fmt.Println(err)
		return nil, err
	}
	return nil, nil
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
func (p *PacientDB) InsertUch(csvslice []string) error {
	stmt := `INSERT INTO uch1
		 (surname, name, patronymic, bday,  enp, uch_zav, tel1, tel2) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8);`

	_, err := p.DB.Exec(context.Background(), stmt, csvslice[0], csvslice[1], csvslice[2], csvslice[3], csvslice[4], csvslice[5], csvslice[6], csvslice[7])

	if err != nil {
		return err
	}
	return nil

}
func (p *PacientDB) InsertPromed(csvslice []string) error {
	stmt := `INSERT INTO promed
		 (card, surname, name, patronymic, birthday, address, live_address, enp) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8);`

	_, err := p.DB.Exec(context.Background(), stmt, csvslice[0], csvslice[1], csvslice[2], csvslice[3], csvslice[4], csvslice[5], csvslice[6], csvslice[7])

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
