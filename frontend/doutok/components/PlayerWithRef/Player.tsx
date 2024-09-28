"use client";

import React, {
  MutableRefObject,
  useEffect,
  useRef,
  useState
} from "react";
import "plyr-react/plyr.css";
import "./Player.css";
import { SourceInfo } from "plyr";
import { APITypes, PlyrInstance, PlyrProps, usePlyr } from "plyr-react";

export interface PlayerProps {
  src?: string;
  sources?: SourceInfo;
  title: string;
  onRef: (ref: MutableRefObject<APITypes>) => void;
}

const DouTokPlayer = React.forwardRef<APITypes, PlyrProps>((props, ref) => {
  const { source, options } = props;
  const raptorRef = usePlyr(ref, { options, source });

  // Do all api access here, it is guaranteed to be called with the latest plyr instance
  React.useEffect(() => {
    /**
     * Fool react for using forward ref as normal ref
     * NOTE: in a case you don't need the forward mechanism and handle everything via props
     * you can create the ref inside the component by yourself
     */
    const { current } = ref as React.MutableRefObject<APITypes>;
    if (current.plyr.source === null) return;

    /* This code is accessing the Plyr instance API and adding event listeners to it. */
    const api = current as { plyr: PlyrInstance };
    // api.plyr.on("ready", () => console.log("I'm ready"));
    api.plyr.on("canplay", () => {
      // NOTE: browser may pause you from doing so:  https://goo.gl/xX8pDD
      api.plyr.play();
      console.log("duration of audio is", api.plyr.duration);
    });
    // api.plyr.on("ended", () => console.log("I'm Ended"));
  });

  return (
    <video
      ref={raptorRef as React.MutableRefObject<HTMLVideoElement>}
      className="plyr-react plyr"
    />
  );
});

DouTokPlayer.displayName = "DouTokPlayer";

export function Player(props: PlayerProps) {
  const [haveSource, setHaveSource] = useState(false);

  useEffect(() => {
    setHaveSource(true);
  }, []);

  const ref = useRef<APITypes>(null);
  useEffect(() => {
    props.onRef(ref as MutableRefObject<APITypes>);
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
      <DouTokPlayer
        ref={ref}
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
          autoplay: true
        }}
      />
    </div>
  );
}
