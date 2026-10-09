const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 8080 });

wss.on('connection', (ws) => {
  ws.on('message', (message) => {
    console.log('Received: %s', message);
    // Handle incoming messages here
  });

  ws.send('Hey there!');
});
