package service

import (
	"kasir-api/models"
	"kasir-api/repository"
)

type TransactionService interface {
	CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error)
	GetDailyReport() (*models.DailyReport, error)
	GetReportByDateRange(startDate, endDate string) (*models.DailyReport, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{
		repo: repo,
	}
}

func (s *transactionService) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *transactionService) GetDailyReport() (*models.DailyReport, error) {
	return s.repo.GetDailyReport()
}

func (s *transactionService) GetReportByDateRange(startDate, endDate string) (*models.DailyReport, error) {
	return s.repo.GetReportByDateRange(startDate, endDate)
}
