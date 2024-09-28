"use client";

import React from "react";
import { RecommendVideoList } from "@/components/RecommendVideoList/RecommendVideoList";
import { RequestComponent } from "@/components/RequestComponent/RequestComponent";

export function Index() {
  return (
    <div>
      <RequestComponent>
        <RecommendVideoList />
      </RequestComponent>
    </div>
  );
}
