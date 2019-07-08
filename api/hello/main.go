package main

import (
	"context"
	"net/http"

	"github.com/micro/go-micro"
	"golang.org/x/net/trace"

	"github.com/qinhan-shu/go-micro-consul-cluster/api/hello/proto"
)

type Hello struct {
}

func (s *Hello) Say(ctx context.Context, req *hello.Reqeust, rsp *hello.Response) error {
	// tr := trace.New("api.v1", "Hotel.Rates")
	// defer tr.Finish()

	// ctx = trace.NewContext(ctx, tr)

	// md, ok := metadata.FromContext(ctx)
	// if !ok {
	// 	md = metadata.Metadata{}
	// }

	// if traceID, err := uuid.NewV4(); err == nil {
	// 	tmd := metadata.Metadata{}
	// 	for k, v := range md {
	// 		tmd[k] = v
	// 	}

	// 	tmd["traceID"] = traceID.String()
	// 	tmd["fromName"] = "api.v1"
	// 	ctx = metadata.NewContext(ctx, tmd)
	// }

	// token, err := getToken(md)
	// if err != nil {
	// 	return merr.Forbidden("api.hotel.rates", err.Error())
	// }

	// authClient := auth.NewAuthService("go.micro.srv.auth", s.Client)
	// if _, err = authClient.VerifyToken(ctx, &auth.Request{AuthToken: token}); err != nil {
	// 	return merr.Unauthorized("api.hotel.rates", "Unauthorized")
	// }

	// inDate, outDate := req.InDate, req.OutDate
	// if inDate == "" || outDate == "" {
	// 	return merr.BadRequest("api.hotel.rates", "Please specify inDate/outDate params")
	// }

	// // finds nearby hotels
	// // TODO(hw): use lat/lon from request params
	// geoClient := geo.NewGeoService("go.micro.srv.geo", s.Client)
	// nearby, err := geoClient.NearBy(ctx, &geo.Request{
	// 	Lat: 51.502973,
	// 	Lon: -0.114723,
	// })

	// if err != nil {
	// 	return merr.InternalServerError("api.hotel.rates...", err.Error())
	// }

	// profileCh := getHotelProfiles(s.Client, ctx, nearby.HotelIds)
	// rateCh := getRatePlans(s.Client, ctx, nearby.HotelIds, inDate, outDate)

	// profileReply := <-profileCh
	// if err := profileReply.err; err != nil {
	// 	return merr.InternalServerError("api.hotel.rates!!!", err.Error())
	// }

	// rateReply := <-rateCh
	// if err := rateReply.err; err != nil {
	// 	return merr.InternalServerError("api.hotel.rates????", err.Error())
	// }

	// rsp.Hotels = profileReply.hotels
	// rsp.RatePlans = rateReply.ratePlans
	rsp.Str = "ceshi "
	return nil
}

func main() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	service := micro.NewService(
		micro.Name("go.micro.api.hello"),
	)

	service.Init()
	hello.RegisterHelloHandler(service.Server(), new(Hello))
	service.Run()
}
