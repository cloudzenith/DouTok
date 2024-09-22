import { ProForm, ProFormText, ProFormTextArea, ProFormUploadButton, ProFormUploadDragger } from "@ant-design/pro-form";
import {
  UserServiceUpdateUserInfoResponse,
  useUserServiceUpdateUserInfo
} from "@/api/svapi/api";
import { Form, Modal, notification, Upload } from "antd";
import { useState } from "react";
import { FileType } from "next/dist/lib/file-exists";
import { RcFile } from "antd/es/upload/interface";
import { SimpleUpload } from "@/components/SimpleUpload/SimpleUpload";

export interface UpdateUserInfoFormProps {
  open?: boolean;
  onCancel?: () => void;
  name?: string;
  signature?: string;
  avatar?: string;
  backgroundImage?: string;
}

export function UpdateUserInfoForm(props: UpdateUserInfoFormProps) {
  const [open, setOpen] = useState(props.open);
  const [name, setName] = useState(props.name);
  const [signature, setSignature] = useState(props.signature);
  const [avatar, setAvatar] = useState(props.avatar ? props.avatar : "no-login.svg");
  const [backgroundImage, setBackgroundImage] = useState(props.backgroundImage ? props.backgroundImage : "no-login.svg");

  const [formRef] = Form.useForm();

  const updateUserInfo = useUserServiceUpdateUserInfo({});

  const submit4ModifyUserInfo = (formData: Record<string, string>) => {
    updateUserInfo
      .mutate({
        name: formData?.name,
        signature: formData?.signature,
        avatar: formData?.avatar,
      })
      .then((r: UserServiceUpdateUserInfoResponse) => {
        setOpen(false);

        if (r?.code !== 0) {
          notification.error({
            message: "修改失败",
            description:
              r?.msg === undefined ? "请检查修改信息是否正确" : r?.msg
          });
          return;
        }

        window.location.reload();
      });
  };

  return (
    <Modal open={open} onCancel={props.onCancel} footer={null}>
      <ProForm
        form={formRef}
        onFinish={submit4ModifyUserInfo}
        submitter={{
          submitButtonProps: {},
          resetButtonProps: {
            style: {
              display: 'none',
            },
          },
        }}
      >
        <SimpleUpload
          name={"background"}
          accept={".jpg,.jpeg,.png"}
          listType={"picture-card"}
          className={"avatar-uploader"}
          showUploadList={false}
        >
          {avatar ? <img src={backgroundImage} alt={"background"} /> : undefined}
        </SimpleUpload>
        <ProFormText
          name="name"
          label="用户昵称"
          initialValue={name}
          placeholder={"给自己一个好听的名字"}
        />
        <ProFormTextArea
          name={"signature"}
          label={"签名"}
          initialValue={signature}
          placeholder={"介绍一下你自己"}
        />
      </ProForm>
    </Modal>
  );
}
