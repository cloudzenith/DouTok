import React from "react"
import {Button, Card, Divider} from "antd";
import Avatar from "antd/es/avatar/avatar";
import Meta from "antd/es/card/Meta";

import "./UserCard.css"

export interface UserCardProps {
  avatar: string | undefined,
  username: string | undefined,
  following: number | undefined,
  fans: number | undefined,
  likes: number | undefined,
  doutok_id: string | undefined
};

export function UserCard(props: UserCardProps) {
  const {avatar, username, following, fans, likes, doutok_id} = props;

  return (
    <div className={"user-card-container"}>
      <Card>
        <Meta
          avatar={<Avatar
            src={avatar != undefined && avatar.length != 0 ? avatar : "no-login.svg"}
            size={150}
          />}
          title={<span className={"username"}>
            {username != undefined ? username : "未登录"}
          </span>}
          description={
            <>
              <div className={"user-stats"}>
                <div className={"following-num"}>
                  <span>
                    关注: {following != undefined ? following : "-"}
                  </span>
                </div>
                <Divider className={"divider"} type={"horizontal"}/>
                <div className={"fans-num"}>
                  <span>
                    粉丝: {fans != undefined ? following : "-"}
                  </span>
                </div>
                <Divider className={"divider"} type={"horizontal"}/>
                <div className={"likes-num"}>
                  <span>
                    获赞: {likes != undefined ? following : "-"}
                  </span>
                </div>
              </div>
            </>
          }
        />
        {
          doutok_id != undefined && (
            <>
              <div className={"buttons"}>
                <div className={"doutok-id"}>
                  DouTok号: {doutok_id != undefined ? doutok_id : "-"}
                </div>
                <div className={"edit-button"}>
                  <Button>编辑资料</Button>
                </div>
              </div>
            </>
          )
        }
        {
          doutok_id == undefined && (
            <>
              <div className={"buttons"}>
                <div className={"edit-button"}>
                  <Button>登录DouTok</Button>
                </div>
              </div>
            </>
          )
        }
      </Card>
    </div>
  );
}
