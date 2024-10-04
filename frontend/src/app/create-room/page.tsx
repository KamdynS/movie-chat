'use client';

import { useUser } from "@clerk/nextjs";
import { useRouter } from 'next/navigation';
import CreateRoomForm from "@/components/CreateRoomForm";

export default function CreateRoom() {
  const { isSignedIn, isLoaded } = useUser();
  const router = useRouter();

  if (!isLoaded) {
    return <div>Loading...</div>;
  }

  if (!isSignedIn) {
    router.push('/');
    return null;
  }

  return (
    <div className="min-h-screen bg-slate-900 text-slate-100">
      <div className="container mx-auto py-8">
        <h1 className="text-3xl font-bold mb-6 text-blue-300">Create a New Room</h1>
        <div className="max-w-md mx-auto">
          <CreateRoomForm />
        </div>
      </div>
    </div>
  );
}
