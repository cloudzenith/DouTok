"use client";

import { Button, Card, Divider, message, notification } from "antd";
import Avatar from "antd/es/avatar/avatar";
import Meta from "antd/es/card/Meta";
import { useEffect, useState } from "react";

import {
  UserServiceGetUserInfoResponse,
  useUserServiceGetUserInfo,
  useUserServiceUpdateUserInfo
} from "@/api/svapi/api";
import "./UserCard.css";
import { LoginModal } from "@/components/LoginModalProvider/LoginModal/LoginModal";
import useUserStore from "@/components/UserStore/useUserStore";
import { UpdateUserInfoForm } from "@/components/UpdateUserInfoForm/UpdateUserInfoForm";
import { RcFile } from "antd/es/upload/interface";
import { SimpleUpload } from "@/components/SimpleUpload/SimpleUpload";

const failedGetUserInfo = () => {
  message.error("获取用户信息失败，请重新登录");
};

export function UserCard() {
  const [openLoginModal, setOpenLoginModal] = useState<boolean>(false);

  const removeToken = useUserStore(state => state.removeToken);

  const [avatar, setAvatar] = useState<string>("no-login.svg");
  const [username, setUsername] = useState<string | undefined>();
  const [following, setFollowing] = useState<string | undefined>();
  const [fans, setFans] = useState<string | undefined>();
  const [likes, setLikes] = useState<string | undefined>();
  const [doutokId, setDouTokId] = useState<string | undefined>();
  const [signature, setSignature] = useState<string>();

  const [avatarBase64, setAvatarBase64] = useState<string>();

  const avatarState = useUserStore(state => state.avatar);
  const setAvatarState = useUserStore(state => state.setAvatar);

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
        failedGetUserInfo();
        return resp;
      }

      // TODO: 暂时写死，未来整理成读取配置
      setAvatar(
        data.user?.avatar !== undefined
          ? "http://localhost:9000/shortvideo/" + data.user.avatar
          : "no-login.svg"
      );
      setUsername(data.user?.name);
      setFollowing(data.user?.followCount ? data.user?.followCount : "0");
      setFans(data.user?.followerCount ? data.user?.followerCount : "0");
      setDouTokId(data.user?.id);
      setLikes("0");
      setSignature(
        data.user?.signature ? data.user?.signature : "这个人很懒，什么都没写"
      );
      return resp;
    }
  });

  const [oepnEditUserInfoModal, setOpenEditUserInfoModal] =
    useState<boolean>(false);

  const updateUserInfoMutate = useUserServiceUpdateUserInfo({});
  const [avatarObjectName, setAvatarObjectName] = useState<string>();
  useEffect(() => {
    if (avatarObjectName === undefined) {
      return;
    }

    updateUserInfoMutate
      .mutate({
        avatar: avatarObjectName
      })
      .then(() => {
        notification.success({
          message: "上传成功",
          description: "头像上传成功"
        });
      });
  }, [avatarObjectName]);

  return (
    <div className={"user-card-container"}>
      <Card>
        <Meta
          avatar={
            <SimpleUpload
              name={"avatar"}
              accept={".jpg,.jpeg,.png"}
              listType={"picture-circle"}
              className={"avatar-uploader"}
              showUploadList={false}
              onFilePreSigned={(file: RcFile) => {
                const reader = new FileReader();
                reader.readAsDataURL(file);
                reader.onload = () => {
                  setAvatarBase64(reader.result as string);
                };
              }}
              setParentComponentFileObjectName={(objectName: string) => {
                setAvatarObjectName(objectName);
              }}
            >
              {avatarBase64 && (
                <Avatar src={avatarBase64} alt={"avatar"} size={100} />
              )}
              {!avatarBase64 && avatar && (
                <Avatar src={avatar} alt={"avatar"} size={100} />
              )}
            </SimpleUpload>
          }
          title={
            <>
              <span className={"username"}>
                {username != undefined ? username : "未登录"}
              </span>
              <div className={"signature"}>
                {signature != undefined ? signature : "这个人很懒，什么都没写"}
              </div>
            </>
          }
          description={
            <>
              <div className={"user-stats"}>
                <div className={"following-num"}>
                  <span>关注: {following != undefined ? following : "-"}</span>
                </div>
                <Divider className={"divider"} type={"vertical"} />
                <div className={"fans-num"}>
                  <span>粉丝: {fans != undefined ? following : "-"}</span>
                </div>
                <Divider className={"divider"} type={"vertical"} />
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
              <div className={"logout-button"}>
                <Button
                  onClick={() => {
                    removeToken();
                    localStorage.removeItem("token");
                    window.location.reload();
                  }}
                >
                  退出登陆
                </Button>
              </div>
              <div className={"edit-button"}>
                <Button onClick={() => setOpenEditUserInfoModal(true)}>
                  编辑资料
                </Button>
              </div>
            </div>
          </>
        )}
        {doutokId == undefined && (
          <>
            <div className={"buttons"}>
              <div className={"edit-button"}>
                <Button onClick={() => setOpenLoginModal(true)}>
                  登录DouTok
                </Button>
              </div>
            </div>
          </>
        )}
      </Card>
      <LoginModal
        open={openLoginModal}
        onCancel={() => {
          setOpenLoginModal(false);
        }}
        type={"login"}
      />
      {oepnEditUserInfoModal && (
        <UpdateUserInfoForm
          open={oepnEditUserInfoModal}
          onCancel={() => setOpenEditUserInfoModal(false)}
          avatar={avatar}
          name={username}
          signature={signature}
        />
      )}
    </div>
  );
}
