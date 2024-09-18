import { notification, Tabs, theme } from "antd";

import { LoginForm, ProFormText } from "@ant-design/pro-form";
import React, { useState } from "react";
import {
  LockOutlined,
  MailOutlined,
  MobileOutlined,
  UserOutlined
} from "@ant-design/icons";

import { UserServiceLoginResponse, useUserServiceLogin } from "@/api/svapi/api";

import "./LoginComponent.css";
import useUserStore from "@/components/UserStore/useUserStore";

type LoginType = "phone" | "account" | "email";

export interface LoginComponentProps {
  switcher: React.ReactNode;
}

export function LoginComponent(props: LoginComponentProps) {
  const { token } = theme.useToken();

  const setToken = useUserStore(state => state.setToken);

  const [loginType, setLoginType] = useState<LoginType>("phone");
  const loginMutate = useUserServiceLogin({});

  const submit4Login = (formData: Record<string, string>) => {
    loginMutate
      .mutate({
        mobile: formData?.phone,
        email: formData?.email,
        password: formData?.password
      })
      .then((r: UserServiceLoginResponse) => {
        if (r?.code !== 0) {
          notification.error({
            message: "登录失败",
            description: "请检查登录信息是否正确"
          });
        }

        if (r?.data?.token !== undefined) {
          window.localStorage.setItem("token", r.data.token);
          setToken(r.data.token);
          window.location.reload();
        }
      });
  };

  return (
    <LoginForm
      className={"login-form"}
      title={"DouTok"}
      subTitle={"Cloudzneith"}
      onFinish={submit4Login}
    >
      <Tabs
        centered
        activeKey={loginType}
        onChange={key => setLoginType(key as LoginType)}
      >
        <Tabs.TabPane key={"account"} tab={"账号登录"} />
        <Tabs.TabPane key={"phone"} tab={"手机登录"} />
        <Tabs.TabPane key={"email"} tab={"邮箱登录"} />
      </Tabs>
      {loginType === "account" && (
        <>
          <ProFormText
            name={"account"}
            fieldProps={{
              size: "large",
              prefix: <UserOutlined className={"prefixIcon"} />
            }}
            placeholder={"DouTok ID"}
            rules={[
              {
                required: true,
                message: "请输入DouTok ID"
              }
            ]}
          />
        </>
      )}
      {loginType === "phone" && (
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
      {loginType === "email" && (
        <>
          <ProFormText
            name={"phone"}
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
            message: "请输入密码！"
          }
        ]}
      />
      {props.switcher}
    </LoginForm>
  );
}
