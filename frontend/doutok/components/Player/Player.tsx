"use client";

import React, { RefObject, useEffect, useRef, useState } from "react";
import dynamic from "next/dynamic";
import "plyr-react/plyr.css";
import "./Player.css";
import { SourceInfo } from "plyr";
import { Button, message } from "antd";
import Avatar from "antd/es/avatar/avatar";
import {
  CheckOutlined,
  HeartFilled,
  HeartOutlined,
  MessageOutlined,
  ShareAltOutlined, StarFilled,
  StarOutlined
} from "@ant-design/icons";
import {
  FavoriteServiceAddFavoriteResponse,
  FollowServiceAddFollowResponse, SvapiVideo, useFavoriteServiceAddFavorite, useFavoriteServiceRemoveFavorite,
  useFollowServiceAddFollow,
  useFollowServiceRemoveFollow
} from "@/api/svapi/api";
import { APITypes, PlyrProps, usePlyr } from "plyr-react";

const CustomPlyrInstance = React.forwardRef<
  APITypes,
  PlyrProps
>((props, ref) => {
  const { source, options = null } = props;
  const raptorRef = usePlyr(ref, {
    source,
    options
  }) as React.MutableRefObject<HTMLVideoElement>;

  useEffect(() => {
    raptorRef.current.play()
  }, []);

  return <video ref={raptorRef} className="plyr-react plyr" />;
});
CustomPlyrInstance.displayName = "CustomPlyrInstance";

interface CorePlayerProps {
  title: string,
  sources?: SourceInfo,
  src: string
}

const CorePlayer = (props: CorePlayerProps, ref) => {
  return (
    <CustomPlyrInstance
      ref={ref}
      source={{
        type: "video",
        title: props.title,
        sources:
          props.sources !== undefined
            ? props.sources
            : [
              {
                src: props.src
              }
            ]
      }}
      options={{
        ratio: "16:9",
        autoplay: false,
        hideControls: false
      }}
    />
  );
}

const CorePlayerMemorized = React.memo(React.forwardRef(CorePlayer));

export interface PlayerProps {
  src?: string;
  sources?: SourceInfo;
  title: string;
  avatar?: string;
  username?: string;
  description?: string;
  userId: string;
  isCouldFollow?: boolean;
  videoInfo: SvapiVideo;
  displaying: boolean;
}

const Plyr = dynamic(() => import("plyr-react"), { ssr: false });

