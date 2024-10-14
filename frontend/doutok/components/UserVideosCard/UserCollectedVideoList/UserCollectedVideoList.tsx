import {
  CollectionServiceListCollectionResponse,
  FavoriteServiceListFavoriteVideoResponse, SvapiListCollectionResponse,
  SvapiVideo,
  useCollectionServiceListCollection,
  useCollectionServiceListVideo4Collection,
  useFavoriteServiceListFavoriteVideo
} from "@/api/svapi/api";
import { message } from "antd";
import React, { useEffect } from "react";
import { UserVideosList } from "@/components/UserVideosList/UserVideosList";

export function UserCollectedVideoList() {
  const [total, setTotal] = React.useState(1);
  const [data, setData] = React.useState<SvapiVideo[]>([]);
  const [loading, setLoading] = React.useState(false);
  const [page, setPage] = React.useState(1);

  const [defaultCollectionId, setDefaultCollectionId] = React.useState<string>();

  useCollectionServiceListCollection({
    resolve: (result: CollectionServiceListCollectionResponse) => {
      if (result?.code !== 0) {
        message.error("获取收藏夹失败");
        return result;
      }

      const collection = result?.data?.collections?.[0];
      if (collection) {
        setDefaultCollectionId(collection.id);
      }

      return result;
    }
  });

  const listFavoriteVideoMutate = useCollectionServiceListVideo4Collection({});
  const loadData = () => {
    if (loading) {
      return;
    }

    setLoading(true);
    listFavoriteVideoMutate
      .mutate({
        page: page,
        size: 10
      })
      .then((result: FavoriteServiceListFavoriteVideoResponse) => {
        if (result?.code !== 0) {
          message.error("获取视频列表失败");
          return;
        }

        setData([...data, ...(result.data?.videos ?? [])]);
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
      domId={"favorite-list"}
      loadData={loadData}
      total={total}
      data={data}
      loading={loading}
    />
  );
}
