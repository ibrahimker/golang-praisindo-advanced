package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/assignment-1-risk-assessment/service"
)

// ISubmissionHandler mendefinisikan interface untuk handler submission
type ISubmissionHandler interface {
	CreateSubmission(c *gin.Context)
	GetSubmission(c *gin.Context)
	GetAllSubmissions(c *gin.Context)
	DeleteSubmission(c *gin.Context)
}

// NewSubmissionHandler membuat instance baru dari SubmissionHandler
func NewSubmissionHandler(submissionService service.ISubmissionService) ISubmissionHandler {
	return &SubmissionHandler{
		submissionService: submissionService,
	}
}

type SubmissionHandler struct {
	submissionService service.ISubmissionService
}

func (s *SubmissionHandler) CreateSubmission(c *gin.Context) {
	//TODO implement meb
	panic("implement me")
}

func (s *SubmissionHandler) GetSubmission(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *SubmissionHandler) DeleteSubmission(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *SubmissionHandler) GetAllSubmissions(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
