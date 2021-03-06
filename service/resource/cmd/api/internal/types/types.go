// Code generated by goctl. DO NOT EDIT.
package types

type Oss struct {
	TenantId   string `json:"tenantId,optional"`   // 租户ID
	Category   int    `json:"category,optional"`   // 所属分类
	OssCode    string `json:"ossCode,optional"`    // 资源编号
	Endpoint   string `json:"endpoint,optional"`   // 资源地址
	AccessKey  string `json:"accessKey,optional"`  // accessKey
	SecretKey  string `json:"secretKey,optional"`  // secretKey
	BucketName string `json:"bucketName,optional"` // 空间名
	AppId      string `json:"appId,optional"`      // 应用ID TencentCOS需要
	Region     string `json:"region,optional"`     // 地域简称 TencentCOS需要
	Remark     string `json:"remark,optional"`     // 所属分类
	Status     int    `json:"status,optional"`     // 状态
}

type OssCreate struct {
	TenantId   string `json:"tenantId,optional"`        // 租户ID
	Category   int    `json:"category,options=1|2|3|4"` // 所属分类
	OssCode    string `json:"ossCode,optional"`         // 资源编号
	Endpoint   string `json:"endpoint,optional"`        // 资源地址
	AccessKey  string `json:"accessKey,optional"`       // accessKey
	SecretKey  string `json:"secretKey,optional"`       // secretKey
	BucketName string `json:"bucketName,optional"`      // 空间名
	AppId      string `json:"appId,optional"`           // 应用ID TencentCOS需要
	Region     string `json:"region,optional"`          // 地域简称 TencentCOS需要
	Remark     string `json:"remark,optional"`          // 所属分类
	Status     int    `json:"status,optional"`          // 状态
}

type OssUpdate struct {
	Id         uint   `json:"id"`                                // 主键ID
	TenantId   string `json:"tenantId,optional"`                 // 租户ID
	Category   int    `json:"category,options=1|2|3|4,optional"` // 所属分类
	OssCode    string `json:"ossCode,optional"`                  // 资源编号
	Endpoint   string `json:"endpoint,optional"`                 // 资源地址
	AccessKey  string `json:"accessKey,optional"`                // accessKey
	SecretKey  string `json:"secretKey,optional"`                // secretKey
	BucketName string `json:"bucketName,optional"`               // 空间名
	AppId      string `json:"appId,optional"`                    // 应用ID TencentCOS需要
	Region     string `json:"region,optional"`                   // 地域简称 TencentCOS需要
	Remark     string `json:"remark,optional"`                   // 所属分类
	Status     int    `json:"status,optional"`                   // 状态
}

type OssDelete struct {
	Id uint `json:"id"` // 主键ID
}

type OssListReq struct {
	PageInfo
	Oss
}

type OssFile struct {
	Link        string `json:"link"`        // 文件地址
	Name        string `json:"name"`        // 文件名
	Length      int64  `json:"length"`      // 文件大小
	PutTime     string `json:"putTime"`     // 文件上传时间
	ContentType string `json:"contentType"` // 文件contentType
}

type MakeBucketReq struct {
	TenantId   string `json:"tenantId,optional"` // 租户ID
	Code       string `json:"code,optional"`     // 资源编号
	BucketName string `json:"bucketName"`        // 存储桶名称
}

type RemoveBucketReq struct {
	TenantId   string `json:"tenantId,optional"` // 租户ID
	Code       string `json:"code,optional"`     // 资源编号
	BucketName string `json:"bucketName"`        // 存储桶名称
}

type StatFileReq struct {
	TenantId   string `json:"tenantId,optional"`   // 租户ID
	Code       string `json:"code,optional"`       // 资源编号
	BucketName string `json:"bucketName,optional"` // 存储桶名称
	FileName   string `json:"fileName"`            // 文件名
}

type PutFileReq struct {
	TenantId   string `form:"tenantId,optional"`   // 租户ID
	Code       string `form:"code,optional"`       // 资源编号
	BucketName string `form:"bucketName,optional"` // 存储桶名称
}

type GetFileReq struct {
	TenantId   string `json:"tenantId,optional"`   // 租户ID
	Code       string `json:"code,optional"`       // 资源编号
	BucketName string `json:"bucketName,optional"` // 存储桶名称
	FileName   string `json:"fileName"`            // 文件名
}

type RemoveFileReq struct {
	TenantId   string `json:"tenantId,optional"`   // 租户ID
	Code       string `json:"code,optional"`       // 资源编号
	BucketName string `json:"bucketName,optional"` // 存储桶名称
	FileName   string `json:"fileName"`            // 文件名
}

type RemoveFilesReq struct {
	TenantId   string   `json:"tenantId,optional"`   // 租户ID
	Code       string   `json:"code,optional"`       // 资源编号
	BucketName string   `json:"bucketName,optional"` // 存储桶名称
	FileNames  []string `json:"fileNames"`           // 文件名集合
}

type File struct {
	Link         string `json:"link"`               // 文件地址
	Domain       string `json:"domain"`             // 域名地址
	Name         string `json:"name"`               // 文件名
	OriginalName string `json:"originalName"`       // 初始文件名
	AttachId     string `json:"attachId,omitempty"` // 附件表ID
}

type EmptyResult struct {
}

type BaseResult struct {
	Id uint `json:"id"` // 主键ID
}

type TenantResult struct {
	BaseResult
	TenantId uint `json:"tenantId"` // 租户ID
}

type PageInfo struct {
	Page     int `json:"page,range=[1:100],default=1"`      // 页码
	PageSize int `json:"pageSize,range=[1:100],default=10"` // 每页大小
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`     // 页码
	PageSize int         `json:"pageSize"` // 每页大小
}
