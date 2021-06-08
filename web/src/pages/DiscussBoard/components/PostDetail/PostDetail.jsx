import React from 'react';
import { connect } from 'react-redux'
import { withRouter } from "react-router-dom";
import { Card, Layout, Space, Button, Modal, Comment, notification } from 'antd';
import { CommentOutlined } from '@ant-design/icons';
import { postActions } from '../../../../redux/actions/postAction';
import CreateComment from '../CreateComment/CreateComment';
import { ACTION_TYPES } from '../../../../redux/actions/actionType';
import './PostDetail.css';

const { Content } = Layout;
const IconText = ({ icon, text, that }) => {    
    return (
        <div>
            <Space>
                <a onClick={() => iconClick(that)}>{React.createElement(icon)} {text}</a>                
            </Space>
        </div>
    );    
};

const iconClick = (that) => {
    that.props.getComments(that.state.postId);
    that.setState({ readComment: true });
};

const ShowCommentList = ({ that }) => {
    return (        
        that.state.comments.map((comment) => {            
            return (
                <Comment
                    author={comment.userName}
                    key={`${comment.id} ${comment.text}`}
                    content={
                        <p>{comment.text}</p>
                    }
                />
            );
        })
    );
}

const openNotificationWithIcon = (type, title, msg) => {
    notification[type]({
        message: title,
        description: msg        
    });
};

class PostDefail extends React.Component {
    state = { 
        post: {}, 
        comments: [], 
        readComment: false, 
        loading: false, 
        visible: false,
        postId: -1
    };

    componentDidMount() {
        const { url } = this.props.match;
        const lastSlashIdx = url.lastIndexOf('/');
        if (lastSlashIdx + 1 < url.length) {
            const postId = url.substring(lastSlashIdx + 1);
            this.props.getPost(postId);
            this.setState({ postId });
        }        
    }

    componentDidUpdate() {        
        if (this.props.post.data.postId && this.state.post.postId != this.state.postId) { // first time load post
            this.setState({ post: this.props.post.data });
        }         
        else {            
            if (this.props.newComments.data.length > this.state.comments.length ||
                this.props.comments.data.length > this.state.comments.length) {
                if (this.props.newComments.data.length > this.state.comments.length) {
                    openNotificationWithIcon(this.props.newComments.type, this.props.newComments.title, this.props.newComments.msg);
                }
                const latestComments = this.props.comments.data.length > this.props.newComments.data.length ? this.props.comments.data : this.props.newComments.data;
                this.setState({ comments: [...latestComments], readComment: true });
            }
        }        
    }

    componentWillUnmount() {
        this.props.dispatch({ type: ACTION_TYPES.RESET_NEW_COMMENT });
    }

    openCommentModal() {        
        this.setState({ visible: true });
    }

    handleOk = () => {        
        this.props.dispatch({ type: ACTION_TYPES.CREATE_COMMENT_NOTIFICATION });
        this.setState({ visible: false });
    };
    
    handleCancel = () => {        
        this.props.dispatch({ type: ACTION_TYPES.RESET_NEW_COMMENT });
        this.setState({ visible: false });
    };

    render() {
        return (
            <Layout className="discuss-main-layout">
                <Content className="site-layout-background discuss-main-content">                    
                    <Card 
                        title={this.props.post.data.postName} 
                        extra={<Button type="primary" onClick={() => this.openCommentModal()}>Create Comment</Button>} 
                        style={{ width: '100%' }}>
                        <p>{this.props.post.data.description}</p>                         
                        {
                            this.state.readComment ? <ShowCommentList that={this} /> :
                            <IconText icon={ CommentOutlined } text={this.props.post.data.commentCount} that={this} key="comments" />                      
                        }                     
                    </Card>
                    <Modal
                        visible={this.state.visible}
                        title="Create Comment"
                        onOk={this.handleOk}
                        onCancel={this.handleCancel}
                        width={800}
                        destroyOnClose={true}
                        footer={[
                            <Button key="submit" type="primary" onClick={this.handleOk}>
                                Create
                            </Button>,
                            <Button key="back" onClick={this.handleCancel}>
                                Cancel
                            </Button>                        
                        ]}>
                        <CreateComment postId={this.state.postId} />
                    </Modal>
                </Content>
            </Layout>
        );
    }
}

const mapStateToProps = (state) => {
    return {
        post: state.getSinglePostReducer,
        comments: state.getPostCommentsReducer,
        newComments: state.newCommentsReducer
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        getPost: (postId) => dispatch(postActions.getSinglePost(postId)),
        getComments: (postId) => dispatch(postActions.getPostComments(postId)),        
        dispatch
    };
}

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(PostDefail));