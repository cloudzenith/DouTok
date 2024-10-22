import { Player } from "@/components/Player/Player";
import { SvapiVideo } from "@/api/svapi/api";
import * as React from "react";
import { Drawer, FloatButton } from "antd";
import { CommentComponent } from "@/components/Player/CommentComponent/CommentComponent";
import useUserStore from "@/components/UserStore/useUserStore";
import "./PlayerModal.css";
import { CloseOutlined, LeftOutlined, RightOutlined } from "@ant-design/icons";
import { useEffect } from "react";

export interface PlayerModalProps {
  open?: boolean;
  onCancel?: (e: React.MouseEvent<HTMLButtonElement>) => void;
  playUrl: string;
  username: string;
  description: string;
  videoInfo: SvapiVideo;
  onClose?: (e: React.MouseEvent | React.KeyboardEvent) => void;
  onLastOne?: () => void;
  onNextOne?: () => void;
}

export const PlayerModal = React.memo(CorePlayerModal);

export function CorePlayerModal(props: PlayerModalProps) {
  const [openCommentDrawer, setOpenCommentDrawer] = React.useState(false);
  const currentUserId: string = useUserStore(state => state.currentUserId);
  const [videoInfo, setVideoInfo] = React.useState<SvapiVideo>(props.videoInfo);

  useEffect(() => {
    setVideoInfo(props.videoInfo);
  }, [props.videoInfo]);

  return (
    <>
      <Drawer
        placement={"right"}
        open={props.open}
        onClose={props.onClose}
        width={"80vw"}
        height={"90vh"}
        zIndex={1000}
        destroyOnClose={true}
        className={"player-modal"}
        closeIcon={
          <CloseOutlined
            style={{
              color: "white"
            }}
          />
        }
      >
        <Player
          src={"http://localhost:9000/shortvideo/" + videoInfo.play_url}
          avatar={
            "http://localhost:9000/shortvideo/" + videoInfo.author?.avatar
          }
          username={videoInfo?.author?.name}
          description={videoInfo.title}
          title={videoInfo.title}
          userId={videoInfo.author?.id}
          isCouldFollow={currentUserId !== videoInfo.author?.id}
          videoInfo={videoInfo}
          useExternalCommentDrawer={true}
          onOpenExternalCommentDrawer={() => {
            setOpenCommentDrawer(!openCommentDrawer);
          }}
        />
        <FloatButton
          icon={<RightOutlined />}
          style={{ insetInlineEnd: 24 }}
          tooltip={"下一个视频"}
          onClick={() => props.onNextOne?.()}
        />
        <FloatButton
          icon={<LeftOutlined />}
          style={{ insetInlineEnd: 94 }}
          tooltip={"上一个视频"}
          onClick={() => props.onLastOne?.()}
        />
      </Drawer>
      <Drawer
        title={"评论"}
        open={openCommentDrawer}
        placement={"left"}
        closable={true}
        onClose={() => {
          setOpenCommentDrawer(false);
        }}
        mask={false}
        destroyOnClose={true}
      >
        <CommentComponent videoId={videoInfo.id} />
      </Drawer>
    </>
  );
}
