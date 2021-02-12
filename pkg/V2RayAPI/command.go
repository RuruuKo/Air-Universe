package V2RayAPI

import (
	"google.golang.org/grpc"
	v2Cmd "v2ray.com/core/app/proxyman/command"
	v2StatsService "v2ray.com/core/app/stats/command"
)

type V2rayController struct {
	HsClient *v2Cmd.HandlerServiceClient
	SsClient *v2StatsService.StatsServiceClient
	CmdConn  *grpc.ClientConn
}