package serverhandler

import (
	"context"
	"errors"
	"strings"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"google.golang.org/grpc/metadata"
)

// TestSetData - тест для функции SetData.
func (s *HandlerSuite) TestSetData() {
	s.server()
	defer s.closer()

	tests := map[string]struct {
		in  *pb.Value
		err error
	}{
		"OK_Login_Password_Data": {
			in: &pb.Value{
				Kind: &pb.Value_LoginPassword{
					LoginPassword: &pb.LoginPasswordData{
						Login:       "test",
						Password:    "test",
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"OK_Text_Data": {
			in: &pb.Value{
				Kind: &pb.Value_Text{
					Text: &pb.TextData{
						Data:        "test",
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"OK_Binary_Data": {
			in: &pb.Value{
				Kind: &pb.Value_BinData{
					BinData: &pb.BinaryData{
						Data:        []byte("test"),
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"OK_Card_Data": {
			in: &pb.Value{
				Kind: &pb.Value_CardData{
					CardData: &pb.BankCardData{
						Number:      "test",
						ValidThru:   "test",
						Cvv:         "test",
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"Service error": {
			in: &pb.Value{
				Kind: &pb.Value_Text{
					Text: &pb.TextData{
						Data:        "",
						Description: "test",
					},
				},
			},
			err: errors.New("rpc error: code = Internal desc = server error"),
		},
	}

	for name, tt := range tests {
		s.Run(name, func() {
			if strings.Contains(name, "Login") {
				s.svc.EXPECT().SetLoginPasswordData("test", models.LoginPasswordData{
					Login:       "test",
					Password:    "test",
					Description: "test",
				}).Return(nil)
			} else if strings.Contains(name, "Text") {
				s.svc.EXPECT().SetTextData("test", models.TextData{
					Description: "test",
					Data:        "test",
				}).Return(nil)
			} else if strings.Contains(name, "Binary") {
				s.svc.EXPECT().SetBinaryData("test", models.BinaryData{
					Description: "test",
					Data:        []byte("test"),
				}).Return(nil)
			} else if strings.Contains(name, "Card") {
				s.svc.EXPECT().SetBankCardData("test", models.BankCardData{
					Description: "test",
					Number:      "test",
					ValidThru:   "test",
					CVV:         "test",
				}).Return(nil)
			} else if name == "Service error" {
				s.svc.EXPECT().SetTextData("test", models.TextData{
					Description: "test",
					Data:        "",
				}).Return(errors.New("service error"))
			}

			ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"authentication": s.authToken}))

			_, err := s.storeClient.SetData(ctx, tt.in)
			if err != nil {
				s.Require().Equal(tt.err.Error(), err.Error())
			}
		})
	}
}

// TestSync - тест для функции Sync.
func (s *HandlerSuite) TestSync() {
	s.server()
	defer s.closer()

	tests := map[string]struct {
		expected *pb.Value
		err      error
	}{
		"OK_Login_Password_Data": {
			expected: &pb.Value{
				Kind: &pb.Value_LoginPassword{
					LoginPassword: &pb.LoginPasswordData{
						Login:       "test",
						Password:    "test",
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"OK_Text_Data": {
			expected: &pb.Value{
				Kind: &pb.Value_Text{
					Text: &pb.TextData{
						Data:        "test",
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"OK_Binary_Data": {
			expected: &pb.Value{
				Kind: &pb.Value_BinData{
					BinData: &pb.BinaryData{
						Data:        []byte("test"),
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"OK_Card_Data": {
			expected: &pb.Value{
				Kind: &pb.Value_CardData{
					CardData: &pb.BankCardData{
						Number:      "test",
						ValidThru:   "test",
						Cvv:         "test",
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"Service error": {
			expected: &pb.Value{
				Kind: &pb.Value_Text{
					Text: &pb.TextData{
						Data:        "",
						Description: "test",
					},
				},
			},
			err: errors.New("rpc error: code = Unknown desc = service error"),
		},
	}

	for name, tt := range tests {
		s.Run(name, func() {
			if strings.Contains(name, "Login") {
				s.svc.EXPECT().Sync("test").Return([]any{models.LoginPasswordData{
					Login:       "test",
					Password:    "test",
					Description: "test",
				}}, nil)
			} else if strings.Contains(name, "Text") {
				s.svc.EXPECT().Sync("test").Return([]any{models.TextData{
					Description: "test",
					Data:        "test",
				}}, nil)
			} else if strings.Contains(name, "Binary") {
				s.svc.EXPECT().Sync("test").Return([]any{models.BinaryData{
					Description: "test",
					Data:        []byte("test"),
				}}, nil)
			} else if strings.Contains(name, "Card") {
				s.svc.EXPECT().Sync("test").Return([]any{models.BankCardData{
					Description: "test",
					Number:      "test",
					ValidThru:   "test",
					CVV:         "test",
				}}, nil)
			} else if name == "Service error" {
				s.svc.EXPECT().Sync("test").Return([]any{}, errors.New("service error"))
			}

			ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"authentication": s.authToken}))

			dataArray, err := s.storeClient.Sync(ctx, &pb.Empty{})
			if err != nil {
				s.Require().Equal(tt.err.Error(), err.Error())
			}

			if err == nil {
				s.Require().Equal(tt.expected, dataArray.Values[0])
			}
		})
	}
}

// TestUpdateData - тест для функции UpdateData.
func (s *HandlerSuite) TestUpdateData() {
	s.server()
	defer s.closer()

	tests := map[string]struct {
		in  *pb.Value
		err error
	}{
		"OK_Login_Password_Data": {
			in: &pb.Value{
				Kind: &pb.Value_LoginPassword{
					LoginPassword: &pb.LoginPasswordData{
						Login:       "test",
						Password:    "test",
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"OK_Text_Data": {
			in: &pb.Value{
				Kind: &pb.Value_Text{
					Text: &pb.TextData{
						Data:        "test",
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"OK_Binary_Data": {
			in: &pb.Value{
				Kind: &pb.Value_BinData{
					BinData: &pb.BinaryData{
						Data:        []byte("test"),
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"OK_Card_Data": {
			in: &pb.Value{
				Kind: &pb.Value_CardData{
					CardData: &pb.BankCardData{
						Number:      "test",
						ValidThru:   "test",
						Cvv:         "test",
						Description: "test",
					},
				},
			},
			err: nil,
		},
		"Service error": {
			in: &pb.Value{
				Kind: &pb.Value_Text{
					Text: &pb.TextData{
						Data:        "",
						Description: "test",
					},
				},
			},
			err: errors.New("rpc error: code = Internal desc = server error"),
		},
	}

	for name, tt := range tests {
		s.Run(name, func() {
			if strings.Contains(name, "Login") {
				s.svc.EXPECT().UpdateLoginPasswordData("test", models.LoginPasswordData{
					Login:       "test",
					Password:    "test",
					Description: "test",
				}).Return(nil)
			} else if strings.Contains(name, "Text") {
				s.svc.EXPECT().UpdateTextData("test", models.TextData{
					Description: "test",
					Data:        "test",
				}).Return(nil)
			} else if strings.Contains(name, "Binary") {
				s.svc.EXPECT().UpdateBinaryData("test", models.BinaryData{
					Description: "test",
					Data:        []byte("test"),
				}).Return(nil)
			} else if strings.Contains(name, "Card") {
				s.svc.EXPECT().UpdateBankCardData("test", models.BankCardData{
					Description: "test",
					Number:      "test",
					ValidThru:   "test",
					CVV:         "test",
				}).Return(nil)
			} else if name == "Service error" {
				s.svc.EXPECT().UpdateTextData("test", models.TextData{
					Description: "test",
					Data:        "",
				}).Return(errors.New("service error"))
			}

			ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"authentication": s.authToken}))

			_, err := s.storeClient.UpdateData(ctx, tt.in)
			if err != nil {
				s.Require().Equal(tt.err.Error(), err.Error())
			}
		})
	}
}

// TestDeleteData - тесты для функции DeleteData.
func (s *HandlerSuite) TestDeleteData() {
	s.server()
	defer s.closer()

	tests := map[string]struct {
		in  *pb.Key
		err error
	}{
		"OK": {
			in:  &pb.Key{Key: "test"},
			err: nil,
		},

		"Service error": {
			in:  &pb.Key{Key: "test"},
			err: errors.New("rpc error: code = Unknown desc = service error"),
		},
	}

	for name, tt := range tests {
		s.Run(name, func() {
			if name == "OK" {
				s.svc.EXPECT().DeleteData(tt.in.Key, "test").Return(nil)
			} else {
				s.svc.EXPECT().DeleteData(tt.in.Key, "test").Return(errors.New("service error"))
			}
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"authentication": s.authToken}))

			_, err := s.storeClient.DeleteData(ctx, tt.in)
			if err != nil {
				s.Require().Equal(tt.err.Error(), err.Error())
			}
		})
	}
}
