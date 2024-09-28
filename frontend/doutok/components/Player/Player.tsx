"use client";

import React, { useEffect, useState } from "react";
import dynamic from "next/dynamic";
import "plyr-react/plyr.css";
import "./Player.css";
import { SourceInfo } from "plyr";
import { Button } from "antd";
import Avatar from "antd/es/avatar/avatar";
import { HeartOutlined, MessageOutlined, ShareAltOutlined, StarOutlined } from "@ant-design/icons";

export interface PlayerProps {
  src?: string;
  sources?: SourceInfo;
  title: string;
  avatar?: string;
  username?: string;
  description?: string;
}

const Plyr = dynamic(() => import("plyr-react"), { ssr: false });

export function Player(props: PlayerProps) {
  const [haveSource, setHaveSource] = useState(false);

  useEffect(() => {
    setHaveSource(true);
  }, []);

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
            <div>
              <Button
                size={"small"}
                block={true}
                style={{
                  pointerEvents: "all"
                }}
              >
                关注
              </Button>
            </div>
          </div>
            <div className={"mask-button-container"}>
              <Button
                className={"mask-button"}
                ghost={true}
                block={true}
              >
                <HeartOutlined
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
        <Plyr
          source={{
            type: "video",
          title: props.title,
          sources:
            props.sources !== undefined
              ? props.sources
              : [
                {
                  src: props.src as string
                }
              ]
        }}
        options={{
          ratio: "16:9",
          autoplay: true,
          hideControls: false,
        }}
      />
    </div>
  );
}
