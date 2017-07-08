
var starwarsMessages = require('./services_pb/starwars_pb');
var starwarsServices = require('./services_pb/starwars_grpc_pb');

var grpc = require('grpc');

var client = new starwarsServices.StarwarsQuotesClient('localhost:50051', grpc.credentials.createInsecure());
var request = new starwarsMessages.Action;

request.setStart(1);

var call = client.quotes(request);

call.on('data', function(text){
    console.log(text.getText());
});
call.on('end', function() {
    // The server has finished sending
});