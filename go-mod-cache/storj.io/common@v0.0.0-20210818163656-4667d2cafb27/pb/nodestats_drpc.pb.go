// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.23
// source: nodestats.proto

package pb

import (
	bytes "bytes"
	context "context"
	errors "errors"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"

	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_nodestats_proto struct{}

func (drpcEncoding_File_nodestats_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_nodestats_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_nodestats_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_nodestats_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCNodeStatsClient interface {
	DRPCConn() drpc.Conn

	GetStats(ctx context.Context, in *GetStatsRequest) (*GetStatsResponse, error)
	DailyStorageUsage(ctx context.Context, in *DailyStorageUsageRequest) (*DailyStorageUsageResponse, error)
	PricingModel(ctx context.Context, in *PricingModelRequest) (*PricingModelResponse, error)
}

type drpcNodeStatsClient struct {
	cc drpc.Conn
}

func NewDRPCNodeStatsClient(cc drpc.Conn) DRPCNodeStatsClient {
	return &drpcNodeStatsClient{cc}
}

func (c *drpcNodeStatsClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcNodeStatsClient) GetStats(ctx context.Context, in *GetStatsRequest) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/GetStats", drpcEncoding_File_nodestats_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodeStatsClient) DailyStorageUsage(ctx context.Context, in *DailyStorageUsageRequest) (*DailyStorageUsageResponse, error) {
	out := new(DailyStorageUsageResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/DailyStorageUsage", drpcEncoding_File_nodestats_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodeStatsClient) PricingModel(ctx context.Context, in *PricingModelRequest) (*PricingModelResponse, error) {
	out := new(PricingModelResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/PricingModel", drpcEncoding_File_nodestats_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCNodeStatsServer interface {
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	DailyStorageUsage(context.Context, *DailyStorageUsageRequest) (*DailyStorageUsageResponse, error)
	PricingModel(context.Context, *PricingModelRequest) (*PricingModelResponse, error)
}

type DRPCNodeStatsUnimplementedServer struct{}

func (s *DRPCNodeStatsUnimplementedServer) GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

func (s *DRPCNodeStatsUnimplementedServer) DailyStorageUsage(context.Context, *DailyStorageUsageRequest) (*DailyStorageUsageResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

func (s *DRPCNodeStatsUnimplementedServer) PricingModel(context.Context, *PricingModelRequest) (*PricingModelResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

type DRPCNodeStatsDescription struct{}

func (DRPCNodeStatsDescription) NumMethods() int { return 3 }

func (DRPCNodeStatsDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/nodestats.NodeStats/GetStats", drpcEncoding_File_nodestats_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeStatsServer).
					GetStats(
						ctx,
						in1.(*GetStatsRequest),
					)
			}, DRPCNodeStatsServer.GetStats, true
	case 1:
		return "/nodestats.NodeStats/DailyStorageUsage", drpcEncoding_File_nodestats_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeStatsServer).
					DailyStorageUsage(
						ctx,
						in1.(*DailyStorageUsageRequest),
					)
			}, DRPCNodeStatsServer.DailyStorageUsage, true
	case 2:
		return "/nodestats.NodeStats/PricingModel", drpcEncoding_File_nodestats_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeStatsServer).
					PricingModel(
						ctx,
						in1.(*PricingModelRequest),
					)
			}, DRPCNodeStatsServer.PricingModel, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterNodeStats(mux drpc.Mux, impl DRPCNodeStatsServer) error {
	return mux.Register(impl, DRPCNodeStatsDescription{})
}

type DRPCNodeStats_GetStatsStream interface {
	drpc.Stream
	SendAndClose(*GetStatsResponse) error
}

type drpcNodeStats_GetStatsStream struct {
	drpc.Stream
}

func (x *drpcNodeStats_GetStatsStream) SendAndClose(m *GetStatsResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_nodestats_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodeStats_DailyStorageUsageStream interface {
	drpc.Stream
	SendAndClose(*DailyStorageUsageResponse) error
}

type drpcNodeStats_DailyStorageUsageStream struct {
	drpc.Stream
}

func (x *drpcNodeStats_DailyStorageUsageStream) SendAndClose(m *DailyStorageUsageResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_nodestats_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodeStats_PricingModelStream interface {
	drpc.Stream
	SendAndClose(*PricingModelResponse) error
}

type drpcNodeStats_PricingModelStream struct {
	drpc.Stream
}

func (x *drpcNodeStats_PricingModelStream) SendAndClose(m *PricingModelResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_nodestats_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
