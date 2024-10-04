'use client'

import { useState, useRef, useEffect } from 'react'
import { useUser } from "@clerk/nextjs";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { ScrollArea } from "@/components/ui/scroll-area"
import { createAuthenticatedWebSocket } from '@/utils/websocket';

interface Message {
  content: string
  room_id: string
  username: string
}

export default function Chatroom({ params }: { params: { id: string } }) {
  const [messages, setMessages] = useState<Message[]>([])
  const [inputMessage, setInputMessage] = useState('')
  const scrollAreaRef = useRef<HTMLDivElement>(null)
  const { user } = useUser();
  const [socket, setSocket] = useState<WebSocket | null>(null)

  useEffect(() => {
    const ws = createAuthenticatedWebSocket(`ws://localhost:8080/ws?roomId=${params.id}`);

    ws.onopen = () => {
      console.log('Connected to WebSocket');
    };

    ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      setMessages((prevMessages) => [...prevMessages, message]);
    };

    ws.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    ws.onclose = () => {
      console.log('Disconnected from WebSocket');
    };

    setSocket(ws);

    return () => {
      ws.close();
    };
  }, [params.id]);

  const handleSendMessage = () => {
    if (inputMessage.trim() && socket && socket.readyState === WebSocket.OPEN) {
      const message: Message = {
        content: inputMessage,
        room_id: params.id,
        username: user?.username || 'Anonymous',
      };
      socket.send(JSON.stringify(message));
      setInputMessage('');
    }
  };

  useEffect(() => {
    if (scrollAreaRef.current) {
      const scrollContainer = scrollAreaRef.current.querySelector('[data-radix-scroll-area-viewport]');
      if (scrollContainer) {
        scrollContainer.scrollTop = scrollContainer.scrollHeight;
      }
    }
  }, [messages]);

  return (
    <div className="min-h-screen bg-slate-900 flex items-center justify-center py-8">
      <Card className="w-full max-w-md bg-slate-800 border-slate-700">
        <CardHeader>
          <CardTitle className="text-blue-300">Chatroom {params.id}</CardTitle>
        </CardHeader>
        <CardContent>
          <ScrollArea className="h-[400px] w-full pr-4" ref={scrollAreaRef}>
            {messages.map((message, index) => (
              <div key={index} className={`mb-2 ${message.username === user?.username ? 'text-right' : 'text-left'}`}>
                <p className="text-xs text-slate-400 mb-1">{message.username}</p>
                <p className={`p-2 rounded-lg inline-block ${
                  message.username === user?.username 
                    ? 'bg-purple-600 text-white' 
                    : 'bg-slate-700 text-slate-300'
                }`}>
                  {message.content}
                </p>
              </div>
            ))}
          </ScrollArea>
        </CardContent>
        <CardFooter>
          <div className="flex w-full space-x-2">
            <Input
              type="text"
              placeholder="Type your message..."
              value={inputMessage}
              onChange={(e) => setInputMessage(e.target.value)}
              onKeyPress={(e) => e.key === 'Enter' && handleSendMessage()}
              className="bg-slate-700 text-slate-100 placeholder-slate-400 border-slate-600 focus:border-blue-400 focus:ring-blue-400"
            />
            <Button onClick={handleSendMessage} className="bg-purple-700 hover:bg-purple-600 text-slate-100">Send</Button>
          </div>
        </CardFooter>
      </Card>
    </div>
  )
}