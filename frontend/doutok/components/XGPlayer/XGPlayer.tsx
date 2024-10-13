import Player from 'xgplayer'
import 'xgplayer/dist/index.min.css';
import { useEffect, useRef } from "react";

export interface XGPlayerProps {
  url: string;
}

export function XGPlayer(props: XGPlayerProps) {
  const ref = useRef();

  useEffect(() => {
    const player = new Player({
      el: ref.current,
      url: props.url,
    });
  }, []);

  return (
    <div
      ref={ref}
    >

    </div>
  );
}