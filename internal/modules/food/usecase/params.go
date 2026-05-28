package usecase

import "mime/multipart"

type AnalyzeAndSaveFoodParams struct {
	File     multipart.File
	Filename string
	Comment  string
}

type FoodAnalyzerAnalyzeFoodParams struct {
	ImageUrl string
	Comment  string
}

type ImageStorageUploadParams struct {
	File     multipart.File
	Filename string
}
