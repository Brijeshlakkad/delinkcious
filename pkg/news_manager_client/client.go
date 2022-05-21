package news_manager_client

import (
	"github.com/Brijeshlakkad/delinkcious/pb/news_service/pb"
	om "github.com/Brijeshlakkad/delinkcious/pkg/object_model"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type DisconnectFunc func() error

func NewClient(grpcAddr string) (cli om.NewsManager, disconnectFunc DisconnectFunc, err error) {
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	disconnectFunc = func() (err error) {
		if conn == nil {
			return
		}

		err = conn.Close()
		return
	}

	if err != nil {
		return
	}
	var getNewsEndpoint = grpctransport.NewClient(
		conn, "pb.News", "GetNews",
		encodeGetNewsRequest,
		decodeGetNewsResponse,
		pb.GetNewsResponse{},
	).Endpoint()

	cli = EndpointSet{
		GetNewsEndpoint: getNewsEndpoint,
	}
	return
}