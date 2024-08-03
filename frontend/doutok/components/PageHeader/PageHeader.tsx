import React from "react";

import "./PageHeader.css";
import { Header } from "antd/es/layout/layout";
import Avatar from "antd/es/avatar/avatar";
import { UserOutlined } from "@ant-design/icons";
import { MainSearch } from "@/components/MainSearch/MainSearch";
import { Image } from "antd";

export function PageHeader() {
  return (
    <Header id={"header-container"}>
      <div className={"header"}>
        <div className={"logo"}>
          <Image src={"logo.png"} preview={false} alt={"DouTok Logo"} />
        </div>
        <div className={"search"}>
          <MainSearch />
        </div>
        <div className={"header-menu"}>
          <Avatar icon={<UserOutlined />} className={"user-avatar"} />
        </div>
      </div>
    </Header>
  );
}
