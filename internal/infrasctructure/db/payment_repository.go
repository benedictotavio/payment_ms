package db

import "github.com/benedictotavio/payment_ms/internal/domain"

type PaymentRepository struct {
	conn *Psql
}

// TODO: adicionar variaveis de ambiente
func NewPaymentRepository() (*PaymentRepository, error) {
	conn, err := NewDB(
		ConfigPsql{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			DBName:   "payment",
		},
	)

	if err != nil {
		return nil, err
	}

	return &PaymentRepository{
		conn: conn,
	}, nil
}

func (p* PaymentRepository) Create(payment domain.Payment) error {
	
	_, err  := p.conn.DB.Exec("INSERT INTO payments (order_id) VALUES ($1)", payment.ID)

	if err != nil {
		return err
	}

	defer p.conn.Close()

	return nil
}

func (p* PaymentRepository) Get(id int) (domain.Payment, error) {
	var payment domain.Payment
	row := p.conn.DB.QueryRow("SELECT * FROM payments WHERE id = $1", id)
	if err := row.Scan(&payment.ID, &payment.Amount, &payment.Method); err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}