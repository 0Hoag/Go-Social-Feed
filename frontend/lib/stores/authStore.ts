import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import { User } from '../api/users';

interface AuthState {
    token: string | null;
    user: User | null;
    isAuthenticated: boolean;
    setToken: (token: string) => void;
    setUser: (user: User) => void;
    logout: () => void;
}

export const useAuthStore = create<AuthState>()(
    persist(
        (set) => ({
            token: null,
            user: null,
            isAuthenticated: false,

            setToken: (token: string) =>
                set({ token, isAuthenticated: true }),

            setUser: (user: User) =>
                set({ user }),

            logout: () =>
                set({ token: null, user: null, isAuthenticated: false }),
        }),
        {
            name: 'auth-storage',
            partialize: (state) => ({ token: state.token, user: state.user, isAuthenticated: state.isAuthenticated }),
        }
    )
);
