#if !defined(GPB_GRPC_PROTOCOL_ONLY) || !GPB_GRPC_PROTOCOL_ONLY
#import "Thegame.pbrpc.h"
#import "Thegame.pbobjc.h"
#import <ProtoRPC/ProtoRPC.h>
#import <RxLibrary/GRXWriter+Immediate.h>


@implementation BaseGame

// Designated initializer
- (instancetype)initWithHost:(NSString *)host {
  self = [super initWithHost:host
                 packageName:@"protobuf"
                 serviceName:@"BaseGame"];
  return self;
}

// Override superclass initializer to disallow different package and service names.
- (instancetype)initWithHost:(NSString *)host
                 packageName:(NSString *)packageName
                 serviceName:(NSString *)serviceName {
  return [self initWithHost:host];
}

#pragma mark - Class Methods

+ (instancetype)serviceWithHost:(NSString *)host {
  return [[self alloc] initWithHost:host];
}

#pragma mark - Method Implementations

#pragma mark CreateRoom(CreateRoomReq) returns (stream RoomState)

- (void)createRoomWithRequest:(CreateRoomReq *)request eventHandler:(void(^)(BOOL done, RoomState *_Nullable response, NSError *_Nullable error))eventHandler{
  [[self RPCToCreateRoomWithRequest:request eventHandler:eventHandler] start];
}
// Returns a not-yet-started RPC object.
- (GRPCProtoCall *)RPCToCreateRoomWithRequest:(CreateRoomReq *)request eventHandler:(void(^)(BOOL done, RoomState *_Nullable response, NSError *_Nullable error))eventHandler{
  return [self RPCToMethod:@"CreateRoom"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[RoomState class]
        responsesWriteable:[GRXWriteable writeableWithEventHandler:eventHandler]];
}
#pragma mark JoinRoom(Player) returns (stream RoomState)

- (void)joinRoomWithRequest:(Player *)request eventHandler:(void(^)(BOOL done, RoomState *_Nullable response, NSError *_Nullable error))eventHandler{
  [[self RPCToJoinRoomWithRequest:request eventHandler:eventHandler] start];
}
// Returns a not-yet-started RPC object.
- (GRPCProtoCall *)RPCToJoinRoomWithRequest:(Player *)request eventHandler:(void(^)(BOOL done, RoomState *_Nullable response, NSError *_Nullable error))eventHandler{
  return [self RPCToMethod:@"JoinRoom"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[RoomState class]
        responsesWriteable:[GRXWriteable writeableWithEventHandler:eventHandler]];
}
@end
#endif
