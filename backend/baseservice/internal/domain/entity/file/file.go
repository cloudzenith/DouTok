package file

import (
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"strings"
	"time"
)

type File struct {
	ID         int64
	DomainName string
	BizName    string
	Hash       string
	FileType   string
	FileSize   int64
	Uploaded   bool
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
		Uploaded:   m.Uploaded,
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
		Uploaded:   f.Uploaded,
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

func (f *File) CheckHash(hash string) bool {
	// 将两个哈希值都转换为小写
	fHashLower := strings.ToLower(f.Hash)
	hashLower := strings.ToLower(hash)

	// 比较小写哈希值
	return fHashLower == hashLower
}

func (f *File) SetUploaded() *File {
	f.Uploaded = true
	return f
}
