package internal

import (
	"database/sql"
	"net/url"
	"time"
)

type TransactionModel struct {
	DB *sql.DB
}

func (t *TransactionModel) All() ([]Transaction, error) {

	stmt := `SELECT id, txn_date, who, description, payee, amount, category 
			 FROM transactions 
			 ORDER BY id DESC
			`

	rows, err := t.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	transactions := []Transaction{}

	for rows.Next() {
		p := Transaction{}
		err := rows.Scan(&p.Id, &p.Txn_date, &p.Who, &p.Description, &p.Payee, &p.Amount, &p.Category)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *TransactionModel) ByDate(startDate, endDate time.Time) ([]Transaction, error) {

	stmt := `SELECT id, txn_date, who, description, payee, amount, category 
			 FROM transactions 
			 WHERE txn_date BETWEEN ? AND ?
			 ORDER BY id DESC
			`

	rows, err := t.DB.Query(stmt, startDate, endDate)
	if err != nil {
		return nil, err
	}

	transactions := []Transaction{}

	for rows.Next() {
		p := Transaction{}
		err := rows.Scan(&p.Id, &p.Txn_date, &p.Who, &p.Description, &p.Payee, &p.Amount, &p.Category)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *TransactionModel) Insert(data url.Values) error {

	stmt := `INSERT INTO transactions (txn_date, who, description, payee, amount, category,  inserted_at, updated_at)
    		 VALUES (?, ?, ?, ?, ?, ?, datetime('now'), datetime('now'))
  			`
	_, err := t.DB.Exec(stmt,
		data.Get("txn_date"),
		data.Get("who"),
		data.Get("description"),
		data.Get("payee"),
		data.Get("amount"),
		data.Get("category"),
	)

	return err
}
