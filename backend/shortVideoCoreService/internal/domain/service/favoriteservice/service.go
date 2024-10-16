package favoriteservice

import (
	"context"
	"fmt"
	"github.com/TremblingV5/box/dbtx"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/favoriteserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils/pageresult"
	"github.com/go-kratos/kratos/v2/log"
)

type Service struct {
	favorite repoiface.FavoriteRepository
}

func New(favorite repoiface.FavoriteRepository) *Service {
	return &Service{
		favorite: favorite,
	}
}

func (s *Service) AddFavorite(ctx context.Context, dto *favoriteserviceiface.WriteOpDTO) (err error) {
	if err := dto.Check(); err != nil {
		log.Context(ctx).Fatalf("invalid dto: %v, err: %v", dto, err)
		return err
	}

	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	err = s.favorite.AddFavorite(ctx, dto.UserId, dto.TargetId, int32(dto.TargetType), int32(dto.FavoriteType))
	if err != nil {
		log.Context(ctx).Fatalf("add favorite failed: %v", err)
		return err
	}

	return nil
}

func (s *Service) RemoveFavorite(ctx context.Context, dto *favoriteserviceiface.WriteOpDTO) error {
	if err := dto.Check(); err != nil {
		log.Context(ctx).Fatalf("invalid dto: %v, err: %v", dto, err)
		return err
	}

	if err := s.favorite.RemoveFavorite(ctx, dto.UserId, dto.TargetId, int32(dto.TargetType), int32(dto.FavoriteType)); err != nil {
		log.Context(ctx).Fatalf("remove favorite failed: %v", err)
		return err
	}

	return nil
}

func (s *Service) ListFavorite(ctx context.Context, dto *favoriteserviceiface.AggOpDTO, limit, offset int) (*pageresult.R[int64], error) {
	if err := dto.Check(); err != nil {
		log.Context(ctx).Fatalf("invalid dto: %v, err: %v", dto, err)
		return nil, err
	}

	idList, err := s.favorite.ListFavorite(ctx, dto.BizId, int32(dto.AggType), int32(dto.FavoriteType), limit, offset)
	if err != nil {
		log.Context(ctx).Errorf("list favorite failed: %v", err)
		return nil, err
	}

	countResult, err := s.favorite.CountFavorite(ctx, []int64{dto.BizId}, int32(dto.AggType), int32(dto.FavoriteType))
	if err != nil {
		log.Context(ctx).Errorf("count favorite failed: %v", err)
		return nil, err
	}

	if len(countResult) == 0 {
		return pageresult.New(idList, 0), nil
	}

	return pageresult.New(idList, countResult[0].Cnt), nil
}

func (s *Service) CountFavorite(ctx context.Context, dto *favoriteserviceiface.AggOpDTO) ([]*v1.CountFavoriteResponseItem, error) {
	if err := dto.Check(); err != nil {
		log.Context(ctx).Fatalf("invalid dto: %v, err: %v", dto, err)
		return nil, err
	}

	countResult, err := s.favorite.CountFavorite(ctx, dto.BizIdList, int32(dto.AggType), int32(dto.FavoriteType))
	if err != nil {
		log.Context(ctx).Fatalf("count favorite failed: %v", err)
		return nil, err
	}

	var res []*v1.CountFavoriteResponseItem
	for _, item := range countResult {
		res = append(res, &v1.CountFavoriteResponseItem{
			BizId: item.Id,
			Count: item.Cnt,
		})
	}

	return res, nil
}

func (s *Service) getIsFavoriteRecordKey(userId, bizId int64) string {
	return fmt.Sprintf("%d_%d", userId, bizId)
}

func (s *Service) IsFavorite(ctx context.Context, dto []*v1.IsFavoriteRequestItem) ([]*v1.IsFavoriteResponseItem, error) {
	var userIds, bizIds []int64
	for _, item := range dto {
		userIds = append(userIds, item.UserId)
		bizIds = append(bizIds, item.BizId)
	}

	data, err := s.favorite.Get4IsFavorite(ctx, userIds, bizIds)
	if err != nil {
		log.Context(ctx).Fatalf("get data for judge is favorite failed: %v", err)
		return nil, err
	}

	resultMap := make(map[string]bool)
	for _, item := range data {
		resultMap[s.getIsFavoriteRecordKey(item.UserID, item.TargetID)] = true
	}

	var result []*v1.IsFavoriteResponseItem
	for _, item := range dto {
		key := s.getIsFavoriteRecordKey(item.UserId, item.BizId)
		respItem := &v1.IsFavoriteResponseItem{}
		respItem.UserId = item.UserId
		respItem.BizId = item.BizId
		if _, ok := resultMap[key]; ok {
			respItem.IsFavorite = true
		} else {
			respItem.IsFavorite = false
		}
		result = append(result, respItem)
	}

	return result, nil
}

var _ favoriteserviceiface.FavoriteService = (*Service)(nil)
