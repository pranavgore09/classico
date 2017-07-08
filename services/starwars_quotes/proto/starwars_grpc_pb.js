// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var starwars_pb = require('./starwars_pb.js');

function serialize_starwars_quotes_Action(arg) {
  if (!(arg instanceof starwars_pb.Action)) {
    throw new Error('Expected argument of type starwars_quotes.Action');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_starwars_quotes_Action(buffer_arg) {
  return starwars_pb.Action.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_starwars_quotes_Text(arg) {
  if (!(arg instanceof starwars_pb.Text)) {
    throw new Error('Expected argument of type starwars_quotes.Text');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_starwars_quotes_Text(buffer_arg) {
  return starwars_pb.Text.deserializeBinary(new Uint8Array(buffer_arg));
}


var StarwarsQuotesService = exports.StarwarsQuotesService = {
  quotes: {
    path: '/starwars_quotes.StarwarsQuotes/Quotes',
    requestStream: false,
    responseStream: true,
    requestType: starwars_pb.Action,
    responseType: starwars_pb.Text,
    requestSerialize: serialize_starwars_quotes_Action,
    requestDeserialize: deserialize_starwars_quotes_Action,
    responseSerialize: serialize_starwars_quotes_Text,
    responseDeserialize: deserialize_starwars_quotes_Text,
  },
};

exports.StarwarsQuotesClient = grpc.makeGenericClientConstructor(StarwarsQuotesService);
