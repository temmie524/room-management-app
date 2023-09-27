import { create } from 'zustand'

type EditedReservation = {
    id: number
    purpose: string
    start_time: string
    end_time: string
}

type State = {
    editedReservation: EditedReservation
    updateEditedReservation: (payload: EditedReservation) => void
    resetEditedReservation: () => void
}

const useStore = create<State>((set) => ({
    editedReservation: {
        id: 0,
        purpose: '',
        start_time: '',
        end_time: '',
    },
    updateEditedReservation: (payload) =>
        set({
            editedReservation: payload,
        }),
        resetEditedReservation: () =>
        set({
            editedReservation: {
                id: 0,
                purpose: '',
                start_time: '',
                end_time: '',
            }}),
}))

export default useStore