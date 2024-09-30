## Guide Page Plans

### 1. JSON Files
Okay so basically we have json files in public/guide. 

prelimenary_guide_data.json - This is the main data file that contains the schedule for all channels. 

prelimenary_show_data.json - This file contains the shows that are available to watch. 

These json files have a set list of shows to play at certain times. When a time has a null value (""), that means that we can choose any random show and random episode. 

The format for the prelimenary_guide_data.json is as follows: 

```
{
    "channels": [
        {
         "id": 1,
        "name": "Sitcoms",
        "schedule": [
          {
            "day": [times starting from 00:00 to 23:30]
            "Sunday": ["1", "1", "2", "2", "3", "3", "4", "4", "5", "5", "6", "6", "7", "7", "8", "8"],
            "Monday": ["9", "9", "10", "10", "11", "11", "12", "12", "1", "1", "2", "2", "3", "3", "4", "4"],
            "Tuesday": ["5", "5", "6", "6", "7", "7", "8", "8", "9", "9", "10", "10", "11", "11", "12", "12"],
            "Wednesday": ["1", "1", "2", "2", "3", "3", "4", "4", "5", "5", "6", "6", "7", "7", "8", "8"],
            "Thursday": ["9", "9", "10", "10", "11", "11", "12", "12", "1", "1", "2", "2", "3", "3", "4", "4"],
            "Friday": ["5", "5", "6", "6", "7", "7", "8", "8", "9", "9", "10", "10", "11", "11", "12", "12"],
            "Saturday": ["1", "1", "2", "2", "3", "3", "4", "4", "5", "5", "6", "6", "7", "7", "8", "8"]
          }
        ]
      },
    ]
}
```

### 2. Page Design


Mock implementation of the guide page: 

This will need to be rewritten to suit our needs better, but this is how AI thinks we should implement the guide page. 

```
'use client';

import React, { useState, useEffect } from 'react'
import { format, addMinutes, setMinutes, startOfHour, parseISO } from 'date-fns'
import Link from 'next/link'
import { Button } from "@/components/ui/button"

interface Show {
  id: number;
  title: string;
  description: string;
  duration: number;
  isMovie: boolean;
  episode: number;
  season: number;
  channelId: number;
}

interface Channel {
  id: number;
  name: string;
  schedule: {
    [day: string]: string[];
  }[];
}

interface TVGuideData {
  channels: Channel[];
  shows: Show[];
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
      <p className="text-xs text-blue-400">S{show.season} E{show.episode}</p>
      {isHovered && (
        <div className="absolute z-10 bg-slate-900 p-4 rounded shadow-lg w-64 border border-blue-400">
          <h4 className="font-bold text-blue-300 mb-2">{show.title}</h4>
          <p className="text-sm mb-2">Season: {show.season}, Episode: {show.episode}</p>
          <p className="text-sm mb-2">Duration: {show.duration} minutes</p>
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

const fetchTVGuideData = async (): Promise<TVGuideData> => {
  // In a real application, you would fetch this data from your API
  // For now, we'll simulate an API call with local imports
  const channelsData = await import('./public/guide/prelimenary_guide_data.json');
  const showsData = await import('./public/guide/prelimenary_show_data.json');
  return {
    channels: channelsData.channels,
    shows: showsData.shows
  };
}

const getDayName = (date: Date): string => {
  return format(date, 'EEEE');
}

export default function TVGuide() {
  const [currentTime, setCurrentTime] = useState(new Date())
  const [guideData, setGuideData] = useState<TVGuideData | null>(null)
  const totalSlots = 16; // 8 hours, 2 slots per hour as per your JSON structure

  useEffect(() => {
    const timer = setInterval(() => setCurrentTime(new Date()), 60000)
    return () => clearInterval(timer)
  }, [])

  useEffect(() => {
    fetchTVGuideData().then(setGuideData)
  }, [])

  const getNearestPastHalfHour = (date: Date) => {
    const minutes = date.getMinutes()
    return minutes < 30 ? startOfHour(date) : setMinutes(date, 30)
  }

  const startTime = getNearestPastHalfHour(currentTime)
  const timeSlots = Array.from({ length: totalSlots }, (_, i) => addMinutes(startTime, i * 30))

  if (!guideData) return <div>Loading...</div>

  const currentDay = getDayName(currentTime);

  return (
    <div className="min-h-screen bg-slate-900 text-slate-100 p-4">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold text-blue-300">TV Guide</h1>
        <p className="text-xl">{format(currentTime, 'h:mm a')}</p>
      </div>
      <div className="overflow-x-auto">
        <div className="grid grid-cols-[auto,repeat(16,minmax(100px,1fr))] min-w-max gap-px bg-slate-700">
          <div className="bg-slate-800 p-2 font-bold sticky left-0 z-10">Channel</div>
          {timeSlots.map((time, index) => (
            <div key={index} className="bg-slate-800 p-2 font-bold">
              {format(time, 'h:mm a')}
            </div>
          ))}

          {guideData.channels.map((channel) => (
            <React.Fragment key={`channel-${channel.id}`}>
              <div className="bg-slate-800 p-2 font-bold sticky left-0 z-10">
                {channel.name}
              </div>
              {channel.schedule[0][currentDay].map((showId, index) => {
                const show = guideData.shows.find(s => s.id === parseInt(showId));
                if (!show) return null;
                return (
                  <TVGuideCard 
                    key={`show-${channel.id}-${index}`} 
                    show={show} 
                    durationInSlots={show.duration / 30}
                  />
                );
              })}
            </React.Fragment>
          ))}
        </div>
      </div>
    </div>
  )
}
```