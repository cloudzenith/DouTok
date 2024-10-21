package collectionservice

import (
	"context"
	"errors"
	"github.com/TremblingV5/box/dbtx"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/collectionserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity/collection"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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

func (s *Service) AddVideo2Collection(ctx context.Context, userId, collectionId, videoId int64) (err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	// 没传collectionId, 检索默认收藏夹
	if collectionId == 0 {
		var coll *model.Collection
		coll, err = s.collection.ListFirstCollection4UserId(ctx, userId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Context(ctx).Errorf("failed to list default collection: %v", err)
			return err
		}

		collectionId = coll.ID
	}

	existedCollection, err := s.collection.GetByIdTx(ctx, collectionId)
	if err != nil {
		log.Context(ctx).Errorf("GetCollectionById error: %v", err)
		return err
	}

	existedRelation, err := s.collection.GetCollectionVideo(ctx, collectionId, videoId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Context(ctx).Errorf("GetCollectionVideo error: %v", err)
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = s.collection.AddVideo2Collection(ctx, existedCollection.UserID, collectionId, videoId)
		if err != nil {
			log.Context(ctx).Errorf("AddVideo2Collection error: %v", err)
			return err
		}

		return nil
	}

	existedRelation.IsDeleted = false
	return s.collection.UpdateCollectionVideoTx(ctx, existedRelation)
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

func (s *Service) RemoveVideo2Collection(ctx context.Context, userId, collectionId, videoId int64) (err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	// 没传collectionId, 检索默认收藏夹
	if collectionId == 0 {
		coll, err := s.collection.ListFirstCollection4UserId(ctx, userId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Context(ctx).Errorf("failed to list default collection: %v", err)
			return err
		}

		collectionId = coll.ID
	}

	return s.collection.RemoveVideoFromCollection(ctx, collectionId, videoId)
}

func (s *Service) ListCollectedVideoByGiven(ctx context.Context, userId int64, videoIdList []int64) ([]int64, error) {
	return s.collection.ListCollectedVideoByGiven(ctx, userId, videoIdList)
}

func (s *Service) GenerateDefaultCollection(ctx context.Context, userId int64) error {
	// TODO 上锁
	collections, err := s.ListCollection(ctx, userId, 1, 0)
	if err != nil {
		log.Context(ctx).Errorf("failed to check existed collections: %v", err)
		return err
	}

	if collections.Count > 0 {
		return nil
	}

	return s.CreateCollection(ctx, userId, "默认收藏夹", "默认收藏夹")
}

func (s *Service) CountCollectedNumber4Video(ctx context.Context, videoId []int64) ([]*collectionserviceiface.CountResult, error) {
	return s.collection.CountByVideoIdList(ctx, videoId)
}

var _ collectionserviceiface.CollectionService = (*Service)(nil)
