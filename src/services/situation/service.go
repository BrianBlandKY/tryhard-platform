package situation

import (
	cfg "tryhard-platform/src/config"
	msg "tryhard-platform/src/messenger"
	service "tryhard-platform/src/service"
)

// Actions
const (
	GENERATE = "GENERATE"
)

// Data
type SituationData struct {
	Situation string  `json:"situation"`
	Solution  float64 `json:"solution"`
}

// Service
type situationService struct {
	service.Service
	situation Situation
}

func (s *situationService) Execute() {
	cmdCh := s.Start()

	// Listen for commands
	for {
		select {
		case cmd := <-cmdCh:
			s.processCommand(cmd)
		}
	}
}

func (s *situationService) processCommand(cmd msg.Command) {
	switch cmd.Action {
	case GENERATE:
		data := s.getSituation()
		cmd.Serialize(data)
		s.Reply(cmd)
	default:
		s.Println("unknown command", cmd)
	}
}

func (s *situationService) getSituation() SituationData {
	s.situation.Generate(2, 1, 2, "+", "-")
	return SituationData{
		Situation: s.situation.GetProblem(),
		Solution:  s.situation.GetSolution(),
	}
}

func GerSituationService(config cfg.Config) service.Service {
	return &situationService{
		Service:   service.DefaultService(config),
		situation: DefaultSituation(),
	}
}
