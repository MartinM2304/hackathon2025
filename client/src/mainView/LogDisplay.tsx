import React, { useState, useEffect } from 'react';
import { SocketService } from '../ServerConnection/ServerConnection.ts';

interface LogDisplayProps {
  socketUrl: string;
}

const LogDisplay: React.FC<LogDisplayProps> = ({ socketUrl }) => {
  const [logs, setLogs] = useState<string[]>([]);
  let socketService: SocketService | null = null;

  useEffect(() => {
    socketService = new SocketService(socketUrl, (log) => {
      setLogs((prevLogs) => [...prevLogs, log]);
    });

    return () => {
      if (socketService) {
        socketService.disconnect();
      }
    };
  }, [socketUrl]);

  return (
    <div className="mt-4">
      <h2 className="mb-2 text-lg font-semibold">Logs:</h2>
      <div className="h-48 overflow-y-auto rounded-lg bg-gray-700 p-3">
        {logs.map((log, index) => (
          <p key={index} className="text-sm text-gray-300">
            {log}
          </p>
        ))}
      </div>
    </div>
  );
};

export default LogDisplay;