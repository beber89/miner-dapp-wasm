const net = require('net');
const port = 8081;
const host = '127.0.0.1';

const server = net.createServer();
server.listen(port, host, () => {
    console.log('TCP Server is running on port ' + port + '.');
});

let sockets = [];

server.on('connection', function(sock) {
    console.log('CONNECTED: ' + sock.remoteAddress + ':' + sock.remotePort);
    sockets.push(sock);

    sock.on('data', function(data) {
        console.log('DATA ' + sock.remoteAddress + ': ' + data);
        let cmp = Buffer.compare(data, Buffer.from('Connect\n'));
        // Write the data back to all the connected, the client will receive it as data from the server
        sockets.forEach(function(s, index, array) {
            if (cmp != 0 && s!= sock) {
                console.log('send data to ' + s.remotePort + ': ' + data);
                s.write(data+'\n');
                // s.write(s.remoteAddress + ':' + s.remotePort + " said " + data + '\n');
            }
        });
    });

    // Add a 'close' event handler to this instance of socket
    sock.on('close', function(data) {
        let index = sockets.findIndex(function(o) {
            return o.remoteAddress === sock.remoteAddress && o.remotePort === sock.remotePort;
        })
        if (index !== -1) sockets.splice(index, 1);
        console.log('CLOSED: ' + sock.remoteAddress + ' ' + sock.remotePort);
    });
});