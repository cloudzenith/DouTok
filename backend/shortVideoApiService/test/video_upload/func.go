package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

// PreSign4UploadVideoRequest 定义了预注册上传视频请求的参数
type PreSign4UploadVideoRequest struct {
	Hash     string `json:"hash"`
	FileType string `json:"file_type"`
	Size     int64  `json:"size"`
	Filename string `json:"filename"`
}

// PreSign4UploadVideoResponse 定义了预注册上传视频响应的结构
type PreSign4UploadVideoResponse struct {
	URL    string `json:"url"`
	FileId string `json:"fileId"`
}

// PreSign4UploadRequest 封面上传请求
type PreSign4UploadRequest struct {
	Hash     string `json:"hash"`
	FileType string `json:"file_type"`
	Size     int64  `json:"size"`
	Filename string `json:"filename"`
}

// PreSign4UploadResponse 封面上传响应
type PreSign4UploadResponse struct {
	URL    string `json:"url"`
	FileId string `json:"fileId"`
}

// getFileSize 获取文件大小
func getFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

// getFileHash 计算文件的 MD5 哈希值
func getFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// getFileType 获取文件的 MIME 类型
func getFileType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	fileType := mime.TypeByExtension(filepath.Ext(filePath))
	if fileType == "" {
		fileType = http.DetectContentType(buffer)
	}
	return fileType, nil
}
