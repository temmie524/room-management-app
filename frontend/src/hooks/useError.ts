import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { CsrfToken } from '../types';
import useStore from '../store';

export const useError = () => {
    const navigate = useNavigate();
    const resetEditedReservation = useStore((state) => state.resetEditedReservation);
    const getCsrfToken = async () => {
        const { data } = await axios.get<CsrfToken>(
            `${process.env.REACT_APP_API_URL}/csrf`
        )
        axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    const switchErrorHandling = (msg: string) => {
			switch (msg) {
				case 'invalid csrf token':
					getCsrfToken()
					alert('セッションが切れました。再度操作してください。')
					break
				case 'invalid or expired jwt':
					alert('access token expired, please login')
					resetEditedReservation()
					navigate('/')
					break
				case 'missing or malformed jwt':
					alert('access token is not valid, please login')
					resetEditedReservation()
					navigate('/')
					break
				case 'duplicated key not allowed':
					alert('email already exist, please use another one')
					break
				case 'crypto/bcrypt: hashedPassword is not the hash of the given password':
					alert('password is not correct')
					break
				case 'record not found':
					alert('email is not correct')
					break
			  default:
					alert(msg)
			}
		}

		return { switchErrorHandling }
}