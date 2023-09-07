import {FC, memo} from 'react'
import { PencilIcon, TrashIcon } from '@heroicons/react/24/solid'
import useStore  from '../store'
import {Reservation} from '../types'
import {useMutateReservation} from '../hooks/useMutateReservation'

const ReservationItemMemo: FC<Omit<Reservation, 'created_at' | 'updated_at'>> = ({
	id,
	purpose,
	start_time,
	end_time,
}) => {
	const updateReservation = useStore((state) => state.updateEditedReservation)
	const { deleteReservationMutation } = useMutateReservation()
	return (
		<li className="my-3">
			<span className="font-bold">{purpose}</span>
			<div className="flex float-right ml-20">
				<PencilIcon
					className="h-5 w-5 mx-1 text-blue-500 cursor-pointer"
					onClick={() => {
						updateReservation({
							id: id,
							purpose: purpose,
							start_time: start_time,
							end_time: end_time,
						})
					}}
				/>
				<TrashIcon
					className="h-5 w-5 text-blue-500 cursor-pointer"
					onClick={() => {
						deleteReservationMutation.mutate(id)
					}}
				/>
			</div>
	  </li>
	)
}
export const ReservationItem = memo(ReservationItemMemo)