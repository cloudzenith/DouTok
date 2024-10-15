import {
  CommentServiceListChildCommentResponse,
  SvapiComment,
  useCommentServiceListChildComment
} from "@/api/svapi/api";
import { List, message } from "antd";
import React from "react";

export interface ChildCommentListProps {
  commentId?: string;
  initialComments: SvapiComment[];
}

export function ChildCommentList(props: ChildCommentListProps) {
  const [comments, setComments] = React.useState(props.initialComments);
  const [expanded, setExpanded] = React.useState(false);
  const [page, setPage] = React.useState(1);
  const [total, setTotal] = React.useState(1);
  const [allLoaded, setAllLoaded] = React.useState(false);

  const loadChildCommentsMutate = useCommentServiceListChildComment({});
  const loadData = () => {
    loadChildCommentsMutate
      .mutate({
        commentId: props.commentId,
        pagination: {
          page: page,
          size: 10
        }
      })
      .then((result: CommentServiceListChildCommentResponse) => {
        if (result?.code !== 0) {
          message.error("获取子评论失败");
          return;
        }

        const filtered = result.data?.comments?.filter(comment => {
          return !comments.some(c => c.id === comment.id);
        });

        setComments([...comments, ...(filtered ?? [])]);
        setTotal(result?.data?.pagination?.total ?? 0);
        setPage(page + 1);
        if (page > total) {
          setAllLoaded(true);
        }
      });
  };

  return (
    <>
      <List
        dataSource={comments}
        bordered={true}
        renderItem={(childComment: SvapiComment) => (
          <List.Item>
            <a>@{childComment.user?.name}</a>对
            <a>@{childComment.reply_user?.name}</a>回复: {childComment.content}{" "}
            <span style={{ color: "gray" }}>{childComment.date}</span>
          </List.Item>
        )}
        footer={
          <>
            {!allLoaded &&
              !expanded &&
              comments?.length !== undefined &&
              comments?.length >= 3 && (
                <a
                  onClick={() => {
                    loadData();
                    setExpanded(true);
                  }}
                >
                  展开所有子评论
                </a>
              )}
            {expanded && !allLoaded && <a onClick={loadData}>加载更多</a>}
            {allLoaded && <>没有更多了</>}
          </>
        }
      />
    </>
  );
}
