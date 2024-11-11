import { create } from "zustand";

interface UserState {
  token: string;
  setToken: (token: string) => void;
  removeToken: () => void;
  avatar: string;
  setAvatar: (avatar: string) => void;
  currentUserId: string;
  setCurrentUserId: (id: string) => void;
  isOpenLoginModal: boolean;
  openLoginModal: () => void;
  closeLoginModal: () => void;
}

const isClient = typeof window !== "undefined";

const localStorageUtil = {
  getItem: (key: string, defaultValue: string) => {
    return isClient ? localStorage.getItem(key) ?? defaultValue : defaultValue;
  },
  setItem: (key: string, value: string) => {
    if (isClient) {
      localStorage.setItem(key, value);
    }
  },
  removeItem: (key: string) => {
    if (isClient) {
      localStorage.removeItem(key);
    }
  }
};

const useUserStore = create<UserState>((set) => ({
  token: localStorageUtil.getItem("token", ""),
  setToken: (token: string) => {
    localStorageUtil.setItem("token", token);
    set({ token });
  },
  removeToken: () => {
    localStorageUtil.removeItem("token");
    set({ token: "" });
  },
  avatar: localStorageUtil.getItem("avatar", ""),
  setAvatar: (avatar: string) => {
    localStorageUtil.setItem("avatar", avatar);
    set({ avatar });
  },
  currentUserId: localStorageUtil.getItem("currentUserId", ""),
  setCurrentUserId: (id: string) => {
    localStorageUtil.setItem("currentUserId", id);
    set({ currentUserId: id });
  },
  isOpenLoginModal: localStorageUtil.getItem("openLoginModal", "false") === "true",
  openLoginModal: () => {
    localStorageUtil.setItem("openLoginModal", "true");
    set({ isOpenLoginModal: true });
  },
  closeLoginModal: () => {
    localStorageUtil.setItem("openLoginModal", "false");
    set({ isOpenLoginModal: false });
  }
}));

export default useUserStore;