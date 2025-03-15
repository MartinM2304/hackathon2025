import { io, Socket } from 'socket.io-client';

export class SocketService {
  private socket: Socket | null = null;
  private onLogReceived: (log: string) => void;

  constructor(serverUrl: string, onLogReceived: (log: string) => void) {
    this.onLogReceived = onLogReceived;
    this.connect(serverUrl);
  }

  private connect(serverUrl: string) {
    this.socket = io(serverUrl);

    this.socket.on('connect', () => {
      console.log('Socket.IO connected');
    });

    this.socket.on('log', (log: string) => {
      this.onLogReceived(log);
    });

    this.socket.on('disconnect', () => {
      console.log('Socket.IO disconnected');
    });

    this.socket.on('connect_error', (err) => {
      console.error('Socket.IO connection error:', err);
    });
  }

  public disconnect() {
    if (this.socket) {
      this.socket.disconnect();
      this.socket = null;
    }
  }
}