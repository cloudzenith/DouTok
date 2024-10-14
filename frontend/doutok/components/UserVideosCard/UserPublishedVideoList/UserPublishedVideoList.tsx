import {
  ShortVideoCoreVideoServiceListPublishedVideoResponse,
  SvapiVideo,
  useShortVideoCoreVideoServiceListPublishedVideo
} from "@/api/svapi/api";
import { message } from "antd";
import React, { useEffect } from "react";
import { UserVideosList } from "@/components/UserVideosList/UserVideosList";

export function UserPublishedVideoList() {
  const [total, setTotal] = React.useState(1);
  const [data, setData] = React.useState<SvapiVideo[]>([]);
  const [loading, setLoading] = React.useState(false);
  const [page, setPage] = React.useState(1);

  const listPublishedMutate = useShortVideoCoreVideoServiceListPublishedVideo(
    {}
  );
  const loadData = () => {
    if (loading) {
      return;
    }

    setLoading(true);
    listPublishedMutate
      .mutate({
        pagination: {
          page: page,
          size: 10
        }
      })
      .then((result: ShortVideoCoreVideoServiceListPublishedVideoResponse) => {
        if (result?.code !== 0) {
          message.error("获取视频列表失败");
          return;
        }

        setData([...data, ...(result.data?.video_list ?? [])]);
        setTotal(result?.data?.pagination?.total ?? 0);
      })
      .finally(() => {
        setLoading(false);
        setPage(page + 1);
      });
  };

  useEffect(() => {
    loadData();
  }, []);

  return (
    <UserVideosList
      domId={"published-list"}
      loadData={loadData}
      total={total}
      data={data}
      loading={loading}
    />
  );
}
