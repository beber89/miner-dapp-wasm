const port = 8081;
const host = '127.0.0.1';

var WebSocketServer = require('websocket').server;
var http = require('http');

var server = http.createServer(function(request, response) {
  // process HTTP request. Since we're writing just WebSockets
  // server we don't have to implement anything.
});

server.listen(port, host, () => {
    console.log('WS Server is running on port ' + port + '.');
});

// create the server
wsServer = new WebSocketServer({
  httpServer: server
});

let sockets = [];


wsServer.on('request', function(request) {
    var connection = request.accept(null, request.origin);
    sockets.push(connection)
    console.log("Connection with node initiated")
    // This is the most important callback for us, we'll handle
    // all messages from users here.
    connection.on('message', function(message) {
      console.log(sockets.length)
      console.log("message received "+message)
        sockets.forEach(function(s, index, array) {
          if (message.type === 'utf8') {
            broadcastData = message.utf8Data
            console.log("it is utf8: "+ broadcastData)
            if(s!= connection) {
              console.log('send data to ' + s.socket.remotePort + ': ' + broadcastData);
              s.sendUTF(broadcastData)
            }
          }
        });
    });
  
    connection.on('close', function(connection) {
      // close user connection
    });
  });