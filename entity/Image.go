package entity

type ImageURL struct {
	ImageURI          string      `xorm:"image_uri default '''::character varying' VARCHAR(256)"`
	//ThumbnailImageURI string      `xorm:"thumbnail_image_uri default '''::character varying' VARCHAR(256)"`
	CutboardImageURI  string      `xorm:"cutboard_image_uri default '''::character varying' VARCHAR(256)"`
}
