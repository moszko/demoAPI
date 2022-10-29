package infrastructure

import (
	"bufio"
	"database/sql"
	"os"
	"vlserver/domain"

	_ "github.com/lib/pq"
)

const (
	Host   = "db"
	Port   = 5432
	User   = "postgres"
	Dbname = "postgres"
	maxD   = 3
)

var Password string = getPassword()

func getPassword() string {
	file, err := os.Open("/usr/src/app/pg_pass")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return scanner.Text()
}

type trademarkRepository struct {
	db *sql.DB
}

func NewTrademarkRepository(db *sql.DB) domain.TrademarkRepository {
	return &trademarkRepository{db: db}
}

func (p *trademarkRepository) GetByName(input string) (*domain.Trademark, error) {
	sqlStmt := `SELECT * FROM trademarks WHERE name=$1`
	row := p.db.QueryRow(sqlStmt, input)
	trademark := &domain.Trademark{}
	err := row.Scan(&trademark.Id, &trademark.Name, &trademark.StatusCode, &trademark.StatusDate)
	if err != nil {
		return nil, err
	}

	return trademark, nil
}

func (p *trademarkRepository) GetSimilarByName(input string) ([]string, error) {
	sqlStmt := `SELECT name FROM trademarks WHERE levenshtein_less_equal(lower(name), lower($1), $2) <= $2`
	rows, err := p.db.Query(sqlStmt, input, maxD)
	if err != nil {
		return nil, err
	}
	trademarkNames := []string{}
	tempTrademarkName := ""
	for rows.Next() {
		err := rows.Scan(&tempTrademarkName)
		if err != nil {
			break
		}
		trademarkNames = append(trademarkNames, tempTrademarkName)
	}

	return trademarkNames, nil
}
