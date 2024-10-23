import { Player } from "@/components/Player/Player";
import { SvapiVideo } from "@/api/svapi/api";
import { FloatButton, message } from "antd";
import React, { useEffect } from "react";
import useUserStore from "@/components/UserStore/useUserStore";
import { CloseOutlined, DownOutlined, UpOutlined } from "@ant-design/icons";

export interface RecommendPageVideoProps {
  domId: string;
  // 是否提供关闭功能
  couldCancel: boolean;
  // 点击关闭按钮后的回调
  onCancel?: () => void;
  // 初始数据
  data?: SvapiVideo[];
  // 初始当前视频
  initialCurrent?: number;
  // 加载数据方法
  loadData: () => void;
}

export function RecommendPageVideo(props: RecommendPageVideoProps) {
  const currentUserId: string = useUserStore(state => state.currentUserId);
  const [current, setCurrent] = React.useState<number>(
    props.initialCurrent || 0
  );
  const [data, setData] = React.useState<SvapiVideo[]>(props.data || []);

  useEffect(() => {
    setData(props.data || []);
  }, [props.data]);

  return (
    <div>
      <FloatButton.Group shape={"square"}>
        <FloatButton
          icon={<UpOutlined />}
          tooltip={"上一个视频"}
          onClick={() => {
            if (current === 0) {
              message.info("已经是第一个了");
              return;
            }

            setCurrent(current - 1);
          }}
        />
        <FloatButton
          icon={<DownOutlined />}
          tooltip={"下一个视频"}
          onClick={() => {
            if (current === data.length - 2) {
              props.loadData();
            }

            if (current === data.length - 1) {
              message.info("已经是最后一个了");
              return;
            }

            setCurrent(current + 1);
          }}
        />
        {props.couldCancel && (
          <FloatButton
            icon={<CloseOutlined />}
            tooltip={"关闭"}
            onClick={() => {
              props.onCancel?.();
            }}
          />
        )}
      </FloatButton.Group>
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
            useExternalCommentDrawer={false}
          />
        </div>
      ))}
    </div>
  );
}
