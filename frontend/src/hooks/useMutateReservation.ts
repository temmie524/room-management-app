import axios from 'axios';
import { useQueryClient, useMutation } from '@tanstack/react-query';
import { Reservation } from '../types';
import useStore from '../store';
import { useError } from './useError';

export const useMutateReservation = () => {
	const queryClient = useQueryClient();
	const { switchErrorHandling }= useError();
	const resetEditedReservation = useStore((state) => state.resetEditedReservation)

	const createReservationMutation = useMutation(
		(reservation: Omit<Reservation, 'id' | 'user_id' | 'room_id' | 'created_at' | 'updated_at'>) =>
			axios.post(`${process.env.REACT_APP_API_URL}/reservations/new`, reservation),
		{
			onSuccess: (res) => {
				const previousReservations = queryClient.getQueryData<Reservation[]>(['reservations']);
				if (previousReservations) {
					queryClient.setQueryData(['reservations'], [...previousReservations, res.data]);
				}
				resetEditedReservation();
			},
			onError: (err: any) => {
				if (err.response.data.message) {
					switchErrorHandling(err.response.data.message)
				} else {
					switchErrorHandling(err.message.data)
				}
			},
		}
	)

	const updateReservationMutation = useMutation(
		(reservation: Omit<Reservation, 'user_id'|'room_id'|'created_at' | 'updated_at'>) =>
		  axios.put<Reservation>(`${process.env.REACT_APP_API_URL}/reservations/${reservation.id}`,
			{
				purpose: reservation.purpose,
				start_time: reservation.start_time,
				end_time: reservation.end_time
			}),
		{
			onSuccess: (res, variables) => {
				const previousReservations = queryClient.getQueryData<Reservation[]>(['reservations'])
				if (previousReservations) {
					queryClient.setQueryData<Reservation[]>(
						['reservations'],
						previousReservations.map((reservation) =>
							reservation.id === variables.id ? res.data : reservation
						)
					)
				}
				resetEditedReservation();
			},
			onError: (err: any) => {
				if (err.response.data.message) {
					switchErrorHandling(err.response.data.message)
				} else {
					switchErrorHandling(err.message.data)
				}
			},
		}
	)

	const deleteReservationMutation = useMutation(
		(id: number) =>
		  axios.delete(`${process.env.REACT_APP_API_URL}/reservations/${id}`),
		{
			onSuccess: (_, variables) => {
				const previousReservations = queryClient.getQueryData<Reservation[]>(['reservations'])
				if (previousReservations) {
					queryClient.setQueryData<Reservation[]>(
						['reservations'],
						previousReservations.filter((reservation) => reservation.id !== variables)
					)
				}
				resetEditedReservation()
			},
			onError: (err: any) => {
				if (err.response.data.message) {
					switchErrorHandling(err.response.data.message)
				} else {
					switchErrorHandling(err.message.data)
				}
			},
		}
	)
	return {
		createReservationMutation,
		updateReservationMutation,
		deleteReservationMutation,
	}
}