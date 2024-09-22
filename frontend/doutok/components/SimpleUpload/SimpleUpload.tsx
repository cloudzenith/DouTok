import { RequestComponent } from "@/components/RequestComponent/RequestComponent";
import { message, notification, Upload, UploadProps } from "antd";
import React from "react";
import { RcFile, UploadChangeParam, UploadListType } from "antd/es/upload/interface";
import {
  ShortVideoCoreVideoServicePreSign4UploadCoverResponse,
  useShortVideoCoreVideoServicePreSign4UploadCover
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
  const [uploadUrl, setUploadUrl] = React.useState<string>();
  const [fileHash, setFileHash] = React.useState<string>();

  const preSignUploadMutate = useShortVideoCoreVideoServicePreSign4UploadCover({})

  const beforeUpload: UploadProps['beforeUpload'] = (file: RcFile) => {
    const fileReader = new FileReader();
    fileReader.readAsArrayBuffer(file);
    fileReader.onload = (event: ProgressEvent<FileReader>) => {
      if (event === null || event.target === null) {
        return ;
      }

      const hashHandle = new SparkMD5();
      hashHandle.append(event.target.result as string);
      setFileHash(hashHandle.end());
    }

    preSignUploadMutate.mutate({
      hash: fileHash,
      fileType: file.type,
      size: file.size.toString(),
      filename: file.name
    }).then((result: ShortVideoCoreVideoServicePreSign4UploadCoverResponse) => {
      if (result?.code !== 0 || result.data === undefined) {
        message.error("上传失败，请重试")
        return ;
      }

      setFileId(result.data.fileId);
      setUploadUrl(result.data.url);

      if (props.onFilePreSigned) {
        props.onFilePreSigned(file);
      }
    })

    return false;
  }

  const onChange = (info: UploadChangeParam) => {
    console.log("info", info);
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