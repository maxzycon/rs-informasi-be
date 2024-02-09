package dto

type PayloadDeleteFile struct {
	Data  PayloadDeleteS3Path `json:"data"`
	Error *string             `json:"error"`
}

type PayloadDeleteS3Path struct {
	Path string `json:"path"`
}
