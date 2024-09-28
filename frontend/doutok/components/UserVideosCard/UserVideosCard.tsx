"use client";

import React from "react";
import { Card, Tabs } from "antd";
import { UserPublishedVideoList } from "@/components/UserPublishedVideoList/UserPublishedVideoList";

type PageType = "published" | "liked" | "collected" | "history" | "want";

export function UserVideosCard() {
  const [pageType, setPageType] = React.useState<PageType>("published");

  return (
    <>
      <Card>
        <Tabs
          activeKey={pageType}
          onChange={(key: string) => setPageType(key as PageType)}
        >
          <Tabs.TabPane key={"published"} tab={"作品"} />
          <Tabs.TabPane key={"liked"} tab={"喜欢"} />
          <Tabs.TabPane key={"collected"} tab={"收藏"} />
          <Tabs.TabPane key={"history"} tab={"观看历史"} />
          <Tabs.TabPane key={"want"} tab={"稍后再看"} />
        </Tabs>
        {pageType === "published" && (
          <>
            <UserPublishedVideoList />
          </>
        )}
      </Card>
    </>
  );
}
