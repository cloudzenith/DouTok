import {
  CommentServiceCreateCommentResponse,
  CommentServiceListComment4VideoResponse,
  SvapiComment,
  useCommentServiceCreateComment,
  useCommentServiceListComment4Video
} from "@/api/svapi/api";
import React, { useEffect } from "react";
import {
  Button,
  Card,
  Divider,
  Input,
  List,
  message,
  Skeleton,
  Space
} from "antd";
import InfiniteScroll from "react-infinite-scroll-component";
import Avatar from "antd/es/avatar/avatar";
import {
  CloseOutlined,
  DislikeOutlined,
  LikeOutlined,
  MessageOutlined
} from "@ant-design/icons";
import Meta from "antd/es/card/Meta";
import { ChildCommentList } from "@/components/Player/CommentComponent/ChildCommentList/ChildCommentList";

export interface CommentListProps {
  videoId?: string;
}

export function CommentComponent(props: CommentListProps) {
  const [page, setPage] = React.useState(1);
  const [data, setData] = React.useState<SvapiComment[]>([]);
  const [total, setTotal] = React.useState<number>(1);
  const [loading, setLoading] = React.useState(false);
  const [newComment, setNewComment] = React.useState<string>();
  const [newCommentParentId, setNewCommentParentId] = React.useState<string>();
  const [newCommentReplyUserId, setNewCommentReplyUserId] =
    React.useState<string>();
  const [newCommentReplyUserName, setNewCommentReplyUserName] =
    React.useState<string>();

  const addCommentMutate = useCommentServiceCreateComment({});
  const addCommentHandle = () => {
    addCommentMutate
      .mutate({
        videoId: props.videoId,
        content: newComment,
        parentId: newCommentParentId,
        replyUserId: newCommentReplyUserId
      })
      .then((result: CommentServiceCreateCommentResponse) => {
        if (result?.code !== 0) {
          message.error("评论失败");
          return;
        }

        if (result?.data?.comment !== undefined && result.data.comment.parentId === undefined) {
          setData([result.data.comment, ...data]);
        }

        setNewComment(undefined);
        setNewCommentParentId(undefined);
        setNewCommentReplyUserId(undefined);
        setNewCommentReplyUserName(undefined);
        message.info("评论成功");
      });
  };

  const listCommentMutate = useCommentServiceListComment4Video({});

  const loadData = () => {
    if (loading) {
      return;
    }

    setLoading(true);
    listCommentMutate
      .mutate({
        videoId: props.videoId,
        pagination: {
          page: page,
          size: 10
        }
      })
      .then((result: CommentServiceListComment4VideoResponse | null) => {
        if (result?.code !== 0) {
          message.error("获取评论失败");
        }

        setData(prevData => [...prevData, ...(result?.data?.comments ?? [])]);
        setTotal(result?.data?.pagination?.count ?? 0);
        setPage(page + 1);
        return result;
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    setPage(1);
    loadData();
  }, [props.videoId]);

  return (
    <>
      <Space.Compact
        style={{
          width: "100%"
        }}
      >
        <Input
          placeholder={"发表你的看法"}
          value={newComment}
          onChange={e => {
            setNewComment(e.target.value);
          }}
          prefix={
            newCommentReplyUserName !== undefined ? (
              <>
                <Button
                  size={"small"}
                  onClick={() => {
                    setNewCommentReplyUserId(undefined);
                    setNewCommentParentId(undefined);
                    setNewCommentReplyUserName(undefined);
                  }}
                >
                  <CloseOutlined />
                </Button>
                <>{"回复-" + newCommentReplyUserName}</>
              </>
            ) : undefined
          }
        />
        <Button onClick={addCommentHandle}>评论</Button>
      </Space.Compact>
      <div id={"comment-list"}>
        <InfiniteScroll
          next={loadData}
          hasMore={data.length < total}
          loader={<Skeleton paragraph={{ rows: 1 }} />}
          endMessage={<Divider plain>暂时没有更多了</Divider>}
          dataLength={data.length}
          scrollableTarget={"comment-list"}
        >
          <List
            dataSource={data}
            renderItem={(item: SvapiComment) => (
              <>
                <List.Item key={item.id}>
                  <Card
                    style={{
                      width: "100%"
                    }}
                    actions={[
                      <Button
                        key={"reply"}
                        onClick={() => {
                          setNewCommentReplyUserId(item.user?.id);
                          setNewCommentParentId(item.id);
                          setNewCommentReplyUserName(item.user?.name);
                        }}
                      >
                        <MessageOutlined />
                      </Button>,
                      <Button key={"favorite"}>
                        <LikeOutlined />
                      </Button>,
                      <Button key={"dislike"}>
                        <DislikeOutlined />
                      </Button>
                    ]}
                  >
                    <Meta
                      title={<div>{item.user?.name}</div>}
                      avatar={
                        <Avatar
                          src={
                            "http://localhost:9000/shortvideo/" +
                            item.user?.avatar
                          }
                        />
                      }
                      description={item.date}
                    />
                    <br />
                    {item.content}
                    {item?.comments?.length !== undefined &&
                      item?.comments?.length > 0 && (
                        <ChildCommentList
                          commentId={item.id}
                          initialComments={item.comments}
                        />
                      )}
                  </Card>
                </List.Item>
              </>
            )}
          />
        </InfiniteScroll>
      </div>
    </>
  );
}
