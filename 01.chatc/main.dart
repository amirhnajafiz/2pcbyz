// ignore_for_file: avoid_print
import 'dart:io';
import 'package:web_socket_client/web_socket_client.dart';
import 'package:dotenv/dotenv.dart';

void main() async {
  var env = DotEnv(includePlatformEnvironment: true)..load();
  
  // Create a WebSocket client.
  final uri = Uri.parse(env['SERVER_ADDRESS']);
  const backoff = ConstantBackoff(Duration(seconds: 1));
  final socket = WebSocket(uri, backoff: backoff);

  // Listen for changes in the connection state.
  socket.connection.listen((state) => print('state: "$state"'));

  // Listen for incoming messages.
  socket.messages.listen((message) {
    print('message: "$message"');

    // Send a message to the server.
    socket.send('ping');
  });

  await Future<void>.delayed(const Duration(seconds: 3));

  // Close the connection.
  socket.close();
}
