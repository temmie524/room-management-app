import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import useStore from "../store";
import { Credential } from "../types";
import { useError } from "./useError";

export const useMutateAuth = () => {
	const navigate = useNavigate();
	const resetEditedReservation = useStore((state) => state.resetEditedReservation);
	const { switchErrorHandling } = useError();
	const loginMutation = useMutation(
		async (user: Credential) =>
		await axios.post(`${process.env.REACT_APP_API_URL}/login`, user),
		{
			onSuccess: () => {
				navigate("/mypage");
			},
			onError: (err: any) => {
				if (err.response.data.message) {
					switchErrorHandling(err.response.data.message);
				} else {
					switchErrorHandling(err.message.data);
				}
			},
		}
	)
	const registerMutation = useMutation(
		async (user: Credential) =>
		  await axios.post(`${process.env.REACT_APP_API_URL}/signup`, user),
		{
			onError: (err: any) => {
				if (err.response.data.message) {
					switchErrorHandling(err.response.data.message);
				} else {
					switchErrorHandling(err.message.data);
				}
			},
		}
	)
	const logoutMutation = useMutation(
		async () => await axios.post(`${process.env.REACT_APP_API_URL}/logout`),
		{
			onSuccess: () => {
				resetEditedReservation();
				navigate("/");
			},
			onError: (err: any) => {
				if (err.response.data.message) {
					switchErrorHandling(err.response.data.message);
				} else {
					switchErrorHandling(err.message.data);
				}
			},
		}
	)
	return { loginMutation, registerMutation, logoutMutation }
}

