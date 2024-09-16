import {Modal, Tabs} from "antd";

import "./LoginModal.css"
import {LoginForm, ProFormText} from "@ant-design/pro-form";
import React, {useState} from "react";
import {MailOutlined, MobileOutlined, UserOutlined} from "@ant-design/icons";
import {PasswordInput} from "@/components/PasswordInput/PasswordInput";
import {
  UserServiceLoginResponse,
  useUserServiceLogin,
} from "@/api/svapi/api";

type LoginType = 'phone' | 'account' | 'email';

export interface LoginModalProps {
  open?: boolean
  onCancel?: (e: React.MouseEvent<HTMLButtonElement>) => void
}

export function LoginModal(props: LoginModalProps) {
  const [loginType, setLoginType] = useState<LoginType>('phone');
  const loginMutate = useUserServiceLogin({});

  const submit4Login = (formData: Record<string, string>) => {
    loginMutate.mutate({
      mobile: formData?.phone,
      email: formData?.email,
      password: formData?.password
    }).then((r: UserServiceLoginResponse) => {
      console.log(r);
    });
  }

  return (
     <Modal
      open={props.open}
      onCancel={props.onCancel}
      footer={null}
     >
       <LoginForm
         className={"login-form"}
         title={"DouTok"}
         subTitle={"Cloudzneith"}
         onFinish={submit4Login}
       >
         <Tabs
           centered
           activeKey={loginType}
           onChange={(key) => setLoginType(key as LoginType)}
         >
           <Tabs.TabPane key={'account'} tab={'账号登录'} />
           <Tabs.TabPane key={'phone'} tab={'手机登录'} />
           <Tabs.TabPane key={'email'} tab={'邮箱登录'} />
         </Tabs>
         {loginType === 'account' && (
           <>
             <ProFormText
               name={"account"}
               fieldProps={{
                 size: 'large',
                 prefix: <UserOutlined className={'prefixIcon'} />,
               }}
               placeholder={'DouTok ID'}
               rules={[
                 {
                   required: true,
                   message: '请输入DouTok ID',
                 },
               ]}
             />
             <PasswordInput />
           </>
         )}
         {loginType === 'phone' && (
           <>
             <ProFormText
               name={"phone"}
               fieldProps={{
                 size: 'large',
                 prefix: <MobileOutlined  className={'prefixIcon'} />,
               }}
               placeholder={'手机号码'}
               rules={[
                 {
                   required: true,
                   message: '请输入手机号',
                 },
               ]}
             />
             <PasswordInput />
           </>
         )}
         {loginType === 'email' && (
           <>
             <ProFormText
               name={"phone"}
               fieldProps={{
                 size: 'large',
                 prefix: <MailOutlined  className={'prefixIcon'} />,
               }}
               placeholder={'邮箱地址'}
               rules={[
                 {
                   required: true,
                   message: '请输入邮箱',
                 },
               ]}
             />
             <PasswordInput />
           </>
         )}
       </LoginForm>
     </Modal>
  );
}