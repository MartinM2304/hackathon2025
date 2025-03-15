export class SocketService {
    private socket: WebSocket | null = null;
    private onLogReceived: (log: string) => void;
  
    constructor(serverUrl: string, onLogReceived: (log: string) => void) {
      this.onLogReceived = onLogReceived;
      this.connect(serverUrl);
    }
  
    private connect(serverUrl: string) {
      this.socket = new WebSocket(serverUrl);
  
      this.socket.onopen = () => {
        console.log('WebSocket connected');
      };
  
      this.socket.onmessage = (event) => {
        const log = event.data as string;
        console.log('WebSocket message received:', log);
        this.onLogReceived(log);
      };
  
      this.socket.onclose = () => {
        console.log('WebSocket disconnected');
      };
  
      this.socket.onerror = (error) => {
        console.error('WebSocket error:', error);
      };
    }
  
    public disconnect() {
      if (this.socket) {
        this.socket.close();
        this.socket = null;
      }
    }
  }