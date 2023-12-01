package storages

type Storage interface {
	UploadFile(localPath, remotePath string) error
	DownloadFile(remotePath, localPath string) error
}
