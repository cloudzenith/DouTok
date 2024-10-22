import React, { useEffect } from "react";
import { LoginModal } from "@/components/LoginModalProvider/LoginModal/LoginModal";
import useUserStore from "@/components/UserStore/useUserStore";

export interface LoginModalProviderProps {
  open?: boolean;
  onCancel: (e: React.MouseEvent<HTMLButtonElement>) => void;
}

export function LoginModalProvider() {
  const isOpenLoginModal = useUserStore(state => state.isOpenLoginModal);
  const openLoginModal = useUserStore(state => state.openLoginModal);
  const closeLoginModal = useUserStore(state => state.closeLoginModal);
  const [open, setOpen] = React.useState<boolean>(false);

  useEffect(() => {
    if (!isOpenLoginModal && isOpenLoginModal !== "true") {
      openLoginModal();
      setOpen(true);
    }
  }, []);

  return (
    <>
      {open && (
        <LoginModal
          open={open}
          type={"login"}
          onCancel={() => {
            setOpen(false);
            closeLoginModal();
          }}
        />
      )}
    </>
  );
}
