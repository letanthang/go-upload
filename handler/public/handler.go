package handler_public

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"

	"cloud.google.com/go/storage"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world! I love you so much. :3:3:3:4")
}

func Upload(c echo.Context) error {
	r := c.Request()
	message := r.URL.Path
	fmt.Printf("We have a request at url: %s\n", message)

	var data map[string]interface{}

	f, fh, err := r.FormFile("file")
	if err != nil {
		log.Error(err)

		data["err_code"] = http.StatusBadRequest
		data["err_msg"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	defer f.Close()

	if StorageBucket == nil {
		err = errors.New("storage bucket is missing - check config.go")
		log.Error(err)
		data["err_code"] = http.StatusInternalServerError
		data["err_msg"] = err.Error()
		return c.JSON(http.StatusInternalServerError, data)
	}

	// random filename, retaining existing extension.
	name := uuid.Must(uuid.NewV4()).String() + path.Ext(fh.Filename)

	ctx := context.Background()
	writer := StorageBucket.Object(name).NewWriter(ctx)

	// Warning: storage.AllUsers gives public read access to anyone.
	writer.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	writer.ContentType = fh.Header.Get("Content-Type")

	// Entries are immutable, be aggressive about caching (1 day).
	writer.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(writer, f); err != nil {
		log.Error(err)
		data["err_code"] = http.StatusInternalServerError
		data["err_msg"] = err.Error()
		return c.JSON(http.StatusInternalServerError, data)
	}
	if err := writer.Close(); err != nil {
		log.Error(err)
		data["err_code"] = http.StatusInternalServerError
		data["err_msg"] = err.Error()
		return c.JSON(http.StatusInternalServerError, data)
	}

	const publicURL = "https://storage.googleapis.com/%s/%s"
	imageUrl := fmt.Sprintf(publicURL, StorageBucketName, name)

	msg := imageUrl
	return c.String(http.StatusOK, msg)
}
