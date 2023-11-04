package handlers

import (
	"net/http"
	"visio/repositories"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type FacesHandlerv1 struct {
  logger *logrus.Logger
  faces_repo *repositories.FacesRepo
}

func NewFacesHandlerV1(logger *logrus.Logger, faces_repo *repositories.FacesRepo) *FacesHandlerv1 {
  return &FacesHandlerv1{
    logger: logger,
    faces_repo: faces_repo,
  }
}

func (h *FacesHandlerv1) GetFaces(w http.ResponseWriter, r *http.Request){
}
func (h *FacesHandlerv1) CreateFace(w http.ResponseWriter, r *http.Request){
}
func (h *FacesHandlerv1) UpdateFace(w http.ResponseWriter, r *http.Request){
}
func (h *FacesHandlerv1) DeleteFace(w http.ResponseWriter, r *http.Request){
}
func (h *FacesHandlerv1) CompareFaces(w http.ResponseWriter, r *http.Request){
}
func (h *FacesHandlerv1) DetectFace(w http.ResponseWriter, r *http.Request){
}

func (h *FacesHandlerv1) RegisterRoutes(r chi.Router){
  r.Get("/v1/faces", h.GetFaces)
  r.Post("/v1/faces", h.CreateFace)
  r.Put("/v1/faces/{face}", h.UpdateFace)
  r.Delete("/v1/faces/{face}", h.DeleteFace)

  r.Post("/v1/faces/detect", nil)
  r.Post("/v1/faces/compare", nil)
}
