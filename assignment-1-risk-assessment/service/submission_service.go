package service

import (
	"context"
	"fmt"
	"github.com/ibrahimker/golang-praisindo-advanced/assignment-1-risk-assessment/entity"
)

// ISubmissionService mendefinisikan interface untuk layanan submission
type ISubmissionService interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) error
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
	GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.Submission, error)
}

// ISubmissionRepository mendefinisikan interface untuk repository submission
type ISubmissionRepository interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
	GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.Submission, error)
}

// submissionService adalah implementasi dari ISubmissionService yang menggunakan ISubmissionRepository
type submissionService struct {
	submissionRepo ISubmissionRepository
}

// NewSubmissionService membuat instance baru dari submissionService
func NewSubmissionService(submissionRepo ISubmissionRepository) ISubmissionService {
	return &submissionService{submissionRepo: submissionRepo}
}

func (s *submissionService) CreateSubmission(ctx context.Context, submission *entity.Submission) error {
	//TODO implement me
	// input validation

	// calculate risk profile
	score, category, definition := calculateProfileRiskFromAnswers(submission.Answers)
	fmt.Println(score, category, definition)

	// insert into repository
	panic("implement me")
}

func (s *submissionService) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	//TODO implement me
	panic("implement me")
}

func (s *submissionService) GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error) {
	//TODO implement me
	panic("implement me")
}

func (s *submissionService) DeleteSubmission(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *submissionService) GetAllSubmissions(ctx context.Context) ([]entity.Submission, error) {
	//TODO implement me
	panic("implement me")
}

// TODO: implement logic for profile risk calculation based on score mapping from entity.RiskMapping
// calculateProfileRiskFromAnswers will be used on submission creation
func calculateProfileRiskFromAnswers(answers []entity.Answer) (score int, category entity.ProfileRiskCategory, definition string) {
	// TODO: calculate total score from answers

	// TODO: get category and definition based on total score
	return score, category, definition
}
