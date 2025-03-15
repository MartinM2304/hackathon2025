import React, { useState, useEffect } from 'react';

interface LogDisplayProps {
  socketUrl: string;
}

const LogDisplay: React.FC<LogDisplayProps> = ({ socketUrl }) => {
  const [logs, setLogs] = useState<string[]>([]);

  useEffect(() => {
    const socket = new WebSocket(socketUrl);

    socket.onopen = () => {
      console.log('WebSocket connected');
    };

    socket.onmessage = (event) => {
      const log = event.data as string;
      console.log('WebSocket message received:', log);
      setLogs((prevLogs) => [...prevLogs, log]);
    };

    socket.onclose = () => {
      console.log('WebSocket disconnected');
    };

    socket.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    return () => {
      socket.close();
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