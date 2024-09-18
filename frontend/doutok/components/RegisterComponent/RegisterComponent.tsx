import { Form, notification, Tabs, theme } from "antd";

import { LoginForm, ProFormCaptcha, ProFormText } from "@ant-design/pro-form";
import React, { useState } from "react";
import { LockOutlined, MailOutlined, MobileOutlined } from "@ant-design/icons";
import {
  UserServiceGetVerificationCodeResponse,
  UserServiceRegisterResponse,
  useUserServiceGetVerificationCode,
  useUserServiceRegister
} from "@/api/svapi/api";

import "./RegisterComponent.css";

type RegisterType = "phone" | "email";

export interface RegisterComponentProps {
  switcher: React.ReactNode;
}

export function RegisterComponent(props: RegisterComponentProps) {
  const { token } = theme.useToken();

  const [formRef] = Form.useForm();

  const [registerType, setRegisterType] = useState<RegisterType>("phone");
  const registerMutate = useUserServiceRegister({});

  const [captchaId, setCaptchaId] = useState<string | undefined>();

  const submit4Register = (formData: Record<string, string>) => {
    registerMutate
      .mutate({
        mobile: formData?.phone,
        email: formData?.email,
        password: formData?.password,
        codeId: captchaId,
        code: formData?.captcha
      })
      .then((r: UserServiceRegisterResponse) => {
        if (r?.code !== 0) {
          notification.error({
            message: "注册失败",
            description:
              r?.msg === undefined ? "请检查注册信息是否正确" : r?.msg
          });
          return;
        }
      });
  };

  const mutate = useUserServiceGetVerificationCode({});

  const getCaptcha = (data: string) => {
    mutate
      .mutate({
        mobile: registerType === "phone" ? data : undefined,
        email: registerType === "email" ? data : undefined
      })
      .then((r: UserServiceGetVerificationCodeResponse) => {
        console.log(r);
        if (r?.code !== 0) {
          return;
        }

        setCaptchaId(r?.data?.code_id);
      });
  };

  return (
    <LoginForm
      className={"login-form"}
      title={"DouTok"}
      subTitle={"Cloudzneith"}
      onFinish={submit4Register}
      form={formRef}
    >
      <Tabs
        centered
        activeKey={registerType}
        onChange={key => setRegisterType(key as RegisterType)}
      >
        <Tabs.TabPane key={"phone"} tab={"手机注册"} />
        <Tabs.TabPane key={"email"} tab={"邮箱注册"} />
      </Tabs>
      {registerType === "phone" && (
        <>
          <ProFormText
            name={"phone"}
            fieldProps={{
              size: "large",
              prefix: <MobileOutlined className={"prefixIcon"} />
            }}
            placeholder={"手机号码"}
            rules={[
              {
                required: true,
                message: "请输入手机号"
              }
            ]}
          />
        </>
      )}
      {registerType === "email" && (
        <>
          <ProFormText
            name={"email"}
            fieldProps={{
              size: "large",
              prefix: <MailOutlined className={"prefixIcon"} />
            }}
            placeholder={"邮箱地址"}
            rules={[
              {
                required: true,
                message: "请输入邮箱"
              }
            ]}
          />
        </>
      )}
      <ProFormText.Password
        name="password"
        fieldProps={{
          size: "large",
          prefix: <LockOutlined className={"prefixIcon"} />,
          strengthText:
            "Password should contain numbers, letters and special characters, at least 8 characters long.",
          statusRender: value => {
            const getStatus = () => {
              if (value && value.length > 12) {
                return "ok";
              }
              if (value && value.length > 6) {
                return "pass";
              }
              return "poor";
            };
            const status = getStatus();
            if (status === "pass") {
              return <div style={{ color: token.colorWarning }}>强度：中</div>;
            }
            if (status === "ok") {
              return <div style={{ color: token.colorSuccess }}>强度：强</div>;
            }
            return <div style={{ color: token.colorError }}>强度：弱</div>;
          }
        }}
        placeholder={"密码"}
        rules={[
          {
            required: true,
            message: "请输入密码1！"
          }
        ]}
      />
      <ProFormText.Password
        name="ensure"
        fieldProps={{
          size: "large",
          prefix: <LockOutlined className={"prefixIcon"} />,
          strengthText:
            "Password should contain numbers, letters and special characters, at least 8 characters long.",
          statusRender: value => {
            const getStatus = () => {
              if (value && value.length > 12) {
                return "ok";
              }
              if (value && value.length > 6) {
                return "pass";
              }
              return "poor";
            };
            const status = getStatus();
            if (status === "pass") {
              return <div style={{ color: token.colorWarning }}>强度：中</div>;
            }
            if (status === "ok") {
              return <div style={{ color: token.colorSuccess }}>强度：强</div>;
            }
            return <div style={{ color: token.colorError }}>强度：弱</div>;
          }
        }}
        placeholder={"确认密码"}
        rules={[
          {
            required: true,
            message: "请再次输入密码！"
          },
          {
            validator: async (_, value) => {
              if (!value || formRef.getFieldValue("password") === value) {
                return Promise.resolve();
              }
              return Promise.reject(new Error("两次输入的密码不匹配！"));
            }
          }
        ]}
      />
      <ProFormCaptcha
        fieldProps={{
          size: "large",
          prefix: <LockOutlined className={"prefixIcon"} />
        }}
        captchaProps={{
          size: "large"
        }}
        placeholder={"请输入验证码"}
        captchaTextRender={(timing, count) => {
          if (timing) {
            return `${count} ${"获取验证码"}`;
          }
          return "获取验证码";
        }}
        name="captcha"
        rules={[
          {
            required: true,
            message: "请输入验证码！"
          }
        ]}
        phoneName={registerType}
        onGetCaptcha={async data => {
          getCaptcha(data);
        }}
      />
      {props.switcher}
    </LoginForm>
  );
}
