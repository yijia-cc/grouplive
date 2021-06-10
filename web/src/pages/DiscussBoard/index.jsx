import React from 'react';
import { Layout, Menu, Breadcrumb } from 'antd';
import DiscussHome from './components/DiscussHome/DiscussHome';
import DiscussSider from './components/DiscussSider/DiscussSider';



const DiscussBoard = () => {
  return (
    <Layout>      
      <Layout>
        <DiscussSider />
        <DiscussHome />
      </Layout>
    </Layout>
  );
};

export default DiscussBoard;
