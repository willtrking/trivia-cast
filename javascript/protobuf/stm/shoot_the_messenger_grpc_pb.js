// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var stm_shoot_the_messenger_pb = require('../stm/shoot_the_messenger_pb.js');

function serialize_stm_BeginGameReq(arg) {
  if (!(arg instanceof stm_shoot_the_messenger_pb.BeginGameReq)) {
    throw new Error('Expected argument of type stm.BeginGameReq');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_stm_BeginGameReq(buffer_arg) {
  return stm_shoot_the_messenger_pb.BeginGameReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stm_GameState(arg) {
  if (!(arg instanceof stm_shoot_the_messenger_pb.GameState)) {
    throw new Error('Expected argument of type stm.GameState');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_stm_GameState(buffer_arg) {
  return stm_shoot_the_messenger_pb.GameState.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_stm_PlayerFrame(arg) {
  if (!(arg instanceof stm_shoot_the_messenger_pb.PlayerFrame)) {
    throw new Error('Expected argument of type stm.PlayerFrame');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_stm_PlayerFrame(buffer_arg) {
  return stm_shoot_the_messenger_pb.PlayerFrame.deserializeBinary(new Uint8Array(buffer_arg));
}


var ShootTheMessengerService = exports.ShootTheMessengerService = {
  beginGame: {
    path: '/stm.ShootTheMessenger/BeginGame',
    requestStream: false,
    responseStream: false,
    requestType: stm_shoot_the_messenger_pb.BeginGameReq,
    responseType: stm_shoot_the_messenger_pb.GameState,
    requestSerialize: serialize_stm_BeginGameReq,
    requestDeserialize: deserialize_stm_BeginGameReq,
    responseSerialize: serialize_stm_GameState,
    responseDeserialize: deserialize_stm_GameState,
  },
  playerStream: {
    path: '/stm.ShootTheMessenger/PlayerStream',
    requestStream: true,
    responseStream: true,
    requestType: stm_shoot_the_messenger_pb.PlayerFrame,
    responseType: stm_shoot_the_messenger_pb.PlayerFrame,
    requestSerialize: serialize_stm_PlayerFrame,
    requestDeserialize: deserialize_stm_PlayerFrame,
    responseSerialize: serialize_stm_PlayerFrame,
    responseDeserialize: deserialize_stm_PlayerFrame,
  },
};

exports.ShootTheMessengerClient = grpc.makeGenericClientConstructor(ShootTheMessengerService);
