import { useAuth } from "@clerk/nextjs";

export function createAuthenticatedWebSocket(url: string): WebSocket {
  const ws = new WebSocket(url);

  ws.onopen = async () => {
    try {
      const token = await getAuthToken();
      if (token) {
        ws.send(JSON.stringify({ type: 'authenticate', token }));
      }
    } catch (error) {
      console.error("Failed to authenticate WebSocket:", error);
    }
  };

  return ws;
}
