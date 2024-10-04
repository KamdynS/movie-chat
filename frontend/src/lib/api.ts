interface CreateRoomParams {
  movieName: string;
  duration: number;
  platform: string;
}

export async function createRoom(params: CreateRoomParams) {
  const response = await fetch('localhost:8080/ws/createRoom', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params),
  });

  if (!response.ok) {
    console.error('Error response:', await response.text());
    throw new Error('Failed to create room');
  }

  return response.json();
}
