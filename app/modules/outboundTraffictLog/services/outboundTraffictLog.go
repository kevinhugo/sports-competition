package services

import (
	outboundTraffictLogModel "sports-competition/app/modules/outboundTraffictLog/models"
	outboundTraffictLogRepository "sports-competition/app/modules/outboundTraffictLog/repositories"
)

type OutboundTraffictLogService struct {
	outboundTraffictLogRepository outboundTraffictLogRepository.OutboundTraffictLogRepository
}

func NewOutboundTraffictLogService() *OutboundTraffictLogService {
	return &OutboundTraffictLogService{
		outboundTraffictLogRepository: outboundTraffictLogRepository.NewOutboundTraffictLogRepository(),
	}
}

func (h *OutboundTraffictLogService) SaveOutboundTraffictLog(outboundTraffictLogData map[string]interface{}) (*outboundTraffictLogModel.OutboundTraffictLog, error) {
	outboundTraffictLog, err := h.outboundTraffictLogRepository.SaveOutboundTraffictLog(outboundTraffictLogData)
	return outboundTraffictLog, err
}

func (h *OutboundTraffictLogService) UpdateOutboundTraffictLog(existingData *outboundTraffictLogModel.OutboundTraffictLog, outboundTraffictLogData map[string]interface{}) (*outboundTraffictLogModel.OutboundTraffictLog, error) {
	var err error = h.outboundTraffictLogRepository.UpdateOutboundTraffictLog(existingData, outboundTraffictLogData)
	return existingData, err
}
