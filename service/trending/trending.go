package trending

import (
	"context"

	"github.com/jasona122/ecommerce-search-service/config"
	"github.com/jasona122/ecommerce-search-service/contracts"
	"github.com/jasona122/ecommerce-search-service/db"
)

type Service interface {
	GetTopTrendingQueries(ctx context.Context, serviceAreaID string) ([]contracts.GetTrendingServiceResult, error)
	AddTrendingQuery(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error)
	DeleteTrendingQuery(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error)
	IncrementQueryCount(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error)
}

type service struct {
	store Store
}

func NewService(dbConfig config.DatabaseConfig) Service {
	return &service{
		store: NewStore(db.GetDB(), dbConfig),
	}
}

func (s service) GetTopTrendingQueries(ctx context.Context, serviceAreaID string) ([]contracts.GetTrendingServiceResult, error) {
	dbResults, err := s.store.GetTopTrendingQueries(ctx, serviceAreaID)
	if err != nil {
		return []contracts.GetTrendingServiceResult{}, err
	}

	var results []contracts.GetTrendingServiceResult
	for _, dbResult := range dbResults {
		result := contracts.GetTrendingServiceResult{
			Query:      dbResult.Query,
			QueryCount: dbResult.QueryCount,
		}
		results = append(results, result)
	}

	return results, nil
}

func (s service) AddTrendingQuery(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error) {
	return s.store.AddTrendingQuery(ctx, query, serviceAreaID)
}

func (s service) DeleteTrendingQuery(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error) {
	return s.store.DeleteTrendingQuery(ctx, query, serviceAreaID)
}

func (s service) IncrementQueryCount(ctx context.Context, query string, serviceAreaID string) (contracts.EditTrendingServiceResponse, error) {
	return s.store.IncrementQueryCount(ctx, query, serviceAreaID)
}
