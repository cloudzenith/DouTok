import { Player } from "@/components/Player/Player";
import {
  ShortVideoCoreVideoServiceFeedShortVideoResponse,
  SvapiVideo,
  useShortVideoCoreVideoServiceFeedShortVideo
} from "@/api/svapi/api";
import { Button, message } from "antd";
import React, { useEffect} from "react";
import useUserStore from "@/components/UserStore/useUserStore";

export interface RecommendPageVideoProps {
  domId: string;
}

export function RecommendPageVideo(props: RecommendPageVideoProps) {
  const currentUserId: string = useUserStore(state => state.currentUserId);

  const [data, setData] = React.useState<SvapiVideo[]>([]);
  const [current, setCurrent] = React.useState<number>(0);
  const [loading, setLoading] = React.useState(false);
  const [latestTime, setLatestTime] = React.useState<string>();

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
  }, [props.domId]);

  return (
    <div>
      {data.map((item: SvapiVideo, index: number) => (
        <div
          key={index}
          style={{
            display: current === index ? "block" : "none"
          }}
        >
          <Player
            src={"http://localhost:9000/shortvideo/" + item.play_url}
            avatar={"http://localhost:9000/shortvideo/" + item.author?.avatar}
            username={item.author?.name}
            description={item.title}
            title={"test"}
            userId={item.author?.id}
            isCouldFollow={currentUserId !== item.author?.id}
            videoInfo={item}
            displaying={current === index}
          />
        </div>
      ))}
      <Button
        onClick={() => {
          if (current === 0) {
            message.info("已经是第一个了");
            return;
          }

          setCurrent(current - 1);
        }}
      >
        上一个
      </Button>
      <Button
        onClick={() => {
          if (current === data.length - 2) {
            loadData();
          }

          if (current === data.length - 1) {
            return;
          }

          setCurrent(current + 1);
        }}
      >
        下一个
      </Button>
    </div>
  );
}
