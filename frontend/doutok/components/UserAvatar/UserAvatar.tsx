"use client";

import {
  UserServiceGetUserInfoResponse,
  useUserServiceGetUserInfo
} from "@/api/svapi/api";
import { RequestComponent } from "@/components/RequestComponent/RequestComponent";
import Avatar from "antd/es/avatar/avatar";
import useUserStore from "@/components/UserStore/useUserStore";
import React, { useEffect } from "react";

export function UserAvatar() {
  const avatarState = useUserStore(state => state.avatar);
  const setAvatarState = useUserStore(state => state.setAvatar);

  const [avatar, setAvatar] = React.useState<string>("no-login.svg");

  useEffect(() => {
    setAvatar(avatarState);
  }, [avatarState]);

  useEffect(() => {
    setAvatarState(avatar);
  }, [avatar]);

  useUserServiceGetUserInfo({
    resolve: (resp: UserServiceGetUserInfoResponse) => {
      const { data } = resp;
      if (resp.code !== 0 || data === undefined) {
        return resp;
      }

      // TODO: 暂时写死，未来整理成读取配置
      setAvatar(
        data.user?.avatar !== undefined
          ? "http://localhost:9000/shortvideo/" + data.user.avatar
          : "no-login.svg"
      );
      return resp;
    }
  });

  return (
    <RequestComponent noAuth={false}>
      <Avatar src={avatar} />
    </RequestComponent>
  );
}
