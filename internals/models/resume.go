package models

type Resume struct {
	Id        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	FileName  string `json:"file_name"`
	FilePath  string `json:"file_path"`
	CreatedAT string `json:"created_at"`
}

type ResumeRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	FileName string `json:"file_name" binding:"required,min=2,max=100"`
	FilePath string `json:"file_path" binding:"required,min=2"`
}
type ResumeResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	FileName  string `json:"file_name"`
	FilePath  string `json:"file_path"`
	CreatedAT string `json:"created_at"`
}

func NewResume(userID, name, fileName, filePath string) *Resume {

	return &Resume{
		UserID:   userID,
		Name:     name,
		FileName: fileName,
		FilePath: filePath,
	}

}

func NewResumeResponse(id, name, fileName, filePath, createdAt string) *ResumeResponse {
	return &ResumeResponse{
		Id:        id,
		Name:      name,
		FileName:  fileName,
		FilePath:  filePath,
		CreatedAT: createdAt,
	}
}

func NewResumeRequest(name, fileName, filePath string) *ResumeRequest {
	return &ResumeRequest{
		Name:     name,
		FileName: fileName,
		FilePath: filePath,
	}
}
