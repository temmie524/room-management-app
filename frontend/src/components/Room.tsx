import React, { useEffect, useState } from 'react';
import axios from "axios";
import { useParams, useNavigate } from "react-router-dom";
import { Room } from '../types';
import { ArrowRightOnRectangleIcon } from '@heroicons/react/24/solid'
import { useMutateAuth } from '../hooks/useMutateAuth'

export const ShowRoom = () => {
    const {id} = useParams();
    const [room, setRoom] = useState<Room | null>(null);
    const navi = useNavigate()
    const { logoutMutation } = useMutateAuth()
    const logout = () => {
        logoutMutation.mutateAsync()
    }

    const GoBack = () => {navi(-1)}
	const GoReservations = () => {navi(`/reservations`)}

    useEffect(() => {
        axios.get<Room>(`${process.env.REACT_APP_API_URL}/rooms/${id}`).then((response) => {
            setRoom(response.data);
        });
}, []);

if (!room) {
    return <div>Loading...</div>;
}

  return (
	<>
		<div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
			<h1>Room Info	✨</h1>
			<h1>部屋番号:{room.room_num}</h1>
			<h1>館(別館OR本館):{room.building}</h1>
			
			<button 
				className="disabled:opacity-40 py-1 px-4 rounded text-white bg-indigo-600"
				onClick={() => GoReservations()}>予約に進む</button>
			<button 
				className="disabled:opacity-40 py-1 px-4 rounded text-white bg-indigo-600"
				onClick={() => GoBack()}>一覧に戻る</button>
            <ArrowRightOnRectangleIcon
                onClick={logout}
                className="h-6 w-6 text-blue-500 cursor-pointer"
            />
		</div>	
	</>

  );
}