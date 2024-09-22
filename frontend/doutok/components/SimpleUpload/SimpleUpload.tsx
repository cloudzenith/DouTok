import { RequestComponent } from "@/components/RequestComponent/RequestComponent";
import { message, notification, Upload, UploadProps } from "antd";
import React, { useEffect } from "react";
import { RcFile, UploadChangeParam, UploadListType } from "antd/es/upload/interface";
import {
  FileServiceReportPublicFileUploadedResponse,
  ShortVideoCoreVideoServicePreSign4UploadCoverResponse,
  useFileServicePreSignUploadingPublicFile,
  useFileServiceReportPublicFileUploaded,
  useUserServiceUpdateUserInfo
} from "@/api/svapi/api";
import SparkMD5 from "spark-md5";

export interface SimpleUploadProps {
  className?: string;
  name: string;
  accept?: string;
  listType?: UploadListType;
  showUploadList?: boolean;
  children: React.ReactNode;
  onFilePreSigned?: (file: RcFile) => void;
}


export function SimpleUpload(props: SimpleUploadProps) {
  const [fileId, setFileId] = React.useState<string>();
  const [reportFileId, setReportFileId] = React.useState<string>();
  const [uploadFile, setUploadFile] = React.useState<RcFile>();
  const [uploadUrl, setUploadUrl] = React.useState<string>();
  const [fileHash, setFileHash] = React.useState<string>();
  const [fileBin, setFileBin] = React.useState<ArrayBuffer>();

  const preSignUploadMutate = useFileServicePreSignUploadingPublicFile({});
  const reportUploadedMutate = useFileServiceReportPublicFileUploaded({});
  const updateUserInfoMutate = useUserServiceUpdateUserInfo({});

  const beforeUpload: UploadProps['beforeUpload'] = (file: RcFile) => {
    const fileReader = new FileReader();
    fileReader.readAsArrayBuffer(file);
    fileReader.onload = (event: ProgressEvent<FileReader>) => {
      if (event === null || event.target === null) {
        return ;
      }

      const hashHandle = new SparkMD5.ArrayBuffer();
      hashHandle.append(event.target.result as ArrayBuffer);
      setFileHash(hashHandle.end());
      setUploadFile(file);
    }

    return false;
  }

  useEffect(() => {
    if (fileHash === undefined || uploadFile === undefined) {
      return ;
    }

    preSignUploadMutate.mutate({
      hash: fileHash,
      fileType: uploadFile.type,
      size: uploadFile.size.toString(),
    }).then((result: ShortVideoCoreVideoServicePreSign4UploadCoverResponse) => {
      if (result?.code !== 0 || result.data === undefined) {
        message.error("上传失败，请重试")
        return ;
      }

      setFileId(result.data.fileId);
      setUploadUrl(result.data.url);

      if (props.onFilePreSigned) {
        props.onFilePreSigned(uploadFile);
      }

      if (result.data.url === undefined) {
        // 触发秒传
        setReportFileId(result.data.file_id);
        return ;
      }

      fetch(result.data.url as string, {
        method: "PUT",
        body: uploadFile,
      }).then((response) => {
        console.log(response)
        if (response.status !== 200) {
          message.error("上传失败，请重试");
          return ;
        }

        setReportFileId(result.data?.file_id);
      })
    })
  }, [fileHash, uploadFile]);

  useEffect(() => {
    if (reportFileId === undefined) {
      return ;
    }

    reportUploadedMutate.mutate({
      fileId: reportFileId,
    }).then((result: FileServiceReportPublicFileUploadedResponse) => {
      if (result?.code !== 0 || result?.data === undefined) {
        message.error("上传失败，请重试")
        return ;
      }

      updateUserInfoMutate.mutate({
        avatar: result.data.object_name,
      }).then(() => {
        notification.success({
          message: "上传成功",
          description: "头像已更新",
        })
      })
    })
  }, [reportFileId]);

  return (
    <RequestComponent
      noAuth={false}
    >
      <Upload
        action={uploadUrl}
        className={props.className}
        name={props.name}
        accept={props.accept}
        listType={props.listType}
        showUploadList={props.showUploadList}
        beforeUpload={beforeUpload}
      >
        {props.children}
      </Upload>
    </RequestComponent>
  );
}