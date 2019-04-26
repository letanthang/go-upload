package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"

	"cloud.google.com/go/storage"
	"github.com/gofrs/uuid"
	"github.com/labstack/gommon/log"
)

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

func formHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path

	fmt.Printf("We have a request at url: %s\n", message)
	http.ServeFile(w, r, "./form.html")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path

	fmt.Printf("We have a request at url: %s\n", message)
	msg := "hello world! I love you so much."
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	fmt.Printf("We have a request at url: %s\n", message)

	// name := c.FormValue("name")

	f, fh, err := r.FormFile("file")
	if err != nil {
		log.Error(err)
		w.Write([]byte(err.Error()))
		return
	}

	defer f.Close()

	if StorageBucket == nil {
		err = errors.New("storage bucket is missing - check config.go")
		log.Error(err)
		w.Write([]byte(err.Error()))
		return
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
		w.Write([]byte(err.Error()))
		return
	}
	if err := writer.Close(); err != nil {
		log.Error(err)
		w.Write([]byte(err.Error()))
		return
	}

	const publicURL = "https://storage.googleapis.com/%s/%s"
	imageUrl := fmt.Sprintf(publicURL, StorageBucketName, name)

	msg := imageUrl
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func main() {
	fmt.Println("Server listening at 9090")
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
