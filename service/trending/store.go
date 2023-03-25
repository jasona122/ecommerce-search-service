package trending

import (
	"context"

	"github.com/jasona122/ecommerce-search-service/config"
	"github.com/jasona122/ecommerce-search-service/contracts"
	"github.com/jmoiron/sqlx"
)

const (
	resultsLimit    = 5
	queryCountLimit = 2147483647

	getTrendingDBQuery       = "SELECT id, query, query_count, service_area_id FROM trending WHERE service_area_id = $1 ORDER BY query_count DESC LIMIT $2"
	addTrendingDBQuery       = "INSERT INTO trending (query, query_count, service_area_id) VALUES ($1, $2, $3) ON CONFLICT (query, service_area_id) DO UPDATE SET query_count = $4"
	deleteTrendingDBQuery    = "DELETE FROM trending WHERE query = $1 AND service_area_id = $2"
	incrementTrendingDBQuery = "INSERT INTO trending (query, query_count, service_area_id) VALUES ($1, 1, $2) ON CONFLICT (query, service_area_id) DO UPDATE SET query_count = LEAST(trending.query_count + 1, $3)"
)

type Store interface {
	GetTopTrendingQueries(ctx context.Context, serviceAreaID string) ([]contracts.TrendingSchemaDB, error)
	AddTrendingQuery(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error)
	DeleteTrendingQuery(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error)
	IncrementQueryCount(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error)
}

type store struct {
	db     *sqlx.DB
	config config.DatabaseConfig
}

func NewStore(db *sqlx.DB, config config.DatabaseConfig) Store {
	return &store{
		db:     db,
		config: config,
	}
}

func (s store) GetTopTrendingQueries(ctx context.Context, serviceAreaID string) ([]contracts.TrendingSchemaDB, error) {
	ctx, cancel := context.WithTimeout(ctx, s.config.Timeout)
	defer cancel()

	var dbResults []contracts.TrendingSchemaDB

	err := s.db.Select(&dbResults, getTrendingDBQuery, serviceAreaID, resultsLimit)

	if err != nil {
		return []contracts.TrendingSchemaDB{}, err
	}

	return dbResults, nil
}

func (s store) AddTrendingQuery(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.config.Timeout)
	defer cancel()

	_, err := s.db.Exec(addTrendingDBQuery, query, queryCountLimit, serviceAreaID, queryCountLimit)

	if err != nil {
		return contracts.EditTrendingServiceResponse{}, err
	}

	return contracts.DefaultEditSuccessResponse(query, serviceAreaID), nil
}

func (s store) DeleteTrendingQuery(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.config.Timeout)
	defer cancel()

	_, err := s.db.Exec(deleteTrendingDBQuery, query, serviceAreaID)

	if err != nil {
		return contracts.EditTrendingServiceResponse{}, err
	}

	return contracts.DefaultEditSuccessResponse(query, serviceAreaID), nil
}

func (s store) IncrementQueryCount(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.config.Timeout)
	defer cancel()

	_, err := s.db.Exec(incrementTrendingDBQuery, query, serviceAreaID, queryCountLimit)

	if err != nil {
		return contracts.EditTrendingServiceResponse{}, err
	}

	return contracts.DefaultEditSuccessResponse(query, serviceAreaID), nil
}
