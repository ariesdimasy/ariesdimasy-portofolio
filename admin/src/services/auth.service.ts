import api from "@/lib/api";
import { LoginRequest, LoginResponse, RegisterRequest, User } from "@/types";

export const authService = {
  login: async (data: LoginRequest): Promise<LoginResponse> => {
    const res = await api.post<LoginResponse>("/api/auth/login", data);
    return res.data;
  },

  register: async (data: RegisterRequest): Promise<{ message: string }> => {
    const res = await api.post<{ message: string }>("/api/auth/register", data);
    return res.data;
  },
};

export const userService = {
  getAll: async (page = 1, limit = 10) => {
    const res = await api.get<{ data: User[]; total: number }>("/api/users", {
      params: { page, limit },
    });
    return res.data;
  },

  getById: async (id: number) => {
    const res = await api.get<User>(`/api/users/${id}`);
    return res.data;
  },
};
