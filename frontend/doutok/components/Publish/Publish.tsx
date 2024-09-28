"use client";

import { Button, Divider, Form, message, Modal } from "antd";
import {
  ProForm,
  ProFormItem,
  ProFormText,
  ProFormTextArea
} from "@ant-design/pro-form";
import { RequestComponent } from "@/components/RequestComponent/RequestComponent";
import { PlusOutlined } from "@ant-design/icons";
import React from "react";
import {
  ShortVideoCoreVideoServiceReportVideoFinishUploadResponse,
  useShortVideoCoreVideoServiceReportVideoFinishUpload
} from "@/api/svapi/api";
import { SimpleUpload } from "@/components/SimpleUpload/SimpleUpload";

export function Publish() {
  const [open, setOpen] = React.useState(false);

  const [videoFileId, setVideoFileId] = React.useState<string>();
  const [videoFileObjectName, setVideoFileObjectName] =
    React.useState<string>();
  const [coverFileObjectName, setCoverFileObjectName] =
    React.useState<string>();

  const reportVideoUploadedMutate =
    useShortVideoCoreVideoServiceReportVideoFinishUpload({});

  const [formRef] = Form.useForm();
  const reportUploadVideo = (formData: Record<string, string>) => {
    if (videoFileId === undefined) {
      message.error("请上传视频");
      return;
    }

    reportVideoUploadedMutate
      .mutate({
        fileId: videoFileId,
        title: formData?.title,
        videoUrl: videoFileObjectName,
        coverUrl: coverFileObjectName,
        description: formData?.description
      })
      .then(
        (result: ShortVideoCoreVideoServiceReportVideoFinishUploadResponse) => {
          if (result?.code !== 0 || result?.data === undefined) {
            message.error("上传失败，请重试");
            return;
          }
        }
      );

    window.location.reload();
  };

  return (
    <>
      <Button
        className={"publish-button"}
        type={"primary"}
        icon={<PlusOutlined />}
        onClick={() => {
          setOpen(true);
        }}
      >
        发布
      </Button>
      {open && (
        <RequestComponent noAuth={false}>
          <Modal open={open} onCancel={() => setOpen(false)} footer={null}>
            <ProForm form={formRef} onFinish={reportUploadVideo}>
              <ProFormText
                name={"title"}
                label={"标题"}
                placeholder={"请输入标题"}
                rules={[
                  {
                    required: true,
                    message: "请输入视频标题"
                  }
                ]}
              />
              <ProFormTextArea
                name={"description"}
                label={"视频描述"}
                placeholder={"请输入视频描述"}
                rules={[
                  {
                    required: true,
                    message: "请输入视频描述"
                  }
                ]}
              />
              <ProFormItem name={"video"}>
                <SimpleUpload
                  name={"cover"}
                  accept={"video/*"}
                  setParentComponentFileId={(fileId: string) => {
                    setVideoFileId(fileId);
                  }}
                  setParentComponentFileObjectName={(objectName: string) => {
                    setVideoFileObjectName(objectName);
                  }}
                >
                  <Button>上传视频</Button>
                </SimpleUpload>
              </ProFormItem>
              <Divider />
              <ProFormItem name={"cover"}>
                <SimpleUpload
                  name={"cover"}
                  accept={"image/*"}
                  setParentComponentFileObjectName={(objectName: string) => {
                    setCoverFileObjectName(objectName);
                  }}
                >
                  <Button>上传视频封面</Button>
                </SimpleUpload>
              </ProFormItem>
              <Divider />
            </ProForm>
          </Modal>
        </RequestComponent>
      )}
    </>
  );
}
