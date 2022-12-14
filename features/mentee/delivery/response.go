package delivery

import (
	"be12/mentutor/features/mentee"
)

type UpdateResponse struct {
	IdUser uint   `json:"id_user"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Images string `json:"images"`
}

type StatusRespon struct {
	ID       uint            `json:"id_status"`
	Name     string          `json:"name"`
	Images   string          `json:"images"`
	Caption  string          `json:"caption"`
	Comments []CommentRespon `json:"comments"`
}
type CommentRespon struct {
	ID      uint   `json:"id_comment"`
	Caption string `json:"caption"`
	Role    string `json:"role"`
	Name    string `json:"name"`
}
type PostStatusResponse struct {
	ID      uint   `json:"id_status"`
	IdUser  uint   `json:"id_user"`
	Images  string `json:"images"`
	Caption string `json:"caption"`
}
type CommentPostResponse struct {
	IdUser   uint   `json:"id_user"`
	IdStatus uint   `json:"id_status"`
	Caption  string `json:"caption"`
}
type SubResponse struct {
	ID    uint   `json:"id_submission"`
	Title string `json:"title"`
	File  string `json:"file"`
}
type TaskResponse struct {
	ID          uint   `json:"id_task"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      string `json:"images"`
	File        string `json:"file"`
	DueDate     string `json:"due_date"`
	Score       uint   `json:"score"`
	Status      string `json:"status"`
}

func tasksResponse(data []mentee.Task) []TaskResponse {
	var dataCore []TaskResponse
	for i := 0; i < len(data); i++ {
		dataCore = append(dataCore, TaskResponse{
			ID:          data[i].ID,
			Title:       data[i].Title,
			Images:      data[i].Images,
			Description: data[i].Description,
			File:        data[i].File,
			DueDate:     data[i].DueDate.Format("2006-01-02 15:04 MST"),
			Score:       data[i].Score,
			Status:      data[i].Status,
		})
	}
	return dataCore
}
func ToResponse(data mentee.Status) PostStatusResponse {
	return PostStatusResponse{
		ID:      data.ID,
		Caption: data.Caption,
		Images:  data.Images,
		IdUser:  data.IdMentee,
	}
}

func ToResponseComments(data mentee.CommentsCore) CommentPostResponse {
	return CommentPostResponse{
		IdStatus: data.IdStatus,
		IdUser:   data.ID_User,
		Caption:  data.Caption,
	}
}

func ToResponseSub(data mentee.Submission) SubResponse {
	return SubResponse{

		ID:    data.ID,
		Title: data.Title,
		File:  data.File,
	}
}

func ToCoreArray(status []mentee.Status, coment []mentee.CommentsCore, cmnmtr []mentee.CommentsCore) []StatusRespon {
	var res []StatusRespon

	for i, val := range status {
		var comres []CommentRespon
		for j, v := range coment {
			if status[i].ID == coment[j].IdStatus {
				comres = append(comres, CommentRespon{
					ID:      v.ID,
					Caption: v.Caption,
					Role:    v.Role,
					Name:    v.Name,
				})

			}

		}

		for k, v := range cmnmtr {
			if status[i].ID == cmnmtr[k].IdStatus {
				comres = append(comres, CommentRespon{
					ID:      v.ID,
					Caption: v.Caption,
					Role:    v.Role,
					Name:    v.Name,
				})

			}

		}

		res = append(res, StatusRespon{
			ID:       val.ID,
			Name:     val.Name,
			Images:   val.Images,
			Caption:  val.Caption,
			Comments: comres,
		})

	}

	return res
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
