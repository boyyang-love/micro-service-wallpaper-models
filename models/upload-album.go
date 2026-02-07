package models

type UploadAlbum struct {
	UploadId string `form:"upload_id" json:"upload_id"`
	AlbumId  string `form:"album_id" json:"album_id"`
}

func (u *UploadAlbum) TableName() string {
	return "upload_album"
}
