import axios from 'axios';
import { useQuery } from '@tanstack/react-query';
import { Room } from '../types';
import { useError } from './useError';

export const useQueryRooms = () => {
	const { switchErrorHandling } = useError()
	const getRooms = async () => {
		const { data } = await axios.get<Room[]>(
			`${process.env.REACT_APP_API_URL}/rooms`,
			{ withCredentials: true }
		)
		return data;
	}
	return useQuery<Room[], Error>({
		queryKey: ['rooms'],
		queryFn: getRooms,
		onError: (err: any) => {
			if (err.response.data.message) {
				switchErrorHandling(err.response.data.message)
			}else{
				switchErrorHandling(err.response.data)
			}
		},
	})
}