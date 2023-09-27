import axios from 'axios';
import { useQueryClient, useMutation } from '@tanstack/react-query';
import { Room } from '../types';
import useStore from '../store';
import { useError } from './useError';

export const useMutateRooms = () => {
    const queryClient = useQueryClient();
    const switchErrorHandling = useError();
    const resetEditedReservation = useStore((state) => state.resetEditedReservation)

    }