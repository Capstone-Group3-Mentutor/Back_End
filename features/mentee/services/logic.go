package services

import (
	"be12/mentutor/features/mentee"
	"errors"
	"time"
)

type MenteeUsecase struct {
	menteeData mentee.RepoInterface
}

func New(data mentee.RepoInterface) mentee.UseCaseInterface {
	return &MenteeUsecase{
		menteeData: data,
	}
}

func (mu *MenteeUsecase) InsertStatus(data mentee.Status, token int) (mentee.Status, error) {
	if len(data.Caption) < 5 || len(data.Caption) > 120 {
		return mentee.Status{}, errors.New("input not valid")
	}
	// log.Print(data.Images, "INI LOG IMAGES")
	data, err := mu.menteeData.AddStatus(data, token)
	if err != nil {
		return mentee.Status{}, err
	}
	return data, nil
}

func (mu *MenteeUsecase) GetAll() ([]mentee.Status, []mentee.CommentsCore, []mentee.CommentsCore, error) {
	dataStatus, dataCmn, dataMntr, err := mu.menteeData.GetAllPosts()
	if err != nil {
		return nil, nil, nil, errors.New("failed get all data")
	}
	return dataStatus, dataCmn, dataMntr, nil
}

func (mu *MenteeUsecase) GetTask(idClass uint, role string) ([]mentee.Task, error) {
	if role != "mentee" {
		return []mentee.Task{}, errors.New("user not mentee")
	}

	dataTask, err := mu.menteeData.GetAllTask(idClass)
	if err != nil {
		return nil, err
	}
	return dataTask, nil
}

func (mu *MenteeUsecase) Insert(data mentee.CommentsCore) (mentee.CommentsCore, error) {
	if len(data.Caption) < 5 || len(data.Caption) > 120 {
		return mentee.CommentsCore{}, errors.New("failed add your comment check charancter len")
	}

	err := mu.menteeData.GetSingleStatus(data.IdStatus)
	if err != nil {
		return mentee.CommentsCore{}, errors.New("status not found")
	}

	data, err = mu.menteeData.AddComment(data)
	if err != nil {
		return mentee.CommentsCore{}, errors.New("error add comment")
	}
	return data, nil
}

func (mu *MenteeUsecase) InsertSub(data mentee.Submission) (mentee.Submission, error) {

	res, err := mu.menteeData.GetSingleTask(data.ID_Tasks)
	if err != nil {
		return mentee.Submission{}, errors.New("task not found")
	}

	timeRes := res.DueDate
	now := time.Now()
	if now.After(timeRes) == true {
		return mentee.Submission{}, errors.New("melewati due date")
	}

	data, err = mu.menteeData.AddSub(data)
	if err != nil {
		return mentee.Submission{}, err
	}
	return data, nil
}

func (mu *MenteeUsecase) AddToken(data mentee.Token) (mentee.Token, error) {
	res, err := mu.menteeData.InsertToken(data)
	if err != nil {
		return mentee.Token{}, err
	}
	return res, nil
}

func (mu *MenteeUsecase) GetTokenMentee(idMentee uint) (mentee.Token, error) {

	res, err := mu.menteeData.GetTokenMentee(idMentee)
	if err != nil {
		return mentee.Token{}, errors.New("error get token")
	}
	return res, nil
}

func (mu * MenteeUsecase) GetSingleTask(idTask uint) (mentee.Task, error) {
	res, err := mu.menteeData.GetSingleTask(idTask)
	if err != nil {
		return mentee.Task{}, errors.New("task not found")
	}
	return res, nil
}

func (mu *MenteeUsecase) GetMentee(idUser uint) (mentee.MenteeCore, error) {
	res, err := mu.menteeData.GetMentee(idUser)
	if err != nil {
		return mentee.MenteeCore{}, err
	}
	return res, nil
}

func (mu *MenteeUsecase) GetMentor(idUser uint) (mentee.MentorCore, error) {
	res, err := mu.menteeData.GetMentor(idUser)
	if err != nil {
		return mentee.MentorCore{}, err
	}
	return res, nil
}

func (mu *MenteeUsecase) GetSub(idUser uint, idTask uint) (int, error) {
	count, err := mu.menteeData.GetSub(idUser, idTask)
	if err != nil {
		return 0, errors.New("submission has been, submited")
	}
	return count, nil
}