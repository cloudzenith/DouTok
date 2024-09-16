"use client";

import React from "react";
import Search from "antd/es/input/Search";
import { SearchProps } from "antd/lib/input";

import "./MainSearch.css";
import { SearchOutlined } from "@ant-design/icons";

const onSearch: SearchProps["onSearch"] = (value, _e, info) => {
  console.log(info?.source, value);
};

export function MainSearch() {
  return (
    <Search
      className={"search"}
      placeholder="搜索你的兴趣"
      allowClear
      enterButton={
        <div className={"search-button"}>
          <SearchOutlined /> 搜索
        </div>
      }
      size="large"
      onSearch={onSearch}
    />
  );
}
