import 'dart:convert';

import 'package:websocket/websocket.dart' as websocket;

void main(List<String> arguments) {
  // Create the WebSocket channel
  final channel = WebSocketChannel.connect(
    Uri.parse('wss://ws-feed.pro.coinbase.com'),
  );

  channel.sink.add(
    jsonEncode(
      {
        "type": "subscribe",
        "channels": [
          {
            "name": "ticker",
            "product_ids": [
              "BTC-EUR",
            ]
          }
        ]
      },
    ),
  );

  // Listen for all incoming data
  channel.stream.listen(
        (data) {
      print(data);
    },
    onError: (error) => print(error),
  );
}
