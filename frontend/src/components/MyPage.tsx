import React, { useEffect, useState } from 'react';
import axios from "axios";
import { useParams, useNavigate } from "react-router-dom";
import { User } from '../types';

export const MyPage = () => {
    const {id} = useParams();
    const baseUrl = process.env.REACT_APP_BE_URL;
    const [user, setUser] = useState<User | null>(null);
    const navi = useNavigate();

    const GoBack = () => {navi(-1)}
		const GoRooms = () => {navi(`/rooms`)}

    useEffect(() => {
        axios.get<User>(`${baseUrl}/users/${id}`).then((response) => {
            setUser(response.data);
        });
}, []);

if (!user) {
    return <div>Loading...</div>;
}

  return (
	<>
		<div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
			<h1>User Info	✨</h1>
			<h1>YourID:{user.id}</h1>
			<h1>LASTNAME:{user.last_name}</h1>
			<h1>FIRSTNAME:{user.first_name}</h1>
			<h1>AGE:{user.age}</h1>
			<h1>ROLE:{user.role}</h1>
			<h1>IDNUMBER:{user.id_number}</h1>
			<button 
				className="disabled:opacity-40 py-1 px-4 rounded text-white bg-indigo-600"
				onClick={() => GoRooms()}>予約に進む</button>
			<button 
				className="disabled:opacity-40 py-1 px-4 rounded text-white bg-indigo-600"
				onClick={() => GoBack()}>一覧に戻る</button>
		</div>	
	</>

  );
}
