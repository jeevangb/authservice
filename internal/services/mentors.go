package services

import (
	"context"
	"errors"

	"github.com/jeevangb/authservice/internal/grpc/proto"
	"github.com/jeevangb/authservice/internal/models"
	"github.com/rs/zerolog/log"
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

func (s *GrpcImpl) UpdateProject(ctx context.Context, req *proto.UpdateProjectRequest) (*proto.Project, error) {
	project := &models.Project{
		Title:           req.Name,
		Description:     req.Description,
		TechnologyStack: req.TechnologyStack,
		MentorName:      req.MentorName,
		Status:          req.Status,
	}
	existingProject, err := s.usrsrc.GetProjectByTitle(ctx, project.Title, project)
	if err != nil {
		return &proto.Project{}, errors.New("project not exists to update")
	}

	if req.Description != "" {
		existingProject.Description = req.Description
	}
	if req.Name != "" {
		existingProject.Title = req.Name
	}
	if req.Status != "" {
		existingProject.Status = req.Status
	}
	if len(req.TechnologyStack) > 0 {
		existingProject.TechnologyStack = req.TechnologyStack
	}
	project, err = s.usrsrc.SaveProject(existingProject)
	if err != nil {
		return nil, err
	}
	return &proto.Project{
		Name:            project.Title,
		Description:     project.Description,
		TechnologyStack: project.TechnologyStack,
		MentorName:      project.MentorName,
		Status:          project.Status,
	}, nil
}

func (s *GrpcImpl) DeleteProject(ctx context.Context, req *proto.DeleteProjectRequest) (*proto.DeleteProjectResponse, error) {
	if err := s.usrsrc.DeleteProjectByTitle(ctx, req.Name); err != nil {
		log.Error().Err(err).Msg("failed to delete project")
		return &proto.DeleteProjectResponse{Success: false}, errors.New("failed to delete project")
	}
	return &proto.DeleteProjectResponse{
		Success: true,
	}, nil
}
