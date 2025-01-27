package supabase

import (
	"errors"
	"log"
	"mime/multipart"
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

type Interface interface {
	Upload(file *multipart.FileHeader) (string, error)
	Delete(file string) error
}

type SupabaseStorage struct {
	client *storage_go.Client
}

func Init() (Interface, error) {
	url := os.Getenv("SUPABASE_URL")
	token := os.Getenv("SUPABASE_TOKEN")
	if url == "" || token == "" {
		return nil, errors.New("missing required environment variable for Supabase")
	}
	storageClient := storage_go.NewClient(url, token, nil)
	return &SupabaseStorage{
		client: storageClient,
	}, nil
}

func (s *SupabaseStorage) Upload(file *multipart.FileHeader) (string, error) {
	buff, err := file.Open()
	if err != nil {
		return "", err
	}
	defer buff.Close()

	bucket := os.Getenv("SUPABASE_BUCKET")
	_, err = s.client.UpdateFile(bucket, file.Filename, buff)
	if err != nil {
		log.Printf("Detailed error: %v\n", err)
		return "", err
	}

	link := s.client.GetPublicUrl(bucket, file.Filename).SignedURL
	return link, nil
}

func (s *SupabaseStorage) Delete(file string) error {
	bucket := os.Getenv("SUPABASE_BUCKET")
	_, err := s.client.RemoveFile(bucket, []string{file})
	if err != nil {
		return err
	}
	return nil
}
