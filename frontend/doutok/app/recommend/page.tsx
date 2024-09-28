"use client";

import React from "react";
import { RecommendPageVideo } from "@/components/RecommendPageVideo/RecommendPageVideo";
import { RequestComponent } from "@/components/RequestComponent/RequestComponent";

const Recommend = () => {
  return (
    <div>
      <RequestComponent noAuth={true}>
        <RecommendPageVideo domId={"recommend-page-video"} />
      </RequestComponent>
    </div>
  );
};

export default Recommend;
