import React, { useEffect } from "react";
import {
  ShortVideoCoreVideoServiceFeedShortVideoResponse,
  SvapiVideo,
  useShortVideoCoreVideoServiceFeedShortVideo,
} from "@/api/svapi/api";
import { message } from "antd";
import { UserVideosList } from "@/components/UserVideosList/UserVideosList";

export function RecommendVideoList() {
  const [total, setTotal] = React.useState(0);
  const [data, setData] = React.useState<SvapiVideo[]>([]);
  const [loading, setLoading] = React.useState(false);
  const [latestTime] = React.useState<string>();

  const feedMutate = useShortVideoCoreVideoServiceFeedShortVideo({});
  const loadData = () => {
    if (loading) {
      return;
    }

    setLoading(true);
    feedMutate
      .mutate({
        latestTime: latestTime,
        feedNum: "10"
      })
      .then((result: ShortVideoCoreVideoServiceFeedShortVideoResponse) => {
        if (result?.code !== 0) {
          message.error("获取视频列表失败");
          return;
        }

        setData([...data, ...(result.data?.videos ?? [])]);
        setTotal(total + (result.data?.videos?.length ?? 0));
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    loadData();
  }, []);

  return (
    <UserVideosList
      domId={"recommend-list"}
      loadData={loadData}
      total={total}
      data={data}
      loading={loading}
    />
  );
}
