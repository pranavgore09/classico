var messages = require('./proto/starwars_pb');
var services = require('./proto/starwars_grpc_pb');
var starwars = require('starwars');
var grpc = require('grpc');

/**
 * Implements the GetQuote RPC method.
 */
function StreamQuotes(call) {
    var total = 10;
    interval = setInterval(function(){
      var reply = new messages.Text();
      reply.setText(starwars());
      call.write(reply);
      total --;
      if (total <= 0) {
        clearInterval(interval);
        call.end();
      }
    }, 10000);
}

/**
 * Starts an RPC server that receives requests for the Greeter service at the
 * sample server port
 */
function main() {
  var server = new grpc.Server();
  server.addService(services.StarwarsQuotesService, {quotes: StreamQuotes});
  server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure());
  console.log("starting server on 0.0.0.0:50051")
  server.start();
}

main();