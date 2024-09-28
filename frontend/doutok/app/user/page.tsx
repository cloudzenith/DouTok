import React from "react";
import { UserCard } from "@/components/UserCard/UserCard";
import { RequestComponent } from "@/components/RequestComponent/RequestComponent";
import { UserVideosCard } from "@/components/UserVideosCard/UserVideosCard";
import { Divider } from "antd";

const User = () => {
  return (
    <RequestComponent noAuth={true}>
      <div className={"user-page"}>
        <UserCard />
        <Divider />
        <UserVideosCard />
      </div>
    </RequestComponent>
  );
};

export default User;
