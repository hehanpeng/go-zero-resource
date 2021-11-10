package ossx

import (
	"github.com/minio/minio-go"
	"github.com/tal-tech/go-zero/core/fx"
	"go-zero-resource/service/resource/model/gorm_model"
	"mime/multipart"
)

type MinioTemplate struct {
	client        *minio.Client // Minio客户端
	ossProperties OssProperties // 配置参数
	ossRule       OssRule
}

func (m MinioTemplate) MakeBucket(tenantId, bucketName string) error {
	return m.client.MakeBucket(m.ossRule.bucketName(tenantId, bucketName), "")
}

func (m MinioTemplate) RemoveBucket(tenantId, bucketName string) error {
	return m.client.RemoveBucket(m.ossRule.bucketName(tenantId, bucketName))
}

func (m MinioTemplate) BucketExists(tenantId, bucketName string) (bool, error) {
	return m.client.BucketExists(m.ossRule.bucketName(tenantId, bucketName))
}

func (m MinioTemplate) PutFile(tenantId, bucketName string, file *multipart.FileHeader) (*File, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fileName := m.ossRule.fileName(file.Filename)
	if len(bucketName) == 0 {
		bucketName = m.ossProperties.BucketName
	}
	_, err = m.client.PutObject(m.ossRule.bucketName(tenantId, bucketName),
		fileName, f, file.Size, minio.PutObjectOptions{
			ContentType: file.Header.Get("content-type"),
		})
	if err != nil {
		return nil, err
	} else {
		return &File{
			Link:         m.fileLink(tenantId, bucketName, fileName),
			Domain:       m.getOssHost(tenantId, bucketName),
			Name:         fileName,
			OriginalName: file.Filename,
		}, nil
	}
}

func (m MinioTemplate) RemoveFile(tenantId, bucketName, fileName string) error {
	return m.client.RemoveObject(m.ossRule.bucketName(tenantId, bucketName), m.ossRule.fileName(fileName))
}

func (m MinioTemplate) RemoveFiles(tenantId string, bucketName string, fileNames []string) error {
	objectsCh := make(chan string)
	go func() {
		defer close(objectsCh)
		fx.Just(fileNames).ForEach(func(item interface{}) {
			objectsCh <- m.ossRule.fileName(item.(string))
		})
	}()
	errorCh := m.client.RemoveObjects(m.ossRule.bucketName(tenantId, bucketName), objectsCh)
	select {
	case resp, flag := <-errorCh:
		if flag {
			return resp.Err
		}
	}
	return nil
}

func (m MinioTemplate) getOssHost(tenantId, bucketName string) string {
	return m.ossProperties.Endpoint + "/" + m.ossRule.bucketName(tenantId, bucketName)
}

func (m MinioTemplate) fileLink(tenantId, bucketName, fileName string) string {
	return m.ossProperties.Endpoint + "/" + m.ossRule.bucketName(tenantId, bucketName) + "/" + fileName
}

func NewMinioTemplate(Oss gorm_model.ResourceOss, ossRule OssRule) *MinioTemplate {
	ossProperties := OssProperties{
		Endpoint:   Oss.Endpoint,
		AccessKey:  Oss.AccessKey,
		SecretKey:  Oss.SecretKey,
		BucketName: Oss.BucketName,
		Args:       nil,
	}
	// 初使化 minio client对象。
	minioClient, _ := minio.New(Oss.Endpoint, Oss.AccessKey, Oss.SecretKey, false)
	return &MinioTemplate{
		client:        minioClient,
		ossProperties: ossProperties,
		ossRule:       ossRule,
	}
}