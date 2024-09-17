import {Modal, notification, Tabs} from "antd";

import "./LoginModal.css"
import {LoginForm, ProFormText} from "@ant-design/pro-form";
import React, {useEffect, useState} from "react";
import {MailOutlined, MobileOutlined, UserOutlined} from "@ant-design/icons";
import {PasswordInput} from "@/components/PasswordInput/PasswordInput";
import {
  UserServiceLoginResponse,
  useUserServiceLogin,
} from "@/api/svapi/api";
import {LoginComponent} from "@/components/LoginComponent/LoginComponent";
import {RegisterComponent} from "@/components/RegisterComponent/RegisterComponent";

type ModalType = 'login' | 'register';

export interface LoginModalProps {
  open?: boolean
  onCancel?: (e: React.MouseEvent<HTMLButtonElement>) => void
  type: string
}

export function LoginModal(props: LoginModalProps) {
  const [modalType, setModalType] = useState<ModalType>('login');

  useEffect(() => {
    setModalType(props.type as ModalType)
  }, []);

  return (
     <Modal
      open={props.open}
      onCancel={props.onCancel}
      footer={null}
     >
       {modalType === 'login' && (<LoginComponent switcher={(
         <a
           style={{
             float: 'right',
             marginBottom: '10px'
           }}
           onClick={() => setModalType('register')}
         >还没账号？点击注册</a>
       )} />)}
       {modalType === 'register' && (<RegisterComponent switcher={(
         <a
           style={{
             float: 'right',
             marginBottom: '10px'
          }}
           onClick={() => setModalType('login')}
         >已有账号？点击登录</a>
       )} />)}
     </Modal>
  );
}