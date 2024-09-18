import { create } from "zustand";

const useUserStore = create(set => ({
  token: localStorage.getItem("token") || "",
  setToken: () => set((state: { token: string }) => ({ token: state.token })),
  removeToken: () => set({ token: "" })
}));

export default useUserStore;
