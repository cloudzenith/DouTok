import React from "react";
import "./PageSider.css";
import Sider from "antd/es/layout/Sider";
import { Menu } from "antd";
import { CaretRightFilled } from "@ant-design/icons";
import Link from "next/link";

export function PageSider() {
  const fixedItems = [
    {
      key: "index",
      icon: React.createElement(CaretRightFilled),
      label: <Link href={"/"}>首页</Link>
    },
    {
      key: "recommend",
      icon: React.createElement(CaretRightFilled),
      label: <Link href={"/recommend"}>推荐</Link>
    },
    {
      key: "followed",
      icon: React.createElement(CaretRightFilled),
      label: <Link href={"/followed"}>关注</Link>
    },
    {
      key: "friend",
      icon: React.createElement(CaretRightFilled),
      label: <Link href={"/friend"}>朋友</Link>
    },
    {
      key: "user",
      icon: React.createElement(CaretRightFilled),
      label: <Link href={"/user"}>我的</Link>
    }
  ];

  return (
    <Sider id={"page-sider"}>
      <Menu
        id={"sider-menu"}
        theme={"dark"}
        mode={"inline"}
        items={fixedItems}
      />
    </Sider>
  );
}
