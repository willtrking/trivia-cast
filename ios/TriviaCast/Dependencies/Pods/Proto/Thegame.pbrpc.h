#if !defined(GPB_GRPC_FORWARD_DECLARE_MESSAGE_PROTO) || !GPB_GRPC_FORWARD_DECLARE_MESSAGE_PROTO
#import "Thegame.pbobjc.h"
#endif

#if !defined(GPB_GRPC_PROTOCOL_ONLY) || !GPB_GRPC_PROTOCOL_ONLY
#import <ProtoRPC/ProtoService.h>
#import <ProtoRPC/ProtoRPC.h>
#import <RxLibrary/GRXWriteable.h>
#import <RxLibrary/GRXWriter.h>
#endif

@class CreateRoomReq;
@class Player;
@class RoomState;

#if !defined(GPB_GRPC_FORWARD_DECLARE_MESSAGE_PROTO) || !GPB_GRPC_FORWARD_DECLARE_MESSAGE_PROTO
#endif

@class GRPCProtoCall;


NS_ASSUME_NONNULL_BEGIN

@protocol BaseGame <NSObject>

#pragma mark CreateRoom(CreateRoomReq) returns (stream RoomState)

- (void)createRoomWithRequest:(CreateRoomReq *)request eventHandler:(void(^)(BOOL done, RoomState *_Nullable response, NSError *_Nullable error))eventHandler;

- (GRPCProtoCall *)RPCToCreateRoomWithRequest:(CreateRoomReq *)request eventHandler:(void(^)(BOOL done, RoomState *_Nullable response, NSError *_Nullable error))eventHandler;


#pragma mark JoinRoom(Player) returns (stream RoomState)

- (void)joinRoomWithRequest:(Player *)request eventHandler:(void(^)(BOOL done, RoomState *_Nullable response, NSError *_Nullable error))eventHandler;

- (GRPCProtoCall *)RPCToJoinRoomWithRequest:(Player *)request eventHandler:(void(^)(BOOL done, RoomState *_Nullable response, NSError *_Nullable error))eventHandler;


@end


#if !defined(GPB_GRPC_PROTOCOL_ONLY) || !GPB_GRPC_PROTOCOL_ONLY
/**
 * Basic service implementation, over gRPC, that only does
 * marshalling and parsing.
 */
@interface BaseGame : GRPCProtoService<BaseGame>
- (instancetype)initWithHost:(NSString *)host NS_DESIGNATED_INITIALIZER;
+ (instancetype)serviceWithHost:(NSString *)host;
@end
#endif

NS_ASSUME_NONNULL_END

