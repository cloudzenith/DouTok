import React from "react";

import "./PageHeader.css";
import { Header } from "antd/es/layout/layout";
import { MainSearch } from "@/components/MainSearch/MainSearch";
import { Divider, Image } from "antd";
import { UserAvatar } from "@/components/UserAvatar/UserAvatar";
import { Publish } from "@/components/Publish/Publish";
import { RequestComponent } from "@/components/RequestComponent/RequestComponent";

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
          <RequestComponent noAuth={false}>
            <Publish />
          </RequestComponent>
          <Divider type={"vertical"} />
          <RequestComponent noAuth={false}>
            <UserAvatar />
          </RequestComponent>
        </div>
      </div>
    </Header>
  );
}
