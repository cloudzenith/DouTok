package warmup

import (
	"github.com/go-kratos/kratos/v2/log"
	"os/exec"
)

func InitMinioPublicDirectory() {
	entrypoint := "http://minio:9000"
	username := "root"
	password := "rootroot"
	bucketNameList := []string{"shortvideo"}
	if output, err := exec.Command("./mc", "config", "host", "add", "minio", entrypoint, username, password).CombinedOutput(); err != nil {
		log.Errorf("mc config host add err: %v, output: %s", err, output)
		panic(err)
	}
	for _, bucketName := range bucketNameList {
		if output, err := exec.Command("./mc", "anonymous", "set", "public", "minio/"+bucketName+"/public").CombinedOutput(); err != nil {
			log.Errorf("mc anonymous set err: %v, output: %s", err, output)
			panic(err)
		}
	}
}
