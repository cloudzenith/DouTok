import { Modal } from "antd";
import { Player } from "@/components/Player/Player";

export interface PlayerModalProps {
  open?: boolean;
  onCancel?: (e: React.MouseEvent<HTMLButtonElement>) => void;
  playUrl: string;
  username: string;
  description: string;
}

export function PlayerModal(props: PlayerModalProps) {
  return (
    <Modal
      open={props.open}
      onCancel={props.onCancel}
      footer={null}
      style={{
        minWidth: "60vw",
        maxHeight: "40vh"
      }}
    >
      <Player src={props.playUrl} title={props.username} />
    </Modal>
  );
}
