"use client"

import {RestfulProvider} from "restful-react";
import {useState} from "react";
import {useRouter} from "next/router";
import {LoginModal} from "@/components/LoginModal/LoginModal";

export interface RequestProps {
  children: React.ReactNode
  noAuth?: boolean
}

export function RequestComponent(props: RequestProps) {
  const [headers, setHeaders] = useState<
    {authorization: string} | undefined
  >();

  // const router = useRouter();

  return (
    //@ts-ignore
    <RestfulProvider
      base={"http://localhost:22000"}
      requestOptions={{ headers: { ...headers } }}
    >
      {(headers || props.noAuth ) && props.children}
      {!(headers || props.noAuth ) && (
        <LoginModal open={true} type={"login"} />
      )}
    </RestfulProvider>
  );
}