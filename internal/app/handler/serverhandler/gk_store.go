package serverhandler

import (
	"context"
	"fmt"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type StorageServer struct {
	pb.UnimplementedStorageServer

	Service KeeperService
}

func (s *StorageServer) SetData(ctx context.Context, in *pb.Value) (*pb.ResponseStatus, error) {
	var userID string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			userID = values[0]
		}
	}
	if userID == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	switch tp := in.Kind.(type) {
	case *pb.Value_LoginPassword:
		if err := s.Service.SetLoginPasswordData(userID, models.LoginPasswordData{
			Description: tp.LoginPassword.Description,
			Login:       tp.LoginPassword.Login,
			Password:    tp.LoginPassword.Password,
		}); err != nil {
			log.Error().Err(err).Send()
			return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, status.Error(codes.Internal, "server error")
		}
	case *pb.Value_Text:
		if err := s.Service.SetTextData(userID, models.TextData{
			Description: tp.Text.Description,
			Data:        tp.Text.Data,
		}); err != nil {
			log.Error().Err(err).Send()
			return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, status.Error(codes.Internal, "server error")
		}
	case *pb.Value_BinData:
		if err := s.Service.SetBinaryData(userID, models.BinaryData{
			Description: tp.BinData.Description,
			Data:        tp.BinData.Data,
		}); err != nil {
			log.Error().Err(err).Send()
			return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, status.Error(codes.Internal, "server error")
		}
	case *pb.Value_CardData:
		if err := s.Service.SetBankCardData(userID, models.BankCardData{
			Description: tp.CardData.Description,
			Number:      tp.CardData.Number,
			ValidThru:   tp.CardData.ValidThru,
			CVV:         tp.CardData.Cvv,
		}); err != nil {
			log.Error().Err(err).Send()
			return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, status.Error(codes.Internal, "server error")
		}
	}

	return &pb.ResponseStatus{
		Status:  pb.ResponseStatus_OK,
		Comment: "data seted",
	}, nil
}

/*func (s *StorageServer) GetData(ctx context.Context, in *pb.Key) (*pb.Value, error) {
	var userID string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			userID = values[0]
		}
	}
	if userID == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	data, err := s.Service.GetData(in.Key, userID)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	switch t := data.(type) {
	case models.LoginPasswordData:
		return &pb.Value{
			Kind: &pb.Value_LoginPassword{
				LoginPassword: &pb.LoginPasswordData{
					Login:       t.Login,
					Password:    t.Password,
					Description: t.Description,
				},
			},
		}, nil
	case models.TextData:
		return &pb.Value{
			Kind: &pb.Value_Text{
				Text: &pb.TextData{
					Data:        t.Data,
					Description: t.Description,
				},
			},
		}, nil
	case models.BinaryData:
		return &pb.Value{
			Kind: &pb.Value_BinData{
				BinData: &pb.BinaryData{
					Data:        t.Data,
					Description: t.Description,
				},
			},
		}, nil
	case models.BankCardData:
		return &pb.Value{
			Kind: &pb.Value_CardData{
				CardData: &pb.BankCardData{
					Number:      t.Number,
					ValidThru:   t.ValidThru,
					Cvv:         t.CVV,
					Description: t.Description,
				},
			},
		}, nil
	}

	return nil, status.Error(codes.Internal, "server error")
}*/

func (s *StorageServer) GetDataArray(ctx context.Context, in *pb.Empty) (*pb.DataArray, error) {
	var userID string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			userID = values[0]
		}
	}
	if userID == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	dataArray, err := s.Service.GetDataArray(userID)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	result := make([]*pb.Value, len(dataArray))
	for i := range dataArray {
		var tmp *pb.Value
		switch t := dataArray[i].(type) {
		case models.LoginPasswordData:
			tmp = &pb.Value{
				Kind: &pb.Value_LoginPassword{
					LoginPassword: &pb.LoginPasswordData{
						Login:       t.Login,
						Password:    t.Password,
						Description: t.Description,
					},
				},
			}
		case models.TextData:
			tmp = &pb.Value{
				Kind: &pb.Value_Text{
					Text: &pb.TextData{
						Data:        t.Data,
						Description: t.Description,
					},
				},
			}
		case models.BinaryData:
			tmp = &pb.Value{
				Kind: &pb.Value_BinData{
					BinData: &pb.BinaryData{
						Data:        t.Data,
						Description: t.Description,
					},
				},
			}
		case models.BankCardData:
			tmp = &pb.Value{
				Kind: &pb.Value_CardData{
					CardData: &pb.BankCardData{
						Number:      t.Number,
						ValidThru:   t.ValidThru,
						Cvv:         t.CVV,
						Description: t.Description,
					},
				},
			}
		}
		result[i] = tmp
	}

	return &pb.DataArray{Values: result}, nil
}

func (s *StorageServer) UpdateData(ctx context.Context, in *pb.Value) (*pb.ResponseStatus, error) {
	var userID string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			userID = values[0]
		}
	}
	if userID == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	switch tp := in.Kind.(type) {
	case *pb.Value_LoginPassword:
		if err := s.Service.UpdateLoginPasswordData(userID, models.LoginPasswordData{
			Description: tp.LoginPassword.Description,
			Login:       tp.LoginPassword.Login,
			Password:    tp.LoginPassword.Password,
		}); err != nil {
			log.Error().Err(err).Send()
			return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, status.Error(codes.Internal, "server error")
		}
	case *pb.Value_Text:
		if err := s.Service.UpdateTextData(userID, models.TextData{
			Description: tp.Text.Description,
			Data:        tp.Text.Data,
		}); err != nil {
			log.Error().Err(err).Send()
			return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, status.Error(codes.Internal, "server error")
		}
	case *pb.Value_BinData:
		if err := s.Service.UpdateBinaryData(userID, models.BinaryData{
			Description: tp.BinData.Description,
			Data:        tp.BinData.Data,
		}); err != nil {
			log.Error().Err(err).Send()
			return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, status.Error(codes.Internal, "server error")
		}
	case *pb.Value_CardData:
		if err := s.Service.UpdateBankCardData(userID, models.BankCardData{
			Description: tp.CardData.Description,
			Number:      tp.CardData.Number,
			ValidThru:   tp.CardData.ValidThru,
			CVV:         tp.CardData.Cvv,
		}); err != nil {
			log.Error().Err(err).Send()
			return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, status.Error(codes.Internal, "server error")
		}
	}

	return &pb.ResponseStatus{
		Status:  pb.ResponseStatus_OK,
		Comment: "data updated",
	}, nil
}

func (s *StorageServer) DeleteData(ctx context.Context, in *pb.Key) (*pb.ResponseStatus, error) {
	var userID string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			userID = values[0]
		}
	}
	if userID == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	if err := s.Service.DeleteData(in.Key, userID); err != nil {
		log.Error().Err(err).Send()
		return &pb.ResponseStatus{
			Status: pb.ResponseStatus_FAIL,
		}, err
	}

	return &pb.ResponseStatus{
		Status:  pb.ResponseStatus_OK,
		Comment: fmt.Sprintf("data with key %s deleted", in.Key),
	}, nil
}
