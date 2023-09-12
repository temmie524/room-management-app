import { useEffect } from 'react';
import { Auth } from './components/Auth';
import { Reservation } from './components/Reservation'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import axios from 'axios'
import { CsrfToken } from './types';
import { Rooms }  from './components/Rooms';
import { ShowRoom }  from './components/Room';
import { NotFound } from './components/NotFound';
import { MyPage } from './components/MyPage';

function App() {
	useEffect(() => {
		axios.defaults.withCredentials = true
		const getCsrfToken = async () => {
			const  { data } = await axios.get<CsrfToken>(
					`${process.env.REACT_APP_API_URL}/csrf`
			)
			axios.defaults.headers['X-CSRF-TOKEN'] = data.csrf_token
		}
		getCsrfToken()
	},[])
  return (
    <BrowserRouter>
			<Routes>
				<Route path="/" element={<Auth />} />
				<Route path="/reservations" element={<Reservation />} />
				<Route path="/rooms" element={<Rooms />} />
				<Route path="/rooms/:id" element={<ShowRoom />} />
				<Route path="/mypage" element={<MyPage />} />
				<Route path="*" element={<NotFound />} />
			</Routes>
    </BrowserRouter>
  );
}

export default App;

