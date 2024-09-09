package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func uploadCover() {
	// 替换为你的视频文件路径
	filePath := "./test/video_upload/cover.png"

	// 获取文件信息
	fileSize, err := getFileSize(filePath)
	if err != nil {
		log.Fatalf("Error getting file size: %v", err)
	}

	fileHash, err := getFileHash(filePath)
	if err != nil {
		log.Fatalf("Error getting file hash: %v", err)
	}

	fileType, err := getFileType(filePath)
	if err != nil {
		log.Fatalf("Error getting file type: %v", err)
	}

	filename := filepath.Base(filePath)

	// 构建请求参数
	requestData := PreSign4UploadRequest{
		Hash:     fileHash,
		FileType: fileType,
		Size:     fileSize,
		Filename: filename,
	}

	// 将请求参数序列化为 JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalf("Error marshalling request body: %v", err)
	}

	fmt.Printf("Request body: %s\n", requestBody)

	// 构建 HTTP 请求
	url := "http://127.0.0.1:22000/cover/upload" // 替换为你的服务器地址
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxODI5NzgyMjk1NDk0NTMzMTIwfQ.4OxzdoYsgIu6AMvwi_6iEc_CAl--Y2alRDg0vSoLQfE")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected response status code: %d, resp body: %v", resp.StatusCode, resp.Body)
	}

	fmt.Printf("resp body: %v\n", resp.Body)
	// 读取响应体
	var response PreSign4UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}

	// 输出获取到的上传 URL
	fmt.Printf("Upload URL: %s, FileId: %s\n", response.URL, response.FileId)

	// 接下来，你可以使用获取到的 UploadURL 上传文件
	file, err := os.Open(filePath) // 替换为你的视频文件路径
	if err != nil {
		log.Fatalf("Failed to open video file: %v", err)
	}
	defer file.Close()

	putReq, err := http.NewRequest("PUT", response.URL, file)
	putReq.ContentLength = fileSize
	if err != nil {
		log.Fatalf("Failed to create put request: %v", err)
	}
	putReq.Header.Set("Content-Type", "application/octet-stream")

	putResp, err := client.Do(putReq)
	if err != nil {
		log.Fatalf("Failed to upload cover: %v", err)
	}
	defer putResp.Body.Close()

	if putResp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to upload cover: %d", putResp.StatusCode)
	}

	fmt.Println("cover uploaded successfully")
}
