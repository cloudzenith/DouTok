import {
  CollectionServiceListCollectionResponse, CollectionServiceListVideo4CollectionResponse,
  SvapiVideo,
  useCollectionServiceListCollection,
  useCollectionServiceListVideo4Collection,

} from "@/api/svapi/api";
import { message } from "antd";
import React, { useEffect } from "react";
import { UserVideosList } from "@/components/UserVideosList/UserVideosList";
import { RequestComponent } from "@/components/RequestComponent/RequestComponent";

export function UserCollectedVideoList() {
  const [total, setTotal] = React.useState(1);
  const [data, setData] = React.useState<SvapiVideo[]>([]);
  const [loading, setLoading] = React.useState(false);
  const [page, setPage] = React.useState(1);

  const [defaultCollectionId, setDefaultCollectionId] = React.useState<string>();

  useCollectionServiceListCollection({
    queryParams: {
      "pagination.page": 1,
      "pagination.size": 1
    },
    resolve: (result: CollectionServiceListCollectionResponse) => {
      if (result?.code !== 0) {
        message.error("获取收藏夹失败");
        return result;
      }

      const collection = result?.data?.collections?.[0];
      if (collection !== undefined && collection.id !== undefined) {
        setDefaultCollectionId(collection.id);
      }

      return result;
    }
  });

  useEffect(() => {
    if (defaultCollectionId === undefined) {
      return;
    }

    loadData();
  }, [defaultCollectionId]);

  const { refetch: doFetchCollectionVideo } = useCollectionServiceListVideo4Collection({
    lazy: true
  });

  const loadData = () => {
    if (loading) {
      return;
    }

    if (defaultCollectionId === undefined) {
      return ;
    }

    setLoading(true);

    doFetchCollectionVideo({
      queryParams: {
        collectionId: defaultCollectionId,
        "pagination.page": page,
        "pagination.size": 10
      }
    }).then((result: CollectionServiceListVideo4CollectionResponse | null) => {
      if (result?.code !== 0) {
        message.error("获取视频列表失败");
        return;
      }

      setData([...data, ...(result.data?.videos ?? [])]);
      setTotal(result?.data?.pagination?.total ?? 0);
    }).finally(() => {
      setLoading(false);
      setPage(page + 1);
    });
  };

  useEffect(() => {
    loadData();
  }, []);

  return (
    <>
      <UserVideosList
        domId={"favorite-list"}
        loadData={loadData}
        total={total}
        data={data}
        loading={loading}
      />
    </>
  );
}
