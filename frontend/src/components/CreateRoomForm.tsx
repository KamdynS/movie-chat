import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from './ui/select';
import { createRoom } from '@/lib/api';

export default function CreateRoomForm() {
  const [movieName, setMovieName] = useState('');
  const [duration, setDuration] = useState('');
  const [platform, setPlatform] = useState('');
  const [error, setError] = useState('');
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    try {
      const room = await createRoom({ movieName, duration: parseInt(duration), platform });
      console.log('Room created successfully:', room);
      // Assuming the backend returns the room ID in the response
      router.push(`/chatroom/${room.id}`);
    } catch (err) {
      console.error('Error creating room:', err);
      setError('Failed to create room. Please try again.');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div>
        <label htmlFor="movieName" className="block text-sm font-medium text-slate-100">Movie Name</label>
        <Input
          id="movieName"
          value={movieName}
          onChange={(e) => setMovieName(e.target.value)}
          required
          className="mt-1"
        />
      </div>
      <div>
        <label htmlFor="duration" className="block text-sm font-medium text-slate-100">Duration (minutes)</label>
        <Input
          id="duration"
          type="number"
          value={duration}
          onChange={(e) => setDuration(e.target.value)}
          required
          className="mt-1"
        />
      </div>
      <div>
        <label htmlFor="platform" className="block text-sm font-medium text-slate-100">Platform</label>
        <Select onValueChange={setPlatform} required>
          <SelectTrigger className="mt-1">
            <SelectValue placeholder="Select a platform" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="Netflix">Netflix</SelectItem>
            <SelectItem value="Hulu">Hulu</SelectItem>
            <SelectItem value="Amazon Prime">Amazon Prime</SelectItem>
            <SelectItem value="Disney+">Disney+</SelectItem>
            <SelectItem value="HBO Max">HBO Max</SelectItem>
          </SelectContent>
        </Select>
      </div>
      {error && <p className="text-red-500">{error}</p>}
      <Button type="submit" className="w-full">Create Room</Button>
    </form>
  );
}
