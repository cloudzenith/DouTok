package file

import (
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"time"
)

type File struct {
	ID         int64
	DomainName string
	BizName    string
	Hash       string
	FileType   string
	FileSize   int64
	CreateTime time.Time
	UpdateTime time.Time
}

func NewFile(options ...Option) *File {
	f := &File{}
	for _, option := range options {
		option(f)
	}

	return f
}

func NewWithModel(m *models.File) *File {
	return &File{
		ID:         m.ID,
		DomainName: m.DomainName,
		BizName:    m.BizName,
		Hash:       m.Hash,
		FileType:   m.FileType,
		FileSize:   m.FileSize,
		CreateTime: m.CreateTime,
		UpdateTime: m.UpdateTime,
	}
}

func NewWithFileContext(fc *api.FileContext) *File {
	return &File{
		ID:         fc.FileId,
		DomainName: fc.Domain,
		BizName:    fc.BizName,
		Hash:       fc.Hash,
		FileType:   fc.FileType,
		FileSize:   fc.Size,
	}
}

func (f *File) ToModel() *models.File {
	return &models.File{
		ID:         f.ID,
		DomainName: f.DomainName,
		BizName:    f.BizName,
		Hash:       f.Hash,
		FileType:   f.FileType,
		FileSize:   f.FileSize,
		CreateTime: f.CreateTime,
		UpdateTime: f.UpdateTime,
	}
}

func (f *File) SetId() {
	f.ID = utils.GetSnowflakeId()
}

func (f *File) GetObjectName() string {
	return fmt.Sprintf("%s/%d", f.BizName, f.ID)
}
