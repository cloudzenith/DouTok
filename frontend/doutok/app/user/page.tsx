import React from "react";
import { UserCard } from "@/components/UserCard/UserCard";
import {RequestComponent} from "@/components/RequestComponent/RequestComponent";

const User = () => {
  return (
    <RequestComponent noAuth={true}>
      <UserCard/>
    </RequestComponent>
  );
};

export default User;
