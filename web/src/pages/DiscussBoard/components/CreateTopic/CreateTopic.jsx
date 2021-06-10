import React from 'react';
import { withRouter } from 'react-router';
import { Form, Input, Button, notification } from 'antd';
import { connect } from 'react-redux'
import { topicActions } from '../../../../redux/actions/topicAction';
import { ACTION_TYPES } from '../../../../redux/actions/actionType';
import './CreateTopic.css';

const layout = {
    labelCol: { span: 6 },
    wrapperCol: { span: 14 },
};

const tailLayout = {    
    wrapperCol: { offset: 10, span: 16 },
}

const validateMessages = {
    required: '${label} is required!'
};

const PAYLOAD = {
    title: 'name',
    description: 'description'
};

const openNotificationWithIcon = (type, title, msg) => {
    notification[type]({
        message: title,
        description: msg        
    });
};

class CreateTopic extends React.Component {
    formRef = React.createRef();
    state = {
        title: '',
        description: ''
    }

    onReset = () => {
        this.formRef.current.resetFields();     
    }

    onFormChange(val, type) {
        this.setState({ [type]: val});        
    }

    submitTopic() {
        if (this.isValid()) {
            const redirectTo = '/discussion/postList';
            const payload = {
                [PAYLOAD.title]: this.state.title,
                [PAYLOAD.description]: this.state.description
            };
            this.props.createTopic(payload, this.props.history, redirectTo);
        }
    }

    isValid() {
        return this.state.title !== '';
    }

    componentDidUpdate() {
        if (this.props.newTopic.type === 'failed' && this.props.newTopic.msg !== '') {
            openNotificationWithIcon(this.props.newTopic.type, this.props.newTopic.title, this.props.newTopic.msg);
        }
    }

    componentWillUnmount() {
        this.props.dispatch({ type: ACTION_TYPES.RESET_NEW_TOPIC });
    }

    render() {        
        return (
            <div className="topicForm-container">
                <Form {...layout} ref={this.formRef} name="topicForm" 
                    className="topicForm" validateMessages={validateMessages}>
                    <Form.Item name='title' label="Title" rules={[{ required: true }]}>
                        <Input onChange={(e) => this.onFormChange(e.target.value, 'title')}/>
                    </Form.Item>                       
                    <Form.Item name='description' label="Description">
                        <Input.TextArea onChange={(e) => this.onFormChange(e.target.value, 'description')}/>
                    </Form.Item>
                    <Form.Item {...tailLayout}>
                        <Button type="primary" htmlType="submit" onClick={() => this.submitTopic()}>Submit</Button>
                        <Button htmlType="button" onClick={this.onReset}>Reset</Button>
                    </Form.Item>
                </Form>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return {
        newTopic: state.newTopicReducer
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        createTopic: (payload, history, redirectTo) => dispatch(topicActions.createTopic(payload, history, redirectTo)),
        dispatch        
    }
}

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(CreateTopic));
