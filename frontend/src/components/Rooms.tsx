import React, { useEffect, useState } from 'react';
import axios from "axios";
import {useNavigate} from "react-router-dom";
import {useMutateAuth} from "../hooks/useMutateAuth";
import { Room } from '../types';
import {
	ArrowRightOnRectangleIcon,
} from '@heroicons/react/24/solid'

export const Rooms = () => {
	const baseUrl = process.env.REACT_APP_BE_URL;
	const [data, setData] = useState<Room[] | null>(null);
	const navi = useNavigate();
	const { logoutMutation } = useMutateAuth()
	const logout = () => {
			logoutMutation.mutateAsync()
	}

	//Page遷移の処理
	const handleClick = (id: any) => {
			navi(`/rooms/${id}`)
	}
	const GoTop = () =>{
			navi(`/rooms`)
	}

  //Roomsの配列を取得
	useEffect(() => {
		axios.get<Room[]>('http://localhost:8080/rooms').then((response) => {
				setData(response.data);
		})
	}, [baseUrl]);

  //error
	if (!data) {
		return <div>Loading...</div>;
	}

    //受け取った配列を取り出す
	const roomList = data.map((room: any) => {
		return (
			<li key={room.id}>
				<button 
					onClick={() => handleClick(room.id)}
					className="bg-blue-500 text-white px-2 py-1 rounded-md hover:bg-blue-600 focus:-none focus:ring focus:ring-blue-300"
					>詳細</button>
				<a>{room.building}</a>
				<a>{room.room_num}</a>
			</li>
			
		);
	});

	//表示
	return (
		<>
			<div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
				<ArrowRightOnRectangleIcon
					onClick={logout}
					className="h-6 w-6 text-blue-500 cursor-pointer"
				/>
				<h1 className="text-center text-3xl font-extrabold">部屋リスト</h1>
		
				<ul>
					{roomList}
				</ul>
				<button
					onClick={GoTop}
					className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 focus:-none focus:ring focus:ring-blue-300"
				>
					Topに戻る
				</button>
			</div>
		</>

	)
}
