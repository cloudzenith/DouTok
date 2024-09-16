"use client"

import { Button, Card, Divider, message } from "antd";
import Avatar from "antd/es/avatar/avatar";
import Meta from "antd/es/card/Meta";
import { useState } from "react";

import { UserServiceGetUserInfoResponse, useUserServiceGetUserInfo } from "@/api/svapi/api";
import "./UserCard.css";
import {LoginModal} from "@/components/LoginModal/LoginModal";

export interface UserCardProps {
  avatar?: string | "no-login.svg";
  username?: string | undefined;
  following?: number | undefined;
  fans?: number | undefined;
  likes?: number | undefined;
  doutok_id?: string | undefined;
}

const failedGetUserInfo = () => {
  message.error("获取用户信息失败，请重新登录");
}

export function UserCard(props: UserCardProps) {
  const [openLoginModal, setOpenLoginModal] = useState<boolean>(false);

  const [avatar, setAvatar] = useState<string | undefined>("no-login.svg");
  const [username, setUsername] = useState<string | undefined>();
  const [following, setFollowing] = useState<string | undefined>();
  const [fans, setFans] = useState<string | undefined>();
  const [likes, setLikes] = useState<string | undefined>();
  const [doutokId, setDouTokId] = useState<string | undefined>();

  useUserServiceGetUserInfo({
    resolve: (resp: UserServiceGetUserInfoResponse) => {
      const {data} = resp;
      if (resp.code !== 0 || data === undefined) {
        failedGetUserInfo();
        return resp;
      }

      setAvatar(data.user?.avatar);
      setUsername(data.user?.name);
      setFollowing(data.user?.followCount);
      setFans(data.user?.followerCount);
      setDouTokId(data.user?.id);
      return resp;
    }
  });

  return (
    <div className={"user-card-container"}>
      <Card>
        <Meta
          avatar={
            <Avatar
              src={
                avatar != undefined && avatar.length != 0
                  ? avatar
                  : "no-login.svg"
              }
              size={150}
            />
          }
          title={
            <span className={"username"}>
              {username != undefined ? username : "未登录"}
            </span>
          }
          description={
            <>
              <div className={"user-stats"}>
                <div className={"following-num"}>
                  <span>关注: {following != undefined ? following : "-"}</span>
                </div>
                <Divider className={"divider"} type={"horizontal"} />
                <div className={"fans-num"}>
                  <span>粉丝: {fans != undefined ? following : "-"}</span>
                </div>
                <Divider className={"divider"} type={"horizontal"} />
                <div className={"likes-num"}>
                  <span>获赞: {likes != undefined ? following : "-"}</span>
                </div>
              </div>
            </>
          }
        />
        {doutokId != undefined && (
          <>
            <div className={"buttons"}>
              <div className={"doutok-id"}>
                DouTok号: {doutokId != undefined ? doutokId : "-"}
              </div>
              <div className={"edit-button"}>
                <Button>编辑资料</Button>
              </div>
            </div>
          </>
        )}
        {doutokId == undefined && (
          <>
            <div className={"buttons"}>
              <div className={"edit-button"}>
                <Button
                  onClick={() => setOpenLoginModal(true)}
                >
                  登录DouTok
                </Button>
              </div>
            </div>
          </>
        )}
      </Card>
      { openLoginModal && <LoginModal
          open={openLoginModal}
          onCancel={() => setOpenLoginModal(false)}
      />}
    </div>
  );
}
