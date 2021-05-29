package service_test

import (
	"context"
	"testing"

	"github.com/jirawan-chuapradit/grpc-golang-pcbook/pb"
	"github.com/jirawan-chuapradit/grpc-golang-pcbook/sample"
	"github.com/jirawan-chuapradit/grpc-golang-pcbook/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	// case laptop without id
	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	// case laptop invalid id
	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	// case duplicate id
	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := service.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)
	require.Nil(t, err)



	testCase := []struct {
		name string
		laptop *pb.Laptop
		store service.LaptopStore
		code codes.Code
	}{
		{
			name: "success_with_id",
			laptop: sample.NewLaptop(),
			store: service.NewInMemoryLaptopStore(),
			code: codes.OK,
		},
		{
			name: "success_without_id",
			laptop: laptopNoID,
			store: service.NewInMemoryLaptopStore(),
			code: codes.OK,
		},
		{
			name: "failure_invalid_id",
			laptop: laptopInvalidID,
			store: service.NewInMemoryLaptopStore(),
			code: codes.InvalidArgument,
		},
		{
			name: "failure_duplicate_id",
			laptop: laptopDuplicateID,
			store: storeDuplicateID,
			code: codes.AlreadyExists,
		},
	}

	for i := range testCase {
		tc := testCase[i] // avoid concurrency error

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := service.NewLaptopServer(tc.store, nil)
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t,err)
				require.NotNil(t,res)
				require.NotEmpty(t,res.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			}else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError (err)
				require.True(t, ok)
				require.Equal(t,tc.code, st.Code())
			}
		})



	}
} 