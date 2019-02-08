// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var thegame_pb = require('./thegame_pb.js');

function serialize_protobuf_CreateRoomReq(arg) {
  if (!(arg instanceof thegame_pb.CreateRoomReq)) {
    throw new Error('Expected argument of type protobuf.CreateRoomReq');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_protobuf_CreateRoomReq(buffer_arg) {
  return thegame_pb.CreateRoomReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protobuf_Player(arg) {
  if (!(arg instanceof thegame_pb.Player)) {
    throw new Error('Expected argument of type protobuf.Player');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_protobuf_Player(buffer_arg) {
  return thegame_pb.Player.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protobuf_RoomState(arg) {
  if (!(arg instanceof thegame_pb.RoomState)) {
    throw new Error('Expected argument of type protobuf.RoomState');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_protobuf_RoomState(buffer_arg) {
  return thegame_pb.RoomState.deserializeBinary(new Uint8Array(buffer_arg));
}


var BaseGameService = exports.BaseGameService = {
  createRoom: {
    path: '/protobuf.BaseGame/CreateRoom',
    requestStream: false,
    responseStream: true,
    requestType: thegame_pb.CreateRoomReq,
    responseType: thegame_pb.RoomState,
    requestSerialize: serialize_protobuf_CreateRoomReq,
    requestDeserialize: deserialize_protobuf_CreateRoomReq,
    responseSerialize: serialize_protobuf_RoomState,
    responseDeserialize: deserialize_protobuf_RoomState,
  },
  joinRoom: {
    path: '/protobuf.BaseGame/JoinRoom',
    requestStream: false,
    responseStream: true,
    requestType: thegame_pb.Player,
    responseType: thegame_pb.RoomState,
    requestSerialize: serialize_protobuf_Player,
    requestDeserialize: deserialize_protobuf_Player,
    responseSerialize: serialize_protobuf_RoomState,
    responseDeserialize: deserialize_protobuf_RoomState,
  },
};

exports.BaseGameClient = grpc.makeGenericClientConstructor(BaseGameService);
