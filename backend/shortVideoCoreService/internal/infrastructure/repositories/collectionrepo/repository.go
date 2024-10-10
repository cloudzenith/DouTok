package collectionrepo

import (
	"context"
	"github.com/TremblingV5/box/dbtx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/snowflakeutil"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
)

type PersistRepository struct {
}

func New() *PersistRepository {
	return &PersistRepository{}
}

func (p *PersistRepository) Create(ctx context.Context, collection *model.Collection) error {
	return query.Q.WithContext(ctx).Collection.Create(collection)
}

func (p *PersistRepository) GetById(ctx context.Context, id int64) (*model.Collection, error) {
	return query.Q.WithContext(ctx).Collection.Where(query.Q.Collection.ID.Eq(id)).First()
}

func (p *PersistRepository) RemoveById(ctx context.Context, id int64) error {
	_, err := query.Q.WithContext(ctx).
		Collection.
		Where(query.Q.Collection.ID.Eq(id)).
		Update(query.Collection.IsDeleted, true)
	return err
}

func (p *PersistRepository) ListByUserId(ctx context.Context, userId int64, limit, offset int) ([]*model.Collection, error) {
	return query.Q.WithContext(ctx).Collection.Where(query.Q.Collection.UserID.Eq(userId)).Limit(limit).Offset(offset).Find()
}

func (p *PersistRepository) CountByUserId(ctx context.Context, userId int64) (int64, error) {
	return query.Q.WithContext(ctx).Collection.Where(query.Q.Collection.UserID.Eq(userId)).Count()
}

func (p *PersistRepository) Update(ctx context.Context, collection *model.Collection) error {
	_, err := query.Q.WithContext(ctx).
		Collection.
		Where(query.Q.Collection.ID.Eq(collection.ID)).
		Updates(collection)
	return err
}

func (p *PersistRepository) AddVideo2Collection(ctx context.Context, userId, collectionId, videoId int64) error {
	newCollectionVideo := &model.CollectionVideo{
		CollectionID: collectionId,
		VideoID:      videoId,
		UserID:       userId,
		ID:           snowflakeutil.GetSnowflakeId(),
	}
	return query.Q.WithContext(ctx).CollectionVideo.Create(newCollectionVideo)
}

func (p *PersistRepository) RemoveVideoFromCollection(ctx context.Context, collectionId, videoId int64) error {
	_, err := query.Q.WithContext(ctx).CollectionVideo.Where(
		query.Q.CollectionVideo.CollectionID.Eq(collectionId),
		query.Q.CollectionVideo.VideoID.Eq(videoId),
	).Update(query.CollectionVideo.IsDeleted, true)
	return err
}

func (p *PersistRepository) ListCollectionVideo(ctx context.Context, collectionId int64, limit, offset int) ([]*model.CollectionVideo, error) {
	return query.Q.WithContext(ctx).CollectionVideo.Where(query.Q.CollectionVideo.CollectionID.Eq(collectionId)).Limit(limit).Offset(offset).Find()
}

func (p *PersistRepository) CountCollectionVideo(ctx context.Context, collectionId int64) (int64, error) {
	return query.Q.WithContext(ctx).CollectionVideo.Where(query.Q.CollectionVideo.CollectionID.Eq(collectionId)).Count()
}

func (p *PersistRepository) ListCollectedVideoByGiven(ctx context.Context, userId int64, videoIdList []int64) ([]int64, error) {
	data, err := query.Q.WithContext(ctx).CollectionVideo.Select(
		query.Q.CollectionVideo.VideoID,
	).Where(
		query.Q.CollectionVideo.UserID.Eq(userId),
		query.Q.CollectionVideo.VideoID.In(videoIdList...),
	).Find()
	if err != nil {
		return nil, err
	}

	var videoIds []int64
	for _, c := range data {
		videoIds = append(videoIds, c.VideoID)
	}
	return videoIds, nil
}

func (p *PersistRepository) GetCollectionVideo(ctx context.Context, collectionId, videoId int64) (*model.CollectionVideo, error) {
	return dbtx.TxDoGetValue(ctx, func(tx *query.Query) (*model.CollectionVideo, error) {
		return query.Q.WithContext(ctx).CollectionVideo.Where(
			query.Q.CollectionVideo.CollectionID.Eq(collectionId),
			query.Q.CollectionVideo.VideoID.Eq(videoId),
		).First()
	})
}

func (p *PersistRepository) GetByIdTx(ctx context.Context, id int64) (*model.Collection, error) {
	return dbtx.TxDoGetValue(ctx, func(tx *query.QueryTx) (*model.Collection, error) {
		return p.GetById(ctx, id)
	})
}

func (p *PersistRepository) UpdateCollectionVideoTx(ctx context.Context, collectionVideo *model.CollectionVideo) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		_, err := query.Q.WithContext(ctx).CollectionVideo.Updates(collectionVideo)
		return err
	})
}

var _ repoiface.CollectionRepository = (*PersistRepository)(nil)
