'use client';

import { useEffect, useState } from 'react'
import Link from 'next/link'
import { Card, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"

interface Room {
  id: string;
  name: string;
  // Add other properties as needed
}

export default function Home() {
  const [rooms, setRooms] = useState<Room[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    fetchRooms()
  }, [])

  const fetchRooms = async () => {
    try {
      setLoading(true)
      const response = await fetch('/rooms')
      if (!response.ok) {
        throw new Error('Failed to fetch rooms')
      }
      const data = await response.json()
      setRooms(data)
    } catch (err) {
      console.error('Error fetching rooms:', err)
      setError('Failed to load rooms. Please try again later.')
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return <div className="min-h-screen bg-slate-900 text-slate-100 flex items-center justify-center">Loading...</div>
  }

  if (error) {
    return <div className="min-h-screen bg-slate-900 text-slate-100 flex items-center justify-center">{error}</div>
  }

  return (
    <div className="flex flex-col min-h-[100dvh] bg-slate-900 text-slate-100">
      <div className="min-h-screen bg-slate-900 text-slate-100">
        <div className="container mx-auto py-8">
          <h1 className="text-3xl font-bold mb-6 text-blue-300">Movie Rooms</h1>
          {rooms.length === 0 ? (
            <p className="text-center text-slate-400">No rooms available. Create a new room to get started!</p>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {rooms.map((room) => (
                <Card key={room.id} className="bg-slate-800 border-slate-700 hover:border-blue-400 transition-colors">
                  <CardHeader>
                    <CardTitle className="text-blue-200">{room.name}</CardTitle>
                  </CardHeader>
                  <CardFooter>
                    <Link href={`/chatroom/${room.id}`} passHref>
                      <Button className="bg-purple-700 hover:bg-purple-600 text-slate-100">Join Room</Button>
                    </Link>
                  </CardFooter>
                </Card>
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  )
}