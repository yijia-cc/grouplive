import React from 'react';
import { connect } from 'react-redux'
import _ from "lodash";
import { NavLink, withRouter } from "react-router-dom";
import { Layout, Menu, Modal, Button } from 'antd';
import { HomeOutlined, NotificationOutlined, FormOutlined } from '@ant-design/icons';
import CreateTopic from '../CreateTopic/CreateTopic';
import CreatePost from '../CreatePost/CreatePost';
import './DiscussSider.css';
import { postActions } from '../../../../redux/actions/createPost';
import { topicActions } from '../../../../redux/actions/createTopic';

const { SubMenu } = Menu;
const { Sider } = Layout;

class DiscussSideBar extends React.Component {
    state = {
        loading: false,
        visible: false        
    }; 

    showModal = (modalType) => {
        this.setState({
          visible: true,
          modalType
        });
    };
    
    handleOk = () => {
        this.setState({ loading: true });
        if (this.state.modalType === 'Post') {
            if (this.isValidPost()) {                
                const payload = this.props.createpost.payload;
                this.props.getUpdatePostList(payload);
                this.setState({ 
                    loading: false,
                    visible: false
                });
            }
        } 
        else {
            if (this.isValidTopic()) {
                const payload = this.props.createtopic.payload;
                this.props.getUpdateTopicList(payload);
                this.setState({ 
                    loading: false,
                    visible: false
                });               
            }
        }        
    };
    
    handleCancel = () => {
        this.setState({ visible: false });
    };

    isValidPost = () => {
        const postData = _.cloneDeep(this.props.createpost);
        delete postData.payload;
        delete postData.postList;
        return Object.keys(postData).every((key) => this.props.createpost[key].isValid);        
    }

    isValidTopic = () => {
        const topicData = _.cloneDeep(this.props.createtopic);
        delete topicData.payload;        
        const isValid = Object.keys(topicData).every((key) => this.props.createtopic[key].isValid);
        return isValid;
    }    
    
    render() {
        const { path } = this.props.match;
        const { visible, loading } = this.state;
        return (
            <div>
                <Sider width={200} className="discuss-layout-background">          
                    <Menu
                        mode="inline"
                        defaultSelectedKeys={['1']}
                        defaultOpenKeys={['home']}
                        style={{ height: '100%', borderRight: 0 }}>  
                        <Menu.Item icon={<HomeOutlined />} key="'home'">
                            Home
                            <NavLink to={`${path}/home`}></NavLink>
                        </Menu.Item> 
                        <Menu.Item icon={<FormOutlined />} key="'create-topic'" onClick={() => this.showModal('Topic')}>
                            Create Topic
                        </Menu.Item> 
                        <Menu.Item icon={<FormOutlined />} key="'create-post'" onClick={() => this.showModal('Post')}>                        
                            Create Post
                        </Menu.Item>
                        {/* <SubMenu key="sub3" icon={<NotificationOutlined />} title="Browse Topics">
                            <Menu.Item key="9">topic1</Menu.Item>                      
                        </SubMenu> */}
                    </Menu>
                </Sider>
                <Modal
                    visible={visible}
                    title={`Create ${this.state.modalType}`}
                    onOk={this.handleOk}
                    onCancel={this.handleCancel}
                    width={800}
                    footer={[
                        <Button key="submit" type="primary" onClick={this.handleOk}>
                            Create
                        </Button>,
                        <Button key="back" onClick={this.handleCancel}>
                            Cancel
                        </Button>                        
                    ]}>
                        {this.state.modalType == 'Topic' ? <CreateTopic /> : <CreatePost />}  
                </Modal>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return {
        createpost: state.postReducer,
        createtopic: state.topicReducer
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        getUpdatePostList: (payload) => dispatch(postActions.getUpdatePostList(payload)),
        getUpdateTopicList: (payload) => dispatch(topicActions.getUpdateTopicList(payload))
    };
}

const DiscussSider = connect(mapStateToProps, mapDispatchToProps)(DiscussSideBar);

export default withRouter(DiscussSider);