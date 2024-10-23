import { SvapiVideo } from "@/api/svapi/api";
import { Card, Image, List, message, Space } from "antd";
import Meta from "antd/es/card/Meta";
import { LikeOutlined, MessageOutlined, StarOutlined } from "@ant-design/icons";
import React, { useEffect } from "react";
import { PlayerModal } from "@/components/PlayerModal/PlayerModal";

const IconText = ({ icon, text }: { icon: React.FC; text: string }) => (
  <Space>
    {React.createElement(icon)}
    {text}
  </Space>
);

export interface VideoListProps {
  data: SvapiVideo[];
  onNeedLoadData: () => void;
}

export function VideoList(props: VideoListProps) {
  const [openPlayer, setOpenPlayer] = React.useState(false);
  const [playUrl, setPlayUrl] = React.useState("");
  const [videoInfo, setVideoInfo] = React.useState<SvapiVideo>();
  const [publisher, setPublisher] = React.useState("");
  const [description, setDescription] = React.useState("");
  const [current, setCurrent] = React.useState<number>(0);

  const handleLastOne = () => {
    if (current === 0) {
      message.info("已经是第一个了");
      return;
    }

    setCurrent(current - 1);
  };

  const handleNextOne = () => {
    if (current === props.data.length - 2) {
      props.onNeedLoadData();
    }

    if (current === props.data.length - 1) {
      message.info("已经是最后一个了");
      return;
    }

    setCurrent(current + 1);
  };

  useEffect(() => {
    setVideoInfo(props.data[current]);
  }, [current]);

  return (
    <>
      <List
        style={{
          overflow: "hidden"
        }}
        grid={{
          gutter: 16,
          xs: 1,
          sm: 2,
          md: 4,
          lg: 4,
          xl: 6,
          xxl: 3
        }}
        dataSource={props.data}
        renderItem={(item: SvapiVideo, index: number) => (
          <>
            <List.Item key={item.id}>
              <Card
                onClick={() => {
                  setPlayUrl(
                    "http://localhost:9000/shortvideo/" + item.play_url
                  );
                  setPublisher(item.author?.name || "未知用户");
                  setDescription(item.title || "暂无描述");
                  setOpenPlayer(true);
                  setVideoInfo(item);
                  setCurrent(index);
                }}
                hoverable={true}
                cover={
                  <Image
                    src={"http://localhost:9000/shortvideo/" + item.cover_url}
                    preview={false}
                  />
                } // TODO: 未来整理成读取配置
              >
                <Meta title={item.title} />
                <IconText icon={StarOutlined} text="123万" />
                <IconText icon={LikeOutlined} text="156万" />
                <IconText icon={MessageOutlined} text="123万" />
              </Card>
            </List.Item>
            {openPlayer && (
              <PlayerModal
                open={openPlayer}
                onCancel={() => setOpenPlayer(false)}
                onClose={() => setOpenPlayer(false)}
                playUrl={playUrl}
                username={publisher}
                description={description}
                videoInfo={videoInfo as SvapiVideo}
                onLastOne={handleLastOne}
                onNextOne={handleNextOne}
              />
            )}
          </>
        )}
      />
    </>
  );
}
