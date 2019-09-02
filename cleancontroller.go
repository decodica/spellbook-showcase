package main

import (
	"context"
	"decodica.com/flamel"
	"net/http"
)

type CleanController struct{}

func (controller *CleanController) OnDestroy(ctx context.Context) {}

func (controller *CleanController) Process(ctx context.Context, out *flamel.ResponseOutput) flamel.HttpResponse {
	return flamel.HttpResponse{Status: http.StatusNotImplemented}
}
