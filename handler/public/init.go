package handler_public

import (
	"context"

	"cloud.google.com/go/storage"
	"g.ghn.vn/logistic/crm/types"
	"github.com/labstack/gommon/log"
	"github.com/liip/sheriff"
)

var notFoundErrorMessage = types.PayloadResponse("404", "Không tìm thấy thông tin người dùng")
var optionPublic = &sheriff.Options{
	Groups: []string{"public"},
}
var (
	StorageBucketName                       = "thanglab-bucket"
	StorageBucket     *storage.BucketHandle = configureStorage(StorageBucketName)
)

func configureStorage(bucketID string) *storage.BucketHandle {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Error("Config storage bucket err:", err)
		return nil
	}
	return client.Bucket(bucketID)
}

func init() {

}
