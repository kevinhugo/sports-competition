package services

import (
	inboundTraffictLogModel "sports-competition/app/modules/inboundTraffictLog/models"
	inboundTraffictLogRepository "sports-competition/app/modules/inboundTraffictLog/repositories"
)

type InboundTraffictLogService struct {
	inboundTraffictLogRepository inboundTraffictLogRepository.InboundTraffictLogRepository
}

func NewInboundTraffictLogService() *InboundTraffictLogService {
	return &InboundTraffictLogService{
		inboundTraffictLogRepository: inboundTraffictLogRepository.NewInboundTraffictLogRepository(),
	}
}

func (h *InboundTraffictLogService) SaveInboundTraffictLog(inboundTraffictLogData map[string]interface{}) (*inboundTraffictLogModel.InboundTraffictLog, error) {
	inboundTraffictLog, err := h.inboundTraffictLogRepository.SaveInboundTraffictLog(inboundTraffictLogData)
	return inboundTraffictLog, err
}

func (h *InboundTraffictLogService) UpdateInboundTraffictLog(existingData *inboundTraffictLogModel.InboundTraffictLog, inboundTraffictLogData map[string]interface{}) (*inboundTraffictLogModel.InboundTraffictLog, error) {
	var err error = h.inboundTraffictLogRepository.UpdateInboundTraffictLog(existingData, inboundTraffictLogData)
	return existingData, err
}
