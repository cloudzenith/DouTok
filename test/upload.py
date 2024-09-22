import requests

minio_url = "http://localhost:9000/shortvideo/short_video/1824118603822141440?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=root%2F20240815%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240815T161900Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=beb98bf4fb70ddaf6f8475790456e980d796399e7d08bcb8feef70966615c4bb"

# 向minio上传文件
def upload_file(file_path):
    with open(file_path, 'rb') as file_data:
        response = requests.put(
            minio_url,
            data=file_data,
            headers={"Content-Type": "application/octet-stream"}
        )
        print(response)
        return response.status_code

# upload_file("D:\Share/data.pdf")
upload_file("./total.pdf")

