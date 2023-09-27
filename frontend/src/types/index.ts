export type Reservation = {
	id: number
	purpose: string
	start_time: string
	end_time: string
    user_id: number
	room_id: number
	created_at: Date
    updated_at: Date
}

export type Room = {
    id: number
    room_num: string
    building: number
    created_at: Date
    updated_at: Date
}

export type User = {
    id: number
    last_name: string
    first_name: string
    email: string
    password: string
    age: number
    role: string
    id_number: string
    created_at: Date
    updated_at: Date
}

export type CsrfToken = {
    csrf_token: string
}

export type Credential = {
    email: string
    password: string
}
