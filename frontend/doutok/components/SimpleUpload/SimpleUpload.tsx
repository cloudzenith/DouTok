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
  const [uploadFile, setUploadFile] = React.useState<RcFile>();
  const [uploadUrl, setUploadUrl] = React.useState<string>();
  const [fileHash, setFileHash] = React.useState<string>();
  const [fileBin, setFileBin] = React.useState<ArrayBuffer>();

  const preSignUploadMutate = useFileServicePreSignUploadingPublicFile({});
  const reportUploadedMutate = useFileServiceReportPublicFileUploaded({});
  const updateUserInfoMutate = useUserServiceUpdateUserInfo({});

  const reportUpload = () => {
    reportUploadedMutate.mutate({
      fileId: fileId,
    }).then((result: FileServiceReportPublicFileUploadedResponse) => {
      console.log("result", result);
      if (result?.code !== 0 || result?.data?.objectName === undefined) {
        message.error("上传失败，请重试")
        return ;
      }

      updateUserInfoMutate.mutate({
        avatar: result.data.objectName,
      }).then(() => {
        notification.success({
          message: "上传成功",
          description: "头像已更新",
        })
      })
    })
  };

  const beforeUpload: UploadProps['beforeUpload'] = (file: RcFile) => {
    const fileReader = new FileReader();
    fileReader.readAsArrayBuffer(file);
    fileReader.onload = (event: ProgressEvent<FileReader>) => {
      if (event === null || event.target === null) {
        return ;
      }

      setFileBin(event.target.result as ArrayBuffer);

      const hashHandle = new SparkMD5();
      hashHandle.append(event.target.result as string);
      setFileHash(hashHandle.end());
      setUploadFile(file);
    }

    if (fileHash === undefined) {
      console.log("fileHash is undefined");
      message.error("文件上传失败，请重试");
      return false;
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

      const url = result.data.url as string;
      fetch(url.replace("minio", "localhost"), {
        method: "PUT",
        body: uploadFile,
      }).then((response) => {
        console.log("response", response);
      })
    })
  }, [fileHash, uploadFile]);

  const onChange = (info: UploadChangeParam) => {

  };

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
        onChange={onChange}
      >
        {props.children}
      </Upload>
    </RequestComponent>
  );
}