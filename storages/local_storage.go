package storages

import "os"

type localStorage struct{}

func NewLocalStorage() *localStorage {
	return &localStorage{}
}

func (s *localStorage) UploadFile(localPath, remotePath string) error {
	input, err := os.ReadFile(localPath)
	if err != nil {
		return err
	}

	return os.WriteFile(remotePath, input, os.ModePerm)
}

func (s *localStorage) DownloadFile(remotePath, localPath string) error {
	input, err := os.ReadFile(remotePath)
	if err != nil {
		return err
	}

	return os.WriteFile(localPath, input, os.ModePerm)
}
