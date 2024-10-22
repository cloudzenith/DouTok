import { create } from "zustand";

const useUserStore = create(set => ({
  token: localStorage.getItem("token") || "",
  setToken: () => set((state: { token: string }) => ({ token: state.token })),
  removeToken: () => set({ token: "" }),
  avatar: localStorage.getItem("avatar") || "",
  setAvatar: () =>
    set((state: { avatar: string }) => ({ avatar: state.avatar })),
  currentUserId: localStorage.getItem("currentUserId") || "",
  setCurrentUserId: () =>
    set((state: { currentUserId: string }) => ({
      currentUserId: state.currentUserId
    })),
  isOpenLoginModal: localStorage.getItem("openLoginModal") || false,
  openLoginModal: () => {
    localStorage.setItem("openLoginModal", "true");
    set({ isOpen: true });
  },
  closeLoginModal: () => {
    localStorage.setItem("openLoginModal", "false");
    set({ isOpen: false });
  }
}));

export default useUserStore;
