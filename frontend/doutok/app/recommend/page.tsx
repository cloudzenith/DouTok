"use client";

import React from "react";
import { RequestComponent } from "@/components/RequestComponent/RequestComponent";
import { RecommendPageWrapper } from "@/components/VideoListWrappers/RecommendPageWrapper/RecommendPageWrapper";

const Recommend = () => {
  return (
    <div>
      <RequestComponent noAuth={true}>
        <RecommendPageWrapper />
      </RequestComponent>
    </div>
  );
};

export default Recommend;
