import { RecommendPageVideo } from "@/components/RecommendPageVideo/RecommendPageVideo";
import React, { useEffect } from "react";
import {
  ShortVideoCoreVideoServiceFeedShortVideoResponse,
  SvapiVideo,
  useShortVideoCoreVideoServiceFeedShortVideo
} from "@/api/svapi/api";
import { message } from "antd";

export function RecommendPageWrapper() {
  const [data, setData] = React.useState<SvapiVideo[]>([]);
  const [latestTime, setLatestTime] = React.useState<string>();
  const [loading, setLoading] = React.useState(false);

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

        if (result?.data === undefined || result.data.videos === undefined) {
          message.error("获取视频列表失败");
          return;
        }

        setData([...data, ...result.data.videos]);
        setLatestTime(result.data.nextTime);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    loadData();
  }, []);

  return (
    <RecommendPageVideo
      domId={"recommend-page-video"}
      data={data}
      couldCancel={false}
      loadData={loadData}
    />
  );
}
