import { Divider, Skeleton } from "antd";
import React from "react";
import InfiniteScroll from "react-infinite-scroll-component";
import {
  SvapiVideo,
} from "@/api/svapi/api";
import { VideoList } from "@/components/VideoList/VideoList";

export interface UserVideosListProps {
  domId: string;
  loadData: () => void;
  total: number;
  data: SvapiVideo[];
  loading: boolean;
}

export function UserVideosList(props: UserVideosListProps) {
  return (
    <div id={props.domId}>
      <InfiniteScroll
        next={props.loadData}
        hasMore={props.data.length < props.total}
        loader={<Skeleton paragraph={{ rows: 1 }} />}
        endMessage={<Divider plain>暂时没有更多了</Divider>}
        dataLength={props.data.length}
        scrollableTarget={props.domId}
      >
        <VideoList data={props.data} />
      </InfiniteScroll>
    </div>
  );
}