export function Player(props: PlayerProps) {
  const playerRef = useRef();

  const [haveSource, setHaveSource] = useState(false);
  // 能否关注
  const [isCouldFollow, setIsCouldFollow] = useState(props.isCouldFollow);
  // 是否已关注
  const [isFollowed, setIsFollowed] = useState(false);
  const [isCollected, setIsCollected] = useState(false);
  const [isFavorite, setIsFavorite] = useState(false);
  const [videoSource, setVideoSource] = useState("");
  const [displaying, setDisplaying] = useState(false);

  useEffect(() => {
    setHaveSource(true);
  }, []);

  useEffect(() => {
    setVideoSource(props.src as string);
    setIsCouldFollow(props.isCouldFollow);
    setIsFollowed(props.videoInfo.author?.isFollowing === true);
    setIsCollected(props.videoInfo.isCollected === true);
    setIsFavorite(props.videoInfo.isFavorite === true);
    setDisplaying(props.displaying);
  }, [props.isCouldFollow, props.videoInfo]);

  const addFollowMutate = useFollowServiceAddFollow({});
  const addFollowHandle = () => {
    addFollowMutate.mutate({
      userId: props.userId
    }).then((result: FollowServiceAddFollowResponse) => {
      if (result?.code !== 0) {
        message.error("关注失败");
        return;
      }

      message.info("关注成功");
      setIsFollowed(true);
    });
  };

  const removeFollowMutate = useFollowServiceRemoveFollow({});
  const removeFollowMutateHandle = () => {
    removeFollowMutate.mutate({
      userId: props.userId
    }).then((result: FollowServiceAddFollowResponse) => {
      if (result?.code !== 0) {
        message.error("取消关注失败");
        return;
      }

      message.info("取消关注成功");
      setIsFollowed(false);
    })
  };

  const addFavoriteMutate = useFavoriteServiceAddFavorite({});
  const addFavoriteHandle =() => {
    addFavoriteMutate.mutate({
      id: props.videoInfo.id,
      target: 0,
      type: 0
    }).then((result: FavoriteServiceAddFavoriteResponse) => {
      if (result?.code !== 0) {
        message.error("点赞失败");
        return;
      }

      message.info("点赞成功");
      setIsFavorite(true);
    });
  };

  const removeFavoriteMutate = useFavoriteServiceRemoveFavorite({});
  const removeFavoriteHandle = () => {
    removeFavoriteMutate.mutate({
      id: props.videoInfo.id,
      target: 0,
      type: 0
    }).then((result: FavoriteServiceAddFavoriteResponse) => {
      if (result?.code !== 0) {
        message.error("取消点赞失败");
        return;
      }

      message.info("取消点赞成功");
      setIsFavorite(false);
    });
  };

  return (
    <div
      className={"player"}
      style={
        !haveSource
          ? {
            position: "absolute",
            top: "-100%",
            left: "-100%",
            opacity: 0
          }
          : {}
      }
    >
      <div
        className={"mask"}
        style={{
          color: "white"
        }}
      >
        <div
          className={"down-left"}
        >
          <div
            className={"publish-user-name"}
          >
            <div
              style={{
                fontSize: "30px"
              }}
            >
              @{props.username}
            </div>
            <div
              style={{
                fontSize: "20px"
              }}
            >
              {props.description}
            </div>
          </div>
        </div>
        <div className={"down-right"}>
          <div className={"mask-button-container"}>
            <div>
              <Avatar
                className={"mask-button"}
                src={props.avatar}
                size={70}
              />
            </div>
            <div className={"follow-button"}>
              {isCouldFollow && !isFollowed && <Button
                size={"small"}
                block={true}
                style={{
                  pointerEvents: "all"
                }}
                onClick={addFollowHandle}
              >
                关注
              </Button>}
              {isCouldFollow && isFollowed && <Button
                size={"small"}
                block={true}
                style={{
                  pointerEvents: "all"
                }}
                onClick={removeFollowMutateHandle}
              >
                <CheckOutlined />
              </Button>}
            </div>
          </div>
            <div className={"mask-button-container"}>
              {isFavorite && (
                <Button
                  className={"mask-button"}
                  ghost={true}
                  block={true}
                  onClick={removeFavoriteHandle}
                >
                  <HeartFilled
                    style={{
                      fontSize: "40px"
                    }}
                  />
                </Button>
              )}
              {!isFavorite && (
                <Button
                  className={"mask-button"}
                  ghost={true}
                  block={true}
                  onClick={addFavoriteHandle}
                >
                  <HeartOutlined
                    style={{
                      fontSize: "40px"
                    }}
                  />
                </Button>
              )}
              <div
                className={"number-div"}
              >
                50万
              </div>
            </div>
            <div className={"mask-button-container"}>
              <Button
                className={"mask-button"}
                ghost={true}
              >
                <MessageOutlined
                  style={{
                    fontSize: "40px"
                  }}
                />
              </Button>
              <div
                className={"number-div"}
              >
                50万
              </div>
            </div>
            <div className={"mask-button-container"}>
                {isCollected && (
                  <Button
                    className={"mask-button"}
                    ghost={true}
                  >
                    <StarFilled
                      style={{
                        fontSize: "40px"
                      }}
                    />
                  </Button>
                )}
                {!isCollected && (
                  <Button
                    className={"mask-button"}
                    ghost={true}
                  >
                    <StarOutlined
                      style={{
                        fontSize: "40px"
                      }}
                    />
                  </Button>
                )}
              <div
                className={"number-div"}
              >
                50万
              </div>
            </div>
            <div className={"mask-button-container"}>
              <Button
                className={"mask-button"}
                ghost={true}
              >
                <ShareAltOutlined
                  style={{
                    fontSize: "40px"
                  }}
                />
              </Button>
            </div>
          </div>
        </div>
        <CorePlayerMemorized
          ref={playerRef}
          title={props.title}
          sources={props.sources}
          src={videoSource}
          ref={playerRef}
        />
    </div>
  );
}




