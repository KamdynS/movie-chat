'use client';

import React, { useState, useEffect } from 'react'
import { format, addMinutes, setMinutes, startOfHour } from 'date-fns'
import Link from 'next/link'
import { Button } from "@/components/ui/button"

interface Show {
  id: number;
  title: string;
  platforms: string[];
  viewers: number;
  duration: number;
  startSlot: number;
  year: number;
  description: string;
}

interface TVGuideCardProps {
  show: Show;
  durationInSlots: number;
}

const TVGuideCard: React.FC<TVGuideCardProps> = ({ show, durationInSlots }) => {
  const [isHovered, setIsHovered] = useState(false);

  return (
    <div 
      className={`border-r border-b border-slate-700 p-2 bg-slate-800 text-slate-100 relative`}
      style={{ gridColumn: `span ${durationInSlots}` }}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      <h3 className="font-bold truncate">{show.title}</h3>
      <p className="text-xs text-blue-400">{show.platforms[0]}</p>
      <p className="text-xs text-slate-400">Viewers: {show.viewers}</p>
      {isHovered && (
        <div className="absolute z-10 bg-slate-900 p-4 rounded shadow-lg w-64 border border-blue-400">
          <h4 className="font-bold text-blue-300 mb-2">{show.title}</h4>
          <p className="text-sm mb-2">Year: {show.year}</p>
          <p className="text-sm mb-2">Duration: {show.duration} minutes</p>
          <p className="text-sm mb-2">Platforms: {show.platforms.join(', ')}</p>
          <p className="text-xs text-slate-400 mb-4">{show.description}</p>
          <Link href={`/chatroom/${show.id}`} passHref>
            <Button className="w-full bg-purple-700 hover:bg-purple-600 text-slate-100">
              Join Chatroom
            </Button>
          </Link>
        </div>
      )}
    </div>
  )
}

const generateDummyData = (totalSlots: number) => {
  const channels = ['NBC', 'ABC', 'CBS', 'FOX', 'HBO', 'Netflix', 'Hulu', 'Disney+']
  const shows = [
    { title: 'Stranger Things', possibleDurations: [30, 60, 90], year: 2016, description: "A sci-fi horror drama set in the 1980s." },
    { title: 'The Crown', possibleDurations: [30, 60, 90], year: 2016, description: "A historical drama about the British royal family." },
    { title: 'Game of Thrones', possibleDurations: [60, 90], year: 2011, description: "An epic fantasy series based on George R.R. Martin's novels." },
    { title: 'Breaking Bad', possibleDurations: [30, 60], year: 2008, description: "A chemistry teacher turns to a life of crime." },
    { title: 'Friends', possibleDurations: [30], year: 1994, description: "Six friends navigate life and love in New York City." },
    { title: 'The Office', possibleDurations: [30], year: 2005, description: "A mockumentary about life in a paper company office." },
    { title: 'The Mandalorian', possibleDurations: [30, 60], year: 2019, description: "A lone bounty hunter's adventures in the Star Wars galaxy." },
    { title: 'Black Mirror', possibleDurations: [60, 90], year: 2011, description: "An anthology series exploring the dark side of technology." },
    { title: 'Westworld', possibleDurations: [60], year: 2016, description: "A sci-fi western set in a high-tech amusement park." },
    { title: 'The Witcher', possibleDurations: [60], year: 2019, description: "A fantasy series based on Polish novels and short stories." }
  ]
  const platforms = ['Netflix', 'Hulu', 'Amazon Prime', 'Disney+', 'HBO Max']

  return channels.map(channel => {
    let schedule = [];
    let timeSlot = 0;
    let showId = 1;

    while (timeSlot < totalSlots) {
      const show = shows[Math.floor(Math.random() * shows.length)];
      const duration = show.possibleDurations[Math.floor(Math.random() * show.possibleDurations.length)];
      const durationInSlots = duration / 30;

      if (timeSlot + durationInSlots > totalSlots) {
        schedule.push({
          id: showId++,
          title: 'Short Program',
          platforms: [platforms[Math.floor(Math.random() * platforms.length)]],
          viewers: Math.floor(Math.random() * 1000),
          duration: 30,
          startSlot: timeSlot,
          year: 2023,
          description: "A short filler program."
        });
        timeSlot += 1;
      } else {
        schedule.push({
          id: showId++,
          title: show.title,
          platforms: [platforms[Math.floor(Math.random() * platforms.length)]],
          viewers: Math.floor(Math.random() * 1000),
          duration: duration,
          startSlot: timeSlot,
          year: show.year,
          description: show.description
        });
        timeSlot += durationInSlots;
      }
    }

    return { name: channel, schedule };
  });
}

export default function TVGuide() {
  const [currentTime, setCurrentTime] = useState(new Date())
  interface Channel {
    name: string;
    schedule: Show[];
  }

  const [guideData, setGuideData] = useState<Channel[]>([])
  const totalSlots = 24 * 2; // 24 hours, 2 slots per hour

  useEffect(() => {
    const timer = setInterval(() => setCurrentTime(new Date()), 60000)
    return () => clearInterval(timer)
  }, [])

  useEffect(() => {
    setGuideData(generateDummyData(totalSlots))
  }, [])

  const getNearestPastHalfHour = (date: Date) => {
    const minutes = date.getMinutes()
    return minutes < 30 ? startOfHour(date) : setMinutes(date, 30)
  }

  const startTime = getNearestPastHalfHour(currentTime)
  const timeSlots = Array.from({ length: totalSlots }, (_, i) => addMinutes(startTime, i * 30))

  return (
    <div className="min-h-screen bg-slate-900 text-slate-100 p-4">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold text-blue-300">TV Guide</h1>
        <p className="text-xl">{format(currentTime, 'h:mm a')}</p>
      </div>
      <div className="overflow-x-auto">
        <div className="grid grid-cols-[auto,repeat(48,minmax(100px,1fr))] min-w-max gap-px bg-slate-700">
          <div className="bg-slate-800 p-2 font-bold sticky left-0 z-10">Channel</div>
          {timeSlots.map((time, index) => (
            <div key={index} className="bg-slate-800 p-2 font-bold">
              {format(time, 'h:mm a')}
            </div>
          ))}

          {guideData.map((channel, channelIndex) => (
            <React.Fragment key={`channel-${channelIndex}`}>
              <div className="bg-slate-800 p-2 font-bold sticky left-0 z-10">
                {channel.name}
              </div>
              {channel.schedule.map((show, showIndex) => (
                <TVGuideCard 
                  key={`show-${channelIndex}-${showIndex}`} 
                  show={show} 
                  durationInSlots={show.duration / 30}
                />
              ))}
            </React.Fragment>
          ))}
        </div>
      </div>
    </div>
  )
}