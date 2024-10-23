"use client";

import { RestfulProvider } from "restful-react";
import useUserStore from "@/components/UserStore/useUserStore";
import { LoginModalProvider } from "@/components/LoginModalProvider/LoginModalProvider";

export interface RequestProps {
  children: React.ReactNode;
  noAuth?: boolean;
}

export function RequestComponent(props: RequestProps) {
  const token = useUserStore(state => state.token);

  return (
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    //@ts-ignore
    <RestfulProvider
      base={"http://localhost:22000"}
      resolve={async res => {
        return res;
      }}
      requestOptions={() => ({ headers: { Authorization: `Bearer ${token}` } })}
    >
      {(token || props.noAuth) && props.children}
      {!(token || props.noAuth) && <LoginModalProvider />}
    </RestfulProvider>
  );
}
