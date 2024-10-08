package collectionservice

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/collectionserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity/collection"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type Service struct {
	collection repoiface.CollectionRepository
}

func New(collection repoiface.CollectionRepository) *Service {
	return &Service{
		collection: collection,
	}
}

func (s *Service) CreateCollection(ctx context.Context, userId int64, name, description string) error {
	newCollection := collection.New(
		collection.WithUserId(userId),
		collection.WithTitle(name),
		collection.WithDescription(description),
	)
	newCollection.SetId()
	return s.collection.Create(ctx, newCollection.ToModel())
}

func (s *Service) GetCollectionById(ctx context.Context, collectionId int64) (*collection.Collection, error) {
	c, err := s.collection.GetById(ctx, collectionId)
	if err != nil {
		log.Context(ctx).Errorf("GetCollectionById error: %v", err)
		return nil, err
	}

	return collection.NewWithModel(c), nil
}

func (s *Service) RemoveCollection(ctx context.Context, collectionId int64) error {
	return s.collection.RemoveById(ctx, collectionId)
}

func (s *Service) ListCollection(ctx context.Context, userId int64, limit, offset int) (*collectionserviceiface.ListCollectionResult, error) {
	list, err := s.collection.ListByUserId(ctx, userId, limit, offset)
	if err != nil {
		log.Context(ctx).Errorf("ListCollection error: %v", err)
		return nil, err
	}

	var collections []*collection.Collection
	for _, c := range list {
		collections = append(collections, collection.NewWithModel(c))
	}

	count, err := s.collection.CountByUserId(ctx, userId)
	if err != nil {
		// 弱依赖
		log.Context(ctx).Warnf("failed to count collection: %v", err)
	}

	return &collectionserviceiface.ListCollectionResult{
		Data:  collections,
		Count: count,
	}, nil
}

func (s *Service) UpdateCollection(ctx context.Context, collectionId int64, name, description string) error {
	cModel, err := s.collection.GetById(ctx, collectionId)
	if err != nil {
		log.Context(ctx).Errorf("s.collection.GetById error: %v", err)
		return err
	}

	c := collection.NewWithModel(cModel)
	c.Title = name
	c.Description = description
	return s.collection.Update(ctx, c.ToModel())
}

func (s *Service) AddVideo2Collection(ctx context.Context, collectionId, videoId int64) error {
	return s.collection.AddVideo2Collection(ctx, collectionId, videoId)
}

func (s *Service) ListCollectionVideo(ctx context.Context, collectionId int64, pagination *v1.PaginationRequest) (*collectionserviceiface.ListCollectionVideoResult, error) {
	limit, offset := utils.GetLimitOffset(int(pagination.Page), int(pagination.Size))
	list, err := s.collection.ListCollectionVideo(ctx, collectionId, limit, offset)
	if err != nil {
		log.Context(ctx).Errorf("ListCollectionVideo error: %v", err)
		return nil, err
	}

	var videoIds []int64
	for _, c := range list {
		videoIds = append(videoIds, c.VideoID)
	}

	count, err := s.collection.CountCollectionVideo(ctx, collectionId)
	if err != nil {
		log.Context(ctx).Warnf("failed to count collection video: %v", err)
	}

	return &collectionserviceiface.ListCollectionVideoResult{
		Data:  videoIds,
		Count: count,
	}, nil
}

func (s *Service) RemoveVideo2Collection(ctx context.Context, collectionId, videoId int64) error {
	return s.collection.RemoveVideoFromCollection(ctx, collectionId, videoId)
}

var _ collectionserviceiface.CollectionService = (*Service)(nil)
