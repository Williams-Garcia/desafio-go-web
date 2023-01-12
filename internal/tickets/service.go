package tickets

import "context"

type Service interface {
	GetTotalTickets(ctx context.Context, country string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repo Repository
}

// AverageDestination implements Service
func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0.0, err
	}
	total, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0.0, err
	}

	avg := float64(len(tickets)) / float64(len(total))
	return avg, nil
}

// GetTotalTickets implements Service
func (s *service) GetTotalTickets(ctx context.Context, country string) (int, error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, country)
	if err != nil {
		return 0, err
	}

	return len(tickets), nil
}

func NewService(repo *Repository) Service {
	return &service{repo: *repo}
}
