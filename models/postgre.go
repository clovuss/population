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
	rawsqlFields := make([]string, 0, len(params)+8)
	rawsqlFields = append(rawsqlFields, "main.surname", "main.name", "main.patronymic", "main.gender") //Поля в БД
	destField := make([]interface{}, 0, len(params)+8)
	destField = append(destField, &pctemp.Surname, &pctemp.Name, &pctemp.Patronymic, &pctemp.Gender)
	tables := " FROM main"
	fmt.Println(params)
	//пересмотреть логику добавления джоина, он видимо вставляется много раз
	for key, _ := range params {
		if key == "uch_zav" || key == "phone" || key == "card_num" || key == "live_adress" {
			tables += " JOIN promed ON main.surname=promed.surname AND main.name=promed.name AND main.patronymic=promed.patronymic AND main.birthday=promed.birthday"
		}
	}
	for k, _ := range params {
		switch k {
		case "enp":
			destField = append(destField, &pctemp.Enp)
			rawsqlFields = append(rawsqlFields, "main."+k)
		case "birthday":
			destField = append(destField, &pctemp.Birthday)
			rawsqlFields = append(rawsqlFields, "main."+k)
		case "snils":
			destField = append(destField, &pctemp.Snils)
			rawsqlFields = append(rawsqlFields, "main."+k)
		case "prikreptype":
			destField = append(destField, &pctemp.PrikAuto)
			rawsqlFields = append(rawsqlFields, "main."+k)
		case "prikrepdate":
			destField = append(destField, &pctemp.PrikDate)
			rawsqlFields = append(rawsqlFields, "main."+k)
		case "adress":
			destField = append(destField, &pctemp.City, &pctemp.NasPunkt, &pctemp.Street, &pctemp.House, &pctemp.Korp, &pctemp.Kvart)
			rawsqlFields = append(rawsqlFields, "main.city", "main.naspunkt", "main.street", "main.house", "main.korp", "main.kvart")
		case "uch_zav":
			rawsqlFields = append(rawsqlFields, "promed."+k)
			destField = append(destField, &pctemp.UchZav)
		case "phone":
			rawsqlFields = append(rawsqlFields, "promed."+k)
			destField = append(destField, &pctemp.Phone)
		case "live_adress":
			rawsqlFields = append(rawsqlFields, "promed."+k)
			destField = append(destField, &pctemp.LiveAdress)
		case "card_num":
			rawsqlFields = append(rawsqlFields, "promed."+k)
			destField = append(destField, &pctemp.CardNum)
		case "document":
			destField = append(destField, &pctemp.DocType, &pctemp.DocSeries, &pctemp.DocNumber, &pctemp.DocDate, &pctemp.Docorg)
			rawsqlFields = append(rawsqlFields, "main.doctype", "main.docseries", "main.docnumber", "main.docdate", "main.docorg")
		}
	}
	queryFomSelectToFrom := fmt.Sprintf(`SELECT %v`, strings.Join(rawsqlFields, ", "))
	whereCondition := " WHERE main.snilsdoc=$1"
	snilsparams := make([]interface{}, 0, 2)
	snilsparams = append(snilsparams, snilsdoc[0])
	if len(snilsdoc) == 2 {
		snilsparams = append(snilsparams, snilsdoc[1])
		whereCondition += " OR main.snilsdoc=$2"
	}
	orderBy := " ORDER BY main.surname;"
	query := queryFomSelectToFrom + tables + whereCondition + orderBy
	fmt.Println(query)
	rows, err := p.DB.Query(context.Background(), query, snilsparams...)
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
