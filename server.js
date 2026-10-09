const http = require('http');
const WebSocket = require('ws');
const os = require('os');

const wss = new WebSocket.Server({ port: 8080 });

wss.on('connection', (ws) => {
  ws.on('message', (message) => {
    console.log('Received: %s', message);
    // Handle incoming messages here
  });

  ws.send(`Hello from ${os.hostname()}`);
});

console.log('Server started on port 8080');
