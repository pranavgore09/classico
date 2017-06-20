// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var daily_pb = require('./daily_pb.js');

function serialize_daily_Text(arg) {
  if (!(arg instanceof daily_pb.Text)) {
    throw new Error('Expected argument of type daily.Text');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_daily_Text(buffer_arg) {
  return daily_pb.Text.deserializeBinary(new Uint8Array(buffer_arg));
}


var StarwarsService = exports.StarwarsService = {
  getQuote: {
    path: '/daily.Starwars/GetQuote',
    requestStream: false,
    responseStream: false,
    requestType: daily_pb.Text,
    responseType: daily_pb.Text,
    requestSerialize: serialize_daily_Text,
    requestDeserialize: deserialize_daily_Text,
    responseSerialize: serialize_daily_Text,
    responseDeserialize: deserialize_daily_Text,
  },
};

exports.StarwarsClient = grpc.makeGenericClientConstructor(StarwarsService);
