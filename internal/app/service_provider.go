package app

import (
	"context"
	"log"

	"github.com/BelyaevEI/microservices_chat/internal/client/postgres"
	"github.com/BelyaevEI/microservices_chat/internal/closer"
	"github.com/BelyaevEI/microservices_chat/internal/config"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgClient       postgres.Client
	txManager      db.TxManager
	noteRepository repository.NoteRepository

	noteService service.NoteService

	noteImpl *note.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {

	if s.grpcConfig == nil {

		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PostgresClient(ctx context.Context) postgres.Client {
	if s.pgClient == nil {
		client, err := postgres.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = client.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(client.Close)

		s.pgClient = client
	}

	return s.pgClient
}
