import { FormEvent } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import {
    ArrowRightOnRectangleIcon,
    ShieldCheckIcon,
} from '@heroicons/react/24/solid'
import useStore from '../store'
import { useMutateAuth } from '../hooks/useMutateAuth'
import { useMutateReservation } from '../hooks/useMutateReservation'
import { useQueryReservations } from '../hooks/useQueryReservations'
import { ReservationItem } from './ReservationItem'
import { useNavigate } from 'react-router-dom'

export const Reservation = () => {
	const queryClient = useQueryClient()
	const { editedReservation } = useStore()
	const updateReservation = useStore((state) => state.updateEditedReservation)
	const { data, isLoading } = useQueryReservations()
	const { createReservationMutation, updateReservationMutation } = useMutateReservation()
  const { logoutMutation } = useMutateAuth()
	const navi = useNavigate()
	const GoBack = () => {navi(-1)}
	const submitReservationHandler = (e: FormEvent<HTMLFormElement>) => {
		e.preventDefault()
		if (editedReservation.id === 0)
			createReservationMutation.mutate({
					purpose: editedReservation.purpose,
					start_time: editedReservation.start_time,
					end_time: editedReservation.end_time,
				})
		else{
				updateReservationMutation.mutate(editedReservation)
			}
		}

  const logout = async () => {
  	await logoutMutation.mutateAsync()
		queryClient.removeQueries(['reservations'])
	}
  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
			<div className="flex items-center my-3">
        <ShieldCheckIcon className="h-8 w-8 mr-3 text-indigo-500 cursor-pointer" />
        <span className="text-center text-3xl font-extrabold">
          Room Reservation
        </span>
      </div>

			<form onSubmit={submitReservationHandler}>
        <input
          className="mb-3 mr-3 px-3 py-2 border border-gray-300"
          placeholder="目的"
          type="text"
          onChange={(e) => updateReservation({ ...editedReservation, purpose: e.target.value })}
          value={editedReservation.purpose || ''}
        />
        <input
          className="mb-3 mr-3 px-3 py-2 border border-gray-300"
          placeholder="開始時間"
          type="text"
          onChange={(e) => updateReservation({ ...editedReservation, start_time: e.target.value })}
          value={editedReservation.start_time || ''}
        />
        <input
          className="mb-3 mr-3 px-3 py-2 border border-gray-300"
          placeholder="終了時間"
          type="text"
          onChange={(e) => updateReservation({ ...editedReservation, end_time: e.target.value })}
          value={editedReservation.end_time || ''}
        />
        <button
          className="disabled:opacity-40 mx-3 py-2 px-3 text-white bg-indigo-600 rounded"
          disabled={!editedReservation.purpose || !editedReservation.start_time || !editedReservation.end_time}
        >
          {editedReservation.id === 0 ? 'Create' : 'Update'}
        </button>
      </form>
			<button 
				className="disabled:opacity-40 py-1 px-4 rounded text-white bg-indigo-600"
				onClick={() => GoBack()}>戻る</button>
    </div>
  )
}
