import { ProForm, ProFormText, ProFormTextArea } from "@ant-design/pro-form";
import {
  UserServiceUpdateUserInfoResponse,
  useUserServiceUpdateUserInfo
} from "@/api/svapi/api";
import { Form, Modal, notification } from "antd";
import { useState } from "react";

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
  const [name] = useState(props.name);
  const [signature] = useState(props.signature);

  const [formRef] = Form.useForm();

  const updateUserInfo = useUserServiceUpdateUserInfo({});

  const submit4ModifyUserInfo = (formData: Record<string, string>) => {
    updateUserInfo
      .mutate({
        name: formData?.name,
        signature: formData?.signature,
        avatar: formData?.avatar
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
              display: "none"
            }
          }
        }}
      >
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
