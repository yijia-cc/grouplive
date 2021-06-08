import React from 'react';
import { connect } from 'react-redux'
import { Form, Input } from 'antd';
import { ACTION_TYPES } from '../../../../redux/actions/actionType';
import { commentActions } from '../../../../redux/actions/commentAction';

const layout = {
    labelCol: { span: 4 },
    wrapperCol: { span: 16 },
};

class CreateComment extends React.Component {
    state = { comment: '' };
    onChange(val) {        
        this.setState({ comment: val });
    }

    componentDidUpdate() {
        if (this.props.comment.notificationFlag) {
            const payload = {
                postId: this.props.postId,
                text: this.state.comment
            }
            this.props.createComment(payload);            
        }
    }

    componentWillUnmount() {
        this.props.dispatch({ type: ACTION_TYPES.RESET_NEW_COMMENT });
    }

    render() {
        return (            
            <Form {...layout} name="comments">                               
                <Form.Item name='comment' label="Comment">
                    <Input.TextArea onChange={(e) => this.onChange(e.target.value)}/>
                </Form.Item>
            </Form>
        );
    }
}

const mapStateToProps = (state) => {
    return {
        comment: state.newCommentsReducer
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        createComment: (payload) => dispatch(commentActions.createComment(payload)),
        dispatch
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(CreateComment);