const http = require('http');

const hostname = '127.0.0.1';
const port = 8081;

const server = http.createServer((req, res) => {
  console.log(`received request`);
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/html');
  res.setHeader("Access-Control-Allow-Origin", "*");
  console.log(`writing response`);
  res.write('<html><body><p>This is home Page.</p></body></html>');
  res.end();
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});