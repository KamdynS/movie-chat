'use client'

import { useRef } from 'react'
import Link from 'next/link'
import { Card, CardContent } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { ChevronLeft, ChevronRight } from 'lucide-react'

interface Movie {
    id: number;
    title: string;
    year: number;
}

interface Category {
    name: string;
    movies: Movie[];
}

// Dummy data for movie categories
const categories = [
  {
    name: "Trending Now",
    movies: Array.from({ length: 20 }, (_, i) => ({ id: i + 1, title: `Trending Movie ${i + 1}`, year: 2023 }))
  },
  {
    name: "Action & Adventure",
    movies: Array.from({ length: 20 }, (_, i) => ({ id: i + 21, title: `Action Movie ${i + 1}`, year: 2022 }))
  },
  {
    name: "Comedies",
    movies: Array.from({ length: 20 }, (_, i) => ({ id: i + 41, title: `Comedy ${i + 1}`, year: 2021 }))
  },
  {
    name: "Sci-Fi & Fantasy",
    movies: Array.from({ length: 20 }, (_, i) => ({ id: i + 61, title: `Sci-Fi Movie ${i + 1}`, year: 2020 }))
  },
  {
    name: "Drama",
    movies: Array.from({ length: 20 }, (_, i) => ({ id: i + 81, title: `Drama ${i + 1}`, year: 2019 }))
  }
]

function MovieRow({ category }: { category: Category }) {
  const rowRef = useRef<HTMLDivElement>(null)

  const scroll = (direction: string) => {
    if (rowRef.current) {
      const { current } = rowRef
      const scrollAmount = direction === 'left' ? -current.offsetWidth : current.offsetWidth
      current.scrollBy({ left: scrollAmount, behavior: 'smooth' })
    }
  }

  return (
    <div className="mb-8">
      <h2 className="text-2xl font-bold mb-4 text-slate-100">{category.name}</h2>
      <div className="relative">
        <Button 
          variant="outline" 
          size="icon" 
          className="absolute left-0 top-1/2 -translate-y-1/2 z-10 bg-slate-800/50"
          onClick={() => scroll('left')}
        >
          <ChevronLeft className="h-4 w-4" />
        </Button>
        <div 
          ref={rowRef} 
          className="flex overflow-x-scroll space-x-4 pb-4 hide-scrollbar"
        >
          {category.movies.map((movie) => (
            <Card key={movie.id} className="flex-none w-40 bg-slate-800 border-slate-700 hover:border-blue-400 transition-colors">
              <CardContent className="p-4">
                <Link href={`/chatroom/${movie.id}`}>
                  <div className="aspect-[2/3] bg-slate-700 mb-2 rounded"></div>
                  <h3 className="text-sm font-medium text-blue-200 truncate">{movie.title}</h3>
                  <p className="text-xs text-slate-400">{movie.year}</p>
                </Link>
              </CardContent>
            </Card>
          ))}
        </div>
        <Button 
          variant="outline" 
          size="icon" 
          className="absolute right-0 top-1/2 -translate-y-1/2 z-10 bg-slate-800/50"
          onClick={() => scroll('right')}
        >
          <ChevronRight className="h-4 w-4" />
        </Button>
      </div>
    </div>
  )
}

export default function Guide() {
    return (
      <div className="bg-slate-900 text-slate-100 pb-8">
        <div className="container mx-auto px-4">
          <h1 className="text-3xl font-bold pt-8 mb-8 text-blue-300">Movie Guide</h1>
          {categories.map((category, index) => (
            <MovieRow key={index} category={category} />
          ))}
        </div>
      </div>
    )
  }