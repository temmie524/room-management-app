import axios from 'axios';
import { useQuery } from '@tanstack/react-query';
import { Reservation } from '../types';
import { useError } from './useError';

export const useQueryReservations = () => {
	const { switchErrorHandling } = useError()
	const getReservations = async () => {
		const { data } = await axios.get<Reservation[]>(
			`${process.env.REACT_APP_API_URL}/reservations`,
			{ withCredentials: true }
		)
		return data;
	}
	return useQuery<Reservation[], Error>({
		queryKey: ['reservations'],
		queryFn: getReservations,
		onError: (err: any) => {
			if (err.response.data.message) {
				switchErrorHandling(err.response.data.message)
			}else{
				switchErrorHandling(err.response.data)
			}
		},
	})
}