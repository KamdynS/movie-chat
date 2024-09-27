import Link from 'next/link'
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"

const movieRooms = [
  { id: 1, title: "The Shawshank Redemption", year: 1994 },
  { id: 2, title: "The Godfather", year: 1972 },
  { id: 3, title: "The Dark Knight", year: 2008 },
  { id: 4, title: "12 Angry Men", year: 1957 },
  { id: 5, title: "Schindler's List", year: 1993 },
  { id: 6, title: "The Lord of the Rings: The Return of the King", year: 2003 },
]

export default function Home() {
  return (
    <div className="min-h-screen bg-slate-900 text-slate-100">
      <div className="container mx-auto py-8">
        <h1 className="text-3xl font-bold mb-6 text-blue-300">Movie Rooms</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {movieRooms.map((room) => (
            <Card key={room.id} className="bg-slate-800 border-slate-700 hover:border-blue-400 transition-colors">
              <CardHeader>
                <CardTitle className="text-blue-200">{room.title}</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-slate-400">Year: {room.year}</p>
              </CardContent>
              <CardFooter>
                <Link href={`/chatroom/${room.id}`} passHref>
                  <Button className="bg-purple-700 hover:bg-purple-600 text-slate-100">Join Room</Button>
                </Link>
              </CardFooter>
            </Card>
          ))}
        </div>
      </div>
    </div>
  )
}