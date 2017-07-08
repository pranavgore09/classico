var messages = require('./daily_pb');
var services = require('./daily_grpc_pb');
var starwars = require('starwars');
var grpc = require('grpc');

/**
 * Implements the GetQuote RPC method.
 */
function GetQuote(call, callback) {
    console.log("okay this is the thing")
  var reply = new messages.Text();
  reply.setText(starwars());
  callback(null, reply);
}

/**
 * Starts an RPC server that receives requests for the Greeter service at the
 * sample server port
 */
function main() {
  var server = new grpc.Server();
  server.addService(services.StarwarsService, {getQuote: GetQuote});
  server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure());
  console.log("starting server on 0.0.0.0:50051")
  server.start();
}

main();