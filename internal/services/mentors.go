package services

import (
	"context"
	"errors"

	"github.com/jeevangb/authservice/internal/grpc/proto"
	"github.com/jeevangb/authservice/internal/models"
)

func (s *GrpcImpl) CreateProject(ctx context.Context, req *proto.CreateProjectRequest) (*proto.Project, error) {
	project := models.Project{
		Title:           req.Name,
		Description:     req.Description,
		TechnologyStack: req.TechnologyStack,
		MentorName:      req.MentorName,
		Status:          req.Status,
	}
	existingProject, err := s.usrsrc.GetProjectByTitle(ctx, project.Title, &project)
	if err != nil {
		return &proto.Project{}, errors.New("title already exists")
	}
	if existingProject != nil {
		return &proto.Project{}, errors.New("title already exists")
	}
	if err := s.usrsrc.CreateProject(ctx, project); err != nil {
		return &proto.Project{}, errors.New("failed to create project")
	}
	return &proto.Project{
		Name:            project.Title,
		Description:     project.Description,
		TechnologyStack: project.TechnologyStack,
		MentorName:      project.MentorName,
		Status:          project.Status,
	}, nil
}
