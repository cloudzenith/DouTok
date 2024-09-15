import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { AntdRegistry } from "@ant-design/nextjs-registry";
import Layout, { Content } from "antd/es/layout/layout";
import { PageHeader } from "@/components/PageHeader/PageHeader";
import { PageSider } from "@/components/PageSider/PageSider";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "DouTok",
  description: "DouTok !!!"
};

export default function RootLayout({
  children
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div id={"app"} className="App">
          <AntdRegistry>
            <Layout id={"layout"}>
              <PageHeader />
              <Layout>
                <PageSider />
                <Content id={"content"}>
                  <div className={"content"}>{children}</div>
                </Content>
              </Layout>
            </Layout>
          </AntdRegistry>
        </div>
      </body>
    </html>
  );
}
